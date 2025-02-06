package deps

type ErrString string

func (e ErrString) Error() string {
	return string(e)
}
