package htmlrepository

import (
	"github.com/suifengpiao14/htmltemplate/htmlcomponent"
	"github.com/suifengpiao14/sqlbuilder"
)

type HtmlTemplateApiService struct {
	componentService TemplateSerivce
	slotNameService  SlotService
	attributeService AttributeService
}

func NewHtmlTemplateApiService(dbHander sqlbuilder.Handler, customTableFn func(tableConfig TableConfig) (customedTableConfig TableConfig)) *HtmlTemplateApiService {
	tableConfig := CustomTableConfig(dbHander, customTableFn)
	componentService := newComponentSerivce(tableConfig.Template)
	slotNameService := newSlotService(tableConfig.Slot)
	attributeService := newAttributeService(tableConfig.Attribute)
	return &HtmlTemplateApiService{
		componentService: componentService,
		slotNameService:  slotNameService,
		attributeService: attributeService,
	}
}

func (s HtmlTemplateApiService) Render(componentName string, data map[string]any) (componentHtml string, err error) {
	component, err := s.GetComponent(componentName)
	if err != nil {
		return "", err
	}
	componentHtml, err = component.Render(data)
	if err != nil {
		return "", err
	}
	return componentHtml, nil
}

func (s HtmlTemplateApiService) GetComponent(componentName string) (componentHtml htmlcomponent.Component, err error) {
	slotNames := make([]Slot, 0)
	err = s.slotNameService.ListByComponentName(&slotNames, componentName, func(listParam *sqlbuilder.ListParam) {
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
	templateNames := componentHtml.Slots.TemplateNames()
	templates := make([]Template, 0)
	err = s.componentService.ListByTemplateNames(&templates, templateNames, func(listParam *sqlbuilder.ListParam) {
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
	componentHtml.Templates = ToHtmlComponents(templates...)
	attrs := make([]Attribute, 0)
	err = s.attributeService.ListByTemplateNames(&attrs, templateNames, func(listParam *sqlbuilder.ListParam) {
		listParam.WithCustomFieldsFn(func(fs sqlbuilder.Fields) (customedFs sqlbuilder.Fields) {
			fs.FirstMust().Apply(func(f *sqlbuilder.Field, fs ...*sqlbuilder.Field) {
				f.SetDelayApply(func(f *sqlbuilder.Field, fs ...*sqlbuilder.Field) {
					setSelectColumns := f.GetTable().Columns.FilterByFieldName(
						sqlbuilder.GetFieldName(NewSlotNameField),
						sqlbuilder.GetFieldName(NewTemplateNameField),
						sqlbuilder.GetFieldName(NewTagIdField),
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

type HtmlTemplateAdminService struct {
	Template  TemplateSerivce
	Slot      SlotService
	Attribute AttributeService
}

func NewHtmlTemplateAdminService(dbHander sqlbuilder.Handler, customTableFn func(tableConfig TableConfig) (customedTableConfig TableConfig)) HtmlTemplateAdminService {
	tableConfig := CustomTableConfig(dbHander, customTableFn)
	componentService := newComponentSerivce(tableConfig.Template)
	slotNameService := newSlotService(tableConfig.Slot)
	attributeService := newAttributeService(tableConfig.Attribute)
	return HtmlTemplateAdminService{
		Template:  componentService,
		Slot:      slotNameService,
		Attribute: attributeService,
	}
}
