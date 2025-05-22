package htmlcomponent

import (
	"encoding/json"

	"github.com/spf13/cast"
	"github.com/suifengpiao14/funcs"
)

var XyxzapiPageIndex = Component{
	ComponentName: "xyxzapi/index",
	Template: `
		<div class="flex flex-col w-full ">
		背景色-黄色
		{{{orderTypeTab}}}
		</div>
`,
	DataTpl: `
	<orderTypeTab>{{{orderTypeTab}}}</orderTypeTab>`,
}

var XyxzOrderToolbar = Component{
	ComponentName: "xyxzapi/orderToolbar",
	Template: `
<div class="flex flex-row w-full ">
<div>手动建单</div>
<div>公益回收</div>
<div>订单搜索</div>
<div>收款</div>
</div>
`,
	DataTpl: ``,
}

var Container = Component{
	ComponentName: "suifengpiao14/container",
	Template: `
	<div>
	{{#children}}
	{{{.}}}
	{{/children}}
	</div>
	`,
	DataTpl: `
	{{#children}}
		<children>
			{{{.}}}
		</children>
	{{/children}}
	`,
}

var XyxzOrderListItem = Component{
	ComponentName: "xyxzapi/orderListItem",
	Template: `
	{{#orders}}
	<div data-fullname="{{namespace}}/orders-item">
<div data-fullname="{{namespace}}/orders-item/orders-part1" data-node-key="11111111">
<div data-fullname="{{namespace}}/orders-item/orderType">{{orderType}}</div>
<div data-fullname="{{namespace}}/orders-item/orderId">{{orderId}}</div>
<div data-fullname="{{namespace}}/orders-item/orderTime">{{orderTime}}</div>
</div>

<div data-fullname="{{namespace}}/orders-item/orders-part2">
<img data-fullname="{{namespace}}/orders-item/productPictrue" src="{{productPictrue}}"/>
<div data-fullname="{{namespace}}/orders-item/className">{{className}}   <span data-fullname="{{namespace}}/orders-item/orderStatus">{{orderStatus}}</span></div>
<div data-fullname="{{namespace}}/orders-item/price">{{price}}</div>
<div data-fullname="{{namespace}}/orders-item/operatorName">经办人: {{operatorName}}</div>

<div data-fullname="{{namespace}}/orders-item/userPhone">{{userPhone}}</div>
</div>
</div>
{{/orders}}
`,
	DataTpl: `
	{{#orders}}
		<orders>
			<orderType>{{orderType}}</orderType>
			<orderId>{{orderId}}</orderId>
			<orderTime>{{orderTime}}</orderTime>
			<className>{{className}}</className>
			<productPictrue>{{productPictrue}}</productPictrue>
			<orderStatus>{{orderStatus}}</orderStatus>
			<price>{{price}}</price>
			<operatorName>{{operatorName}}</operatorName>
			<userPhone>{{userPhone}}</userPhone>
		</orders>
	{{/orders}}
`,
	DataExample: `
		{
			"orders": [
				{
					"orderType": "黄金到店",
					"orderId": 16690403,
					"orderTime": "2024-12-26 17:00:07",
					"className": "金条",
					"productPictrue": "/abc.png",
					"orderStatus": "待跟进",
					"price": "Y361121.19",
					"operatorName": "超级管理员",
					"userPhone": "159***2341"
				}
			]
		}
`,
}

var TabComponent = Component{
	ComponentName: "suifengpiao14/tab",
	Template: `
	 	<div data-fullname="{{namespace}}/tab" role="tablist" class="tabs tabs-bordered "
		x-data='$tab({"tab_eventName":"{{eventName}}","tab_activeTabId":"{{activeTabId}}"})'>
		<div data-fullname="{{namespace}}/tab/title" class="w-full border-t-2">
			{{#items}}
			<a data-fullname="{{namespace}}/tab/title/item" tab-for="{{tabId}}" x-bind="bind_tab" role="tab"
				class="tab px-1">{{tabTitle}}</a>
			{{/items}}
		</div>
		<div data-fullname="{{namespace}}/tab/content" class="w-full ">
			{{#items}}
			<div data-fullname="{{namespace}}/tab/content/item" x-cloak id="{{tabId}}" x-bind="bind_tabpanel"
				role="tabpanel" class="tab-content bg-base-100 border-base-300  w-full">
				{{{tabContent}}}
			</div>
			{{/items}}
		</div>
	</div>
	 `,
	DataTpl: `
	<namespace>{{namespace}}</namespace>
	<eventName>{{eventName}}</eventName>
	<activeTabId>{{activeTabId}}</activeTabId>
	{{#items}}
		<items>
			<tabId>{{tabId}}</tabId>
			<tabTitle>{{tabTitle}}</tabTitle>
			<tabContent>{{{tabContent}}}</tabContent>
		</items>
	{{/items}}
	`,
}

