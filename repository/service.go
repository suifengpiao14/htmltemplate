package repository

import (
	"github.com/suifengpiao14/htmltemplate/htmlcomponent"
	"github.com/suifengpiao14/memorytable"
	"github.com/suifengpiao14/sqlbuilder"
)

type _ComponentSerivce[C Component] struct {
	repo        RepoComponent[C] // 依赖数据接口
	repoCommand sqlbuilder.RepositoryCommand
	repoQuery   sqlbuilder.RepositoryQuery[C]
}

func newComponentService[C Component](repo RepoComponent[C], tableConfig sqlbuilder.TableConfig) _ComponentSerivce[C] {
	repoCommand := sqlbuilder.NewRepositoryCommand(tableConfig)
	return _ComponentSerivce[C]{
		repo:        repo,
		repoQuery:   sqlbuilder.NewRepositoryQuery[C](tableConfig),
		repoCommand: repoCommand,
	}
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

func (s HtmlTemplateService[C, A, R]) GetComponent(componentName string) (rootComponent htmlcomponent.RootComponent, err error) {
	assembles, err := s._GetAssemblesByRootComponentName(componentName)
	if err != nil {
		return rootComponent, err
	}
	rootComponent.Assembles = assembles
	componentNames := assembles.ComponentNames()
	componentNames = append(componentNames, componentName)
	componentNames = memorytable.NewTable(componentNames...).Uniqueue(func(row string) (key string) { return key }).ToSlice()
	components, err := s._GetComponentByComponentNames(componentNames)
	if err != nil {
		return rootComponent, err
	}
	rootComponent.Components = components
	attrs, err := s._GetAttributesByRootComponentName(componentName)
	if err != nil {
		return rootComponent, err
	}
	rootComponent.Attributes = attrs
	return rootComponent, nil
}

func (s HtmlTemplateService[C, A, R]) AddComponent(component Component) error {
	return s.componentService._AddComponent(component)
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

func (s _ComponentSerivce[C]) _AddComponent(component Component) error {
	return s.repoCommand.Insert(component)
}
