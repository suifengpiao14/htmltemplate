package htmlrepository

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/pkg/errors"
	"github.com/suifengpiao14/commonlanguage"
	"github.com/suifengpiao14/htmltemplate/htmlcomponent"
	"github.com/suifengpiao14/sqlbuilder"
)

func NewHtmlTemplateServiceByDefaultService[C Component, A Assemble, R Attribute](tableConfig TableConfig) *HtmlTemplateService[C, A, R] {
	return nil

}

type ComponentSerivce[C Component] struct {
	repositoryQueryAny sqlbuilder.RepositoryQuery[htmlcomponent.Component]
	repositoryQuery    sqlbuilder.RepositoryQuery[C]
	repositoryCommand  sqlbuilder.RepositoryCommand
}

func NewComponentSerivce[C Component](tableConfig sqlbuilder.TableConfig) ComponentSerivce[C] {
	repositoryQueryAny := sqlbuilder.NewRepositoryQuery[htmlcomponent.Component](tableConfig)
	repositoryQuery := sqlbuilder.NewRepositoryQuery[C](tableConfig)
	repositoryCommand := sqlbuilder.NewRepositoryCommand(tableConfig)
	return ComponentSerivce[C]{
		repositoryQueryAny: repositoryQueryAny,
		repositoryQuery:    repositoryQuery,
		repositoryCommand:  repositoryCommand,
	}
}

func (s ComponentSerivce[C]) Set(c C, customFn sqlbuilder.CustomFnSetParam) (err error) {
	fields := sqlbuilder.Fields{
		NewComponentNameField(c.GetName()).SetRequired(true).ShieldUpdate(true).AppendWhereFn(sqlbuilder.ValueFnForward),
		NewTemplateField(c.GetTemplate()).SetRequired(true),
		NewDataTplField(c.GetDataTpl()), //对于静态模板，无需数据
	}
	_, _, _, err = s.repositoryCommand.Set(fields, customFn)
	if err != nil {
		return err
	}
	return nil
}
func (s ComponentSerivce[C]) ListByComponentNames(componentNames []string) (models htmlcomponent.Components, err error) {
	fields := sqlbuilder.Fields{
		NewComponentNamesField(componentNames).SetRequired(true).AppendWhereFn(sqlbuilder.ValueFnForward).Apply(func(f *sqlbuilder.Field, fs ...*sqlbuilder.Field) {
			f.SetDelayApply(func(f *sqlbuilder.Field, fs ...*sqlbuilder.Field) {
				columns := f.GetTable().Columns
				componentNameColumn := columns.GetByFieldNameMust(sqlbuilder.GetFieldName(NewComponentNameField))
				templateColumn := columns.GetByFieldNameMust(sqlbuilder.GetFieldName(NewTemplateField))
				dataTplColumn := columns.GetByFieldNameMust(sqlbuilder.GetFieldName(NewDataTplField))
				f.SetSelectColumns(
					goqu.I(componentNameColumn.DbName).As(componentNameColumn.FieldName),
					goqu.I(templateColumn.DbName).As(templateColumn.FieldName),
					goqu.I(dataTplColumn.DbName).As(dataTplColumn.FieldName),
				)
			})

		}),
	}
	models, err = s.repositoryQueryAny.All(fields, nil)
	if err != nil {
		return nil, err
	}
	return models, nil
}

type AssembleService[A Assemble] struct {
	repositoryQuery   sqlbuilder.RepositoryQuery[A]
	repositoryCommand sqlbuilder.RepositoryCommand
}

func NewAssembleService[A Assemble](tableConfig sqlbuilder.TableConfig) AssembleService[A] {
	repositoryQuery := sqlbuilder.NewRepositoryQuery[A](tableConfig)
	repositoryCommand := sqlbuilder.NewRepositoryCommand(tableConfig)
	return AssembleService[A]{
		repositoryQuery:   repositoryQuery,
		repositoryCommand: repositoryCommand,
	}
}

