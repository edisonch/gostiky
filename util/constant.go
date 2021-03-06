package util

var (
	MODE_STICK_UNKNOWN = 0
	MODE_STICK_IMMUTABLE = 1
	MODE_UNSTICK_IMMUTABLE = 2
	MODE_STICK_APPEND = 3
	MODE_UNSTICK_APPEND = 4
	MODE_STICK_IMMUTABLE_RECURSIVE = 5
	MODE_UNSTICK_IMMUTABLE_RECURSIVE = 6
	MODE_STICK_APPEND_IMMUTABLE_RECURSIVE = 7
	MODE_UNSTICK_APPEND_IMMUTABLE_RECURSIVE = 8
)

func ConvertCLIModeToInt(input string) int {
	switch {
	case input == "APPEND":
		return MODE_STICK_APPEND
	case input == "UNAPPEND":
		return MODE_UNSTICK_APPEND
	case input == "IMMUTABLE":
		return MODE_STICK_IMMUTABLE
	case input == "UNIMMUTABLE":
		return MODE_UNSTICK_IMMUTABLE
	case input == "IMMUTE_RECURSIVE":
		return MODE_STICK_IMMUTABLE_RECURSIVE
	case input == "UNIMMUTE_RECURSIVE":
		return MODE_UNSTICK_IMMUTABLE_RECURSIVE
	case input == "APPEND_RECURSIVE":
		return MODE_STICK_APPEND_IMMUTABLE_RECURSIVE
	case input == "UNAPPEND_RECURSIVE":
		return MODE_UNSTICK_APPEND_IMMUTABLE_RECURSIVE
	}
	return MODE_STICK_UNKNOWN
}

func ConvertCLIModeToString(input int) string {
	switch {
	case input == MODE_STICK_APPEND:
		return "APPEND"
	case input == MODE_UNSTICK_APPEND:
		return "UNAPPEND"
	case input == MODE_STICK_IMMUTABLE:
		return "IMMUTABLE"
	case input == MODE_UNSTICK_IMMUTABLE:
		return "UNIMMUTABLE"
	case input == MODE_STICK_IMMUTABLE_RECURSIVE:
		return "IMMUTE_RECURSIVE"
	case input == MODE_UNSTICK_IMMUTABLE_RECURSIVE:
		return "UNIMMUTE_RECURSIVE"
	case input == MODE_STICK_APPEND_IMMUTABLE_RECURSIVE:
		return "APPEND_RECURSIVE"
	case input == MODE_UNSTICK_APPEND_IMMUTABLE_RECURSIVE:
		return "UNAPPEND_RECURSIVE"
	}
	return "MODE STICKY UNKNOWN"
}
