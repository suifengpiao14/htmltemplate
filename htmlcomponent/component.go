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

func (p Component) Render(data map[string]any) (componentHtml string, err error) {
	slots := p.Slots
	templates := p.Templates
	variables, err := slots.Render(templates, p.Attributes, data)
	if err != nil {
		return "", err
	}
	rootSlot := slots.RootSlot()
	val := variables[rootSlot.GetOutputKey()]
	componentHtml = cast.ToString(val)
	return componentHtml, nil
}
