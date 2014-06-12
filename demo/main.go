package main

import (
	"fmt"

	"github.com/elsonwu/i18n"
)

func main() {
	fmt.Println(i18n.T("this is a test"))
	fmt.Println(i18n.T("%s cannot be blank", "email"))
}
