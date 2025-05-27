package htmlrepository

import "github.com/suifengpiao14/sqlbuilder"

type TableConfig struct {
	Component sqlbuilder.TableConfig
	Slot      sqlbuilder.TableConfig
	Attribute sqlbuilder.TableConfig
}

var table_component = sqlbuilder.NewTableConfig("t_component").AddColumns(
	sqlbuilder.NewColumnConfig("Fcomponent_name", sqlbuilder.GetFieldName(NewComponentNameField)),
	sqlbuilder.NewColumnConfig("Ftemplate", sqlbuilder.GetFieldName(NewTemplateField)),
	sqlbuilder.NewColumnConfig("Fdata_tpl", sqlbuilder.GetFieldName(NewDataTplField)),
).AddIndexs(sqlbuilder.Index{
	Unique: true,
	ColumnNames: func(tableColumns sqlbuilder.ColumnConfigs) (columnNames []string) {
		columnNames = tableColumns.FieldName2ColumnName(
			sqlbuilder.GetFieldName(NewComponentNameField),
		)
		return columnNames
	},
})

var table_slotName = sqlbuilder.NewTableConfig("t_slotName").AddColumns(
	sqlbuilder.NewColumnConfig("Froot_component_name", sqlbuilder.GetFieldName(NewRootComponentNameField)),
	sqlbuilder.NewColumnConfig("Fcomponent_name", sqlbuilder.GetFieldName(NewComponentNameField)),
	sqlbuilder.NewColumnConfig("FslotName_name", sqlbuilder.GetFieldName(NewSlotNameField)),
	sqlbuilder.NewColumnConfig("Fdata_tpl", sqlbuilder.GetFieldName(NewDataTplField)),
).AddIndexs(sqlbuilder.Index{
	Unique: true,
	ColumnNames: func(tableColumns sqlbuilder.ColumnConfigs) (columnNames []string) {
		columnNames = tableColumns.FieldName2ColumnName(
			sqlbuilder.GetFieldName(NewRootComponentNameField),
			sqlbuilder.GetFieldName(NewSlotNameField),
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
	Unique: true,
	ColumnNames: func(tableColumns sqlbuilder.ColumnConfigs) (columnNames []string) {
		columnNames = tableColumns.FieldName2ColumnName(
			sqlbuilder.GetFieldName(NewRootComponentNameField),
			sqlbuilder.GetFieldName(NewNodeIdField),
			sqlbuilder.GetFieldName(NewAttributeNameField),
		)
		return columnNames
	},
})

// CustomTableConfig 初始化表配置信息
func CustomTableConfig(dbHandler sqlbuilder.Handler, configFn func(table TableConfig) (configedTable TableConfig)) TableConfig {
	var tableConfig = TableConfig{
		Component: table_component,
		Slot:      table_slotName,
		Attribute: table_attribute,
	}

	tableConfig.Component = tableConfig.Component.WithHandler(dbHandler)
	tableConfig.Slot = tableConfig.Slot.WithHandler(dbHandler)
	tableConfig.Attribute = tableConfig.Attribute.WithHandler(dbHandler)
	if configFn != nil {
		tableConfig = configFn(tableConfig)
	}
	return tableConfig
}
