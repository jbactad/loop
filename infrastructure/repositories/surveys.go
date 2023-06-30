package repositories

import (
	"context"
	"time"

	"github.com/jbactad/loop/domain"
)

type SurveyData struct {
	ID          string `gorm:"primaryKey"`
	Name        string
	Description string
	Question    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewSurveyData(
	s *domain.Survey,
) (obj *SurveyData) {
	obj = &SurveyData{
		ID:          s.ID(),
		Name:        s.Name(),
		Description: s.Description(),
		Question:    s.Question(),
		CreatedAt:   s.CreatedAt(),
		UpdatedAt:   s.UpdatedAt(),
	}
	return
}

func (s SurveyData) ToDomain() *domain.Survey {
	return domain.NewSurvey(s.ID, s.Name, s.Description, s.Question, s.CreatedAt, s.UpdatedAt)
}

type SurveyRepository struct {
	db Database
}

func NewSurveyRepository(dbConn Database) *SurveyRepository {
	return &SurveyRepository{
		db: dbConn,
	}
}

func (repo *SurveyRepository) GetSurveys(
	ctx context.Context, limit int, offset int,
) (res []*domain.Survey, err error) {
	var ss []*SurveyData
	err = repo.db.
		Table("surveys").
		Limit(limit).
		Offset(offset).
		Find(&ss).
		Error()
	if err != nil {
		return nil, err
	}

	for _, s := range ss {
		res = append(res, s.ToDomain())
	}

	return res, nil
}

func (repo *SurveyRepository) CreateSurvey(ctx context.Context, survey *domain.Survey) error {
	s := NewSurveyData(survey)

	return repo.db.Table("surveys").Create(s).Error()
}

func (repo *SurveyRepository) GetSurvey(ctx context.Context, id string) (*domain.Survey, error) {
	panic("implement me")
}
