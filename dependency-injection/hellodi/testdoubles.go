package hellodi

type spyMessageWriter struct {
	WrittenMessage string
	MessageCount   int
}

func (s *spyMessageWriter) Write(message string) {
	s.WrittenMessage += message
	s.MessageCount++
}

type testIdentity struct {
	isAuthenticated bool
}

func (t *testIdentity) IsAuthenticated() bool {
	return t.isAuthenticated
}
