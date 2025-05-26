package htmltemplate

import (
	"github.com/suifengpiao14/htmltemplate/htmlcomponent"
	"github.com/suifengpiao14/htmltemplate/htmlenhance"
)

var SetNodeIdAndAttrHolder = htmlenhance.SetNodeIdAndAttrHolder
var MergeClassAttrs = htmlenhance.MergeClassAttrs

type Component = htmlcomponent.Component
type Assemble = htmlcomponent.ComponentNode
type Assembles = htmlcomponent.ComponentNodes
type Attribute = htmlcomponent.Attribute
type Attributes = htmlcomponent.Attributes
