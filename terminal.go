package c80

var (
	Terminal Terminator
)

func init() {
	Terminal = machine.Terminal()
}
