package hardware

import (
	"testing"
)

func TestCPUAdd8(t *testing.T) {
	cpu := GetCPU()

	// result is 3 and both carry flags are 0
	res, carry := cpu.add8(0x01, 0x01, 1)
	if res != 0x03 || carry != 0b00000000 {
		t.Fatalf("got result %02x and carry %02x, wanted 03 and 00", res, carry)
	}

	res, carry = cpu.add8(0x0f, 0x01, 0)
	if res != 0x10 || carry != 0b00100000 {
		t.Fatalf("got result %02x and carry %02x, wanted 10 and 20", res, carry)
	}

	res, carry = cpu.add8(0x0f, 0, 1)
	if res != 0x10 || carry != 0b00100000 {
		t.Fatalf("got result %02x and carry %02x, wanted 10 and 20", res, carry)
	}

	res, carry = cpu.add8(0xff, 0xff, 1)
	if res != 0xff || carry != 0b00110000 {
		t.Fatalf("got result %02x and carry %02x, wanted ff and 30", res, carry)
	}

	res, carry = cpu.add8(0xff, 0, 1)
	if res != 0x00 || carry != 0b10110000 {
		t.Fatalf("got result %02x and carry %02x, wanted 00 and B0", res, carry)
	}

	res, carry = cpu.add8(0xe0, 0xf1, 0)
	if res != 0xd1 || carry != 0b00010000 {
		t.Fatalf("got result %02x and carry %02x, wanted d1 and 10", res, carry)
	}

	res, carry = cpu.add8(0x00, 0x00, 0)
	if res != 0x00 || carry != 0b10000000 {
		t.Fatalf("got result %02x and carry %02x, wanted 00 and 80", res, carry)
	}
}

func TestCPUSub8(t *testing.T) {
	cpu := GetCPU()

	res, carry := cpu.sub8(0x01, 0x01, 1)
	if res != 0xff || carry != 0b01110000 {
		t.Fatalf("got result %02x and carry %02x, wanted ff and 70", res, carry)
	}

	res, carry = cpu.sub8(0x0f, 0x01, 0)
	if res != 0x0e || carry != 0b01000000 {
		t.Fatalf("got result %02x and carry %02x, wanted 0e and 40", res, carry)
	}

	res, carry = cpu.sub8(0x0f, 0, 1)
	if res != 0x0e || carry != 0b01000000 {
		t.Fatalf("got result %02x and carry %02x, wanted 0e and 40", res, carry)
	}

	res, carry = cpu.sub8(0xff, 0xff, 0)
	if res != 0x00 || carry != 0b11000000 {
		t.Fatalf("got result %02x and carry %02x, wanted 00 and c0", res, carry)
	}

	res, carry = cpu.sub8(0xf0, 0, 1)
	if res != 0xef || carry != 0b01100000 {
		t.Fatalf("got result %02x and carry %02x, wanted ef and 00", res, carry)
	}

	res, carry = cpu.sub8(0x22, 0xf1, 0)
	if res != 0x31 || carry != 0b01010000 {
		t.Fatalf("got result %02x and carry %02x, wanted 31 and 50", res, carry)
	}

	res, carry = cpu.sub8(0x00, 0x00, 0)
	if res != 0x00 || carry != 0b11000000 {
		t.Fatalf("got result %02x and carry %02x, wanted 00 and C0", res, carry)
	}
}

func TestCPUADDB(t *testing.T) {
	cpu := GetCPU()

	cpu.A = 0xff
	cpu.B = 0xff

	cpu.ADD_B()
	if cpu.A != 0xfe || cpu.F != 0b00110000 {
		t.Fatalf("got result %02x and carry %02x, wanted fe and 30", cpu.A, cpu.F)
	}

	cpu.A = 0xff
	cpu.B = 0x01

	cpu.ADD_B()
	if cpu.A != 0x00 || cpu.F != 0b10110000 {
		t.Fatalf("got result %02x and carry %02x, wanted 00 and B0", cpu.A, cpu.F)
	}
}

func TestCPUINCB(t *testing.T) {
	cpu := GetCPU()

	cpu.B = 0x00
	cpu.F = 0x00
	cpu.INC_B()
	if cpu.B != 0x01 || cpu.F != 0 {
		t.Fatalf("got result %02x and carry %02x, wanted 01 and 00", cpu.B, cpu.F)
	}

	cpu.B = 0x00
	cpu.F = 0x10
	cpu.INC_B()
	if cpu.B != 0x01 || cpu.F != 0x10 {
		t.Fatalf("got result %02x and carry %02x, wanted 01 and 10", cpu.B, cpu.F)
	}

	cpu.B = 0xff
	cpu.INC_B()
	if cpu.B != 0x00 || cpu.F != 0xB0 {
		t.Fatalf("got result %02x and carry %02x, wanted 01 and 10", cpu.B, cpu.F)
	}
}

func TestAddByteToStackPointer(t *testing.T) {
	cpu := GetCPU()

	cpu.SP = 0xff00
	cpu.Bus[cpu.PC+1] = 0x01
	cpu.ADD_SP_e8()
	if cpu.SP != 0xff01 || cpu.F != 0x00 {
		t.Fatalf("Stack pointer is %04x and flags are %02x, wanted 0xff01 and 0x00", cpu.SP, cpu.F)
	}

	cpu.SP = 0xff00
	cpu.Bus[cpu.PC+1] = 0xff // -1
	cpu.ADD_SP_e8()
	if cpu.SP != 0xfeff || cpu.F != 0x00 {
		t.Fatalf("Stack pointer is %04x and flags are %02x, wanted 0xff01 and 0x00", cpu.SP, cpu.F)
	}

	cpu.SP = 0x00ff
	cpu.Bus[cpu.PC+1] = 0x01
	cpu.ADD_SP_e8()
	if cpu.SP != 0x0100 || cpu.F != 0x30 {
		t.Fatalf("Stack pointer is %04x and flags are %02x, wanted 0x0100 and 0x30", cpu.SP, cpu.F)
	}

	cpu.SP = 0x000e
	cpu.Bus[cpu.PC+1] = 0x03
	cpu.ADD_SP_e8()
	if cpu.SP != 0x0011 || cpu.F != 0x20 {
		t.Fatalf("Stack pointer is %04x and flags are %02x, wanted 0x0011 and 0x20", cpu.SP, cpu.F)
	}

	cpu.SP = 0x00f0
	cpu.Bus[cpu.PC+1] = 0x16
	cpu.ADD_SP_e8()
	if cpu.SP != 0x0106 || cpu.F != 0x10 {
		t.Fatalf("Stack pointer is %04x and flags are %02x, wanted 0x0106 and 0x10", cpu.SP, cpu.F)
	}

	cpu.SP = 0x00f0
	cpu.Bus[cpu.PC+1] = 0xef // -17
	cpu.ADD_SP_e8()
	if cpu.SP != 0x00df || cpu.F != 0x10 {
		t.Fatalf("Stack pointer is %04x and flags are %02x, wanted 0x00df and 0x10", cpu.SP, cpu.F)
	}
}
