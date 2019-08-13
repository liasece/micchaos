package boxes

type Session map[string]string

func (this *Session) Get(key string) string {
	if v, ok := (*this)[key]; ok {
		return v
	}
	return ""
}

func (this *Session) HasKey(key string) bool {
	_, ok := (*this)[key]
	return ok
}
