package ui

type IUserInterface interface {
	Serve() error
	Close() error
}
