package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"

	"github.com/jbactad/loop/application/queries"
	"github.com/jbactad/loop/graph/generated"
	"github.com/jbactad/loop/graph/mappers"
	"github.com/jbactad/loop/graph/models"
	"github.com/mehdihadeli/go-mediatr"
)

// Surveys is the resolver for the surveys field.
func (r *queryResolver) Surveys(ctx context.Context, limit *int, page *int) ([]*models.Survey, error) {
	if limit == nil {
		limit = new(int)
		*limit = 10
	}
	if page == nil {
		page = new(int)
		*page = 0
	}
	result, err := mediatr.Send[queries.GetSurveysQuery, queries.GetSurveysQueryResponse](ctx, queries.GetSurveysQuery{
		Limit: *limit,
		Page:  *page,
	})
	if err != nil {
		return nil, err
	}
	surveys := mappers.SurveysToSurveysResponse(result.Surveys)

	return surveys, nil
}

// Survey is the resolver for the survey field.
func (r *queryResolver) Survey(ctx context.Context, id string) (*models.Survey, error) {
	panic(fmt.Errorf("not implemented: Survey - survey"))
}

// SurveyResponses is the resolver for the surveyResponses field.
func (r *queryResolver) SurveyResponses(ctx context.Context) ([]*models.SurveyResponse, error) {
	panic(fmt.Errorf("not implemented: SurveyResponses - surveyResponses"))
}

// SurveyResponse is the resolver for the surveyResponse field.
func (r *queryResolver) SurveyResponse(ctx context.Context, id string) (*models.SurveyResponse, error) {
	panic(fmt.Errorf("not implemented: SurveyResponse - surveyResponse"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }