package controller

import (
	"fmt"

	mystore "github.com/gopperin/sme-mini/types/mariadb"

	"github.com/gopperin/sme-mini/gateway/persist"
)

// GetLangContent GetLangContent
func GetLangContent(code, lang, defaultStr string) string {
	if len(code) == 0 {
		return defaultStr
	}

	if len(lang) == 0 {
		lang = "cn"
	}

	key := lrucacheKeyMsgTempl + code + "." + lang
	value, ok := lruCache.Get(key)
	if ok {
		fmt.Println("get MessageTemplate from lru:", key)
		return value.(mystore.GudpMessageTemplateBase).Detail
	}

	templ, err := persist.GMariadb.GetGudpSystemTemplateInfo(code, lang)
	if err != nil {
		return defaultStr
	}

	lruCache.Add(key, templ)

	return templ.Detail
}
