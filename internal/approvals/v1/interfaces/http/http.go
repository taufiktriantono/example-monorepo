package http

import "go.uber.org/fx"

type ApprovalHandler interface{}

type approval struct{}

type Params struct {
	fx.In
}

func ProvideHandler(p Params) ApprovalHandler {
	return &approval{}
}
