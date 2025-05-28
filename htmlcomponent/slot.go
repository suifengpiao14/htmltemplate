package htmlcomponent

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/suifengpiao14/htmltemplate/xmldata"
	"github.com/suifengpiao14/memorytable"
)

type Slot struct {
	ComponentName string   `json:"componentName"`
	TemplateName  string   `json:"templateName"`
	SlotName      string   `json:"partialName"`
	DataTpl       string   `json:"dataTpl"`    // 支持json和xml模板格式
	Attributes    string   `json:"attributes"` // 属性,json格式
	DataExample   string   `json:"dataExample"`
	dependences   []string //依赖组件(所有的input key)

}

func (r Slot) GetInputKey() string {
	return fmt.Sprintf(`%sInput`, r.SlotName)
}
func (r Slot) GetOutputKey() string {
	return fmt.Sprintf(`%sOutput`, r.SlotName)
}

func replacePlaceholder(s string, data map[string]any) any {
	if strings.HasPrefix(s, "{{") {
		placeholder := strings.Trim(strings.Trim(s, "{"), "}")
		if dataValue, ok := data[placeholder]; ok {
			return dataValue
		}
	}
	return s
}

// ReplacePlaceholder 递归地遍历并替换 JSON 数据中的占位符
func ReplacePlaceholder(jsonPlaceholder any, data map[string]any) (newData any) {
	// 根据 jsonPlaceholder 的类型进行处理
	switch v := jsonPlaceholder.(type) {
	case string:
		// 如果是字符串，替换占位符
		return replacePlaceholder(v, data)
	case map[string]any:
		// 如果是嵌套的 map，递归替换
		newMap := map[string]any{}
		for key, value := range v {
			newMap[key] = ReplacePlaceholder(value, data)
		}
		return newMap
	case []any:
		// 如果是数组，遍历数组中的每个元素
		newArray := make([]any, len(v))
		for i, item := range v {
			newArray[i] = ReplacePlaceholder(item, data)
		}
		return newArray
	}
	return jsonPlaceholder
}

func (r Slot) DecodeData(data map[string]any) (newData map[string]any, err error) {
	newData, err = xmldata.DecodeTplData([]byte(r.DataTpl), data)
	if err != nil {
		return nil, errors.Wrap(err, "Slot.DecodeData")
	}
	newData = MergeMap(data, newData)
	return newData, nil
}

func (r Slot) GetDependence() (dependences []string) {
	dependences = make([]string, 0)
	if r.DataTpl == "" {
		dependences = memorytable.NewTable(dependences...).FilterEmpty()
		return dependences
	}
	regexp := regexp.MustCompile(`\{\{\{?([\w\.\-]+)\}\}\}?`)
	matches := regexp.FindAllStringSubmatch(r.DataTpl, -1)
	for _, match := range matches {
		slotName := strings.TrimSuffix(match[1], "Output")
		slotName = strings.TrimSuffix(slotName, "Input")
		dependences = append(dependences, slotName)
	}
	dependences = memorytable.NewTable(dependences...).FilterEmpty().Uniqueue(func(row string) (key string) {
		return row
	}).ToSlice()
	return dependences
}

// func (r Slot) Render(tpl string, data map[string]any) (key string, value string) {
// 	value = pkg.MustacheMustRender(tpl, data)
// 	return r.GetOutputKey(), value
// }

type Slots []Slot

var Error_not_found = errors.New("not found")

func (as Slots) First() (slotName Slot, err error) {
	if len(as) == 0 {
		return slotName, Error_not_found
	}
	slotName = as[0]
	return slotName, nil

}

