//+build tarm

package serial

import (
	"io"
	"log"
	"time"

	"github.com/tarm/serial"
)

func init() {
	log.Println("using tarm/serial")
}

func New(device string, opts ...foption) (io.Reader, error) {
	c := &serial.Config{
		Name: device,
		Baud: 9600,
	}

	for _, f := range opts {
		f(c)
	}

	port, err := serial.OpenPort(c)
	if err != nil {
		return nil, err
	}

	return port, err
}

type foption func(*serial.Config)

func setTimeout(t time.Duration) foption {
	return func(c *serial.Config) {
		c.ReadTimeout = t
	}
}
