package i18n

import "fmt"

var local string

type LanguagePackage map[string]string

// package example:
// {
//     "en": {
//         {
//             "this is a test": "this is a test",
//             "%s cannot be blank": "% cannot be blank"
//         }
//     },
//     "zh-cn": {
//         {
//             "this is a test": "这是一个测试",
//             "%s cannot be blank": "% 不能为空"
//         }
//     },
//     ...
// }
var languagePackages map[string]LanguagePackage

func T(err string, vars ...interface{}) string {
	//replace the translate language
	if local != "" {
		if lp := pack(); lp != nil {
			if e, ok := (*lp)[err]; ok {
				err = e
			}
		}
	}

	return fmt.Sprintf(err, vars...)
}

func SetLocal(l string) {
	local = l
}

func AddLanguagePackage(localName string, lp LanguagePackage) {
	if languagePackages == nil {
		languagePackages = map[string]LanguagePackage{}
	}

	languagePackages[localName] = lp
}

func pack() *LanguagePackage {
	if languagePackages == nil {
		return nil
	}

	lp, ok := languagePackages[local]
	if ok {
		return &lp
	}

	return nil
}
