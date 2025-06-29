package repository

import (
	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/domain"
	"github.com/taufiktriantono/api-first-monorepo/pkg/repository"
	"go.uber.org/fx"
)

var Module = fx.Module("repositories",
	fx.Provide(
		repository.ProvideStore[domain.ApprovalTemplate],
		repository.ProvideStore[domain.ApprovalTemplateStep],
		repository.ProvideStore[domain.Approval],
		repository.ProvideStore[domain.ApprovalStep],
	),
	fx.Provide(
		NewApprovalTemplateRepository,
		NewApprovalTemplateStepRepository,
		NewApprovalRepository,
		NewApprovalStepRepository,
	),
)