var TableComponent = Component{
	ComponentName: "suifengpiao14/table",
	Template: `
	<table class="min-w-full table-auto bg-white shadow-md rounded-lg overflow-hidden">
            <thead class="bg-indigo-600 text-white">
                <tr>
                    {{#headers}}
                    <th class="px-6 py-3 text-left text-sm font-medium" data-column="{{column}}" {{{attrs}}}>{{{title}}}</th>
                    {{/headers}}

                </tr>
            </thead>
            <tbody>
                {{#rows}}
                <tr class="border-b hover:bg-gray-50">
                    {{#columns}}
                    <td class="px-6 py-4 text-sm" data-column="{{column}}" {{{attrs}}}>{{{value}}}</td>
                    {{/columns}}
                </tr>
                {{/rows}}
            </tbody>
        </table>`,
	DataTpl: `
	{{#headers}}
	<headers>
		<column>{{{column}}}</column>
		<title>{{title}}</title>
	</headers>
	{{/headers}}

	{{#rows}}
	<rows>
	{{#columns}}
	<columns>
		<column>{{{column}}}</column>
		<value>{{{value}}}</value>
		<attrs>{{{attrs}}}</attrs>
	</columns>
	{{/columns}}
	</rows>
	{{/rows}}
`,
	DataExample: `
{
	"headers": [
		{"column":"id","title":"ID"},
		{"column":"title","title":"名称"}
	],
	"rows":[
{
	"columns":[{"column":"id","value":"1"},{"column":"value","value":"张三","attrs":"class=\"text-red\""}]
}

	]
}
`,
}

var SearchFormComponent = Component{
	ComponentName: "suifengpiao14/searchForm",
	Template: `
 <form hx-post="{{hxpost}}" hx-target="{{hxtarget}}" hx-ext="json-enc-custom"
        class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6 p-6 bg-white rounded-lg shadow-md">
		{{#hiddenFields}}
        <input type="hidden" name="{{name}}" value="{{value}}"/>
		{{/hiddenFields}}
		{{#selectFields}}
		<div class="flex flex-col gap-2">
			<label for="{{name}}" class="inline text-sm font-medium text-gray-700">{{{title}}}</label>
			<select id="{{name}}" name="{{name}}" class="inline  mt-1 px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:outline
			focus:ring-2 focus:ring-indigo-500">
			{{#options}}
			<option value="{{value}}" {{#selected}} selected="selected" {{/selected}} >{{title}}</option>
			{{/options}}
		</select>
		</div>
		{{/selectFields}}

        {{#inputFields}}
        <div>
            <label for="{{name}}" class="inline text-sm font-medium text-gray-700">{{{title}}}</label>
            <input type="text" id="{{name}}" name="{{name}}" value="{{value}}"
                class="inline  mt-1 px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-indigo-500"
                placeholder="请输入{{title}}"/>
        </div>
		{{/inputFields}}

        <div class="flex justify-between gap-4">
            <!-- 提交按钮 -->
            <button type="submit"
                class="sm:w-auto px-6 py-2 bg-indigo-600 text-white rounded-md hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500">
                搜索
            </button>
            <!-- 重置按钮 -->
            <button type="reset"
                class="sm:w-auto px-6 py-2 bg-gray-200 text-gray-700 rounded-md hover:bg-gray-300 focus:outline-none focus:ring-2 focus:ring-gray-500">
                重置
            </button>
        </div>
    </form>
	`,
	DataTpl: `
	<hxpost>{{{hxpost}}}</hxpost>
	<hxtarget>{{{hxtarget}}}</hxtarget>
        {{#hiddenFields}}
	<hiddenFields>
		<name>{{name}}</name>
		<value>{{value}}</value>
	</hiddenFields>
        {{/hiddenFields}}

{{#inputFields}}
	<inputFields>
		<name>{{name}}</name>
		<title>{{title}}</title>
		<value>{{value}}</value>
	</inputFields>
{{/inputFields}}

{{#selectFields}}
	<selectFields>
		<name>{{name}}</name>
		<title>{{title}}</title>
{{#options}}
		<options>
			<value>{{value}}</value>
			<title>{{title}}</title>
			<selected>{{selected}}</selected>
		</options>
{{/options}}
	</selectFields>
{{/selectFields}}
`,
}

