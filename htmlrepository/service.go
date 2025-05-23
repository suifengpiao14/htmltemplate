package htmlrepository

import (
	"github.com/suifengpiao14/commonlanguage"
	"github.com/suifengpiao14/htmltemplate/htmlcomponent"
	"github.com/suifengpiao14/memorytable"
	"github.com/suifengpiao14/sqlbuilder"
)

type HtmlTemplateApiService struct {
	componentService _ComponentSerivce[Component]
	assembleService  _AssembleService[Assemble]
	attributeService _AttributeService[Attribute]
}

func NewHtmlTemplateApiService(dbHander sqlbuilder.Handler, customTableFn func(tableConfig TableConfig) (customedTableConfig TableConfig)) *HtmlTemplateApiService {
	tableConfig := customTableConfig(dbHander, customTableFn)
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

func (s HtmlTemplateApiService) GetComponent(componentRootName string) (rootComponentHtml htmlcomponent.RootComponent, err error) {
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
	rootComponentHtml.Assembles = ToHtmlAssembles(assembles...)
	componentNames := rootComponentHtml.Assembles.ComponentNames()
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
	rootComponentHtml.Components = ToHtmlComponents(components...)
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
	componentService _ComponentSerivce[C]
	assembleService  _AssembleService[A]
	attributeService _AttributeService[R]
}

func NewHtmlTemplateAdminService[C any, A any, R any](dbHander sqlbuilder.Handler, customTableFn func(tableConfig TableConfig) (customedTableConfig TableConfig)) HtmlTemplateAdminService[C, A, R] {
	tableConfig := customTableConfig(dbHander, customTableFn)
	componentService := newComponentSerivce[C](tableConfig.Component)
	assembleService := newAssembleService[A](tableConfig.Assemble)
	attributeService := newAttributeService[R](tableConfig.Attribute)
	return HtmlTemplateAdminService[C, A, R]{
		componentService: componentService,
		assembleService:  assembleService,
		attributeService: attributeService,
	}
}

func (s HtmlTemplateAdminService[C, A, R]) ListByComponentNames(componentNames []string) (components []C, err error) {
	return s.componentService.ListByComponentNames(componentNames, nil)
}

func (s HtmlTemplateAdminService[C, A, R]) ComponentSet(c Component, customFn sqlbuilder.CustomFnSetParam) (err error) {
	return s.componentService.Set(c, customFn)
}

func (s HtmlTemplateAdminService[C, A, R]) ComponentFirst(fields sqlbuilder.Fields, customFn sqlbuilder.CustomFnFirstParam) (model C, exists bool, err error) {
	return s.componentService.First(fields, customFn)
}

func (s HtmlTemplateAdminService[C, A, R]) ComponentPagination(pageIndex, pageSize int, customFn sqlbuilder.CustomFnPaginationParam) (models []C, total int64, err error) {
	fields := sqlbuilder.Fields{
		commonlanguage.NewPageIndex(pageIndex),
		commonlanguage.NewPageSize(pageSize),
	}
	return s.componentService.Pagination(fields, customFn)
}
func (s HtmlTemplateAdminService[C, A, R]) ComponentAll(fields sqlbuilder.Fields, customFn sqlbuilder.CustomFnListParam) (models []C, err error) {
	return s.componentService.All(fields, customFn)
}

func (s HtmlTemplateAdminService[C, A, R]) AssembleSet(assemble Assemble, customFn sqlbuilder.CustomFnSetParam) (err error) {
	return s.assembleService.Set(assemble, customFn)
}

func (s HtmlTemplateAdminService[C, A, R]) AssembleGetAllByRootComponentName(rootComponentName string, customFn sqlbuilder.CustomFnListParam) ([]A, error) {

	return s.assembleService.ListByRootComponentName(rootComponentName, customFn)
}
func (s HtmlTemplateAdminService[C, A, R]) AssembleDelete(assemble Assemble, customFn sqlbuilder.CustomFnDeleteParam) (err error) {

	return s.assembleService.Delete(assemble, customFn)
}

func (s HtmlTemplateAdminService[C, A, R]) AttributeSet(attribute Attribute, customFn sqlbuilder.CustomFnSetParam) (err error) {

	return s.attributeService.Set(attribute, customFn)
}

func (s HtmlTemplateAdminService[C, A, R]) AttributeGetAllByRootComponentName(rootComponentName string, customFn sqlbuilder.CustomFnListParam) (models []R, err error) {
	return s.attributeService.ListByRootComponentName(rootComponentName, customFn)
}
func (s HtmlTemplateAdminService[C, A, R]) AttributeDelete(attribute Attribute, customFn sqlbuilder.CustomFnDeleteParam) (err error) {
	return s.attributeService.Delete(attribute, customFn)
}
