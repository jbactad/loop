package ports

import (
	"context"

	"github.com/jbactad/loop/domain"
)

//go:generate mockery --name=SurveyResponseProvider --output=./mocks --outpkg=mocks --with-expecter
type SurveyResponseProvider interface {
	GetSurveyResponses(ctx context.Context, limit int, offset int) ([]*domain.SurveyResponse, error)
}

//go:generate mockery --name=SurveyResponseCreator --output=./mocks --outpkg=mocks --with-expecter
type SurveyResponseCreator interface {
	CreateSurveyResponse(ctx context.Context, survey *domain.SurveyResponse) error
}

//go:generate mockery --name=SurveyResponseCreatorProvider --output=./mocks --outpkg=mocks --with-expecter
type SurveyResponseCreatorProvider interface {
	SurveyResponseCreator
	SurveyResponseProvider
}
