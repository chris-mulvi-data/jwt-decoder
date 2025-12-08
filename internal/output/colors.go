package output

type Color string

const (
	Default Color = "\033[0m"
	Black   Color = "\033[30m"
	Gray    Color = "\033[90m"
	Red     Color = "\033[91m"
	Green   Color = "\033[92m"
	Yellow  Color = "\033[93m"
	Blue    Color = "\033[94m"
	Magenta Color = "\033[95m"
	Cyan    Color = "\033[96m"
	White   Color = "\033[97m"
)
