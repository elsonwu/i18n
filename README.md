i18n
===

#install

    go get github.com/elsonwu/i18n

#usage

    import (
        "fmt"

        "github.com/elsonwu/i18n"
    )

    func main() {
        fmt.Println(i18n.T("This is a test"))
        fmt.Println(i18n.T("%s cannot be blank", "email"))
    }


#todo

    add SetLocal method
    support Language package