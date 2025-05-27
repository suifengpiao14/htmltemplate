package htmltemplate

import (
	"github.com/suifengpiao14/htmltemplate/htmlcomponent"
	"github.com/suifengpiao14/htmltemplate/htmlenhance"
)

var SetNodeIdAndAttrHolder = htmlenhance.SetNodeIdAndAttrHolder
var MergeClassAttrs = htmlenhance.MergeClassAttrs

type Component = htmlcomponent.ComponentTemplate
type Assemble = htmlcomponent.Slot
type Assembles = htmlcomponent.Slots
type Attribute = htmlcomponent.Attribute
type Attributes = htmlcomponent.Attributes
