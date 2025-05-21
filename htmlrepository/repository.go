package htmlrepository

type RepoComponent[Entity Component] interface {
	ListByComponentNames(componentNames []string) ([]Entity, error)
}

type RepoAssemble[Entity Assemble] interface {
	ListByRootComponentName(rootComponentName string) ([]Entity, error)
}
type RepoAttribute[Entity Attribute] interface {
	ListByRootComponentName(rootComponentName string) ([]Entity, error)
}
