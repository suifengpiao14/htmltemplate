package htmltemplate

import (
	"github.com/pkg/errors"
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
var HtmlTemplateService *repository.HtmlTemplateService[repository.Component, repository.Assemble, repository.Attribute]

func PageHtml(rootComponentName string, data map[string]any) (rootComponentHtml string, err error) {
	if HtmlTemplateService == nil {
		err = errors.Errorf("HtmlTemplateService uninitialized")
		return "", err
	}
	rootComponent, err := HtmlTemplateService.GetRootComponent(rootComponentName)
	if err != nil {
		return "", err
	}
	rootComponentHtml, err = rootComponent.ToHtml(data)
	if err != nil {
		return "", err
	}
	return rootComponentHtml, nil
}
