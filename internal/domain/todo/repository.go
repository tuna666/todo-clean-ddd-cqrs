package todo

// Repository は集約 Todo の保存/取得ポート
type Repository interface {
	Create(t *Todo) error
	List() ([]Todo, error)
}
