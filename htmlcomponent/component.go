package htmlcomponent

import (
	"github.com/spf13/cast"
)

type Component struct {
	Name       string
	Nodes      Slots
	Templates  ComponentTemplates
	Attributes Attributes
}

func NewComponentTree(name string, slotNames Slots, components ComponentTemplates, attributes Attributes) *Component {
	return &Component{
		Name:       name,
		Nodes:      slotNames,
		Templates:  components,
		Attributes: attributes,
	}
}

func (p Component) Render(data map[string]any) (rootComponentHtml string, err error) {
	componentName := p.Name
	nodes := p.Nodes
	templates := p.Templates
	variables, err := nodes.RenderTemplate(templates, p.Attributes, data)
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
