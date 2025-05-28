package htmlcomponent

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/suifengpiao14/htmltemplate/htmlenhance"
	"github.com/suifengpiao14/memorytable"
)

/*
动态设置元素属性，主要用于动态设置class属性

*/

type Attribute struct { // Attribute 属于template的固定数据，不一定属于某个Component，所以不能添加componentName属性
	SlotName       string `json:"slotName"`     // component node id(同一个template 在一个组件中可能用于多次,比如按钮-确定/取消按钮,所以需要携带NodeId区分不同节点)
	TemplateName   string `json:"templateName"` // 模板名称，用于区分不同组件的相同节点
	TagId          string `json:"tagId"`        // html标签id
	AttributeName  string `json:"key"`
	AttributeValue string `json:"value"`
	sort           int
}

func (a Attribute) String() string {
	b, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	return string(b)

}

type Attributes []Attribute

func (as Attributes) MapData() (attrMap map[string]any) {
	attrMap = make(map[string]any)
	for _, a := range as {
		attrMap[htmlenhance.AttrPlaceholderName(a.TagId)] = as.GetByTagID(a.TagId).String()
	}
	return attrMap
}

func (as Attributes) GetByTagID(TagId string) Attributes {
	attrs := memorytable.NewTable(as...).Where(func(record Attribute) bool {
		return record.TagId == TagId
	})
	return attrs.ToSlice()
}

func (as Attributes) GetByNodeID(nodeId string) Attributes {
	attrs := memorytable.NewTable(as...).Where(func(record Attribute) bool {
		return record.SlotName == nodeId
	})
	return attrs.ToSlice()
}

func (a *Attributes) Remove(names ...string) {
	arr := Attributes{}
	for _, attr := range *a {
		ignore := false
		for _, name := range names {
			if strings.EqualFold(name, attr.AttributeName) {
				ignore = true
				break
			}
		}
		if ignore {
			continue
		}
		arr = append(arr, attr)
	}
	*a = arr
}

func (a Attributes) Sort() Attributes {
	a.initSort()
	sort.Slice(a, func(i, j int) bool { return a[i].sort < a[j].sort })
	return a
}

var Attr_sort = []string{"_id", "name", "class", "value", "type", "data-nid"}

func (a *Attributes) initSort() {
	for i, v := range *a {
		for j, sortKey := range Attr_sort {
			if v.AttributeName == sortKey {
				(*a)[i].sort = j
			}
		}

	}
}

func (as Attributes) GetByAttributeName(AttributeName string) Attribute {
	attr := memorytable.NewTable(as...).GetOneWithDefault(func(row Attribute) bool {
		return row.AttributeName == AttributeName
	})
	for _, attr := range as {
		if attr.AttributeName == AttributeName {
			return attr
		}
	}
	attr.AttributeName = AttributeName
	return attr
}
func (as *Attributes) ResetByKey(newAttr Attribute) {
	if *as == nil {
		*as = Attributes{
			newAttr,
		}
		return
	}
	for i := range *as {
		if (*as)[i].AttributeName == newAttr.AttributeName {
			(*as)[i] = newAttr
		}
	}
}

func (a *Attributes) Append(attrs ...Attribute) {
	tmp := memorytable.NewTable(*a...).Set(func(t Attribute) (identity string) { return t.TagId }, attrs...)
	*a = Attributes(tmp)
}

func (a Attributes) String() string {
	attrs := make([]string, 0)
	for _, attr := range a {
		attrs = append(attrs, fmt.Sprintf(`%s="%s"`, attr.AttributeName, attr.AttributeValue))
	}
	out := strings.Join(attrs, " ")
	return out
}

func ParseAttributes(attrString string) (attrs Attributes, err error) {
	div := fmt.Sprintf("<div %s></div>", attrString)
	root, _, err := htmlenhance.ParseHTML(div)
	if err != nil {
		return nil, err
	}
	divNode := htmlquery.FindOne(root, "//div")
	for _, attr := range divNode.Attr {
		attrs = append(attrs, Attribute{
			AttributeName:  attr.Key,
			AttributeValue: attr.Val,
		})
	}
	return attrs, nil
}
