package repository

import (
	"github.com/suifengpiao14/htmltemplate/htmlcomponent"
	"github.com/suifengpiao14/memorytable"
)

type _ComponentSerivce[C Component] struct {
	repo RepoComponent[C] // 依赖数据接口
}

type _AssembleService[A Assemble] struct {
	repo RepoAssemble[A] // 依赖数据接口
}

type _AttributeService[R Attribute] struct {
	repo RepoAttribute[R] // 依赖数据接口
}

type HtmlTemplateService[C Component, A Assemble, R Attribute] struct {
	componentService _ComponentSerivce[C]
	assembleService  _AssembleService[A]
	attributeService _AttributeService[R]
}

func NewHtmlTemplateService[C Component, A Assemble, R Attribute](repoComponent RepoComponent[C], repoAssemble RepoAssemble[A], repoAttribute RepoAttribute[R]) *HtmlTemplateService[C, A, R] {
	componentService := _ComponentSerivce[C]{
		repo: repoComponent,
	}
	assembleService := _AssembleService[A]{
		repo: repoAssemble,
	}
	attributeService := _AttributeService[R]{
		repo: repoAttribute,
	}
	return &HtmlTemplateService[C, A, R]{
		componentService: componentService,
		assembleService:  assembleService,
		attributeService: attributeService,
	}
}

func (s HtmlTemplateService[C, A, R]) GetHtmlPage(pageName string) (htmlPage htmlcomponent.HtmlPage, err error) {
	assembles, err := s._GetAssemblesByPageName(pageName)
	if err != nil {
		return htmlPage, err
	}
	htmlPage.Assembles = assembles
	componentNames := assembles.ComponentNames()
	componentNames = append(componentNames, pageName)
	componentNames = memorytable.NewTable(componentNames...).Uniqueue(func(row string) (key string) { return key }).ToSlice()
	components, err := s._GetComponentByComponentNames(componentNames)
	if err != nil {
		return htmlPage, err
	}
	htmlPage.Components = components
	attrs, err := s._GetAttributesByPageName(pageName)
	if err != nil {
		return htmlPage, err
	}
	htmlPage.Attributes = attrs
	return htmlPage, nil
}

func (s HtmlTemplateService[C, A, R]) _GetAssemblesByPageName(pageName string) (assembles htmlcomponent.Assembles, err error) {
	assembles, err = s.assembleService._GetByPageName(pageName)
	if err != nil {
		return nil, err
	}
	return assembles, nil
}
func (s HtmlTemplateService[C, A, R]) _GetComponentByComponentNames(componentNames []string) (components htmlcomponent.Components, err error) {
	components, err = s.componentService._GetByComponentNames(componentNames)
	if err != nil {
		return nil, err
	}
	return components, nil
}

func (s HtmlTemplateService[C, A, R]) _GetAttributesByPageName(pageName string) (assembles htmlcomponent.Attributes, err error) {
	assembles, err = s.attributeService._GetByPageName(pageName)
	if err != nil {
		return nil, err
	}
	return assembles, nil
}

func (s _AssembleService[A]) _GetByPageName(pageName string) (assembles htmlcomponent.Assembles, err error) {
	if pageName == "" {
		return assembles, nil
	}
	assembleModels, err := s.repo.GetByPageName(pageName)
	if err != nil {
		return nil, err
	}
	for _, assembleModel := range assembleModels {
		assemble := htmlcomponent.Assemble{
			PageName:      assembleModel.PageName(),
			ComponentName: assembleModel.ComponentName(),
			AssembleName:  assembleModel.AssembleName(),
			DataTpl:       assembleModel.DataTpl(),
		}
		assembles = append(assembles, assemble)
	}
	return assembles, nil
}

func (s _AttributeService[R]) _GetByPageName(pageName string) (attributes htmlcomponent.Attributes, err error) {
	if pageName == "" {
		return attributes, nil
	}
	attrModels, err := s.repo.GetByPageName(pageName)
	if err != nil {
		return nil, err
	}
	for _, attrModel := range attrModels {
		attribute := htmlcomponent.Attribute{
			NodeId: attrModel.GetNodeId(),
			Key:    attrModel.GetKey(),
			Value:  attrModel.GetValue(),
		}
		attributes = append(attributes, attribute)
	}
	return attributes, nil
}

func (s _ComponentSerivce[C]) _GetByComponentNames(componentNames []string) (components htmlcomponent.Components, err error) {
	if len(componentNames) == 0 {
		return components, nil
	}
	componentModels, err := s.repo.GetByComponentNames(componentNames)
	if err != nil {
		return nil, err
	}

	for _, componentModel := range componentModels {
		component := htmlcomponent.Component{
			Name:     componentModel.Name(),
			Template: componentModel.Template(),
			DataTpl:  componentModel.DataTpl(),
		}
		components = append(components, component)
	}
	return components, nil
}
