package htmlrepository

import (
	"context"

	"github.com/pkg/errors"
	"github.com/suifengpiao14/htmltemplate/htmlcomponent"
	"github.com/suifengpiao14/sqlbuilder"
)

type _ComponentSerivce[C any] struct {
	repositoryQuery   sqlbuilder.RepositoryQuery[C]
	repositoryCommand sqlbuilder.RepositoryCommand
}

func newComponentSerivce[C any](tableConfig sqlbuilder.TableConfig) _ComponentSerivce[C] {
	repositoryQuery := sqlbuilder.NewRepositoryQuery[C](tableConfig)
	repositoryCommand := sqlbuilder.NewRepositoryCommand(tableConfig)
	return _ComponentSerivce[C]{
		repositoryQuery:   repositoryQuery,
		repositoryCommand: repositoryCommand,
	}
}

func (s _ComponentSerivce[C]) Set(c htmlcomponent.Component, customFn sqlbuilder.CustomFnSetParam) (err error) {
	fields := sqlbuilder.Fields{
		NewComponentNameField(c.ComponentName).SetRequired(true).ShieldUpdate(true).AppendWhereFn(sqlbuilder.ValueFnForward),
		NewTemplateField(c.Template).SetRequired(true),
		NewDataTplField(c.DataTpl), //对于静态模板，无需数据
	}
	_, _, _, err = s.repositoryCommand.Set(fields, customFn)
	if err != nil {
		return err
	}
	return nil
}
func (s _ComponentSerivce[C]) ListByComponentNames(componentNames []string, customFn sqlbuilder.CustomFnListParam) (models []C, err error) {
	fields := sqlbuilder.Fields{
		NewComponentNamesField(componentNames).SetRequired(true).AppendWhereFn(sqlbuilder.ValueFnForward),
	}
	models, err = s.repositoryQuery.All(fields, customFn)
	if err != nil {
		return nil, err
	}
	return models, nil
}

type _AssembleService[A any] struct {
	repositoryQuery   sqlbuilder.RepositoryQuery[A]
	repositoryCommand sqlbuilder.RepositoryCommand
}

func newAssembleService[A any](tableConfig sqlbuilder.TableConfig) _AssembleService[A] {
	repositoryQuery := sqlbuilder.NewRepositoryQuery[A](tableConfig)
	repositoryCommand := sqlbuilder.NewRepositoryCommand(tableConfig)
	return _AssembleService[A]{
		repositoryQuery:   repositoryQuery,
		repositoryCommand: repositoryCommand,
	}
}

func (s _AssembleService[A]) Set(assemble htmlcomponent.Assemble, customFn sqlbuilder.CustomFnSetParam) (err error) {
	fields := sqlbuilder.Fields{
		NewRootComponentNameField(assemble.RootComponentName).SetRequired(true).ShieldUpdate(true).AppendWhereFn(sqlbuilder.ValueFnForward),
		NewAssembleNameField(assemble.AssembleName).SetRequired(true).ShieldUpdate(true).AppendWhereFn(sqlbuilder.ValueFnForward),
		NewComponentNameField(assemble.ComponentName),
		NewDataTplField(assemble.DataTpl), //对于静态模板，无需数据
	}
	_, _, _, err = s.repositoryCommand.Set(fields, customFn)
	if err != nil {
		return err
	}
	return nil
}
func (s _AssembleService[A]) ListByRootComponentName(rootComponentName string, customFn sqlbuilder.CustomFnListParam) (models []A, err error) {
	fields := sqlbuilder.Fields{
		NewRootComponentNameField(rootComponentName).SetRequired(true).AppendWhereFn(sqlbuilder.ValueFnForward),
	}
	models, err = s.repositoryQuery.All(fields, customFn)
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (s _AssembleService[A]) Delete(assemble htmlcomponent.Assemble, customFn sqlbuilder.CustomFnDeleteParam) (err error) {
	ctx := context.Background()
	_, err = s.repositoryCommand.GetTableConfig().MergeTableLevelFields(ctx).DeletedAt()
	if err != nil {
		err = errors.WithMessage(err, "should set sorft delete field by tableConfig.TableLevelFieldsHook")
		return err
	}
	fields := sqlbuilder.Fields{
		NewRootComponentNameField(assemble.RootComponentName).SetRequired(true).AppendWhereFn(sqlbuilder.ValueFnForward),
		NewAssembleNameField(assemble.AssembleName).SetRequired(true).AppendWhereFn(sqlbuilder.ValueFnForward),
	}
	err = s.repositoryCommand.Delete(fields, customFn)
	if err != nil {
		return err
	}
	return nil
}

type _AttributeService[R any] struct {
	repositoryQuery   sqlbuilder.RepositoryQuery[R]
	repositoryCommand sqlbuilder.RepositoryCommand
}

func newAttributeService[R any](tableConfig sqlbuilder.TableConfig) _AttributeService[R] {
	repositoryQuery := sqlbuilder.NewRepositoryQuery[R](tableConfig)
	repositoryCommand := sqlbuilder.NewRepositoryCommand(tableConfig)
	return _AttributeService[R]{
		repositoryQuery:   repositoryQuery,
		repositoryCommand: repositoryCommand,
	}
}

func (s _AttributeService[R]) Set(attribute htmlcomponent.Attribute, customFn sqlbuilder.CustomFnSetParam) (err error) {
	fields := sqlbuilder.Fields{
		NewNodeIdField(attribute.NodeId).SetRequired(true).ShieldUpdate(true).AppendWhereFn(sqlbuilder.ValueFnForward),
		NewAttributeNameField(attribute.AttributeName).SetRequired(true).ShieldUpdate(true).AppendWhereFn(sqlbuilder.ValueFnForward),
		NewAttributeValueField(attribute.AttributeValue),
	}
	_, _, _, err = s.repositoryCommand.Set(fields, customFn)
	if err != nil {
		return err
	}
	return nil
}

func (s _AttributeService[R]) ListByRootComponentName(rootComponentName string, customFn sqlbuilder.CustomFnListParam) (models []R, err error) {
	fields := sqlbuilder.Fields{
		NewRootComponentNameField(rootComponentName).SetRequired(true).AppendWhereFn(sqlbuilder.ValueFnForward),
	}
	models, err = s.repositoryQuery.All(fields, customFn)
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (s _AttributeService[R]) Delete(attribute htmlcomponent.Attribute, customFn sqlbuilder.CustomFnDeleteParam) (err error) {
	ctx := context.Background()
	_, err = s.repositoryCommand.GetTableConfig().MergeTableLevelFields(ctx).DeletedAt()
	if err != nil {
		err = errors.WithMessage(err, "should set sorft delete field by tableConfig.TableLevelFieldsHook")
		return err
	}
	fields := sqlbuilder.Fields{
		NewNodeIdField(attribute.NodeId).SetRequired(true).ShieldUpdate(true).AppendWhereFn(sqlbuilder.ValueFnForward),
		NewAttributeNameField(attribute.AttributeName).SetRequired(true).ShieldUpdate(true).AppendWhereFn(sqlbuilder.ValueFnForward),
	}
	err = s.repositoryCommand.Delete(fields, customFn)
	if err != nil {
		return err
	}
	return nil
}

type TableConfig struct {
	Component sqlbuilder.TableConfig
	Assemble  sqlbuilder.TableConfig
	Attribute sqlbuilder.TableConfig
}
