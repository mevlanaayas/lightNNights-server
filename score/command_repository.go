package score

type CommandRepository interface {
	Save(score Score) error
	Update(score Score) (int64, error)
}
