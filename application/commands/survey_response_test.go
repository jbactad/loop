package commands_test

import (
	"context"
	"testing"
	"time"

	"github.com/jbactad/loop/application/commands"
	"github.com/jbactad/loop/application/ports/mocks"
	"github.com/jbactad/loop/domain"
	"github.com/stretchr/testify/assert"
)

func TestCommands_CreateSurveyResponse(t *testing.T) {
	now := time.Date(2023, 1, 27, 0, 0, 0, 0, time.UTC)
	defaultCtx := context.Background()
	type args struct {
		ctx context.Context
		cmd commands.CreateSurveyResponseCommand
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.SurveyResponse
		wantErr assert.ErrorAssertionFunc
		setup   func(scp *mocks.SurveyCreatorProvider, srcp *mocks.SurveyResponseCreatorProvider, ug *mocks.UUIDGenerator, tp *mocks.TimeProvider)
	}{
		{
			name: "given a valid command, then return a survey response",
			args: args{
				ctx: defaultCtx,
				cmd: commands.CreateSurveyResponseCommand{
					SurveyID: "test-survey-id",
					Answer:   "Test Answer",
					Rating:   5,
				},
			},
			want: func() *domain.SurveyResponse {
				s := domain.NewSurvey(
					"test-survey-id", "Test Survey", "Test Description",
					"Test Question", now.Add(-1*time.Second), now.Add(-1*time.Second))

				return domain.NewSurveyResponse("test-uuid", s, "Test Answer", 5, now.Add(2*time.Second), now.Add(2*time.Second))
			}(),
			setup: func(scp *mocks.SurveyCreatorProvider, srcp *mocks.SurveyResponseCreatorProvider, ug *mocks.UUIDGenerator, tp *mocks.TimeProvider) {
				uid := "test-uuid"
				ug.EXPECT().Generate().Return(uid).Once()
				tp.EXPECT().Now().Return(now.Add(2 * time.Second)).Once()

				s := domain.NewSurvey(
					"test-survey-id", "Test Survey", "Test Description",
					"Test Question", now.Add(-1*time.Second), now.Add(-1*time.Second))

				scp.EXPECT().GetSurvey(defaultCtx, "test-survey-id").Return(s, nil).Once()

				sr := domain.NewSurveyResponse(uid, s, "Test Answer", 5, now.Add(2*time.Second), now.Add(2*time.Second))

				srcp.EXPECT().CreateSurveyResponse(defaultCtx, sr).Return(nil).Once()
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scp := mocks.NewSurveyCreatorProvider(t)
			srcp := mocks.NewSurveyResponseCreatorProvider(t)
			tp := mocks.NewTimeProvider(t)
			uig := mocks.NewUUIDGenerator(t)
			if tt.setup != nil {
				tt.setup(scp, srcp, uig, tp)
			}

			cs := commands.New(scp, srcp, uig, tp)

			got, err := cs.CreateSurveyResponse(tt.args.ctx, tt.args.cmd)

			if !tt.wantErr(t, err, "Commands.CreateSurveyResponse() error = %v, wantErr %v", err, tt.wantErr) ||
				err != nil {
				return
			}

			assert.EqualValuesf(t, tt.want, got, "Commands.CreateSurveyResponse() = %v, want %v", got, tt.want)
		})
	}
}