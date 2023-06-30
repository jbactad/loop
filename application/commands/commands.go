package commands

import (
	"context"

	"github.com/jbactad/loop/application/ports"
	"github.com/jbactad/loop/domain"
)

//go:generate mockery --name UseCases --output ./mocks --filename commands.go --with-expecter
type UseCases interface {
	CreateSurvey(ctx context.Context, cmd CreateSurveyCommand) (*domain.Survey, error)
}

type Commands struct {
	surveyCreatorProvider         ports.SurveyCreatorProvider
	surveyResponseCreatorProvider ports.SurveyResponseCreatorProvider
	uuidGenerator                 ports.UUIDGenerator
	timeProvider                  ports.TimeProvider
}

func New(scp ports.SurveyCreatorProvider, scrp ports.SurveyResponseCreatorProvider, uuidGenerator ports.UUIDGenerator, timeProvider ports.TimeProvider) *Commands {
	return &Commands{
		surveyCreatorProvider:         scp,
		surveyResponseCreatorProvider: scrp,
		uuidGenerator:                 uuidGenerator,
		timeProvider:                  timeProvider,
	}
}
