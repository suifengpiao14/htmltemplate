package htmlcomponent

import (
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

type Component struct {
	Name       string
	Nodes      ComponentNodes
	Templates  ComponentTemplates
	Attributes Attributes
}

func NewComponentTree(name string, assembles ComponentNodes, components ComponentTemplates, attributes Attributes) *Component {
	return &Component{
		Name:       name,
		Nodes:      assembles,
		Templates:  components,
		Attributes: attributes,
	}
}

func (p Component) Render(data map[string]any) (rootComponentHtml string, err error) {
	componentName := p.Name
	nodes := p.Nodes
	templates := p.Templates

	attrs := p.Attributes
	data = MergeMap(attrs.MapData(), data)
	variables, err := nodes.RenderTemplate(templates, data)
	if err != nil {
		return "", err
	}

	rootAssembles := nodes.GetByComponentName(componentName)
	first, err := rootAssembles.First()
	if err != nil {
		err = errors.WithMessagef(err, "componentName(same as componentName):%s", componentName)
		return "", err
	}
	val := variables[first.GetOutputKey()]
	rootComponentHtml = cast.ToString(val)
	return rootComponentHtml, nil
}
