package htmlrepository

import "github.com/suifengpiao14/sqlbuilder"

type TableConfig struct {
	Component sqlbuilder.TableConfig
	Assemble  sqlbuilder.TableConfig
	Attribute sqlbuilder.TableConfig
}

var table_component = sqlbuilder.NewTableConfig("t_component").AddColumns(
	sqlbuilder.NewColumnConfig("Fcomponent_name", sqlbuilder.GetFieldName(NewComponentNameField)),
	sqlbuilder.NewColumnConfig("Ftemplate", sqlbuilder.GetFieldName(NewTemplateField)),
	sqlbuilder.NewColumnConfig("Fdata_tpl", sqlbuilder.GetFieldName(NewDataTplField)),
).AddIndexs(sqlbuilder.Index{
	ColumnNames: func(tableColumns sqlbuilder.ColumnConfigs) (columnNames []string) {
		columnNames = tableColumns.FieldName2ColumnName(
			sqlbuilder.GetFieldName(NewComponentNameField),
		)
		return columnNames
	},
})

var table_assemble = sqlbuilder.NewTableConfig("t_assemble").AddColumns(
	sqlbuilder.NewColumnConfig("Froot_component_name", sqlbuilder.GetFieldName(NewRootComponentNameField)),
	sqlbuilder.NewColumnConfig("Fcomponent_name", sqlbuilder.GetFieldName(NewComponentNameField)),
	sqlbuilder.NewColumnConfig("Fassemble_name", sqlbuilder.GetFieldName(NewAssembleNameField)),
	sqlbuilder.NewColumnConfig("Fdata_tpl", sqlbuilder.GetFieldName(NewDataTplField)),
).AddIndexs(sqlbuilder.Index{
	ColumnNames: func(tableColumns sqlbuilder.ColumnConfigs) (columnNames []string) {
		columnNames = tableColumns.FieldName2ColumnName(
			sqlbuilder.GetFieldName(NewRootComponentNameField),
			sqlbuilder.GetFieldName(NewAssembleNameField),
		)
		return columnNames
	},
})

var table_attribute = sqlbuilder.NewTableConfig("t_attribute").AddColumns(
	sqlbuilder.NewColumnConfig("Froot_component_name", sqlbuilder.GetFieldName(NewRootComponentNameField)),
	sqlbuilder.NewColumnConfig("Fnode_id", sqlbuilder.GetFieldName(NewNodeIdField)),
	sqlbuilder.NewColumnConfig("Fattr_name", sqlbuilder.GetFieldName(NewAttributeNameField)),
	sqlbuilder.NewColumnConfig("Fattr_value", sqlbuilder.GetFieldName(NewAttributeValueField)),
).AddIndexs(sqlbuilder.Index{
	ColumnNames: func(tableColumns sqlbuilder.ColumnConfigs) (columnNames []string) {
		columnNames = tableColumns.FieldName2ColumnName(
			sqlbuilder.GetFieldName(NewRootComponentNameField),
			sqlbuilder.GetFieldName(NewNodeIdField),
			sqlbuilder.GetFieldName(NewAttributeNameField),
		)
		return columnNames
	},
})

//CustomTableConfig 初始化表配置信息
func CustomTableConfig(dbHandler sqlbuilder.Handler, configFn func(table TableConfig) (configedTable TableConfig)) TableConfig {
	var tableConfig = TableConfig{
		Component: table_component,
		Assemble:  table_assemble,
		Attribute: table_attribute,
	}

	tableConfig.Component = tableConfig.Component.WithHandler(dbHandler)
	tableConfig.Assemble = tableConfig.Assemble.WithHandler(dbHandler)
	tableConfig.Attribute = tableConfig.Attribute.WithHandler(dbHandler)
	if configFn != nil {
		tableConfig = configFn(tableConfig)
	}
	return tableConfig
}
