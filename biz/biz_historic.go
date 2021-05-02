package biz

func (b BIZ) Historic() []string {
	return b.histo.Liste()
}
