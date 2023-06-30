package infrastructure

import (
	"github.com/jbactad/loop/application/ports"
	"github.com/jbactad/loop/infrastructure/repositories"
	"github.com/samber/do"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ProvideServices(injector *do.Injector) error {
	do.Provide(injector, func(i *do.Injector) (ports.UUIDGenerator, error) {
		return NewGoogleUUIDGenerator(), nil
	})

	do.Provide(injector, func(i *do.Injector) (ports.TimeProvider, error) {
		return NewTimeProvider()
	})

	return nil
}

func ProvideRepositories(injector *do.Injector) {
	do.Provide(injector, func(i *do.Injector) (ports.SurveyResponseCreatorProvider, error) {
		return repositories.NewSurveyResponseRepository(do.MustInvoke[repositories.Database](i)), nil
	})
	do.Provide(injector, func(i *do.Injector) (ports.SurveyResponseProvider, error) {
		return repositories.NewSurveyResponseRepository(do.MustInvoke[repositories.Database](i)), nil
	})
	do.Provide(injector, func(i *do.Injector) (ports.SurveyCreatorProvider, error) {
		return repositories.NewSurveyRepository(do.MustInvoke[repositories.Database](i)), nil
	})
	do.Provide(injector, func(i *do.Injector) (ports.SurveyProvider, error) {
		return repositories.NewSurveyRepository(do.MustInvoke[repositories.Database](i)), nil
	})
}

func ProvideDatabaseConnection(injector *do.Injector) error {
	// TODO: Move this to configuration
	dsn := "host=localhost user=postgres password=postgres dbname=loop port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	do.Provide(injector, func(i *do.Injector) (repositories.Database, error) {
		return NewGormDB(db), nil
	})

	return nil
}
