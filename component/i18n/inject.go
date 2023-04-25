package i18n

import "go.uber.org/dig"

func Inject(container *dig.Container) error {
	return container.Provide(newI18n)
}
