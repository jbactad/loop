package ports

import (
	"context"

	"github.com/jbactad/loop/domain"
)

//go:generate mockery --name=SurveyProvider --output=./mocks --outpkg=mocks --with-expecter
type SurveyProvider interface {
	GetSurveys(ctx context.Context, limit int, offset int) ([]*domain.Survey, error)
	GetSurvey(ctx context.Context, id string) (*domain.Survey, error)
}

//go:generate mockery --name=SurveyCreator --output=./mocks --outpkg=mocks --with-expecter
type SurveyCreator interface {
	CreateSurvey(ctx context.Context, survey *domain.Survey) error
}

//go:generate mockery --name=SurveyCreatorProvider --output=./mocks --outpkg=mocks --with-expecter
type SurveyCreatorProvider interface {
	SurveyCreator
	SurveyProvider
}
