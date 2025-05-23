package htmlcomponent

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/suifengpiao14/htmltemplate/htmlenhance"
	"github.com/suifengpiao14/htmltemplate/xmldata"
)

type Component struct {
	ComponentName string `json:"componentName"`
	Template      string `json:"template"`
	DataTpl       string `json:"dataTpl"`
	DataExample   string `json:"dataExample"` // 示例数据，用于调试
}

func (c Component) Render(data map[string]any) (html string, err error) {
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

func (c Component) DecodeData(data map[string]any) (newData map[string]any, err error) {
	newData, err = xmldata.DecodeTplData([]byte(c.DataTpl), data)
	if err != nil {
		return nil, errors.Wrap(err, "Component.DecodeData")
	}
	newData = MergeMap(data, newData)
	return newData, nil
}

type Components []Component

func (cs Components) GetByName(name string) (c *Component, ok bool) {
	for _, c := range cs {
		if strings.EqualFold(c.ComponentName, name) {
			return &c, true
		}
	}
	return nil, false
}
