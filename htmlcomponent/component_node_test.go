package htmlcomponent_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suifengpiao14/funcs"
	"github.com/suifengpiao14/htmltemplate/htmlcomponent"
)

var allComponent = htmlcomponent.AllComponent
var xyxzApiIndexAssembles = htmlcomponent.ComponentNodes{

	{
		ParentNodeID: "xyxzapi/index",
		TemplateName: "xyxzapi/orderListItem",
		NodeID:       "arriveShop_OrderItems",
		Props: `{
		"namespace":"arriveShop",
		"orders":[{
		"orderType":"黄金到店",
		"orderId":"16690403",
		"orderTime":"2024-12-26 17:00:07",
		"className":"金条",
		"productPictrue":"/abc.png",
		"orderStatus":"待跟进",
		"price":"Y361121.19",
		"operatorName":"超级管理员",
		"userPhone":"159***2341"
	},
	{
		"orderType":"黄金到店",
		"orderId":"16690407",
		"orderTime":"2024-05-18 17:00:07",
		"className":"金条",
		"productPictrue":"/abc.png",
		"orderStatus":"已下单",
		"price":"Y361121.19",
		"operatorName":"超级管理员",
		"userPhone":"159***2333"
	}
		]}`, // 这里用的是上面定义的orderTypeTab的变量
	},
	{
		ParentNodeID: "xyxzapi/index",
		TemplateName: "xyxzapi/orderListItem",
		NodeID:       "arriveHome_orderItems",
		Props: `{
		"namespace":"arriveHome",
		"orders":[{
		"orderType":"黄金上门",
		"orderId":"16690404",
		"orderTime":"2024-12-27 17:00:07",
		"className":"金条",
		"productPictrue":"/abc.png",
		"orderStatus":"待跟进",
		"price":"Y361121.19",
		"operatorName":"超级管理员",
		"userPhone":"159***2341"
	},
	{
		"orderType":"黄金上门",
		"orderId":"26690404",
		"orderTime":"2025-01-01 17:00:07",
		"className":"金条",
		"productPictrue":"/abc.png",
		"orderStatus":"已上门",
		"price":"Y361.00",
		"operatorName":"超级管理员",
		"userPhone":"159***2320"
	}
		]}`, // 这里用的是上面定义的orderTypeTab的变量
	},
	{
		ParentNodeID: "xyxzapi/index",
		TemplateName: "suifengpiao14/container",
		NodeID:       "arriveShopContent",
		Props: `{
		"children":["{{{toolbarOutput}}}","{{{arriveShop_OrderItemsOutput}}}"]
		}`,
	},
	{
		ParentNodeID: "xyxzapi/index",
		TemplateName: "suifengpiao14/container",
		NodeID:       "arriveHomeContent",
		Props: `{
		"children":["{{{toolbarOutput}}}","{{{arriveHome_orderItemsOutput}}}"]
		}`,
	},
	{
		ParentNodeID: "xyxzapi/index",
		TemplateName: "xyxzapi/orderToolbar",
		NodeID:       "toolbar",
	},
	{
		ParentNodeID: "xyxzapi/index",
		TemplateName: "suifengpiao14/tab",
		NodeID:       "orderTypeTab",
		Props: `{
				"namespace":"orderTypeTab",
				"eventName":"",
				"activeTabId":"arrive_shop",
				"items":[
					{"tabId":"arrive_shop","tabTitle":"到店订单","tabContent":"{{{arriveShopContentOutput}}}"},
					{"tabId":"arrive_home","tabTitle":"上门订单","tabContent":"{{{arriveHomeContentOutput}}}"}
				]
	}`,
	},
	{
		ParentNodeID: "xyxzapi/index",
		TemplateName: "xyxzapi/index",
		NodeID:       "index",
		Props: `{
	"orderTypeTab":"{{{orderTypeTabOutput}}}"
	}`,
	},
}

