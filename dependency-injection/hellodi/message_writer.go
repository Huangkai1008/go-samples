package hellodi

import "fmt"

type MessageWriter interface {
	Write(message string)
}

// ConsoleMessageWriter is a MessageWriter that writes to the console.
type ConsoleMessageWriter struct {
}

func (w ConsoleMessageWriter) Write(message string) {
	fmt.Println(message)
}

type SecureMessageWriter struct {
	Writer   MessageWriter
	identity Identity
}

func (w SecureMessageWriter) Write(message string) {
	if w.identity.IsAuthenticated() {
		w.Writer.Write(message)
	}
}
