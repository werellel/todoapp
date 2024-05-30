package main

type ServiceImpl struct{}

func (s *ServiceImpl) Apply(id int) error { return nil }

type OderService interface {
	Apply(int) error
}

type App struct {
	os OderService
}

func NewApp(os OderService) *App {
	return &App{os: os}
}

func (app *App) Apply(id int) error {
	return app.os.Apply(id)
}

func main() {
	app := NewApp(&ServiceImpl{})
	app.Apply(19)
}
