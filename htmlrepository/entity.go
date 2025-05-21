package htmlrepository

type Component interface {
	GetName() string
	GetTemplate() string
	GetDataTpl() string
}

type Assemble interface {
	GetRootComponentName() string
	GetComponentName() string
	GetAssembleName() string
	GetDataTpl() string
}
type Attribute interface {
	GetNodeId() string
	GetAttributeName() string
	GetAttributeValue() string
}
