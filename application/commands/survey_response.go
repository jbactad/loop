package commands

import (
	"context"
	"errors"

	"github.com/jbactad/loop/domain"
)

type CreateSurveyResponseCommand struct {
	SurveyID string
	Answer   string
	Rating   int
}

var (
	ErrInvalidSurveyID = errors.New("invalid survey id")
	ErrSurveyNotFound  = errors.New("survey not found")
	ErrInvalidRating   = errors.New("invalid rating")
)

func (cs *Commands) CreateSurveyResponse(
	ctx context.Context, cmd CreateSurveyResponseCommand,
) (*domain.SurveyResponse, error) {
	if cmd.SurveyID == "" {
		return nil, ErrInvalidSurveyID
	}
	if cmd.Rating < 0 || cmd.Rating > 5 {
		return nil, ErrInvalidRating
	}
	s, err := cs.surveyCreatorProvider.GetSurvey(ctx, cmd.SurveyID)
	if err != nil {
		return nil, ErrSurveyNotFound
	}

	i := cs.uuidGenerator.Generate()
	n := cs.timeProvider.Now()
	sr := domain.NewSurveyResponse(i, s, cmd.Answer, cmd.Rating, n, n)

	err = cs.surveyResponseCreatorProvider.CreateSurveyResponse(ctx, sr)
	if err != nil {
		return nil, err
	}

	return sr, nil
}
