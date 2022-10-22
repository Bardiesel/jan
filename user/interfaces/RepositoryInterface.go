package interfaces

type RepositoryInterface[T any] interface {
	FindAll() []T
	FindById(id int) T
	Insert(data T) T
	Update(data T) T
	Delete(id int) bool
}
