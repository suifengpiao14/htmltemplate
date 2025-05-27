package htmlcomponent_test

import (
	"testing"

	"github.com/suifengpiao14/htmltemplate/htmlcomponent"
)

func TestComponentTemplate(t *testing.T) {
	component := htmlcomponent.Template{
		Name: "xyxzapi/index",
		Template: `
		<div class="flex flex-col w-full ">
		背景色-黄色
		<suifengpiao14:tab config="orderTypeTab" />
		</div>
		`,
		DataTpl: `
	{
		"orderTypeTab":{
				"namespace":"orderTypeTab",
				"eventName":"",
				"activeTabId":"arrive_shop",
				"items":[
					{"tabId":"arrive_shop","tabTitle":"到店订单","tabContent":"{{{arriveShopContentOutput}}}"},
					{"tabId":"arrive_home","tabTitle":"上门订单","tabContent":"{{{arriveHomeContentOutput}}}"}
				]
	}
	}
		`,
	}
	_ = component
}
