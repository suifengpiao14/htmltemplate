package htmlrepository

import (
	"github.com/suifengpiao14/htmltemplate/htmlcomponent"
	"github.com/suifengpiao14/memorytable"
	"github.com/suifengpiao14/sqlbuilder"
)

type HtmlTemplateApiService struct {
	componentService ComponentSerivce[Template]
	slotNameService  SlotService[Slot]
	attributeService AttributeService[Attribute]
}

func NewHtmlTemplateApiService(dbHander sqlbuilder.Handler, customTableFn func(tableConfig TableConfig) (customedTableConfig TableConfig)) *HtmlTemplateApiService {
	tableConfig := CustomTableConfig(dbHander, customTableFn)
	componentService := newComponentSerivce[Template](tableConfig.Component)
	slotNameService := newSlotService[Slot](tableConfig.Slot)
	attributeService := newAttributeService[Attribute](tableConfig.Attribute)
	return &HtmlTemplateApiService{
		componentService: componentService,
		slotNameService:  slotNameService,
		attributeService: attributeService,
	}
}

func (s HtmlTemplateApiService) Render(componentName string, data map[string]any) (componentHtml string, err error) {
	rootComponent, err := s.GetComponent(componentName)
	if err != nil {
		return "", err
	}
	componentHtml, err = rootComponent.Render(data)
	if err != nil {
		return "", err
	}
	return componentHtml, nil
}

func (s HtmlTemplateApiService) GetComponent(componentName string) (componentHtml htmlcomponent.Component, err error) {
	slotNames, err := s.slotNameService.ListByRootComponentName(componentName, func(listParam *sqlbuilder.ListParam) {
		listParam.WithCustomFieldsFn(func(fs sqlbuilder.Fields) (customedFs sqlbuilder.Fields) {
			fs.FirstMust().Apply(func(f *sqlbuilder.Field, fs ...*sqlbuilder.Field) {
				f.SetDelayApply(func(f *sqlbuilder.Field, fs ...*sqlbuilder.Field) {
					setSelectColumns := f.GetTable().Columns.FilterByFieldName(
						sqlbuilder.GetFieldName(NewComponentNameField),
						sqlbuilder.GetFieldName(NewTemplateNameField),
						sqlbuilder.GetFieldName(NewSlotNameField),
						sqlbuilder.GetFieldName(NewDataTplField),
					).DbNameWithAlias()
					f.SetSelectColumns(setSelectColumns.AsAny()...)
				})
			})
			return fs
		})
	})
	if err != nil {
		return componentHtml, err
	}
	componentHtml.Slots = ToHtmlSlots(slotNames...)
	componentNames := componentHtml.Slots.ComponentNames()
	componentNames = append(componentNames, componentName)
	componentNames = memorytable.NewTable(componentNames...).Uniqueue(func(row string) (key string) { return key }).ToSlice()
	components, err := s.componentService.ListByComponentNames(componentNames, func(listParam *sqlbuilder.ListParam) {
		listParam.WithCustomFieldsFn(func(fs sqlbuilder.Fields) (customedFs sqlbuilder.Fields) {
			fs.FirstMust().Apply(func(f *sqlbuilder.Field, fs ...*sqlbuilder.Field) {
				f.SetDelayApply(func(f *sqlbuilder.Field, fs ...*sqlbuilder.Field) {
					setSelectColumns := f.GetTable().Columns.FilterByFieldName(
						sqlbuilder.GetFieldName(NewTemplateNameField),
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
		return componentHtml, err
	}
	componentHtml.Templates = ToHtmlComponents(components...)
	attrs, err := s.attributeService.ListByRootComponentName(componentName, func(listParam *sqlbuilder.ListParam) {
		listParam.WithCustomFieldsFn(func(fs sqlbuilder.Fields) (customedFs sqlbuilder.Fields) {
			fs.FirstMust().Apply(func(f *sqlbuilder.Field, fs ...*sqlbuilder.Field) {
				f.SetDelayApply(func(f *sqlbuilder.Field, fs ...*sqlbuilder.Field) {
					setSelectColumns := f.GetTable().Columns.FilterByFieldName(
						sqlbuilder.GetFieldName(NewNodeIdField),
						sqlbuilder.GetFieldName(NewSlotNameField),
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
		return componentHtml, err
	}
	componentHtml.Name = componentName
	componentHtml.Attributes = ToHtmlAttributes(attrs...)
	return componentHtml, nil
}

type HtmlTemplateAdminService[C any, A any, R any] struct {
	Component ComponentSerivce[C]
	Slot      SlotService[A]
	Attribute AttributeService[R]
}

func NewHtmlTemplateAdminService[C any, A any, R any](dbHander sqlbuilder.Handler, customTableFn func(tableConfig TableConfig) (customedTableConfig TableConfig)) HtmlTemplateAdminService[C, A, R] {
	tableConfig := CustomTableConfig(dbHander, customTableFn)
	componentService := newComponentSerivce[C](tableConfig.Component)
	slotNameService := newSlotService[A](tableConfig.Slot)
	attributeService := newAttributeService[R](tableConfig.Attribute)
	return HtmlTemplateAdminService[C, A, R]{
		Component: componentService,
		Slot:      slotNameService,
		Attribute: attributeService,
	}
}
