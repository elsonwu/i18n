package i18n

import "fmt"

func T(err string, vars ...interface{}) string {
	return fmt.Sprintf(err, vars...)
}
