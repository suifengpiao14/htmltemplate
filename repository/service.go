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

func (s HtmlTemplateService[C, A, R]) GetHtmlPage(rootComponentName string) (htmlPage htmlcomponent.HtmlPage, err error) {
	assembles, err := s._GetAssemblesByRootComponentName(rootComponentName)
	if err != nil {
		return htmlPage, err
	}
	htmlPage.Assembles = assembles
	componentNames := assembles.ComponentNames()
	componentNames = append(componentNames, rootComponentName)
	componentNames = memorytable.NewTable(componentNames...).Uniqueue(func(row string) (key string) { return key }).ToSlice()
	components, err := s._GetComponentByComponentNames(componentNames)
	if err != nil {
		return htmlPage, err
	}
	htmlPage.Components = components
	attrs, err := s._GetAttributesByRootComponentName(rootComponentName)
	if err != nil {
		return htmlPage, err
	}
	htmlPage.Attributes = attrs
	return htmlPage, nil
}

func (s HtmlTemplateService[C, A, R]) _GetAssemblesByRootComponentName(rootComponentName string) (assembles htmlcomponent.Assembles, err error) {
	assembles, err = s.assembleService._GetByRootComponentName(rootComponentName)
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

func (s HtmlTemplateService[C, A, R]) _GetAttributesByRootComponentName(rootComponentName string) (assembles htmlcomponent.Attributes, err error) {
	assembles, err = s.attributeService._GetByRootComponentName(rootComponentName)
	if err != nil {
		return nil, err
	}
	return assembles, nil
}

func (s _AssembleService[A]) _GetByRootComponentName(rootComponentName string) (assembles htmlcomponent.Assembles, err error) {
	if rootComponentName == "" {
		return assembles, nil
	}
	assembleModels, err := s.repo.GetByRootComponentName(rootComponentName)
	if err != nil {
		return nil, err
	}
	for _, assembleModel := range assembleModels {
		assemble := htmlcomponent.Assemble{
			RootComponentName: assembleModel.GetRootComponentName(),
			ComponentName:     assembleModel.GetComponentName(),
			AssembleName:      assembleModel.GetAssembleName(),
			DataTpl:           assembleModel.GetDataTpl(),
		}
		assembles = append(assembles, assemble)
	}
	return assembles, nil
}

func (s _AttributeService[R]) _GetByRootComponentName(rootComponentName string) (attributes htmlcomponent.Attributes, err error) {
	if rootComponentName == "" {
		return attributes, nil
	}
	attrModels, err := s.repo.GetByRootComponentName(rootComponentName)
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
			Name:     componentModel.GetName(),
			Template: componentModel.GetTemplate(),
			DataTpl:  componentModel.GetDataTpl(),
		}
		components = append(components, component)
	}
	return components, nil
}
