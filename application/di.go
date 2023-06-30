package application

import (
	"github.com/jbactad/loop/application/commands"
	"github.com/jbactad/loop/application/ports"
	"github.com/jbactad/loop/application/queries"
	"github.com/samber/do"
)

func ProvideQueryUseCases(injector *do.Injector) {
	do.Provide(injector, func(i *do.Injector) (queries.UseCases, error) {
		sp := do.MustInvoke[ports.SurveyProvider](i)

		return queries.New(sp), nil
	})
}

func ProvideCommandUseCases(injector *do.Injector) {
	do.Provide(injector, func(i *do.Injector) (commands.UseCases, error) {
		scp := do.MustInvoke[ports.SurveyCreatorProvider](i)
		srcp := do.MustInvoke[ports.SurveyResponseCreatorProvider](i)
		ug := do.MustInvoke[ports.UUIDGenerator](i)
		tp := do.MustInvoke[ports.TimeProvider](i)

		return commands.New(scp, srcp, ug, tp), nil
	})
}
