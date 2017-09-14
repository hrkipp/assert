package assert

const (
	NotEnoughArgs    = "not enough arguments provided"
	InvalidArguments = "invalid arguments"
	FalseAssertion   = "assertion was false"
	NoChecker        = "no checker found in the second argument"
)

func Check(args ...interface{}) (bool, string) {
	switch len(args) {
	case 0:
		return false, NotEnoughArgs
	case 1:
		switch arg := args[0].(type) {
		case bool:
			if arg {
				return true, ""
			} else {
				return false, FalseAssertion
			}
		}
	default:
		checker, isChecker := args[1].(Checker)
		if !isChecker {
			return false, NoChecker
		}
		checkerArgs := append(args[:1], args[2:]...)
		if checker.Check(checkerArgs...) {
			return true, ""
		}
		return false, checker.Message(checkerArgs...)

	}
	return false, NotEnoughArgs
}
