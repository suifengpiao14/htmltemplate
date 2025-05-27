package htmlcomponent

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/suifengpiao14/htmltemplate/htmlenhance"
	"github.com/suifengpiao14/htmltemplate/xmldata"
)

type Template struct {
	TemplateName     string `json:"templateName"`      //模板唯一标识
	Template         string `json:"template"`          //HTML 模板内容
	DataTpl          string `json:"dataTpl"`           // 需要的数据结构描述
	DataExample      string `json:"dataExample"`       // 示例数据，用于调试
	AttributeDefault string `json:"attributeDefaults"` // 默认属性值，json 格式

}

func (c Template) Render(data map[string]any) (html string, err error) {
	newData, err := c.DecodeData(data)
	if err != nil {
		return "", err
	}
	html, err = htmlenhance.RenderHtmlTpl(c.Template, newData)
	if err != nil {
		return "", err
	}
	html, err = htmlenhance.MergeClassAttrs(html)
	if err != nil {
		return "", err
	}
	return html, nil
}

func (c Template) DecodeData(data map[string]any) (newData map[string]any, err error) {
	newData, err = xmldata.DecodeTplData([]byte(c.DataTpl), data)
	if err != nil {
		return nil, errors.Wrap(err, "Component.DecodeData")
	}
	newData = MergeMap(data, newData)
	return newData, nil
}

type Templates []Template

func (cs Templates) GetByName(name string) (c *Template, ok bool) {
	for _, c := range cs {
		if strings.EqualFold(c.TemplateName, name) {
			return &c, true
		}
	}
	return nil, false
}
