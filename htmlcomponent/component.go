package htmlcomponent

import (
	"github.com/spf13/cast"
)

type Component struct {
	Name       string
	Slots      Slots
	Templates  Templates
	Attributes Attributes
}

func NewComponent(name string, slotNames Slots, components Templates, attributes Attributes) *Component {
	return &Component{
		Name:       name,
		Slots:      slotNames,
		Templates:  components,
		Attributes: attributes,
	}
}

func (p Component) Render(data map[string]any) (rootComponentHtml string, err error) {
	componentName := p.Name
	nodes := p.Slots
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
