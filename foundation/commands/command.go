package commands

// Runner is used to run a command in the Gimli cli tool.
type Runner interface {
	Run() error
}
