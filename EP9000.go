package ep9000

import (
	"time"

	"go.bug.st/serial"
)

type EP9000 interface {
	Listen(chan []uint8) error
}

type ep9000 struct {
	port serial.Port
}

func New(port string, speed int) (EP9000, error) {
	ep := ep9000{}

	// prepare port
	mode := &serial.Mode{
		BaudRate: speed,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	var err error
	ep.port, err = serial.Open(port, mode)
	if err != nil {
		return nil, err
	}
	ep.port.SetReadTimeout(time.Millisecond * 100)
	ep.port.ResetInputBuffer()
	return &ep, nil
}

func (ep *ep9000) Listen(data chan []uint8) error {
	for {
		// try to read all data
		var err error
		var buffer []uint8
		temp := make([]uint8, 512)
		num := 0
		// while no data, keep reading
		for num == 0 {
			num, err = ep.port.Read(temp)
		}
		// ok, try to read all
		buffer = append(buffer, temp[:num]...)
		if err != nil {
			return err
		}
		for num > 0 {
			num, err = ep.port.Read(temp)
			if err != nil {
				return err
			}
			buffer = append(buffer, temp[:num]...)
		}
		// ok, send data
		data <- buffer
		// empty buffer
		buffer = []uint8{}
	}
}

func (ep *ep9000) Flush() {
	ep.port.ResetInputBuffer()
}
