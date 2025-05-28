package htmlrepository

import (
	"context"

	"github.com/pkg/errors"
	"github.com/suifengpiao14/sqlbuilder"
)

type ComponentSerivce[C any] struct {
	sqlbuilder.RepositoryQuery[C]
	sqlbuilder.RepositoryCommand
}

func newComponentSerivce[C any](tableConfig sqlbuilder.TableConfig) ComponentSerivce[C] {
	repositoryQuery := sqlbuilder.NewRepositoryQuery[C](tableConfig)
	repositoryCommand := sqlbuilder.NewRepositoryCommand(tableConfig)
	return ComponentSerivce[C]{
		RepositoryQuery:   repositoryQuery,
		RepositoryCommand: repositoryCommand,
	}
}

func (s ComponentSerivce[C]) Set(c Template, customFn sqlbuilder.CustomFnSetParam) (err error) {
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

func (s ComponentSerivce[C]) ListByComponentNames(componentNames []string, customFn sqlbuilder.CustomFnListParam) (models []C, err error) {
	fields := sqlbuilder.Fields{
		NewTemplateNamesField(componentNames).SetRequired(true).AppendWhereFn(sqlbuilder.ValueFnForward),
	}
	models, err = s.RepositoryQuery.All(fields, customFn)
	if err != nil {
		return nil, err
	}
	return models, nil
}

type SlotService[A any] struct {
	sqlbuilder.RepositoryQuery[A]
	sqlbuilder.RepositoryCommand
}

func newSlotService[A any](tableConfig sqlbuilder.TableConfig) SlotService[A] {
	repositoryQuery := sqlbuilder.NewRepositoryQuery[A](tableConfig)
	repositoryCommand := sqlbuilder.NewRepositoryCommand(tableConfig)
	return SlotService[A]{
		RepositoryQuery:   repositoryQuery,
		RepositoryCommand: repositoryCommand,
	}
}

func (s SlotService[A]) Set(slotName Slot, customFn sqlbuilder.CustomFnSetParam) (err error) {
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
func (s SlotService[A]) ListByRootComponentName(rootComponentName string, customFn sqlbuilder.CustomFnListParam) (models []A, err error) {
	fields := sqlbuilder.Fields{
		NewComponentNameField(rootComponentName).SetRequired(true).AppendWhereFn(sqlbuilder.ValueFnForward),
	}
	models, err = s.RepositoryQuery.All(fields, customFn)
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (s SlotService[A]) Delete(slotName Slot, customFn sqlbuilder.CustomFnDeleteParam) (err error) {
	ctx := context.Background()

	_, err = s.RepositoryCommand.GetTableConfig().MergeTableLevelFields(ctx).DeletedAt()
	if err != nil {
		err = errors.WithMessage(err, "should set sorft delete field by tableConfig.TableLevelFieldsHook")
		return err
	}
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

type AttributeService[R any] struct {
	sqlbuilder.RepositoryQuery[R]
	sqlbuilder.RepositoryCommand
}

func newAttributeService[R any](tableConfig sqlbuilder.TableConfig) AttributeService[R] {
	repositoryQuery := sqlbuilder.NewRepositoryQuery[R](tableConfig)
	repositoryCommand := sqlbuilder.NewRepositoryCommand(tableConfig)
	return AttributeService[R]{
		RepositoryQuery:   repositoryQuery,
		RepositoryCommand: repositoryCommand,
	}
}

func (s AttributeService[R]) Set(attribute Attribute, customFn sqlbuilder.CustomFnSetParam) (err error) {
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

func (s AttributeService[R]) ListByRootComponentName(rootComponentName string, customFn sqlbuilder.CustomFnListParam) (models []R, err error) {
	fields := sqlbuilder.Fields{
		NewComponentNameField(rootComponentName).SetRequired(true).AppendWhereFn(sqlbuilder.ValueFnForward),
	}
	models, err = s.RepositoryQuery.All(fields, customFn)
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (s AttributeService[R]) Delete(attribute Attribute, customFn sqlbuilder.CustomFnDeleteParam) (err error) {
	ctx := context.Background()
	_, err = s.RepositoryCommand.GetTableConfig().MergeTableLevelFields(ctx).DeletedAt()
	if err != nil {
		err = errors.WithMessage(err, "should set sorft delete field by tableConfig.TableLevelFieldsHook")
		return err
	}
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