func (as Slots) FilterByComponentName(RootComponentName string) (onePageSlots Slots) {
	rows := memorytable.NewTable(as...).Where(func(a Slot) bool {
		return a.ComponentName == RootComponentName
	}).ToSlice()
	return rows
}
func (as Slots) Filter(filterFn func(a Slot) bool) (subSlots Slots) {
	rows := memorytable.NewTable(as...).Where(func(a Slot) bool {
		return filterFn(a)
	}).ToSlice()
	return rows
}
func (as Slots) GetRootNode(componentName string) (node *Slot, err error) {
	node, err = memorytable.NewTable(as...).GetOneWithError(func(a Slot) bool {
		return a.TemplateName == componentName
	})
	if err != nil {
		err = errors.WithMessagef(err, "componentName:%s", componentName)
		return nil, err
	}
	return node, nil

}

func (as Slots) GetBySlotName(slotName string) (slot *Slot, index int) {
	for i, relation := range as {
		if relation.SlotName == slotName {
			return &relation, i
		}
	}
	return nil, -1
}

func (as Slots) TemplateNames() (templateNames []string) {
	templateNames = make([]string, 0)
	for _, a := range as {
		templateNames = append(templateNames, a.TemplateName)
	}
	templateNames = memorytable.NewTable(templateNames...).FilterEmpty().Uniqueue(func(row string) (key string) {
		return row
	})
	return templateNames
}

func (as *Slots) Insert(a Slot, index int) {
	tmp := memorytable.NewTable(*as...).Insert(a, index)
	*as = Slots(tmp)
}

func (as *Slots) InsertBefore(a Slot, index int) {
	as.Insert(a, index-1)
}

// resolveDependence 解析依赖关系，根据组件依赖的变量,以及组件的PlaceHolder,决定渲染顺序
func (slots Slots) resolveDependence() (ordered Slots) {
	ordered = make(Slots, 0)
	maxIndex := len(slots)
	// 构建依赖映射
	for _, a := range slots {
		a.dependences = a.GetDependence()
		_, aIndex := ordered.GetBySlotName(a.SlotName)
		if aIndex < 0 {
			aIndex = maxIndex // 默认增加到最后
		}
		for _, dep := range a.dependences {
			dependence, fullItemsIndex := slots.GetBySlotName(dep)
			if fullItemsIndex < 0 {
				continue
			}
			_, existsIndex := ordered.GetBySlotName(dep)
			if existsIndex < 0 {
				ordered.InsertBefore(*dependence, aIndex)
			}
		}
		if aIndex == maxIndex {
			ordered.Insert(a, aIndex)
		}
	}
	return ordered
}

func (slots Slots) RootSlot() (rootSlot Slot) {
	ordered := slots.resolveDependence()
	if len(ordered) == 0 {
		return Slot{}
	}
	last := ordered[len(ordered)-1]
	return last
}

func (slots Slots) Render(cs Templates, attributes Attributes, data map[string]any) (segments map[string]any, err error) {
	segments = make(map[string]any, 0)
	ordered := slots.resolveDependence()
	for _, node := range ordered {
		componentTpl, ok := cs.GetByName(node.TemplateName)
		if !ok {
			continue
		}
		attrs := attributes.GetByNodeID(node.SlotName)
		data = MergeMap(attrs.MapData(), data)
		subData := data[node.GetInputKey()]
		componentData := make(map[string]any, 0)
		if subData != nil {
			componentData = subData.(map[string]any)
		}
		componentData = MergeMap(data, componentData, segments) // 用指定key值覆盖data中的同名key值，再合并variables中同名key值,属性的数据只能再最早的data[string]any 中定义，所以必须保留最初的 data 的合并

		templateData, err := node.DecodeData(componentData)
		if err != nil {
			return nil, err
		}
		html, err := componentTpl.Render(templateData)
		if err != nil {
			return nil, errors.WithMessagef(err, "render component %s error", componentTpl.TemplateName)
		}
		segments[node.GetOutputKey()] = html
	}
	return segments, nil
}

func MergeMap(maps ...map[string]any) (merged map[string]any) {
	merged = make(map[string]any, 0)
	for _, m := range maps {
		for k, v := range m {
			merged[k] = v
		}
	}
	return merged
}
