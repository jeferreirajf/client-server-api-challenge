package usecase

type Input struct{}

type Output struct{}

type Usecase interface {
	Execute(input Input) (Output, error)
}
