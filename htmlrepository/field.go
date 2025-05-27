package htmlrepository

import "github.com/suifengpiao14/sqlbuilder"

func NewTemplateNameField(componentName string) *sqlbuilder.Field {
	return sqlbuilder.NewStringField(componentName, "templateName", "模版名称", 0)
}
func NewComponentNameField(componentName string) *sqlbuilder.Field {
	return sqlbuilder.NewStringField(componentName, "componentName", "组件名称", 0)
}

func NewTemplateNamesField(componentNames []string) *sqlbuilder.Field {
	componentNameField := NewTemplateNameField("")
	return sqlbuilder.NewStringField(componentNames, componentNameField.Name, componentNameField.Schema.Title, 0)
}

func NewSlotNameField(slotName string) *sqlbuilder.Field {
	return sqlbuilder.NewStringField(slotName, "slotName", "槽名称", 0)
}

func NewTemplateField(template string) *sqlbuilder.Field {
	return sqlbuilder.NewStringField(template, "template", "模板", 0)
}

func NewDataTplField(dataTpl string) *sqlbuilder.Field {
	return sqlbuilder.NewStringField(dataTpl, "dataTpl", "数据模板", 0)
}

func NewNodeIdField(nodeId string) *sqlbuilder.Field {
	return sqlbuilder.NewStringField(nodeId, "nodeId", "节点ID", 0)
}

func NewAttributeNameField(attributeKey string) *sqlbuilder.Field {
	return sqlbuilder.NewStringField(attributeKey, "attributeName", "键名", 0)
}

func NewAttributeValueField(attributeValue string) *sqlbuilder.Field {
	return sqlbuilder.NewStringField(attributeValue, "attributeValue", "键值", 0)
}
