package e

import "fmt"

type Map map[string]interface{}

type Error struct {
	Info  string
	Cause error
	Args  Map
}



func (e Error) Error() string {
	//format := "WRONG ERROR FORMAT"
	//haveCause := e.cause != nil
	//haveInfo := len(e.info) > 0
	//haveArgs := len(e.args) > 0

	return fmt.Sprint(e)
}

func WrapError(info string, cause error, args ...Map) error {
	return wrap(info, cause, args...)
}

func wrap(info string, cause error, args ...Map) Error {
	var argsMap Map
	switch len(args) {
	case 0:
		// Leave as nil
	case 1:
		argsMap = args[0]
	default:
		argsMap = make(Map)
		for _, m := range args {
			for k, v := range m {
				argsMap[k] = v
			}
		}
	}
	return Error{info, cause, argsMap}
}

func Is(err1, err2 error) bool {
	if err1 == err2 {
		return true
	}

	e, ok := err1.(Error)
	if ok {
		return Is(e.Cause, err2)
	}

	return false
}
