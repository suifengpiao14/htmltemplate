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

func NewComponent(name string, slots Slots, components Templates, attributes Attributes) *Component {
	return &Component{
		Name:       name,
		Slots:      slots,
		Templates:  components,
		Attributes: attributes,
	}
}

func (p Component) Render(data map[string]any) (rootComponentHtml string, err error) {
	componentName := p.Name
	slots := p.Slots
	templates := p.Templates
	variables, err := slots.RenderTemplate(templates, p.Attributes, data)
	if err != nil {
		return "", err
	}

	rootNode, err := slots.GetRootNode(componentName)
	if err != nil {

		return "", err
	}
	val := variables[rootNode.GetOutputKey()]
	rootComponentHtml = cast.ToString(val)
	return rootComponentHtml, nil
}
