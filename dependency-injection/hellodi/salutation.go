package hellodi

type Salutation struct {
	writer MessageWriter
}

func (s Salutation) Exclaim() {
	s.writer.Write("Hello DI!")
}
