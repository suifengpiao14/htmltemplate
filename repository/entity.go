package repository

type Component interface {
	Name() string
	Template() string
	DataTpl() string
}

type Assemble interface {
	PageName() string
	ComponentName() string
	AssembleName() string
	DataTpl() string
}
type Attribute interface {
	GetNodeId() string
	GetKey() string
	GetValue() string
}
