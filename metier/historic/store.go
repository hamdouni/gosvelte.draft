package historic

var store Storage

// UseStore use specified storage for historic repository
func UseStore(h Storage) {
	store = h
}

type Storage interface {
	Save(string)
	List() []string
}
