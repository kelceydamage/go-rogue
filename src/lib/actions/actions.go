package actions

type Action interface {
	Execute()
	GetText() string
	GetName() string
}
