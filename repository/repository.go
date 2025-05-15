package repository

type RepoComponent[Entity Component] interface {
	GetByComponentNames(componentNames []string) ([]Entity, error)
}

type RepoAssemble[Entity Assemble] interface {
	GetByPageName(pageName string) ([]Entity, error)
}
type RepoAttribute[Entity Attribute] interface {
	GetByPageName(pageName string) ([]Entity, error)
}
