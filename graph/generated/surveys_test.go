package generated_test

import (
	"testing"
	"time"

	"github.com/99designs/gqlgen/client"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/jbactad/loop/application/ports/mocks"
	"github.com/jbactad/loop/graph/models"
	"github.com/stretchr/testify/assert"
)

func TestSurveys(t *testing.T) {
	type args struct {
		query string
		limit int
		page  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "with valid query, it should return surveys without error",
			args: args{
				query: `query Surveys($limit: Int, $page: Int) {
  surveys(limit: $limit, page: $page) {
    id
    description
    name
    question
  }
}`,
				limit: 1,
				page:  0,
			},
		},
		{
			name: "with invalid query, it should return error",
			args: args{
				query: `query Surveys($limit: Int, $page: Int) {
					  surveys(limit: $limit, page: $page) {
						id
						description
						name
						question

						# This is invalid
						invalid
					  }
					}`,
				limit: 10,
				page:  0,
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewTestClient(t, mocks.NewTimeProvider(t), mocks.NewUUIDGenerator(t))

			var resp struct {
				Surveys []struct {
					ID          string
					Description string
					Name        string
					Question    string
					CreatedAt   string
					UpdatedAt   string
				}
			}

			err := c.Post(tt.args.query,
				&resp,
				client.Var("limit", tt.args.limit),
				client.Var("page", tt.args.page),
			)
			if tt.wantErr != nil {
				tt.wantErr(t, err)
				return
			}

			snaps.MatchJSON(t, resp)
		})
	}
}

func TestCreateSurvey(t *testing.T) {
	type args struct {
		query string
		input models.NewSurvey
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
		setup   func(tp *mocks.TimeProvider, uig *mocks.UUIDGenerator)
	}{
		{
			name: "with valid query, it should return survey without error",
			args: args{
				query: `mutation CreateSurvey($input: NewSurvey!) {
  createSurvey(input: $input) {
    createdAt
    description
    id
    name
    question
    updatedAt
  }
}`,
				input: models.NewSurvey{
					Name:        "Survey 1",
					Description: "Survey 1 description",
					Question:    "Survey 1 question",
				},
			},
			setup: func(tp *mocks.TimeProvider, uig *mocks.UUIDGenerator) {
				tn := time.Date(2023, 1, 27, 12, 0, 0, 0, time.UTC)
				tp.EXPECT().Now().Return(tn)
				uid := "123e4567-e89b-12d3-a456-426614174000"
				uig.EXPECT().Generate().Return(uid)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeProvider := mocks.NewTimeProvider(t)
			uuidGenerator := mocks.NewUUIDGenerator(t)

			c := NewTestClient(t, timeProvider, uuidGenerator)

			if tt.setup != nil {
				tt.setup(timeProvider, uuidGenerator)
			}

			var resp map[string]interface{}
			err := c.Post(tt.args.query, &resp, client.Var("input", tt.args.input))
			if tt.wantErr != nil {
				tt.wantErr(t, err)
				return
			}

			snaps.MatchJSON(t, resp)
		})
	}
}
