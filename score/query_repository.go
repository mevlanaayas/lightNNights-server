package score

type QueryRepository interface {
	Get(id string) error
	List() ([]Score, error)
}
