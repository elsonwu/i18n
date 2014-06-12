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

        // support language package
        i18n.AddLanguagePackage("zh-cn", i18n.LanguagePackage{
            "this is a test":     "这是一个测试",
            "%s cannot be blank": "%s 不能为空",
        })
        i18n.SetLocal("zh-cn")
        fmt.Println(i18n.T("this is a test"))
        fmt.Println(i18n.T("%s cannot be blank", "email"))
    }