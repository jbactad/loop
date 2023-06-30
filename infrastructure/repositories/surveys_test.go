package repositories_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/jbactad/loop/domain"
	"github.com/jbactad/loop/infrastructure/repositories"
	"github.com/jbactad/loop/infrastructure/repositories/mocks"
	"github.com/stretchr/testify/mock"
)

func TestSurveyRepository_GetSurveys(t *testing.T) {
	var testSurveyData []*repositories.SurveyData
	err := faker.FakeData(&testSurveyData)
	if err != nil {
		t.Error(err)
	}

	ctx := context.Background()
	type args struct {
		limit  int
		offset int
	}
	tests := []struct {
		name        string
		args        args
		wantSurveys []*domain.Survey
		wantErr     bool
		setup       func(db *mocks.Database)
	}{
		{
			name: "given a list of surveys, then return a list of surveys",
			setup: func(db *mocks.Database) {
				db.EXPECT().Error().Times(1).Return(nil)
				db.EXPECT().Table("surveys").Times(1).Return(db)
				db.EXPECT().Limit(10).Times(1).Return(db)
				db.EXPECT().Offset(0).Times(1).Return(db)
				db.EXPECT().
					Find(mock.Anything, mock.Anything).
					Run(func(dest interface{}, conds ...interface{}) {
						reflect.ValueOf(dest).Elem().Set(reflect.ValueOf(testSurveyData))
					}).
					Times(1).
					Return(db)
			},
			args: args{
				limit:  10,
				offset: 0,
			},
			wantSurveys: func() []*domain.Survey {
				var surveys []*domain.Survey
				for _, surveyData := range testSurveyData {
					surveys = append(surveys, surveyData.ToDomain())
				}

				return surveys
			}(),
		},
		{
			name: "given db returns an error, then return an error",
			setup: func(db *mocks.Database) {
				db.EXPECT().Error().Times(1).Return(errors.New("error happened"))
				db.EXPECT().Table("surveys").Times(1).Return(db)
				db.EXPECT().Limit(10).Times(1).Return(db)
				db.EXPECT().Offset(0).Times(1).Return(db)
				db.EXPECT().Find(mock.Anything, mock.Anything).Times(1).Return(db)
			},
			args: args{
				limit:  10,
				offset: 0,
			},
			wantSurveys: nil,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := mocks.NewDatabase(t)
			if tt.setup != nil {
				tt.setup(db)
			}

			repo := repositories.NewSurveyRepository(db)

			gotSurveys, err := repo.GetSurveys(ctx, tt.args.limit, tt.args.offset)

			if (err != nil) != tt.wantErr {
				t.Errorf("SurveyRepository.GetSurveys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSurveys, tt.wantSurveys) {
				t.Errorf("SurveyRepository.GetSurveys() = %v, want %v", gotSurveys, tt.wantSurveys)
			}
		})
	}
}

func TestSurveyRepository_CreateSurvey(t *testing.T) {
	var survey domain.Survey
	err := faker.FakeData(&survey)
	if err != nil {
		t.Error(err)
	}
	type args struct {
		ctx    context.Context
		survey *domain.Survey
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		setup   func(db *mocks.Database)
	}{
		{
			name: "given a survey, then create a survey",
			args: args{
				ctx:    context.Background(),
				survey: &survey,
			},
			setup: func(db *mocks.Database) {
				db.EXPECT().Error().Times(1).Return(nil)
				db.EXPECT().Table("surveys").Times(1).Return(db)
				s := repositories.NewSurveyData(&survey)
				db.EXPECT().Create(s).Times(1).Return(db)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := mocks.NewDatabase(t)
			if tt.setup != nil {
				tt.setup(db)
			}

			f := repositories.NewSurveyRepository(db)

			if err := f.CreateSurvey(tt.args.ctx, tt.args.survey); (err != nil) != tt.wantErr {
				t.Errorf("SurveyRepository.CreateSurvey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSurveyRepository_GetSurvey(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Survey
		wantErr bool
		setup   func(db *mocks.Database)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := mocks.NewDatabase(t)
			if tt.setup != nil {
				tt.setup(db)
			}

			repo := repositories.NewSurveyRepository(db)

			got, err := repo.GetSurvey(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("SurveyRepository.GetSurvey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SurveyRepository.GetSurvey() = %v, want %v", got, tt.want)
			}
		})
	}
}
