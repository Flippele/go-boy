package hardware

import (
	"fmt"
	"log"
	"strings"
)

const MEM_SIZE = 0xFFFF + 1

type Memory [MEM_SIZE]byte

var busInstance *Memory

func GetBus() *Memory {
	if busInstance != nil {
		return busInstance
	}

	log.Println("Creating Bus Instance")
	busInstance = &Memory{}
	return busInstance
}

func (m *Memory) Read(addr uint16) byte {
	return m[addr]
}

func (m *Memory) Write(addr uint16, b byte) {
	m[addr] = b
}

func (m *Memory) WriteBytes(code []byte, location uint16) {
	for k, b := range code {
		m[uint16(k)+location] = b
	}
}

func (m *Memory) String() string {
	rep := ""
	for j := 0; j < MEM_SIZE; j += 16 {
		rep += fmt.Sprintf("%04X: ", j)
		row := [16]string{}
		for i := 0; i < 16; i++ {
			row[i] = fmt.Sprintf("%02X", m[j+i])
		}
		rep += fmt.Sprintf("%s\t%s\n", strings.Join(row[:8], " "), strings.Join(row[8:], " "))
	}
	return rep
}
