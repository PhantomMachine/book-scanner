//+build bugst

package serial

import (
	"io"
	"log"
	"time"

	"go.bug.st/serial"
)

func init() {
	log.Println("using go.bug.st/serial")
}

func New(device string, opts ...foption) (io.Reader, error) {
	mode := &serial.Mode{
		BaudRate: 9600,
	}

	port, err := serial.Open(device, mode)
	if err != nil {
		return nil, err
	}

	for _, f := range opts {
		f(port)
	}

	return port, err
}

type foption func(serial.Port)

func setTimeout(t time.Duration) foption {
	return func(p serial.Port) {
		p.SetReadTimeout(t)
	}
}