func (s AssembleService[A]) Set(assemble A, customFn sqlbuilder.CustomFnSetParam) (err error) {
	fields := sqlbuilder.Fields{
		NewRootComponentNameField(assemble.GetRootComponentName()).SetRequired(true).ShieldUpdate(true).AppendWhereFn(sqlbuilder.ValueFnForward),
		NewAssembleNameField(assemble.GetAssembleName()).SetRequired(true).ShieldUpdate(true).AppendWhereFn(sqlbuilder.ValueFnForward),
		NewComponentNameField(assemble.GetAssembleName()),
		NewDataTplField(assemble.GetDataTpl()), //对于静态模板，无需数据
	}
	_, _, _, err = s.repositoryCommand.Set(fields, customFn)
	if err != nil {
		return err
	}
	return nil
}
func (s AssembleService[A]) ListByRootComponentName(rootComponentName string, customFn sqlbuilder.CustomFnListParam) (models []A, err error) {
	fields := sqlbuilder.Fields{
		NewRootComponentNameField(rootComponentName).SetRequired(true).AppendWhereFn(sqlbuilder.ValueFnForward),
	}
	models, err = s.repositoryQuery.All(fields, customFn)
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (s AssembleService[A]) Delete(assemble A, customFn sqlbuilder.CustomFnDeleteParam) (err error) {
	ctx := context.Background()
	_, err = s.repositoryCommand.GetTableConfig().MergeTableLevelFields(ctx).DeletedAt()
	if err != nil {
		err = errors.WithMessage(err, "should set sorft delete field by tableConfig.TableLevelFieldsHook")
		return err
	}
	fields := sqlbuilder.Fields{
		NewRootComponentNameField(assemble.GetRootComponentName()).SetRequired(true).AppendWhereFn(sqlbuilder.ValueFnForward),
		NewAssembleNameField(assemble.GetAssembleName()).SetRequired(true).AppendWhereFn(sqlbuilder.ValueFnForward),
	}
	err = s.repositoryCommand.Delete(fields, customFn)
	if err != nil {
		return err
	}
	return nil
}

type AttributeService[R Attribute] struct {
	repositoryQuery   sqlbuilder.RepositoryQuery[R]
	repositoryCommand sqlbuilder.RepositoryCommand
}

func NewAttributeService[R Attribute](tableConfig sqlbuilder.TableConfig) AttributeService[R] {
	repositoryQuery := sqlbuilder.NewRepositoryQuery[R](tableConfig)
	repositoryCommand := sqlbuilder.NewRepositoryCommand(tableConfig)
	return AttributeService[R]{
		repositoryQuery:   repositoryQuery,
		repositoryCommand: repositoryCommand,
	}
}

func (s AttributeService[R]) Set(attribute R, customFn sqlbuilder.CustomFnSetParam) (err error) {
	fields := sqlbuilder.Fields{
		NewNodeIdField(attribute.GetNodeId()).SetRequired(true).ShieldUpdate(true).AppendWhereFn(sqlbuilder.ValueFnForward),
		NewAttributeNameField(attribute.GetAttributeName()).SetRequired(true).ShieldUpdate(true).AppendWhereFn(sqlbuilder.ValueFnForward),
		NewAttributeValueField(attribute.GetAttributeValue()),
	}
	_, _, _, err = s.repositoryCommand.Set(fields, customFn)
	if err != nil {
		return err
	}
	return nil
}

func (s AttributeService[R]) ListByRootComponentName(rootComponentName string, customFn sqlbuilder.CustomFnListParam) (models []R, err error) {
	fields := sqlbuilder.Fields{
		NewRootComponentNameField(rootComponentName).SetRequired(true).AppendWhereFn(sqlbuilder.ValueFnForward),
	}
	models, err = s.repositoryQuery.All(fields, customFn)
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (s AttributeService[R]) Delete(attribute R, customFn sqlbuilder.CustomFnDeleteParam) (err error) {
	ctx := context.Background()
	_, err = s.repositoryCommand.GetTableConfig().MergeTableLevelFields(ctx).DeletedAt()
	if err != nil {
		err = errors.WithMessage(err, "should set sorft delete field by tableConfig.TableLevelFieldsHook")
		return err
	}
	fields := sqlbuilder.Fields{
		NewNodeIdField(attribute.GetNodeId()).SetRequired(true).ShieldUpdate(true).AppendWhereFn(sqlbuilder.ValueFnForward),
		NewAttributeNameField(attribute.GetAttributeName()).SetRequired(true).ShieldUpdate(true).AppendWhereFn(sqlbuilder.ValueFnForward),
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

func (s HtmlTemplateService[C, A, R]) ListByComponentNames(componentNames []string) (components htmlcomponent.Components, err error) {
	return s.componentService.ListByComponentNames(componentNames)
}

func (s HtmlTemplateService[C, A, R]) ComponentSet(c C, customFn sqlbuilder.CustomFnSetParam) (err error) {
	return s.componentService.Set(c, customFn)
}

func (s HtmlTemplateService[C, A, R]) ComponentPagination(pageIndex, pageSize int, customFn sqlbuilder.CustomFnPaginationParam) (models []C, total int64, err error) {
	fields := sqlbuilder.Fields{
		commonlanguage.NewPageIndex(pageIndex),
		commonlanguage.NewPageSize(pageSize),
	}

	return s.componentService.repositoryQuery.Pagination(fields, customFn)
}

func (s HtmlTemplateService[C, A, R]) AssembleSet(assemble A, customFn sqlbuilder.CustomFnSetParam) (err error) {
	return s.assembleService.Set(assemble, customFn)
}

func (s HtmlTemplateService[C, A, R]) AssembleGetAllByRootComponentName(rootComponentName string, customFn sqlbuilder.CustomFnListParam) ([]A, error) {

	return s.assembleService.ListByRootComponentName(rootComponentName, customFn)
}
func (s HtmlTemplateService[C, A, R]) AssembleDelete(assemble A, customFn sqlbuilder.CustomFnDeleteParam) (err error) {

	return s.assembleService.Delete(assemble, customFn)
}

func (s HtmlTemplateService[C, A, R]) AttributeSet(attribute R, customFn sqlbuilder.CustomFnSetParam) (err error) {

	return s.attributeService.Set(attribute, customFn)
}

func (s HtmlTemplateService[C, A, R]) AttributeGetAllByRootComponentName(rootComponentName string, customFn sqlbuilder.CustomFnListParam) (models []R, err error) {
	return s.attributeService.ListByRootComponentName(rootComponentName, customFn)
}
func (s HtmlTemplateService[C, A, R]) AttributeDelete(attribute R, customFn sqlbuilder.CustomFnDeleteParam) (err error) {

	return s.attributeService.Delete(attribute, customFn)
}
