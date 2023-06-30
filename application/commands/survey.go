package commands

import (
	"context"
	"errors"

	"github.com/jbactad/loop/domain"
)

type CreateSurveyCommand struct {
	Name        string
	Description string
	Question    string
}

var (
	ErrInvalidName        = errors.New("invalid name")
	ErrInvalidDescription = errors.New("invalid description")
	ErrInvalidQuestion    = errors.New("invalid question")
)

func (cs *Commands) CreateSurvey(ctx context.Context, cmd CreateSurveyCommand) (*domain.Survey, error) {
	if cmd.Name == "" {
		return nil, ErrInvalidName
	}
	if cmd.Description == "" {
		return nil, ErrInvalidDescription
	}
	if cmd.Question == "" {
		return nil, ErrInvalidQuestion
	}
	now := cs.timeProvider.Now()
	id := cs.uuidGenerator.Generate()

	s := domain.NewSurvey(id, cmd.Name, cmd.Description, cmd.Question, now, now)

	err := cs.surveyCreatorProvider.CreateSurvey(ctx, s)
	if err != nil {
		return nil, err
	}

	return s, nil
}
