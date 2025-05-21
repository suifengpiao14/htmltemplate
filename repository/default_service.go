package repository

import (
	"context"

	"github.com/pkg/errors"
	"github.com/suifengpiao14/commonlanguage"
	"github.com/suifengpiao14/sqlbuilder"
)

type ComponentSerivce[C Component] struct {
	repositoryQuery   sqlbuilder.RepositoryQuery[C]
	repositoryCommand sqlbuilder.RepositoryCommand
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
func (s AssembleService[A]) GetByRootComponentName(rootComponentName string, customFn sqlbuilder.CustomFnListParam) (models []A, err error) {
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

func (s AttributeService[R]) GetByRootComponentName(rootComponentName string, customFn sqlbuilder.CustomFnListParam) (models []R, err error) {
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

type HtmlTemplateDefaultService[C Component, A Assemble, R Attribute] struct {
	componentService ComponentSerivce[C]
	assembleService  AssembleService[A]
	attributeService AttributeService[R]
}

func NewHtmlTemplateDefaultService[C Component, A Assemble, R Attribute](tableConfig TableConfig) *HtmlTemplateDefaultService[C, A, R] {
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
	return &HtmlTemplateDefaultService[C, A, R]{
		componentService: componentService,
		assembleService:  assembleService,
		attributeService: attributeService,
	}
}

func (s HtmlTemplateDefaultService[C, A, R]) ListByComponentNames(componentNames []string) ([]C, error) {

	return s.componentService.ListByComponentNames(componentNames)
}

func (s HtmlTemplateDefaultService[C, A, R]) ComponentSet(c C, customFn sqlbuilder.CustomFnSetParam) (err error) {
	return s.componentService.Set(c, customFn)
}

func (s HtmlTemplateDefaultService[C, A, R]) ComponentPagination(pageIndex, pageSize int, customFn sqlbuilder.CustomFnPaginationParam) (models []C, total int64, err error) {
	fields := sqlbuilder.Fields{
		commonlanguage.NewPageIndex(pageIndex),
		commonlanguage.NewPageSize(pageSize),
	}

	return s.componentService.repositoryQuery.Pagination(fields, customFn)
}

func (s HtmlTemplateDefaultService[C, A, R]) AssembleSet(assemble A, customFn sqlbuilder.CustomFnSetParam) (err error) {
	return s.assembleService.Set(assemble, customFn)
}

func (s HtmlTemplateDefaultService[C, A, R]) AssembleGetAllByRootComponentName(rootComponentName string, customFn sqlbuilder.CustomFnListParam) ([]A, error) {

	return s.assembleService.GetByRootComponentName(rootComponentName, customFn)
}
func (s HtmlTemplateDefaultService[C, A, R]) AssembleDelete(assemble A, customFn sqlbuilder.CustomFnDeleteParam) (err error) {

	return s.assembleService.Delete(assemble, customFn)
}

func (s HtmlTemplateDefaultService[C, A, R]) AttributeSet(attribute R, customFn sqlbuilder.CustomFnSetParam) (err error) {

	return s.attributeService.Set(attribute, customFn)
}

func (s HtmlTemplateDefaultService[C, A, R]) AttributeGetAllByRootComponentName(rootComponentName string, customFn sqlbuilder.CustomFnListParam) (models []R, err error) {
	return s.attributeService.GetByRootComponentName(rootComponentName, customFn)
}
func (s HtmlTemplateDefaultService[C, A, R]) AttributeDelete(attribute R, customFn sqlbuilder.CustomFnDeleteParam) (err error) {

	return s.attributeService.Delete(attribute, customFn)
}