var HtmlDocumentComponent = Component{
	ComponentName: "suifengpiao14/htmlDocument",
	Template: `
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
    {{#metaList}}
	{{{.}}}
	{{/metaList}}
    <title>{{title}}</title>
	{{#cssList}}
	<link href="{{.}}" rel="stylesheet" type="text/css" />
	{{/cssList}}
	{{#scriptList}}
	<script src="{{src}}" {{#defer}} defer {{/defer}}></script>
	{{/scriptList}}
	{{#inlineStyle}}
	<style type="text/css">
	{{{inlineStyle}}}
	</style>
	{{/inlineStyle}}
	{{{inlineScript}}}
</head>
<body>
{{{content}}}
</body>
</html>
	`,
	DataTpl: `
	{{#metaList}}
	<metaList>{{{.}}}</metaList>
	{{/metaList}}
	<title>{{{title}}}</title>
	{{#cssList}}
	<cssList>{{.}}</cssList>
	{{/cssList}}
	{{#scriptList}}
	<scriptList>
		<src>{{src}}</src>
		<defer>{{defer}}</defer>
	</scriptList>
	{{/scriptList}}
	<inlineStyle>{{{inlineStyle}}}</inlineStyle>
	<inlineScript>{{{inlineScript}}}</inlineScript>
	<content>{{{content}}}</content>
	`,
	DataExample: `
	{
		"title":"html文档标题",
		"metaList":[
		"<meta name=\"htmx-config\" content='{\"selfRequestsOnly\":false}'>"
		],
		"cssList":[
"/static/cdn.jsdelivr.net/npm/daisyui@2.14.0/dist/full.css"
		],
		"scriptList":[
{"src":"/static/cdn.tailwindcss.com"},
{"src":"/static/unpkg.com/htmx.org@2.0.3/dist/htmx.js"},
{"src":"/static/cdn.jsdelivr.net/gh/Emtyloc/json-enc-custom@main/json-enc-custom.js"},
{"src":"/static/cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js"},
{"src":"/static/alpine.component.all.js"}
		],
		"inlineStyle":"[x-cloak] {display: none !important;}",
		"content":"hello world"
	}
	`,
}

var HtmxAlpinejsPageComponent = Component{}

var AllComponent = Components{
	TabComponent,
	XyxzapiPageIndex,
	XyxzOrderListItem,
	XyxzOrderToolbar,
	Container,
	TableComponent,
	SearchFormComponent,
	HtmlDocumentComponent,
}

type TableData struct {
	Headers []TableHeader `json:"headers"`
	Rows    []TableRow    `json:"rows"`
}

type TableRow struct {
	Columns []TableCell `json:"columns"` // xml 不支持[[{},{}]] 格式三维数组，所以改成{"rows":[{"columns":[{},{}]}]}
}

func (td TableData) ToMap() map[string]any {
	m := funcs.Struct2JsonMap(td)
	return m
}

type TableHeader struct {
	Column string `json:"column"`
	Title  string `json:"title"`
	Attrs  string `json:"attrs"`
}

type TableHeaders []TableHeader

type TableCell struct {
	Column string `json:"column"`
	Value  string `json:"value"`
	Attrs  string `json:"attrs"`
}

func Rows2TableData[S ~[]T, T any](tableHeaders TableHeaders, rows S) (tableData TableData) {
	tableData = TableData{
		Headers: tableHeaders,
		Rows:    make([]TableRow, 0),
	}
	if len(rows) == 0 {
		return
	}
	b, err := json.Marshal(rows)
	if err != nil {
		panic(err)
	}
	rowsMap := make([]map[string]any, 0)
	err = json.Unmarshal(b, &rowsMap)
	if err != nil {
		panic(err)
	}

	for _, rowMap := range rowsMap {
		tableRow := TableRow{
			Columns: make([]TableCell, 0),
		}
		for _, tableHeader := range tableHeaders {
			tableCell := TableCell{
				Column: tableHeader.Column,
				Value:  cast.ToString(rowMap[tableHeader.Column]),
			}
			tableRow.Columns = append(tableRow.Columns, tableCell)
		}
		tableData.Rows = append(tableData.Rows, tableRow)
	}

	return tableData
}
