package htmltemplate

import (
	"github.com/suifengpiao14/htmltemplate/htmlcomponent"
	"github.com/suifengpiao14/htmltemplate/htmlenhance"
	"github.com/suifengpiao14/htmltemplate/repository"
)

var SetNodeIdAndAttrHolder = htmlenhance.SetNodeIdAndAttrHolder
var MergeClassAttrs = htmlenhance.MergeClassAttrs

type Component = htmlcomponent.Component
type Assemble = htmlcomponent.Assemble
type Assembles = htmlcomponent.Assembles
type Attribute = htmlcomponent.Attribute
type Attributes = htmlcomponent.Attributes

// HtmlTemplateService 外部调用,必须在初始化时赋值
var HtmlTemplateService repository.HtmlTemplateService[repository.Component, repository.Assemble, repository.Attribute]

func PageHtml(pageName string, data map[string]any) (pageHtml string, err error) {
	htmlPage, err := HtmlTemplateService.GetHtmlPage(pageName)
	if err != nil {
		return "", err
	}
	pageHtml, err = htmlPage.ToHtml(data)
	if err != nil {
		return "", err
	}
	return pageHtml, nil
}
