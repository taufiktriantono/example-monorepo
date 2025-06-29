package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/taufiktriantono/api-first-monorepo/pkg/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("http_server",
	fx.Provide(New),
	fx.Invoke(Run),
)

type Server struct {
	server   *http.Server
	tlsMutex sync.RWMutex
	cert     *tls.Certificate
	certPath string
	keyPath  string
}

type Params struct {
	fx.In
	Config  *config.Config
	Handler *gin.Engine
}

func New(p Params) *Server {
	cfg := p.Config
	srv := &Server{
		server: &http.Server{
			Addr:         fmt.Sprintf(":%s", cfg.Server.Addr),
			Handler:      p.Handler,
			ReadTimeout:  cfg.Server.ReadTimeout,
			WriteTimeout: cfg.Server.WriteTimeout,
			IdleTimeout:  cfg.Server.IdleTimeout,
		},
		certPath: cfg.Server.TLS.CertPath,
		keyPath:  cfg.Server.TLS.KeyPath,
	}

	if cfg.Server.TLS.Enable {
		srv.reloadCert() // initial load
		go srv.watchTLSFiles()

		srv.server.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
			GetCertificate: func(info *tls.ClientHelloInfo) (*tls.Certificate, error) {
				srv.tlsMutex.RLock()
				defer srv.tlsMutex.RUnlock()

				if srv.cert == nil {
					return nil, fmt.Errorf("no TLS cert loaded")
				}

				return srv.cert, nil
			},
		}
	}

	return srv
}

// Reload TLS certificate
func (s *Server) reloadCert() {
	cert, err := tls.LoadX509KeyPair(s.certPath, s.keyPath)
	if err != nil {
		zap.L().Error("failed to reload TLS cert", zap.Error(err))
		return
	}
	s.tlsMutex.Lock()
	s.cert = &cert
	s.tlsMutex.Unlock()
	zap.L().Info("TLS certificate reloaded")
}

// Watch TLS cert/key file
func (s *Server) watchTLSFiles() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		zap.L().Error("failed to create fsnotify watcher", zap.Error(err))
		return
	}
	defer watcher.Close()

	_ = watcher.Add(s.certPath)
	_ = watcher.Add(s.keyPath)

	for retry := 0; retry < 3; retry++ {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Rename) != 0 {
				s.reloadCert()
			}
		case err := <-watcher.Errors:
			zap.L().Error("watcher error", zap.Error(err))
		}
	}
}

func Run(lc fx.Lifecycle, srv *Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if srv.server.TLSConfig != nil {
				zap.L().Info("Starting HTTP server with mtls", zap.String("addr", srv.server.Addr))
				go srv.server.ListenAndServeTLS(srv.certPath, srv.keyPath)
			} else {
				zap.L().Info("Starting HTTP server with non mtls", zap.String("addr", srv.server.Addr))
				go srv.server.ListenAndServe()
			}

			return nil
		},
		OnStop: func(ctx context.Context) error {
			zap.L().Info("Shutting down HTTP server gracefully...")
			return srv.server.Shutdown(ctx)
		},
	})
}
