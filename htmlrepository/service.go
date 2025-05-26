package htmlrepository

import (
	"github.com/suifengpiao14/htmltemplate/htmlcomponent"
	"github.com/suifengpiao14/memorytable"
	"github.com/suifengpiao14/sqlbuilder"
)

type HtmlTemplateApiService struct {
	componentService ComponentSerivce[Component]
	assembleService  AssembleService[Assemble]
	attributeService AttributeService[Attribute]
}

func NewHtmlTemplateApiService(dbHander sqlbuilder.Handler, customTableFn func(tableConfig TableConfig) (customedTableConfig TableConfig)) *HtmlTemplateApiService {
	tableConfig := CustomTableConfig(dbHander, customTableFn)
	componentService := newComponentSerivce[Component](tableConfig.Component)
	assembleService := newAssembleService[Assemble](tableConfig.Assemble)
	attributeService := newAttributeService[Attribute](tableConfig.Attribute)
	return &HtmlTemplateApiService{
		componentService: componentService,
		assembleService:  assembleService,
		attributeService: attributeService,
	}
}

func (s HtmlTemplateApiService) Render(componentRootName string, data map[string]any) (rootComponentHtml string, err error) {
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

func (s HtmlTemplateApiService) GetComponent(componentRootName string) (rootComponentHtml htmlcomponent.Component, err error) {
	assembles, err := s.assembleService.ListByRootComponentName(componentRootName, func(listParam *sqlbuilder.ListParam) {
		listParam.WithCustomFieldsFn(func(fs sqlbuilder.Fields) (customedFs sqlbuilder.Fields) {
			fs.FirstMust().Apply(func(f *sqlbuilder.Field, fs ...*sqlbuilder.Field) {
				f.SetDelayApply(func(f *sqlbuilder.Field, fs ...*sqlbuilder.Field) {
					setSelectColumns := f.GetTable().Columns.FilterByFieldName(
						sqlbuilder.GetFieldName(NewRootComponentNameField),
						sqlbuilder.GetFieldName(NewComponentNameField),
						sqlbuilder.GetFieldName(NewAssembleNameField),
						sqlbuilder.GetFieldName(NewDataTplField),
					).DbNameWithAlias()
					f.SetSelectColumns(setSelectColumns.AsAny()...)
				})
			})
			return fs
		})
	})
	if err != nil {
		return rootComponentHtml, err
	}
	rootComponentHtml.Nodes = ToHtmlAssembles(assembles...)
	componentNames := rootComponentHtml.Nodes.ComponentNames()
	componentNames = append(componentNames, componentRootName)
	componentNames = memorytable.NewTable(componentNames...).Uniqueue(func(row string) (key string) { return key }).ToSlice()
	components, err := s.componentService.ListByComponentNames(componentNames, func(listParam *sqlbuilder.ListParam) {
		listParam.WithCustomFieldsFn(func(fs sqlbuilder.Fields) (customedFs sqlbuilder.Fields) {
			fs.FirstMust().Apply(func(f *sqlbuilder.Field, fs ...*sqlbuilder.Field) {
				f.SetDelayApply(func(f *sqlbuilder.Field, fs ...*sqlbuilder.Field) {
					setSelectColumns := f.GetTable().Columns.FilterByFieldName(
						sqlbuilder.GetFieldName(NewComponentNameField),
						sqlbuilder.GetFieldName(NewTemplateField),
						sqlbuilder.GetFieldName(NewDataTplField),
					).DbNameWithAlias()
					f.SetSelectColumns(setSelectColumns.AsAny()...)
				})
			})
			return fs
		})
	})
	if err != nil {
		return rootComponentHtml, err
	}
	rootComponentHtml.Templates = ToHtmlComponents(components...)
	attrs, err := s.attributeService.ListByRootComponentName(componentRootName, func(listParam *sqlbuilder.ListParam) {
		listParam.WithCustomFieldsFn(func(fs sqlbuilder.Fields) (customedFs sqlbuilder.Fields) {
			fs.FirstMust().Apply(func(f *sqlbuilder.Field, fs ...*sqlbuilder.Field) {
				f.SetDelayApply(func(f *sqlbuilder.Field, fs ...*sqlbuilder.Field) {
					setSelectColumns := f.GetTable().Columns.FilterByFieldName(
						sqlbuilder.GetFieldName(NewNodeIdField),
						sqlbuilder.GetFieldName(NewAttributeNameField),
						sqlbuilder.GetFieldName(NewAttributeValueField),
					).DbNameWithAlias()
					f.SetSelectColumns(setSelectColumns.AsAny()...)
				})
			})
			return fs
		})
	})
	if err != nil {
		return rootComponentHtml, err
	}
	rootComponentHtml.Name = componentRootName
	rootComponentHtml.Attributes = ToHtmlAttributes(attrs...)
	return rootComponentHtml, nil
}

type HtmlTemplateAdminService[C any, A any, R any] struct {
	Component ComponentSerivce[C]
	Assemble  AssembleService[A]
	Attribute AttributeService[R]
}

func NewHtmlTemplateAdminService[C any, A any, R any](dbHander sqlbuilder.Handler, customTableFn func(tableConfig TableConfig) (customedTableConfig TableConfig)) HtmlTemplateAdminService[C, A, R] {
	tableConfig := CustomTableConfig(dbHander, customTableFn)
	componentService := newComponentSerivce[C](tableConfig.Component)
	assembleService := newAssembleService[A](tableConfig.Assemble)
	attributeService := newAttributeService[R](tableConfig.Attribute)
	return HtmlTemplateAdminService[C, A, R]{
		Component: componentService,
		Assemble:  assembleService,
		Attribute: attributeService,
	}
}
