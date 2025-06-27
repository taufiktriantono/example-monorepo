package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/domain"
	"github.com/taufiktriantono/api-first-monorepo/internal/approvals/v1/repository"
	"go.uber.org/mock/gomock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestApprovalService_GetTemplateByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTemplRepo := repository.NewMockApprovalTemplateRepository(ctrl)

	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	svc := approvalservice{
		db:               db,
		approvaltemplate: mockTemplRepo,
	}

	expected := &domain.ApprovalTemplate{
		ID:   "a2dd08ee-62b4-4ac4-8bd4-e7c0514ed007",
		Slug: "my-template",
	}

	tests := []struct {
		name     string
		input    string
		mockFn   func()
		expected *domain.ApprovalTemplate
		wantErr  bool
		errCode  string
	}{
		{
			name:  "OK with uuid",
			input: "a2dd08ee-62b4-4ac4-8bd4-e7c0514ed007",
			mockFn: func() {
				mockTemplRepo.EXPECT().
					FindOne(gomock.Any(), gomock.AssignableToTypeOf(&domain.ApprovalTemplate{})).
					DoAndReturn(func(_ context.Context, tpl *domain.ApprovalTemplate) (*domain.ApprovalTemplate, error) {
						if tpl.ID == expected.ID {
							return expected, nil
						}
						return nil, nil
					})
			},
			expected: expected,
			wantErr:  false,
		},
		{
			name:  "OK with slug",
			input: "my-template",
			mockFn: func() {
				mockTemplRepo.EXPECT().
					FindOne(gomock.Any(), gomock.AssignableToTypeOf(&domain.ApprovalTemplate{})).
					DoAndReturn(func(_ context.Context, tpl *domain.ApprovalTemplate) (*domain.ApprovalTemplate, error) {
						if tpl.Slug == expected.Slug {
							return expected, nil
						}
						return nil, nil
					})
			},
			expected: expected,
			wantErr:  false,
		},
		{
			name:  "TEMPLATE_NOT_FOUND",
			input: "a2dd08ee-62b4-4ac4-8bd4-e7c0514ed999", // UUID valid
			mockFn: func() {
				mockTemplRepo.EXPECT().
					FindOne(gomock.Any(), &domain.ApprovalTemplate{ID: "a2dd08ee-62b4-4ac4-8bd4-e7c0514ed999"}).
					Return(nil, nil)
			},
			expected: nil,
			wantErr:  true,
			errCode:  "TEMPLATE_NOT_FOUND",
		},
		{
			name:  "DATABASE ERROR",
			input: "a2dd08ee-62b4-4ac4-8bd4-e7c0514ed998",
			mockFn: func() {
				mockTemplRepo.EXPECT().
					FindOne(gomock.Any(), &domain.ApprovalTemplate{ID: "a2dd08ee-62b4-4ac4-8bd4-e7c0514ed998"}).
					Return(nil, errors.New("db error"))
			},
			expected: nil,
			wantErr:  true,
			errCode:  "", // bebas
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn() // Set mock expectation

			result, err := svc.GetTemplateByID(context.Background(), tt.input)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expected, result)
		})
	}
}
