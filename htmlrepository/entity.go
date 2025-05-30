package htmlrepository

import (
	"github.com/suifengpiao14/htmltemplate/htmlcomponent"
	"github.com/suifengpiao14/memorytable"
)

// 这里之所以重新声明一遍，主要是解耦 htmlcomponent包,和htmlrepository包之间的耦合关系，这里定义的gorm 必须和NewXXXField.name保持一致

type Template struct {
	TemplateName string `gorm:"column:templateName" json:"templateName"`
	Template     string `gorm:"column:template" json:"template"`
	DataTpl      string `gorm:"column:dataTpl" json:"dataTpl"`
	DataExample  string `gorm:"column:dataExample" json:"dataExample"` // 示例数据，用于调试
}
type Slot struct {
	TemplateName  string `gorm:"column:templateName" json:"templateName"`
	ComponentName string `gorm:"column:componentName" json:"componentName"`
	SlotName      string `gorm:"column:slotName" json:"slotName"`
	DataTpl       string `gorm:"column:dataTpl" json:"dataTpl"`
	DataExample   string `gorm:"column:dataExample" json:"dataExample"`
}

type Attribute struct {
	TemplateName   string `gorm:"column:templateName" json:"templateName"`
	SlotName       string `gorm:"column:slotName" json:"slotName"`
	TagId          string `gorm:"column:tagId" json:"tagId"`
	AttributeName  string `gorm:"column:attributeName" json:"key"`
	AttributeValue string `gorm:"column:attributeValue" json:"value"`
}

func ToHtmlSlot(slotName Slot) htmlcomponent.Slot {
	return htmlcomponent.Slot{
		ComponentName: slotName.ComponentName,
		TemplateName:  slotName.TemplateName,
		SlotName:      slotName.SlotName,
		DataTpl:       slotName.DataTpl,
		DataExample:   slotName.DataExample,
	}
}

func ToHtmlSlots(slotNames ...Slot) htmlcomponent.Slots {
	return memorytable.Map(slotNames, func(item Slot) htmlcomponent.Slot {
		return ToHtmlSlot(item)
	})
}

func ToHtmlAttribute(attribute Attribute) htmlcomponent.Attribute {
	return htmlcomponent.Attribute{
		TemplateName:   attribute.TemplateName,
		SlotName:       attribute.SlotName,
		TagId:          attribute.TagId,
		AttributeName:  attribute.AttributeName,
		AttributeValue: attribute.AttributeValue,
	}
}

func ToHtmlAttributes(attributes ...Attribute) htmlcomponent.Attributes {
	return memorytable.Map(attributes, func(item Attribute) htmlcomponent.Attribute {
		return ToHtmlAttribute(item)
	})
}

func ToHtmlComponent(component Template) htmlcomponent.Template {
	return htmlcomponent.Template{
		TemplateName: component.TemplateName,
		Template:     component.Template,
		DataTpl:      component.DataTpl,
		DataExample:  component.DataExample,
	}
}

func ToHtmlComponents(components ...Template) htmlcomponent.Templates {
	return memorytable.Map(components, func(item Template) htmlcomponent.Template {
		return ToHtmlComponent(item)
	})
}
