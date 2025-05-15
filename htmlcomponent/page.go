package htmlcomponent

import (
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

type HtmlPage struct {
	Name       string
	Assembles  Assembles
	Components Components
	Attributes Attributes
}

func NewHtmlPage(name string, assembles Assembles, components Components, attributes Attributes) *HtmlPage {
	return &HtmlPage{
		Name:       name,
		Assembles:  assembles,
		Components: components,
		Attributes: attributes,
	}
}

func (p HtmlPage) ToHtml(data map[string]any) (pageHtml string, err error) {
	rootComponentName := p.Name
	assembles := p.Assembles
	components := p.Components

	attrs := p.Attributes
	data = MergeMap(attrs.MapData(), data)
	variables, err := assembles.RenderComponent(components, data)
	if err != nil {
		return "", err
	}

	rootAssembles := assembles.GetByComponentName(rootComponentName)
	first, err := rootAssembles.First()
	if err != nil {
		err = errors.WithMessagef(err, "componentName(same as rootComponentName):%s", rootComponentName)
		return "", err
	}
	val := variables[first.GetOutputKey()]
	pageHtml = cast.ToString(val)
	return pageHtml, nil
}
