package htmlrepository

import (
	"context"

	"github.com/pkg/errors"
	"github.com/suifengpiao14/commonlanguage"
	"github.com/suifengpiao14/sqlbuilder"
)

func NewHtmlTemplateServiceByDefaultService[C Component, A Assemble, R Attribute](tableConfig TableConfig) *HtmlTemplateService[C, A, R] {
	return nil

}

type ComponentSerivce[C Component] struct {
	repositoryQuery   sqlbuilder.RepositoryQuery[C]
	repositoryCommand sqlbuilder.RepositoryCommand
}

func NewComponentSerivce[C Component](tableConfig sqlbuilder.TableConfig) ComponentSerivce[C] {
	repositoryQuery := sqlbuilder.NewRepositoryQuery[C](tableConfig)
	repositoryCommand := sqlbuilder.NewRepositoryCommand(tableConfig)
	return ComponentSerivce[C]{
		repositoryQuery:   repositoryQuery,
		repositoryCommand: repositoryCommand,
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
func (s ComponentSerivce[C]) ListByComponentNames(componentNames []string) (models []C, err error) {
	fields := sqlbuilder.Fields{
		NewComponentNamesField(componentNames).SetRequired(true).AppendWhereFn(sqlbuilder.ValueFnForward),
	}
	return s.repositoryQuery.All(fields, nil)
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
	_, err = s.repositoryCommand.GetTableConfig().RunTableLevelFieldsHook(ctx, sqlbuilder.SCENE_SQL_DELETE).DeletedAt()
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
	_, err = s.repositoryCommand.GetTableConfig().RunTableLevelFieldsHook(ctx, sqlbuilder.SCENE_SQL_DELETE).DeletedAt()
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

func (c TableConfig) AddIndex() TableConfig {
	// 组件名称唯一索引
	componentColumn := c.Component.Columns.GetByFieldNameMust(NewComponentNameField("").Name)
	c.Component.Indexs.Append(sqlbuilder.Index{Unique: true, ColumnNames: []string{componentColumn.DbName}}) // 组件唯一索引

	// 组合中同一个rootComponentName 下 assembleName 唯一索引
	rootComponentColumn := c.Assemble.Columns.GetByFieldNameMust(NewRootComponentNameField("").Name)
	AssembleColumn := c.Assemble.Columns.GetByFieldNameMust(NewAssembleNameField("").Name)
	c.Assemble.Indexs.Append(sqlbuilder.Index{Unique: true, ColumnNames: []string{AssembleColumn.DbName, rootComponentColumn.DbName}})

	// 同一个节点，属性名称唯一索引
	nodeIdColumn := c.Attribute.Columns.GetByFieldNameMust(NewNodeIdField("").Name)
	attributeColumn := c.Attribute.Columns.GetByFieldNameMust(NewAttributeNameField("").Name)
	c.Attribute.Indexs.Append(sqlbuilder.Index{Unique: true, ColumnNames: []string{nodeIdColumn.DbName, attributeColumn.DbName}})

	return c
}

func (s HtmlTemplateService[C, A, R]) ListByComponentNames(componentNames []string) ([]C, error) {

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
