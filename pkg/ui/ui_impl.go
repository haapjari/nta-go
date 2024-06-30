package ui

import (
	"log/slog"
)

type CommandLineInterface struct{}

// TODO
func New() (IUserInterface, error) {
	slog.Info("Creating CommandLineInterface")

	c := &CommandLineInterface{}

	return c, nil
}

// TODO
func (c *CommandLineInterface) Serve() error {
	slog.Info("Serving CommandLineInterface")
	return nil
}

// TODO
func (c *CommandLineInterface) Close() error {
	slog.Info("Closing CommandLineInterface")
	return nil
}
