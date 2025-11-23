package htmlrepository

import (
	"github.com/suifengpiao14/sqlbuilder"
)

type TemplateSerivce struct {
	sqlbuilder.RepositoryQuery
	sqlbuilder.RepositoryCommand
}

func newComponentSerivce(tableConfig sqlbuilder.TableConfig) TemplateSerivce {
	repository := tableConfig.Repository()
	return TemplateSerivce{
		RepositoryQuery:   repository.RepositoryQuery,
		RepositoryCommand: repository.RepositoryCommand,
	}
}

func (s TemplateSerivce) Set(c Template, customFn sqlbuilder.CustomFnSetParam) (err error) {
	fields := sqlbuilder.Fields{
		NewTemplateNameField(c.TemplateName).SetRequired(true).ShieldUpdate(true).AppendWhereFn(sqlbuilder.ValueFnForward),
		NewTemplateField(c.Template).SetRequired(true),
		NewDataTplField(c.DataTpl), //对于静态模板，无需数据
	}
	_, _, _, err = s.RepositoryCommand.Set(fields, customFn)
	if err != nil {
		return err
	}
	return nil
}

func (s TemplateSerivce) ListByTemplateNames(models any, componentNames []string, customFn sqlbuilder.CustomFnListParam) (err error) {
	fields := sqlbuilder.Fields{
		NewTemplateNamesField(componentNames).SetRequired(true).AppendWhereFn(sqlbuilder.ValueFnForward),
	}
	err = s.RepositoryQuery.All(models, fields, customFn)
	if err != nil {
		return err
	}
	return nil
}

type SlotService struct {
	sqlbuilder.RepositoryQuery
	sqlbuilder.RepositoryCommand
}

func newSlotService(tableConfig sqlbuilder.TableConfig) SlotService {
	repository := tableConfig.Repository()
	return SlotService{
		RepositoryQuery:   repository.RepositoryQuery,
		RepositoryCommand: repository.RepositoryCommand,
	}
}

func (s SlotService) Set(slotName Slot, customFn sqlbuilder.CustomFnSetParam) (err error) {
	fields := sqlbuilder.Fields{
		NewComponentNameField(slotName.TemplateName).SetRequired(true).ShieldUpdate(true).AppendWhereFn(sqlbuilder.ValueFnForward),
		NewSlotNameField(slotName.SlotName).SetRequired(true).ShieldUpdate(true).AppendWhereFn(sqlbuilder.ValueFnForward),
		NewTemplateNameField(slotName.ComponentName),
		NewDataTplField(slotName.DataTpl), //对于静态模板，无需数据
	}
	_, _, _, err = s.RepositoryCommand.Set(fields, customFn)
	if err != nil {
		return err
	}
	return nil
}
func (s SlotService) ListByComponentName(models any, componentName string, customFn sqlbuilder.CustomFnListParam) (err error) {
	fields := sqlbuilder.Fields{
		NewComponentNameField(componentName).SetRequired(true).AppendWhereFn(sqlbuilder.ValueFnForward),
	}
	err = s.RepositoryQuery.All(models, fields, customFn)
	if err != nil {
		return err
	}
	return nil
}

func (s SlotService) Delete(slotName Slot, customFn sqlbuilder.CustomFnDeleteParam) (err error) {

	fields := sqlbuilder.Fields{
		NewComponentNameField(slotName.TemplateName).SetRequired(true).AppendWhereFn(sqlbuilder.ValueFnForward),
		NewSlotNameField(slotName.SlotName).SetRequired(true).AppendWhereFn(sqlbuilder.ValueFnForward),
	}
	err = s.RepositoryCommand.Delete(fields, customFn)
	if err != nil {
		return err
	}
	return nil
}

type AttributeService struct {
	sqlbuilder.RepositoryQuery
	sqlbuilder.RepositoryCommand
}

func newAttributeService(tableConfig sqlbuilder.TableConfig) AttributeService {
	repository := tableConfig.Repository()
	return AttributeService{
		RepositoryQuery:   repository.RepositoryQuery,
		RepositoryCommand: repository.RepositoryCommand,
	}
}

func (s AttributeService) Set(attribute Attribute, customFn sqlbuilder.CustomFnSetParam) (err error) {
	fields := sqlbuilder.Fields{
		NewTagIdField(attribute.TagId).SetRequired(true).ShieldUpdate(true).AppendWhereFn(sqlbuilder.ValueFnForward),
		NewAttributeNameField(attribute.AttributeName).SetRequired(true).ShieldUpdate(true).AppendWhereFn(sqlbuilder.ValueFnForward),
		NewAttributeValueField(attribute.AttributeValue),
	}
	_, _, _, err = s.RepositoryCommand.Set(fields, customFn)
	if err != nil {
		return err
	}
	return nil
}

func (s AttributeService) ListByTemplateNames(models any, templateNames []string, customFn sqlbuilder.CustomFnListParam) (err error) {
	fields := sqlbuilder.Fields{
		NewTemplateNamesField(templateNames).SetRequired(true).AppendWhereFn(sqlbuilder.ValueFnForward),
	}
	err = s.RepositoryQuery.All(models, fields, customFn)
	if err != nil {
		return err
	}
	return nil
}

func (s AttributeService) Delete(attribute Attribute, customFn sqlbuilder.CustomFnDeleteParam) (err error) {
	fields := sqlbuilder.Fields{
		NewTagIdField(attribute.TagId).SetRequired(true).ShieldUpdate(true).AppendWhereFn(sqlbuilder.ValueFnForward),
		NewAttributeNameField(attribute.AttributeName).SetRequired(true).ShieldUpdate(true).AppendWhereFn(sqlbuilder.ValueFnForward),
	}
	err = s.RepositoryCommand.Delete(fields, customFn)
	if err != nil {
		return err
	}
	return nil
}
