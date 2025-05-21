package htmlrepository

import (
	"github.com/suifengpiao14/htmltemplate/htmlcomponent"
	"github.com/suifengpiao14/memorytable"
	"github.com/suifengpiao14/sqlbuilder"
)

type HtmlTemplateService[C Component, A Assemble, R Attribute] struct {
	componentService ComponentSerivce[C]
	assembleService  AssembleService[A]
	attributeService AttributeService[R]
}

func NewHtmlTemplateService[C Component, A Assemble, R Attribute](tableConfig TableConfig) *HtmlTemplateService[C, A, R] {
	tableConfig = tableConfig.AddIndex()
	componentService := ComponentSerivce[C]{
		repositoryQuery:   sqlbuilder.NewRepositoryQuery[C](tableConfig.Component),
		repositoryCommand: sqlbuilder.NewRepositoryCommand(tableConfig.Component),
	}

	assembleService := AssembleService[A]{
		repositoryQuery: sqlbuilder.NewRepositoryQuery[A](tableConfig.Assemble),

		repositoryCommand: sqlbuilder.NewRepositoryCommand(tableConfig.Assemble),
	}
	attributeService := AttributeService[R]{
		repositoryQuery:   sqlbuilder.NewRepositoryQuery[R](tableConfig.Attribute),
		repositoryCommand: sqlbuilder.NewRepositoryCommand(tableConfig.Attribute),
	}
	return &HtmlTemplateService[C, A, R]{
		componentService: componentService,
		assembleService:  assembleService,
		attributeService: attributeService,
	}
}

func (s HtmlTemplateService[C, A, R]) Render(componentRootName string, data map[string]any) (rootComponentHtml string, err error) {
	rootComponent, err := s.GetComponent(componentRootName)
	if err != nil {
		return "", err
	}
	rootComponentHtml, err = rootComponent.ToHtml(data)
	if err != nil {
		return "", err
	}
	return rootComponentHtml, nil
}

func (s HtmlTemplateService[C, A, R]) GetComponent(componentRootName string) (rootComponent htmlcomponent.RootComponent, err error) {
	assembles, err := s._GetAssemblesByRootComponentName(componentRootName)
	if err != nil {
		return rootComponent, err
	}
	rootComponent.Assembles = assembles
	componentNames := assembles.ComponentNames()
	componentNames = append(componentNames, componentRootName)
	componentNames = memorytable.NewTable(componentNames...).Uniqueue(func(row string) (key string) { return key }).ToSlice()
	components, err := s._GetComponentByComponentNames(componentNames)
	if err != nil {
		return rootComponent, err
	}
	rootComponent.Components = components
	attrs, err := s._GetAttributesByRootComponentName(componentRootName)
	if err != nil {
		return rootComponent, err
	}
	rootComponent.Attributes = attrs
	return rootComponent, nil
}

func (s HtmlTemplateService[C, A, R]) _GetAssemblesByRootComponentName(rootComponentName string) (assembles htmlcomponent.Assembles, err error) {
	models, err := s.assembleService.ListByRootComponentName(rootComponentName, nil)
	if err != nil {
		return nil, err
	}
	for _, model := range models {
		assemble := htmlcomponent.Assemble{
			RootComponentName: model.GetRootComponentName(),
			ComponentName:     model.GetComponentName(),
			AssembleName:      model.GetAssembleName(),
			DataTpl:           model.GetDataTpl(),
		}
		assembles = append(assembles, assemble)
	}
	return assembles, nil
}
func (s HtmlTemplateService[C, A, R]) _GetComponentByComponentNames(componentNames []string) (components htmlcomponent.Components, err error) {
	models, err := s.componentService.ListByComponentNames(componentNames)
	if err != nil {
		return nil, err
	}
	for _, model := range models {
		component := htmlcomponent.Component{
			Name:     model.GetName(),
			Template: model.GetTemplate(),
			DataTpl:  model.GetDataTpl(),
		}
		components = append(components, component)
	}
	return components, nil
}

func (s HtmlTemplateService[C, A, R]) _GetAttributesByRootComponentName(rootComponentName string) (attributes htmlcomponent.Attributes, err error) {
	models, err := s.attributeService.ListByRootComponentName(rootComponentName, nil)
	if err != nil {
		return nil, err
	}
	for _, model := range models {
		attribute := htmlcomponent.Attribute{
			NodeId: model.GetNodeId(),
			Key:    model.GetAttributeName(),
			Value:  model.GetAttributeValue(),
		}
		attributes = append(attributes, attribute)
	}
	return attributes, nil
}
