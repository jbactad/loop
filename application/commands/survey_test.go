package commands_test

import (
	"context"
	"testing"
	"time"

	"github.com/jbactad/loop/application/commands"
	"github.com/jbactad/loop/application/ports/mocks"
	"github.com/jbactad/loop/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCommands_CreateSurvey(t *testing.T) {
	type args struct {
		ctx context.Context
		cmd commands.CreateSurveyCommand
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Survey
		wantErr assert.ErrorAssertionFunc
		setup   func(sm *mocks.SurveyCreatorProvider, ug *mocks.UUIDGenerator, tp *mocks.TimeProvider)
	}{
		{
			name: "given a valid command, then return a survey",
			args: args{
				ctx: context.Background(),
				cmd: commands.CreateSurveyCommand{
					Name:        "Test Survey",
					Description: "Test Description",
					Question:    "Test Question",
				},
			},
			setup: func(sm *mocks.SurveyCreatorProvider, ug *mocks.UUIDGenerator, tp *mocks.TimeProvider) {
				uid := "test-uuid"
				now := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
				ug.EXPECT().Generate().Return(uid).Once()
				tp.EXPECT().Now().Return(now).Once()
				s := domain.NewSurvey(uid, "Test Survey", "Test Description", "Test Question", now, now)
				sm.EXPECT().CreateSurvey(mock.IsType(context.Background()), s).Return(nil).Once()
			},
			want: func() *domain.Survey {
				uid := "test-uuid"
				now := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
				s := domain.NewSurvey(uid, "Test Survey", "Test Description", "Test Question", now, now)
				return s
			}(),
			wantErr: assert.NoError,
		},
		{
			name: "given a command with empty name, then return error",
			args: args{
				ctx: context.Background(),
				cmd: commands.CreateSurveyCommand{
					Name:        "",
					Description: "Test Description",
					Question:    "Test Question",
				},
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "given a command with empty description, then return error",
			args: args{
				ctx: context.Background(),
				cmd: commands.CreateSurveyCommand{
					Name:        "Test Survey",
					Description: "",
					Question:    "Test Question",
				},
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "given a command with empty question, then return error",
			args: args{
				ctx: context.Background(),
				cmd: commands.CreateSurveyCommand{
					Name:        "Test Survey",
					Description: "Test Description",
					Question:    "",
				},
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scm := mocks.NewSurveyCreatorProvider(t)
			ug := mocks.NewUUIDGenerator(t)
			tp := mocks.NewTimeProvider(t)

			if tt.setup != nil {
				tt.setup(scm, ug, tp)
			}

			cs := commands.New(scm, nil, ug, tp)

			got, err := cs.CreateSurvey(tt.args.ctx, tt.args.cmd)
			if !tt.wantErr(t, err, "Commands.CreateSurvey() error = %v, wantErr %v", err, tt.wantErr) {
				return
			}

			assert.Equalf(t, tt.want, got, "Commands.CreateSurvey() = %v, want %v", got, tt.want)
		})
	}
}
