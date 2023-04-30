package historic

func Save(message string) {
	store.Save(message)
}

func Liste() []string {
	return store.List()
}
