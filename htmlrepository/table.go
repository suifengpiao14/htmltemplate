package htmlrepository

import "github.com/suifengpiao14/sqlbuilder"

type TableConfig struct {
	Template  sqlbuilder.TableConfig
	Slot      sqlbuilder.TableConfig
	Attribute sqlbuilder.TableConfig
}

var table_template = sqlbuilder.NewTableConfig("t_template").AddColumns(
	sqlbuilder.NewColumn("Ftemplate_name", sqlbuilder.GetField(NewTemplateNameField)),
	sqlbuilder.NewColumn("Ftemplate", sqlbuilder.GetField(NewTemplateField)),
	sqlbuilder.NewColumn("Fdata_tpl", sqlbuilder.GetField(NewDataTplField)),
).AddIndexs(sqlbuilder.Index{
	Unique: true,
	ColumnNames: func(table sqlbuilder.TableConfig) (columnNames []string) {
		columnNames = table.Columns.FieldName2ColumnName(
			sqlbuilder.GetFieldName(NewTemplateNameField),
		)
		return columnNames
	},
})

var table_slotName = sqlbuilder.NewTableConfig("t_slot").AddColumns(
	sqlbuilder.NewColumn("Fcomponent_name", sqlbuilder.GetField(NewComponentNameField)),
	sqlbuilder.NewColumn("Ftemplate_name", sqlbuilder.GetField(NewTemplateNameField)),
	sqlbuilder.NewColumn("Fslot_name", sqlbuilder.GetField(NewSlotNameField)),
	sqlbuilder.NewColumn("Fdata_tpl", sqlbuilder.GetField(NewDataTplField)),
).AddIndexs(sqlbuilder.Index{
	Unique: true,
	ColumnNames: func(table sqlbuilder.TableConfig) (columnNames []string) {
		columnNames = table.Columns.FieldName2ColumnName(
			sqlbuilder.GetFieldName(NewComponentNameField),
			sqlbuilder.GetFieldName(NewSlotNameField),
		)
		return columnNames
	},
})

var table_attribute = sqlbuilder.NewTableConfig("t_attribute").AddColumns(
	sqlbuilder.NewColumn("Fslot_name", sqlbuilder.GetField(NewSlotNameField)),
	sqlbuilder.NewColumn("Ftemplate_name", sqlbuilder.GetField(NewTemplateNameField)),
	sqlbuilder.NewColumn("Ftag_id", sqlbuilder.GetField(NewTagIdField)),
	sqlbuilder.NewColumn("Fattr_name", sqlbuilder.GetField(NewAttributeNameField)),
	sqlbuilder.NewColumn("Fattr_value", sqlbuilder.GetField(NewAttributeValueField)),
).AddIndexs(sqlbuilder.Index{
	Unique: true,
	ColumnNames: func(table sqlbuilder.TableConfig) (columnNames []string) {
		columnNames = table.Columns.FieldName2ColumnName(
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
