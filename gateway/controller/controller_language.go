package controller

import (
	"fmt"

	mystore "types/mariadb"

	"gateway/persist"
)

// GetLangContent GetLangContent
func GetLangContent(code, lang, _default string) string {
	if len(code) == 0 {
		return _default
	}

	if len(lang) == 0 {
		lang = "cn"
	}

	_key := LrucacheKeyMsgTempl + code + "." + lang
	_value, _ok := lruCache.Get(_key)
	if _ok {
		fmt.Println("get MessageTemplate from lru:", _key)
		return _value.(mystore.GudpMessageTemplateBase).Detail
	}

	_templ, err := persist.GMariadb.GetGudpSystemTemplateInfo(code, lang)
	if err != nil {
		return _default
	}

	lruCache.Add(_key, _templ)

	return _templ.Detail
}
