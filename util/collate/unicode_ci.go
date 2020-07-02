package collate

type unicodeCICollator struct {
}

var mockCollator = &binCollator{}

func (u *unicodeCICollator) Compare(a, b string) int {
	return mockCollator.Compare(a, b)
}

func (u *unicodeCICollator) Key(str string) []byte {
	return mockCollator.Key(str)
}

func (u *unicodeCICollator) Pattern() WildcardPattern {
	return mockCollator.Pattern()
}
