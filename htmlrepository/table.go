package htmlrepository

import "github.com/suifengpiao14/sqlbuilder"

type TableConfig struct {
	Template  sqlbuilder.TableConfig
	Slot      sqlbuilder.TableConfig
	Attribute sqlbuilder.TableConfig
}

var table_template = sqlbuilder.NewTableConfig("t_template").AddColumns(
	sqlbuilder.NewColumnConfig("Ftemplate_name", sqlbuilder.GetFieldName(NewTemplateNameField)),
	sqlbuilder.NewColumnConfig("Ftemplate", sqlbuilder.GetFieldName(NewTemplateField)),
	sqlbuilder.NewColumnConfig("Fdata_tpl", sqlbuilder.GetFieldName(NewDataTplField)),
).AddIndexs(sqlbuilder.Index{
	Unique: true,
	ColumnNames: func(tableColumns sqlbuilder.ColumnConfigs) (columnNames []string) {
		columnNames = tableColumns.FieldName2ColumnName(
			sqlbuilder.GetFieldName(NewTemplateNameField),
		)
		return columnNames
	},
})

var table_slotName = sqlbuilder.NewTableConfig("t_slot").AddColumns(
	sqlbuilder.NewColumnConfig("Fcomponent_name", sqlbuilder.GetFieldName(NewComponentNameField)),
	sqlbuilder.NewColumnConfig("Ftemplate_name", sqlbuilder.GetFieldName(NewTemplateNameField)),
	sqlbuilder.NewColumnConfig("Fslot_name", sqlbuilder.GetFieldName(NewSlotNameField)),
	sqlbuilder.NewColumnConfig("Fdata_tpl", sqlbuilder.GetFieldName(NewDataTplField)),
).AddIndexs(sqlbuilder.Index{
	Unique: true,
	ColumnNames: func(tableColumns sqlbuilder.ColumnConfigs) (columnNames []string) {
		columnNames = tableColumns.FieldName2ColumnName(
			sqlbuilder.GetFieldName(NewComponentNameField),
			sqlbuilder.GetFieldName(NewSlotNameField),
		)
		return columnNames
	},
})

var table_attribute = sqlbuilder.NewTableConfig("t_attribute").AddColumns(
	sqlbuilder.NewColumnConfig("Fslot_name", sqlbuilder.GetFieldName(NewSlotNameField)),
	sqlbuilder.NewColumnConfig("Ftemplate_name", sqlbuilder.GetFieldName(NewTemplateNameField)),
	sqlbuilder.NewColumnConfig("Ftag_id", sqlbuilder.GetFieldName(NewTagIdField)),
	sqlbuilder.NewColumnConfig("Fattr_name", sqlbuilder.GetFieldName(NewAttributeNameField)),
	sqlbuilder.NewColumnConfig("Fattr_value", sqlbuilder.GetFieldName(NewAttributeValueField)),
).AddIndexs(sqlbuilder.Index{
	Unique: true,
	ColumnNames: func(tableColumns sqlbuilder.ColumnConfigs) (columnNames []string) {
		columnNames = tableColumns.FieldName2ColumnName(
			sqlbuilder.GetFieldName(NewSlotNameField),
			sqlbuilder.GetFieldName(NewTemplateNameField),
			sqlbuilder.GetFieldName(NewTagIdField),
			sqlbuilder.GetFieldName(NewAttributeNameField),
		)
		return columnNames
	},
})

// CustomTableConfig 初始化表配置信息
func CustomTableConfig(dbHandler sqlbuilder.Handler, configFn func(table TableConfig) (configedTable TableConfig)) TableConfig {
	var tableConfig = TableConfig{
		Template:  table_template,
		Slot:      table_slotName,
		Attribute: table_attribute,
	}

	tableConfig.Template = tableConfig.Template.WithHandler(dbHandler)
	tableConfig.Slot = tableConfig.Slot.WithHandler(dbHandler)
	tableConfig.Attribute = tableConfig.Attribute.WithHandler(dbHandler)
	if configFn != nil {
		tableConfig = configFn(tableConfig)
	}
	return tableConfig
}
