package htmlcomponent

import (
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

type ComponentTree struct {
	Name           string
	ComponentNodes ComponentNodes
	Components     Components
	Attributes     Attributes
}

func NewComponentTree(name string, assembles ComponentNodes, components Components, attributes Attributes) *ComponentTree {
	return &ComponentTree{
		Name:           name,
		ComponentNodes: assembles,
		Components:     components,
		Attributes:     attributes,
	}
}

func (p ComponentTree) ToHtml(data map[string]any) (rootComponentHtml string, err error) {
	rootComponentName := p.Name
	assembles := p.ComponentNodes
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
