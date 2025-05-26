package htmlrepository

import (
	"github.com/suifengpiao14/htmltemplate/htmlcomponent"
	"github.com/suifengpiao14/memorytable"
)

// 这里之所以重新声明一遍，主要是解耦 htmlcomponent包,和htmlrepository包之间的耦合关系，这里定义的gorm 必须和NewXXXField.name保持一致

type Component struct {
	ComponentName string `gorm:"column:componentName" json:"componentName"`
	Template      string `gorm:"column:template" json:"template"`
	DataTpl       string `gorm:"column:dataTpl" json:"dataTpl"`
	DataExample   string `gorm:"column:dataExample" json:"dataExample"` // 示例数据，用于调试
}
type Assemble struct {
	RootComponentName string `gorm:"column:rootComponentName" json:"rootComponentName"`
	ComponentName     string `gorm:"column:componentName" json:"componentName"`
	AssembleName      string `gorm:"column:assembleName" json:"assembleName"`
	DataTpl           string `gorm:"column:dataTpl" json:"dataTpl"`
	DataExample       string `gorm:"column:dataExample" json:"dataExample"`
}

type Attribute struct {
	TagId          string `gorm:"column:tagId" json:"tagId"`
	NodeId         string `gorm:"column:nodeId" json:"nodeId"`
	AttributeName  string `gorm:"column:attributeName" json:"key"`
	AttributeValue string `gorm:"column:attributeValue" json:"value"`
}

func ToHtmlAssemble(assemble Assemble) htmlcomponent.ComponentNode {
	return htmlcomponent.ComponentNode{
		ComponentName: assemble.RootComponentName,
		TemplateName:  assemble.ComponentName,
		NodeID:        assemble.AssembleName,
		DataTpl:       assemble.DataTpl,
		DataExample:   assemble.DataExample,
	}
}

func ToHtmlAssembles(assembles ...Assemble) htmlcomponent.ComponentNodes {
	return memorytable.Map(assembles, func(item Assemble) htmlcomponent.ComponentNode {
		return ToHtmlAssemble(item)
	})
}

func ToHtmlAttribute(attribute Attribute) htmlcomponent.Attribute {
	return htmlcomponent.Attribute{
		TagId:          attribute.TagId,
		NodeId:         attribute.NodeId,
		AttributeName:  attribute.AttributeName,
		AttributeValue: attribute.AttributeValue,
	}
}

func ToHtmlAttributes(attributes ...Attribute) htmlcomponent.Attributes {
	return memorytable.Map(attributes, func(item Attribute) htmlcomponent.Attribute {
		return ToHtmlAttribute(item)
	})
}

func ToHtmlComponent(component Component) htmlcomponent.ComponentTemplate {
	return htmlcomponent.ComponentTemplate{
		Name:        component.ComponentName,
		Template:    component.Template,
		DataTpl:     component.DataTpl,
		DataExample: component.DataExample,
	}
}

func ToHtmlComponents(components ...Component) htmlcomponent.ComponentTemplates {
	return memorytable.Map(components, func(item Component) htmlcomponent.ComponentTemplate {
		return ToHtmlComponent(item)
	})
}