func TestPage(t *testing.T) {
	rootComponentName := "xyxzapi/index"
	as := xyxzApiIndexAssembles.FilterByRootComponentName(rootComponentName)
	allData := map[string]any{
		"arriveHome_orderItemsInput": map[string]any{
			"namespace": "arriveHome_from_data",
			"orders": []map[string]string{
				{
					"orderType":      "黄金上门",
					"orderId":        "25090404",
					"orderTime":      "2025-01-13 12:00:07",
					"className":      "金条",
					"productPictrue": "/abc.png",
					"orderStatus":    "待跟进",
					"price":          "Y3611.19",
					"operatorName":   "超级管理员",
					"userPhone":      "159***0000",
				},
				{
					"orderType":      "黄金上门",
					"orderId":        "37990154",
					"orderTime":      "2025-11-10 17:00:07",
					"className":      "金条",
					"productPictrue": "/abc.png",
					"orderStatus":    "待上门",
					"price":          "Y361.00",
					"operatorName":   "超级管理员",
					"userPhone":      "159***5274",
				},
			},
		},
	}
	variables, err := as.RenderComponent(allComponent, allData)
	require.NoError(t, err)
	indexHtml := variables["indexOutput"]
	fmt.Println(indexHtml)

}

func TestGetDependence(t *testing.T) {
	a := xyxzApiIndexAssembles[2]
	dependences := a.GetDependence()
	fmt.Println(dependences)
}

func TestRanderTable(t *testing.T) {
	rootComponentName := "html/component"
	var TestHtmlComponentIndexAssembles = htmlcomponent.ComponentNodes{
		{
			ParentNodeID: "html/component",
			TemplateName: "suifengpiao14/table",
			NodeID:       "table",
		},
		{
			ParentNodeID: "test/htmlComponent",
			TemplateName: "suifengpiao14/tab",
		},
	}

	as := TestHtmlComponentIndexAssembles.FilterByRootComponentName(rootComponentName)
	data := rows2TableData()
	rowsMap := funcs.Struct2JsonMap(data)
	allData := map[string]any{
		"tableInput": rowsMap,
	}
	variables, err := as.RenderComponent(allComponent, allData)
	require.NoError(t, err)
	indexHtml := variables["tableOutput"]
	fmt.Println(indexHtml)
}

type userInfo struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Birthday string `json:"birthday"`
}

func TestRanderSubPage(t *testing.T) {

	var TestHtmlComponentIndexAssembles = htmlcomponent.ComponentNodes{
		{
			ParentNodeID: "html/component",
			TemplateName: "suifengpiao14/table",
			NodeID:       "table",
			Props: `
			<headers>
			<column>name</column>
			<title>姓名</title>
			</headers>
			<headers>
			<column>age</column>
			<title>年龄</title>
			</headers>
			<headers>
			<column>birthday</column>
			<title>生日</title>
			</headers>
			<headers>
			<column>action</column>
			<title>操作</title>
			</headers>
			{{#items}}
			<rows>
				<columns>
				<column>name</column>
				<value>{{name}}</value>
				</columns>
				<columns>
				<column>age</column>
				<value>{{age}}</value>
				<attrs>class="text-red"</attrs>
				</columns>
				<columns>
				<column>birthday</column>
				<value>{{birthday}}</value>
				</columns>
				<columns>
				<column>action</column>
				<value><![CDATA[<button>编辑</button><button>删除</button>]]></value>
				</columns>
			</rows>
			{{/items}}
			`,
		},
		{
			ParentNodeID: "test/htmlComponent",
			TemplateName: "suifengpiao14/tab",
		},
	}

	rootComponentName := "html/component"
	as := TestHtmlComponentIndexAssembles.FilterByRootComponentName(rootComponentName)
	tableDataMap := map[string]any{
		"items": []map[string]any{
			{
				"name":     "张三",
				"age":      20,
				"birthday": "1998-01-01",
			},
			{
				"name":     "李四",
				"age":      30,
				"birthday": "1990-01-01",
			},
		},
	}

	allData := map[string]any{
		"tableInput": tableDataMap,
	}
	variables, err := as.RenderComponent(allComponent, allData)
	require.NoError(t, err)
	indexHtml := variables["tableOutput"]
	fmt.Println(indexHtml)
}

func rows2TableData() htmlcomponent.TableData {
	tableHeaders := htmlcomponent.TableHeaders{
		{
			Title:  "姓名",
			Column: "name",
		},
		{
			Title:  "年龄",
			Column: "age",
		},
		{
			Title:  "生日",
			Column: "birthday",
		},
		{
			Title:  "操作",
			Column: "action",
		},
	}

	rows := []userInfo{
		{
			Name:     "张三",
			Age:      20,
			Birthday: "1998-01-01",
		},
		{
			Name:     "李四",
			Age:      30,
			Birthday: "1990-01-01",
		},
	}
	tableData := htmlcomponent.Rows2TableData(tableHeaders, rows)
	return tableData
}
