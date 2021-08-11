package main

import (
	"errors"
	"os"
)

type CommandStrategy interface {
	Execute() (string, error)
}

const (
	INVALID_ARGUMENTS = "Invalid number of arguments"
)

type CreateCommand struct {
	args []string
}

func NewCreateCommand(args []string) *CreateCommand {
	return &CreateCommand{
		args: args,
	}
}

func (cc *CreateCommand) Execute() (string, error) {
	if len(cc.args) < 2 {
		return "ERR\n", errors.New(INVALID_ARGUMENTS)
	}

	dirName := cc.args[1]
	err := os.Mkdir(dirName, os.ModePerm)
	if err != nil {
		return "ERR\n", err
	}

	return "OK\n", nil
}

type RemoveCommand struct {
	args []string
}

func NewRemoveCommand(args []string) *RemoveCommand {
	return &RemoveCommand{
		args: args,
	}
}

func (rc *RemoveCommand) Execute() (string, error) {
	if len(rc.args) < 2 {
		return "ERR\n", errors.New(INVALID_ARGUMENTS)
	}

	dirName := rc.args[1]
	err := os.Remove(dirName)
	if err != nil {
		return "ERR\n", err
	}

	return "OK\n", nil
}

type RenameCommand struct {
	args []string
}

func NewRenameCommand(args []string) *RenameCommand {
	return &RenameCommand{
		args: args,
	}
}

func (rc *RenameCommand) Execute() (string, error) {
	if len(rc.args) < 3 {
		return "ERR\n", errors.New(INVALID_ARGUMENTS)
	}

	oldPath := rc.args[1]
	newPath := rc.args[2]
	err := os.Rename(oldPath, newPath)
	if err != nil {
		return "ERR\n", err
	}

	return "OK\n", nil
}

type ReadDirCommand struct {
	args []string
}

func NewReadDirCommand(args []string) *ReadDirCommand {
	return &ReadDirCommand{
		args: args,
	}
}

func (rrc *ReadDirCommand) Execute() (string, error) {
	dir := "."
	if len(rrc.args) >= 2 {
		dir = rrc.args[1]
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "ERR\n", err
	}
	var response string
	for _, entry := range entries {
		response += entry.Name() + " "
	}

	return response + "\n", nil
}
