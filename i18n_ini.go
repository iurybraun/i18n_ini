/*
 * Copyright © 2016 Iury Braun
 * Copyright © 2017 Weyboo
 */

package i18n_ini

import (
    "github.com/beego/i18n"
    "io/ioutil"
    "strings"
)

type LangString string

func Load_language(lang string) {
    
    b, err := ioutil.ReadFile("lang/" + lang + ".ini")
    if err != nil {
        panic(err.Error())
    }
    
    var set_error error
    langData := []byte(b)
    if set_error = i18n.SetMessageData(lang, langData); set_error != nil {
        panic(set_error.Error())
    }
    
}

func Load_languages(languages_list string) {
    
    var py LangString
	py = LangString(languages_list)
	py.SplitAddLanguages(",")
    
}

func (py LangString) SplitAddLanguages(str string) {
    
	s := strings.Split(string(py), str)

	for i := 0; i < len(s); i += 1 {
		Load_language(s[i])
	}
    
}

func LoadTr(lang, format string, args ...interface{}) string {
    
    return i18n.Tr(lang, format, args)
    
}
