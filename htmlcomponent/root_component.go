package htmlcomponent

import (
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

type RootComponent struct {
	Name       string
	Assembles  Assembles
	Components Components
	Attributes Attributes
}

func NewRootComponent(name string, assembles Assembles, components Components, attributes Attributes) *RootComponent {
	return &RootComponent{
		Name:       name,
		Assembles:  assembles,
		Components: components,
		Attributes: attributes,
	}
}

func (p RootComponent) ToHtml(data map[string]any) (rootComponentHtml string, err error) {
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
	rootComponentHtml = cast.ToString(val)
	return rootComponentHtml, nil
}
