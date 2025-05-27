package htmltemplate

import (
	"github.com/suifengpiao14/htmltemplate/htmlcomponent"
	"github.com/suifengpiao14/htmltemplate/htmlenhance"
)

var SetNodeIdAndAttrHolder = htmlenhance.SetNodeIdAndAttrHolder
var MergeClassAttrs = htmlenhance.MergeClassAttrs

type Component = htmlcomponent.ComponentTemplate
type Slot = htmlcomponent.Slot
type Slots = htmlcomponent.Slots
type Attribute = htmlcomponent.Attribute
type Attributes = htmlcomponent.Attributes
