package persist

import (
	mystore "types/mariadb"
)

// CreateGudpSystemTemplate CreateGudpSystemTemplate Persist
func (maria *Mariadb) CreateGudpSystemTemplate(param mystore.GudpMessageTemplate) error {
	return maria.db.Create(&param).Error
}

// GetGudpSystemTemplateInfo GetGudpSystemTemplateInfo Persist
func (maria *Mariadb) GetGudpSystemTemplateInfo(code, lang string) (mystore.GudpMessageTemplateBase, error) {
	var obj mystore.GudpMessageTemplateBase
	err := maria.db.Table("gudp_message_templates").Where("code = ? and lang = ?", code, lang).First(&obj).Error
	return obj, err
}
