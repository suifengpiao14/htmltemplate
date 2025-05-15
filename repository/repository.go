package repository

type RepoComponent[Entity Component] interface {
	GetByComponentNames(componentNames []string) ([]Entity, error)
}

type RepoAssemble[Entity Assemble] interface {
	GetByRootComponentName(rootComponentName string) ([]Entity, error)
}
type RepoAttribute[Entity Attribute] interface {
	GetByRootComponentName(rootComponentName string) ([]Entity, error)
}
