package htmlcomponent

import (
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

	rootNode, err := nodes.GetRootNode(componentName)
	if err != nil {

		return "", err
	}
	val := variables[rootNode.GetOutputKey()]
	rootComponentHtml = cast.ToString(val)
	return rootComponentHtml, nil
}
