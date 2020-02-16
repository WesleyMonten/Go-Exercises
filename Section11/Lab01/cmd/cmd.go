package cmd

type (
	Command struct {
		Name string
		Args []string
	}
)

func New(name string, args []string) *Command {
	return &Command{Name: name, Args: args}
}
