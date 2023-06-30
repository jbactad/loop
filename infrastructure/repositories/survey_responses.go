package repositories

import (
	"context"
	"time"

	"github.com/jbactad/loop/domain"
)

type SurveyResponsesData struct {
	ID        string      `gorm:"primaryKey"`
	Survey    *SurveyData `gorm:"foreignKey:SurveyID"`
	Answer    string
	Rating    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (sr SurveyResponsesData) ToDomain() *domain.SurveyResponse {
	return domain.NewSurveyResponse(sr.ID, sr.Survey.ToDomain(), sr.Answer, sr.Rating, sr.CreatedAt, sr.UpdatedAt)
}

type SurveyResponseRepository struct {
	db Database
}

func NewSurveyResponseRepository(dbConn Database) *SurveyResponseRepository {
	return &SurveyResponseRepository{
		db: dbConn,
	}
}

func (repo *SurveyResponseRepository) GetSurveyResponses(
	ctx context.Context, limit int, offset int,
) ([]*domain.SurveyResponse, error) {
	panic("not implemented")
}

func (repo *SurveyResponseRepository) CreateSurveyResponse(
	ctx context.Context, survey *domain.SurveyResponse,
) error {
	panic("not implemented")
}
