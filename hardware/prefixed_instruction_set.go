package hardware

func (c *CPU) RLC_B() {
	// 0x00 Rotate register B left with carry
	carry := c.B >> 7
	c.B = c.B<<1 | carry

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(carry << 4)
	if c.B == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RLC_C() {
	// 0x01 Rotate register C left with carry
	carry := c.C >> 7
	c.C = c.C<<1 | carry

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(carry << 4)
	if c.C == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RLC_D() {
	// 0x02 Rotate register D left with carry
	carry := c.D >> 7
	c.D = c.D<<1 | carry

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(carry << 4)
	if c.D == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RLC_E() {
	// 0x03 Rotate register E left with carry
	carry := c.E >> 7
	c.E = c.E<<1 | carry

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(carry << 4)
	if c.E == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RLC_H() {
	// 0x04 Rotate register H left with carry
	carry := c.H >> 7
	c.H = c.H<<1 | carry

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(carry << 4)
	if c.H == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RLC_L() {
	// 0x05 Rotate register L left with carry
	carry := c.L >> 7
	c.L = c.L<<1 | carry

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(carry << 4)
	if c.L == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RLC_ADDR_HL() {
	// 0x06 Rotate register [HL] left with carry
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	carry := b >> 7
	b = b<<1 | carry
	c.Bus.Write(hl, b)

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(carry << 4)
	if b == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RLC_A() {
	// 0x07 Rotate register A left with carry
	carry := c.E >> 7
	c.E = c.E<<1 | carry

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(carry << 4)
	if c.E == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RRC_B() {
	// 0x08 Rotate register B right with carry
	carry := c.B & 0x01
	c.B = c.B >> 1
	c.B |= carry << 7
	c.Flag_set(carry << 4)
	if c.B == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RRC_C() {
	// 0x09 Rotate register C right with carry
	carry := c.C & 0x01
	c.C = c.C >> 1
	c.C |= carry << 7
	c.Flag_set(carry << 4)
	if c.C == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RRC_D() {
	// 0x0A Rotate register D right with carry
	carry := c.D & 0x01
	c.D = c.D >> 1
	c.D |= carry << 7
	c.Flag_set(carry << 4)
	if c.D == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RRC_E() {
	// 0x0B Rotate register E right with carry
	carry := c.E & 0x01
	c.E = c.E >> 1
	c.E |= carry << 7
	c.Flag_set(carry << 4)
	if c.E == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RRC_H() {
	// 0x0C Rotate register H right with carry
	carry := c.H & 0x01
	c.H = c.H >> 1
	c.H |= carry << 7
	c.Flag_set(carry << 4)
	if c.H == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RRC_L() {
	// 0x0D Rotate register L right with carry
	carry := c.L & 0x01
	c.L = c.L >> 1
	c.L |= carry << 7
	c.Flag_set(carry << 4)
	if c.L == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RRC_ADDR_HL() {
	// 0x0E Rotate register [HL] right with carry
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	carry := b & 0x01
	b = b >> 1
	b |= carry << 7
	c.Bus.Write(hl, b)

	c.Flag_set(carry << 4)
	if b == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RRC_A() {
	// 0x0A Rotate register A right with carry
	carry := c.E & 0x01
	c.E = c.E >> 1
	c.E |= carry << 7
	c.Flag_set(carry << 4)
	if c.E == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RL_B() {
	// 0x10 Rotate left B with B[0] = C and C = B[7]
	bit_7 := uint8(c.B & 0x80)
	carry := (c.F & FLAG_C) >> 4
	c.B = (c.B << 1) | carry

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(bit_7 >> 3)
	if c.B == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RL_C() {
	// 0x11 Rotate left C with C[0] = C and C = C[7]
	bit_7 := uint8(c.C & 0x80)
	carry := (c.F & FLAG_C) >> 4
	c.C = (c.C << 1) | carry

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(bit_7 >> 3)
	if c.C == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RL_D() {
	// 0x12 Rotate left D with D[0] = C and C = D[7]
	bit_7 := uint8(c.D & 0x80)
	carry := (c.F & FLAG_C) >> 4
	c.D = (c.D << 1) | carry

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(bit_7 >> 3)
	if c.D == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RL_E() {
	// 0x13 Rotate left E with E[0] = C and C = E[7]
	bit_7 := uint8(c.E & 0x80)
	carry := (c.F & FLAG_C) >> 4
	c.E = (c.E << 1) | carry

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(bit_7 >> 3)
	if c.E == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RL_H() {
	// 0x14 Rotate left H with H[0] = C and C = H[7]
	bit_7 := uint8(c.H & 0x80)
	carry := (c.F & FLAG_C) >> 4
	c.H = (c.H << 1) | carry

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(bit_7 >> 3)
	if c.H == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RL_L() {
	// 0x15 Rotate left A with A[0] = C and C = A[7]
	bit_7 := uint8(c.L & 0x80)
	carry := (c.F & FLAG_C) >> 4
	c.L = (c.L << 1) | carry

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(bit_7 >> 3)
	if c.L == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RL_ADDR_HL() {
	// 0x16 Rotate left [HL] with [HL][0] = C and C = [HL][7]
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	bit_7 := uint8(b & 0x80)
	carry := (c.F & FLAG_C) >> 4
	b = (b << 1) | carry
	c.Bus.Write(hl, b)

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(bit_7 >> 3)
	if b == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RL_A() {
	// 0x17 Rotate left A with A[0] = C and C = A[7]
	bit_7 := uint8(c.E & 0x80)
	carry := (c.F & FLAG_C) >> 4
	c.E = (c.E << 1) | carry

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(bit_7 >> 3)
	if c.E == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RR_B() {
	// 0x18 Rotate Right B with B[0] -> C and C -> B[7]
	bit_0 := uint8(c.B & 0x01)
	carry := (c.F & FLAG_C) << 3

	c.B = (c.B >> 1) | carry
	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(bit_0 << 4)
	if c.B == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RR_C() {
	// 0x19 Rotate Right C with C[0] -> C and C -> C[7]
	bit_0 := uint8(c.C & 0x01)
	carry := (c.F & FLAG_C) << 3

	c.C = (c.C >> 1) | carry
	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(bit_0 << 4)
	if c.C == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RR_D() {
	// 0x1A Rotate Right D with D[0] -> C and C -> D[7]
	bit_0 := uint8(c.D & 0x01)
	carry := (c.F & FLAG_C) << 3

	c.D = (c.D >> 1) | carry
	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(bit_0 << 4)
	if c.D == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RR_E() {
	// 0x1B Rotate Right E with E[0] -> C and C -> E[7]
	bit_0 := uint8(c.E & 0x01)
	carry := (c.F & FLAG_C) << 3

	c.E = (c.E >> 1) | carry
	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(bit_0 << 4)
	if c.E == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RR_H() {
	// 0x1C Rotate Right H with H[0] -> C and C -> H[7]
	bit_0 := uint8(c.H & 0x01)
	carry := (c.F & FLAG_C) << 3

	c.H = (c.H >> 1) | carry
	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(bit_0 << 4)
	if c.H == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RR_L() {
	// 0x1D Rotate Right L with L[0] -> C and C -> L[7]
	bit_0 := uint8(c.L & 0x01)
	carry := (c.F & FLAG_C) << 3

	c.L = (c.L >> 1) | carry
	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(bit_0 << 4)
	if c.L == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RR_ADDR_HL() {
	// 0x1E Rotate Right [HL] with [HL][0] -> C and C -> [HL][7]
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	bit_0 := uint8(b & 0x01)
	carry := (c.F & FLAG_C) << 3
	b = (b >> 1) | carry
	c.Bus.Write(hl, b)

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(bit_0 << 4)
	if b == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RR_A() {
	// 0x1F Rotate Right A with A[0] -> C and C -> A[7]
	bit_0 := uint8(c.A & 0x01)
	carry := (c.F & FLAG_C) << 3

	c.A = (c.A >> 1) | carry
	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(bit_0 << 4)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SLA_B() {
	// 0x20 Shift B Left Arithmetically
	bit_7 := c.B & 0x80
	c.B = (c.B << 1)
	c.F = 0
	c.Flag_set(bit_7 >> 3)
	if c.B == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SLA_C() {
	// 0x21 Shift C Left Arithmetically
	bit_7 := c.C & 0x80
	c.C = c.C << 1
	c.F = 0
	c.Flag_set(bit_7 >> 3)
	if c.C == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SLA_D() {
	// 0x22 Shift D Left Arithmetically
	bit_7 := c.D & 0x80
	c.D = c.D << 1
	c.F = 0
	c.Flag_set(bit_7 >> 3)
	if c.D == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SLA_E() {
	// 0x23 Shift E Left Arithmetically
	bit_7 := c.E & 0x80
	c.E = c.E << 1
	c.F = 0
	c.Flag_set(bit_7 >> 3)
	if c.E == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SLA_H() {
	// 0x24 Shift H Left Arithmetically
	bit_7 := c.H & 0x80
	c.H = c.H << 1
	c.F = 0
	c.Flag_set(bit_7 >> 3)
	if c.H == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SLA_L() {
	// 0x25 Shift L Left Arithmetically
	bit_7 := c.L & 0x80
	c.L = c.L << 1
	c.F = 0
	c.Flag_set(bit_7 >> 3)
	if c.L == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SLA_ADDR_HL() {
	// 0x26 Shift [HL] Left Arithmetically
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	bit_7 := b & 0x80
	b = b << 1
	c.Bus.Write(hl, b)

	c.F = 0
	c.Flag_set(bit_7 >> 3)
	if b == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SLA_A() {
	// 0x27 Shift A Left Arithmetically
	bit_7 := c.A & 0x80
	c.A = c.A << 1
	c.F = 0
	c.Flag_set(bit_7 >> 3)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SRA_B() {
	// 0x28 Shift B Right Arithmetically
	bit_0 := c.B & 0x01
	c.B = (c.B >> 1) | (c.B & 0x80)
	c.F = 0
	c.Flag_set(bit_0 << 4)
	if c.B == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SRA_C() {
	// 0x29 Shift C Right Arithmetically
	bit_0 := c.C & 0x01
	c.C = (c.C >> 1) | (c.C & 0x80)
	c.F = 0
	c.Flag_set(bit_0 << 4)
	if c.C == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SRA_D() {
	// 0x2A Shift D Right Arithmetically
	bit_0 := c.D & 0x01
	c.D = (c.D >> 1) | (c.D & 0x80)
	c.F = 0
	c.Flag_set(bit_0 << 4)
	if c.D == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SRA_E() {
	// 0x2B Shift E Right Arithmetically
	bit_0 := c.E & 0x01
	c.E = (c.E >> 1) | (c.E & 0x80)
	c.F = 0
	c.Flag_set(bit_0 << 4)
	if c.E == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SRA_H() {
	// 0x2C Shift H Right Arithmetically
	bit_0 := c.H & 0x01
	c.H = (c.H >> 1) | (c.H & 0x80)
	c.F = 0
	c.Flag_set(bit_0 << 4)
	if c.H == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SRA_L() {
	// 0x2D Shift L Right Arithmetically
	bit_0 := c.L & 0x01
	c.L = (c.L >> 1) | (c.L & 0x80)
	c.F = 0
	c.Flag_set(bit_0 << 4)
	if c.L == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SRA_ADDR_HL() {
	// 0x2E Shift [HL] Right Arithmetically
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	bit_0 := b & 0x01
	b = (b >> 1) | (b & 0x80)
	c.Bus.Write(hl, b)

	c.F = 0
	c.Flag_set(bit_0 << 4)
	if b == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SRA_A() {
	// 0x2F Shift A Right Arithmetically
	bit_0 := c.A & 0x01
	c.A = (c.A >> 1) | (c.A & 0x80)
	c.F = 0
	c.Flag_set(bit_0 << 4)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SWAP_B() {
	// 0x30 Swap low and high nibble of register B
	c.B = (c.B&0x0f)<<4 | (c.B&0xf0)>>4
	c.F = 0
	if c.B == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SWAP_C() {
	// 0x31 Swap low and high nibble of register C
	c.C = (c.C&0x0f)<<4 | (c.C&0xf0)>>4
	c.F = 0
	if c.C == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SWAP_D() {
	// 0x32 Swap low and high nibble of register D
	c.D = (c.D&0x0f)<<4 | (c.D&0xf0)>>4
	c.F = 0
	if c.D == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SWAP_E() {
	// 0x33 Swap low and high nibble of register E
	c.E = (c.E&0x0f)<<4 | (c.E&0xf0)>>4
	c.F = 0
	if c.E == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SWAP_H() {
	// 0x34 Swap low and high nibble of register H
	c.H = (c.H&0x0f)<<4 | (c.H&0xf0)>>4
	c.F = 0
	if c.H == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SWAP_L() {
	// 0x35 Swap low and high nibble of register L
	c.L = (c.L&0x0f)<<4 | (c.L&0xf0)>>4
	c.F = 0
	if c.L == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SWAP_ADDR_HL() {
	// 0x36 Swap low and high nibble of register [HL]
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	b = (b&0x0f)<<4 | (b&0xf0)>>4
	c.Bus.Write(hl, b)

	c.F = 0
	if b == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SWAP_A() {
	// 0x37 Swap low and high nibble of register A
	c.A = (c.A&0x0f)<<4 | (c.A&0xf0)>>4
	c.F = 0
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SRL_B() {
	// 0x38 Shift B Right Logically
	bit_0 := c.B & 0x01
	c.B = (c.B >> 1)
	c.F = 0
	c.Flag_set(bit_0 << 4)
	if c.B == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SRL_C() {
	// 0x39 Shift C Right Logically
	bit_0 := c.C & 0x01
	c.C = c.C >> 1
	c.F = 0
	c.Flag_set(bit_0 << 4)
	if c.C == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SRL_D() {
	// 0x3A Shift D Right Logically
	bit_0 := c.D & 0x01
	c.D = c.D >> 1
	c.F = 0
	c.Flag_set(bit_0 << 4)
	if c.D == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SRL_E() {
	// 0x3B Shift E Right Logically
	bit_0 := c.E & 0x01
	c.E = c.E >> 1
	c.F = 0
	c.Flag_set(bit_0 << 4)
	if c.E == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SRL_H() {
	// 0x3C Shift H Right Logically
	bit_0 := c.H & 0x01
	c.H = c.H >> 1
	c.F = 0
	c.Flag_set(bit_0 << 4)
	if c.H == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SRL_L() {
	// 0x3D Shift L Right Logically
	bit_0 := c.L & 0x01
	c.L = c.L >> 1
	c.F = 0
	c.Flag_set(bit_0 << 4)
	if c.L == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SRL_ADDR_HL() {
	// 0x3E Shift [HL] Right Logically
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	bit_0 := b & 0x01
	b = b >> 1
	c.Bus.Write(hl, b)

	c.F = 0
	c.Flag_set(bit_0 << 4)
	if b == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) SRL_A() {
	// 0x3F Shift A Right Logically
	bit_0 := c.A & 0x01
	c.A = c.A >> 1
	c.F = 0
	c.Flag_set(bit_0 << 4)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_0_B() {
	// 0x40 Check if bit 0 of register B is set
	bit := c.B & (1 << 0)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_0_C() {
	// 0x41 Check if bit 0 of register C is set
	bit := c.C & (1 << 0)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_0_D() {
	// 0x42 Check if bit 0 of register D is set
	bit := c.D & (1 << 0)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_0_E() {
	// 0x43 Check if bit 0 of register E is set
	bit := c.E & (1 << 0)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_0_H() {
	// 0x44 Check if bit 0 of register H is set
	bit := c.H & (1 << 0)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_0_L() {
	// 0x45 Check if bit 0 of register L is set
	bit := c.L & (1 << 0)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_0_ADDR_HL() {
	// 0x46 Check if bit 0 of register [HL] is set
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	bit := b & (1 << 0)

	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_0_A() {
	// 0x47 Check if bit 0 of register A is set
	bit := c.A & (1 << 0)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_1_B() {
	// 0x48 Check if bit 1 of register B is set
	bit := c.B & (1 << 1)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_1_C() {
	// 0x49 Check if bit 1 of register C is set
	bit := c.C & (1 << 1)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_1_D() {
	// 0x4A Check if bit 1 of register D is set
	bit := c.D & (1 << 1)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_1_E() {
	// 0x4B Check if bit 1 of register E is set
	bit := c.E & (1 << 1)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_1_H() {
	// 0x4C Check if bit 1 of register H is set
	bit := c.H & (1 << 1)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_1_L() {
	// 0x4D Check if bit 1 of register L is set
	bit := c.L & (1 << 1)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_1_ADDR_HL() {
	// 0x4E Check if bit 1 of register [HL] is set
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	bit := b & (1 << 1)

	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_1_A() {
	// 0x4F Check if bit 1 of register A is set
	bit := c.A & (1 << 1)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_2_B() {
	// 0x50 Check if bit 2 of register B is set
	bit := c.B & (1 << 2)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_2_C() {
	// 0x51 Check if bit 2 of register C is set
	bit := c.C & (1 << 2)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_2_D() {
	// 0x52 Check if bit 2 of register D is set
	bit := c.D & (1 << 2)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_2_E() {
	// 0x53 Check if bit 2 of register E is set
	bit := c.E & (1 << 2)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_2_H() {
	// 0x54 Check if bit 2 of register H is set
	bit := c.H & (1 << 2)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_2_L() {
	// 0x55 Check if bit 2 of register L is set
	bit := c.L & (1 << 2)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_2_ADDR_HL() {
	// 0x56 Check if bit 2 of register [HL] is set
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	bit := b & (1 << 2)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_2_A() {
	// 0x57 Check if bit 2 of register A is set
	bit := c.A & (1 << 2)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_3_B() {
	// 0x58 Check if bit 3 of register B is set
	bit := c.B & (1 << 3)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_3_C() {
	// 0x59 Check if bit 3 of register C is set
	bit := c.C & (1 << 3)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_3_D() {
	// 0x5A Check if bit 3 of register D is set
	bit := c.D & (1 << 3)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_3_E() {
	// 0x5B Check if bit 3 of register E is set
	bit := c.E & (1 << 3)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_3_H() {
	// 0x5C Check if bit 3 of register H is set
	bit := c.H & (1 << 3)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_3_L() {
	// 0x5D Check if bit 3 of register L is set
	bit := c.L & (1 << 3)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_3_ADDR_HL() {
	// 0x5E Check if bit 3 of register [HL] is set
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	bit := b & (1 << 3)

	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_3_A() {
	// 0x5F Check if bit 3 of register A is set
	bit := c.A & (1 << 3)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_4_B() {
	// 0x60 Check if bit 4 of register B is set
	bit := c.B & (1 << 4)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_4_C() {
	// 0x61 Check if bit 4 of register C is set
	bit := c.C & (1 << 4)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_4_D() {
	// 0x62 Check if bit 4 of register D is set
	bit := c.D & (1 << 4)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_4_E() {
	// 0x63 Check if bit 4 of register E is set
	bit := c.E & (1 << 4)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_4_H() {
	// 0x64 Check if bit 4 of register H is set
	bit := c.H & (1 << 4)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_4_L() {
	// 0x65 Check if bit 4 of register L is set
	bit := c.L & (1 << 4)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_4_ADDR_HL() {
	// 0x66 Check if bit 4 of register [HL] is set
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	bit := b & (1 << 4)

	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_4_A() {
	// 0x67 Check if bit 4 of register A is set
	bit := c.A & (1 << 4)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_5_B() {
	// 0x68 Check if bit 5 of register B is set
	bit := c.B & (1 << 5)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_5_C() {
	// 0x69 Check if bit 5 of register C is set
	bit := c.C & (1 << 5)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_5_D() {
	// 0x6A Check if bit 5 of register D is set
	bit := c.D & (1 << 5)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_5_E() {
	// 0x6B Check if bit 5 of register E is set
	bit := c.E & (1 << 5)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_5_H() {
	// 0x6C Check if bit 5 of register H is set
	bit := c.H & (1 << 5)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_5_L() {
	// 0x6D Check if bit 5 of register L is set
	bit := c.L & (1 << 5)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_5_ADDR_HL() {
	// 0x6E Check if bit 5 of register [HL] is set
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	bit := b & (1 << 5)

	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_5_A() {
	// 0x6F Check if bit 5 of register A is set
	bit := c.A & (1 << 5)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_6_B() {
	// 0x70 Check if bit 6 of register B is set
	bit := c.B & (1 << 6)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_6_C() {
	// 0x71 Check if bit 6 of register C is set
	bit := c.C & (1 << 6)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_6_D() {
	// 0x72 Check if bit 6 of register D is set
	bit := c.D & (1 << 6)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_6_E() {
	// 0x73 Check if bit 6 of register E is set
	bit := c.E & (1 << 6)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_6_H() {
	// 0x74 Check if bit 6 of register H is set
	bit := c.H & (1 << 6)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_6_L() {
	// 0x75 Check if bit 6 of register L is set
	bit := c.L & (1 << 6)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_6_ADDR_HL() {
	// 0x76 Check if bit 6 of register [HL] is set
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	bit := b & (1 << 6)

	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_6_A() {
	// 0x77 Check if bit 6 of register A is set
	bit := c.A & (1 << 6)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_7_B() {
	// 0x78 Check if bit 7 of register B is set
	bit := c.B & (1 << 7)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_7_C() {
	// 0x79 Check if bit 7 of register C is set
	bit := c.C & (1 << 7)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_7_D() {
	// 0x7A Check if bit 7 of register D is set
	bit := c.D & (1 << 7)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_7_E() {
	// 0x7B Check if bit 7 of register E is set
	bit := c.E & (1 << 7)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_7_H() {
	// 0x7C Check if bit 7 of register H is set
	bit := c.H & (1 << 7)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_7_L() {
	// 0x7D Check if bit 7 of register L is set
	bit := c.L & (1 << 7)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_7_ADDR_HL() {
	// 0x7E Check if bit 7 of register [HL] is set
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	bit := b & (1 << 7)

	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) BIT_7_A() {
	// 0x7F Check if bit 7 of register A is set
	bit := c.A & (1 << 7)
	c.F = (c.F & FLAG_C) | FLAG_H
	if bit != 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RES_0_B() {
	// 0x80 Reset bit 0 of register B
	c.B &= ((1 << 0) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_0_C() {
	// 0x81 Reset bit 0 of register C
	c.C &= ((1 << 0) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_0_D() {
	// 0x82 Reset bit 0 of register D
	c.D &= ((1 << 0) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_0_E() {
	// 0x83 Reset bit 0 of register E
	c.E &= ((1 << 0) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_0_H() {
	// 0x84 Reset bit 0 of register H
	c.H &= ((1 << 0) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_0_L() {
	// 0x85 Reset bit 0 of register L
	c.L &= ((1 << 0) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_0_ADDR_HL() {
	// 0x86 Reset bit 0 of register [HL]
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	b &= ((1 << 0) ^ (0xff))
	c.Bus.Write(hl, b)
	c.PC++
}

func (c *CPU) RES_0_A() {
	// 0x87 Reset bit 0 of register A
	c.A &= ((1 << 0) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_1_B() {
	// 0x88 Reset bit 1 of register B
	c.B &= ((1 << 1) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_1_C() {
	// 0x89 Reset bit 1 of register C
	c.C &= ((1 << 1) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_1_D() {
	// 0x8A Reset bit 1 of register D
	c.D &= ((1 << 1) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_1_E() {
	// 0x8B Reset bit 1 of register E
	c.E &= ((1 << 1) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_1_H() {
	// 0x8C Reset bit 1 of register H
	c.H &= ((1 << 1) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_1_L() {
	// 0x8D Reset bit 1 of register L
	c.L &= ((1 << 1) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_1_ADDR_HL() {
	// 0x8E Reset bit 1 of register [HL]
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	b &= ((1 << 1) ^ (0xff))
	c.Bus.Write(hl, b)
	c.PC++
}

func (c *CPU) RES_1_A() {
	// 0x8F Reset bit 1 of register A
	c.A &= ((1 << 1) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_2_B() {
	// 0x90 Reset bit 2 of register B
	c.B &= ((1 << 2) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_2_C() {
	// 0x91 Reset bit 2 of register C
	c.C &= ((1 << 2) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_2_D() {
	// 0x92 Reset bit 2 of register D
	c.D &= ((1 << 2) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_2_E() {
	// 0x93 Reset bit 2 of register E
	c.E &= ((1 << 2) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_2_H() {
	// 0x94 Reset bit 2 of register H
	c.H &= ((1 << 2) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_2_L() {
	// 0x95 Reset bit 2 of register L
	c.L &= ((1 << 2) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_2_ADDR_HL() {
	// 0x96 Reset bit 2 of register [HL]
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	b &= ((1 << 2) ^ (0xff))
	c.Bus.Write(hl, b)
	c.PC++
}

func (c *CPU) RES_2_A() {
	// 0x97 Reset bit 2 of register A
	c.A &= ((1 << 2) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_3_B() {
	// 0x98 Reset bit 3 of register B
	c.B &= ((1 << 3) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_3_C() {
	// 0x99 Reset bit 3 of register C
	c.C &= ((1 << 3) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_3_D() {
	// 0x9A Reset bit 3 of register D
	c.D &= ((1 << 3) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_3_E() {
	// 0x9B Reset bit 3 of register E
	c.E &= ((1 << 3) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_3_H() {
	// 0x9C Reset bit 3 of register H
	c.H &= ((1 << 3) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_3_L() {
	// 0x9D Reset bit 3 of register L
	c.L &= ((1 << 3) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_3_ADDR_HL() {
	// 0x9E Reset bit 3 of register [HL]
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	b &= ((1 << 3) ^ (0xff))
	c.Bus.Write(hl, b)
	c.PC++
}

func (c *CPU) RES_3_A() {
	// 0x9F Reset bit 3 of register A
	c.A &= ((1 << 3) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_4_B() {
	// 0xA0 Reset bit 4 of register B
	c.B &= ((1 << 4) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_4_C() {
	// 0xA1 Reset bit 4 of register C
	c.C &= ((1 << 4) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_4_D() {
	// 0xA2 Reset bit 4 of register D
	c.D &= ((1 << 4) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_4_E() {
	// 0xA3 Reset bit 4 of register E
	c.E &= ((1 << 4) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_4_H() {
	// 0xA4 Reset bit 4 of register H
	c.H &= ((1 << 4) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_4_L() {
	// 0xA5 Reset bit 4 of register L
	c.L &= ((1 << 4) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_4_ADDR_HL() {
	// 0xA6 Reset bit 4 of register [HL]
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	b &= ((1 << 4) ^ (0xff))
	c.Bus.Write(hl, b)
	c.PC++
}

func (c *CPU) RES_4_A() {
	// 0xA7 Reset bit 4 of register A
	c.A &= ((1 << 4) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_5_B() {
	// 0xA8 Reset bit 5 of register B
	c.B &= ((1 << 5) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_5_C() {
	// 0xA9 Reset bit 5 of register C
	c.C &= ((1 << 5) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_5_D() {
	// 0xAA Reset bit 5 of register D
	c.D &= ((1 << 5) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_5_E() {
	// 0xAB Reset bit 5 of register E
	c.E &= ((1 << 5) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_5_H() {
	// 0xAC Reset bit 5 of register H
	c.H &= ((1 << 5) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_5_L() {
	// 0xAD Reset bit 5 of register L
	c.L &= ((1 << 5) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_5_ADDR_HL() {
	// 0xAE Reset bit 5 of register [HL]
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	b &= ((1 << 5) ^ (0xff))
	c.Bus.Write(hl, b)
	c.PC++
}

func (c *CPU) RES_5_A() {
	// 0xAF Reset bit 5 of register A
	c.A &= ((1 << 5) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_6_B() {
	// 0xB0 Reset bit 6 of register B
	c.B &= ((1 << 6) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_6_C() {
	// 0xB1 Reset bit 6 of register C
	c.C &= ((1 << 6) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_6_D() {
	// 0xB2 Reset bit 6 of register D
	c.D &= ((1 << 6) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_6_E() {
	// 0xB3 Reset bit 6 of register E
	c.E &= ((1 << 6) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_6_H() {
	// 0xB4 Reset bit 6 of register H
	c.H &= ((1 << 6) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_6_L() {
	// 0xB5 Reset bit 6 of register L
	c.L &= ((1 << 6) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_6_ADDR_HL() {
	// 0xB6 Reset bit 6 of register [HL]
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	b &= ((1 << 6) ^ (0xff))
	c.Bus.Write(hl, b)
	c.PC++
}

func (c *CPU) RES_6_A() {
	// 0xB7 Reset bit 6 of register A
	c.A &= ((1 << 6) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_7_B() {
	// 0xB8 Reset bit 7 of register B
	c.B &= ((1 << 7) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_7_C() {
	// 0xB9 Reset bit 7 of register C
	c.C &= ((1 << 7) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_7_D() {
	// 0xBA Reset bit 7 of register D
	c.D &= ((1 << 7) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_7_E() {
	// 0xBB Reset bit 7 of register E
	c.E &= ((1 << 7) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_7_H() {
	// 0xBC Reset bit 7 of register H
	c.H &= ((1 << 7) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_7_L() {
	// 0xBD Reset bit 7 of register L
	c.L &= ((1 << 7) ^ (0xff))
	c.PC++
}

func (c *CPU) RES_7_ADDR_HL() {
	// 0xBE Reset bit 7 of register [HL]
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	b &= ((1 << 7) ^ (0xff))
	c.Bus.Write(hl, b)
	c.PC++
}

func (c *CPU) RES_7_A() {
	// 0xBF Reset bit 7 of register A
	c.A &= ((1 << 7) ^ (0xff))
	c.PC++
}

func (c *CPU) SET_0_B() {
	// 0xC0 Set bit 0 of register B
	c.B |= (1 << 0)
	c.PC++
}

func (c *CPU) SET_0_C() {
	// 0xC1 Set bit 0 of register C
	c.C |= (1 << 0)
	c.PC++
}

func (c *CPU) SET_0_D() {
	// 0xC2 Set bit 0 of register D
	c.D |= (1 << 0)
	c.PC++
}

func (c *CPU) SET_0_E() {
	// 0xC3 Set bit 0 of register E
	c.E |= (1 << 0)
	c.PC++
}

func (c *CPU) SET_0_H() {
	// 0xC4 Set bit 0 of register H
	c.H |= (1 << 0)
	c.PC++
}

func (c *CPU) SET_0_L() {
	// 0xC5 Set bit 0 of register L
	c.L |= (1 << 0)
	c.PC++
}

func (c *CPU) SET_0_ADDR_HL() {
	// 0xC6 Set bit 0 of register [HL]
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	b |= (1 << 0)
	c.Bus.Write(hl, b)
	c.PC++
}

func (c *CPU) SET_0_A() {
	// 0xC7 Set bit 0 of register A
	c.A |= (1 << 0)
	c.PC++
}

func (c *CPU) SET_1_B() {
	// 0xC8 Set bit 1 of register B
	c.B |= (1 << 1)
	c.PC++
}

func (c *CPU) SET_1_C() {
	// 0xC9 Set bit 1 of register C
	c.C |= (1 << 1)
	c.PC++
}

func (c *CPU) SET_1_D() {
	// 0xCA Set bit 1 of register D
	c.D |= (1 << 1)
	c.PC++
}

func (c *CPU) SET_1_E() {
	// 0xCB Set bit 1 of register E
	c.E |= (1 << 1)
	c.PC++
}

func (c *CPU) SET_1_H() {
	// 0xCC Set bit 1 of register H
	c.H |= (1 << 1)
	c.PC++
}

func (c *CPU) SET_1_L() {
	// 0xCD Set bit 1 of register L
	c.L |= (1 << 1)
	c.PC++
}

func (c *CPU) SET_1_ADDR_HL() {
	// 0xCE Set bit 1 of register [HL]
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	b |= (1 << 1)
	c.Bus.Write(hl, b)
	c.PC++
}

func (c *CPU) SET_1_A() {
	// 0xCF Set bit 1 of register A
	c.A |= (1 << 1)
	c.PC++
}

func (c *CPU) SET_2_B() {
	// 0xD0 Set bit 2 of register B
	c.B |= (1 << 2)
	c.PC++
}

func (c *CPU) SET_2_C() {
	// 0xD1 Set bit 2 of register C
	c.C |= (1 << 2)
	c.PC++
}

func (c *CPU) SET_2_D() {
	// 0xD2 Set bit 2 of register D
	c.D |= (1 << 2)
	c.PC++
}

func (c *CPU) SET_2_E() {
	// 0xD3 Set bit 2 of register E
	c.E |= (1 << 2)
	c.PC++
}

func (c *CPU) SET_2_H() {
	// 0xD4 Set bit 2 of register H
	c.H |= (1 << 2)
	c.PC++
}

func (c *CPU) SET_2_L() {
	// 0xD5 Set bit 2 of register L
	c.L |= (1 << 2)
	c.PC++
}

func (c *CPU) SET_2_ADDR_HL() {
	// 0xD6 Set bit 2 of register [HL]
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	b |= (1 << 2)
	c.Bus.Write(hl, b)
	c.PC++
}

func (c *CPU) SET_2_A() {
	// 0xD7 Set bit 2 of register A
	c.A |= (1 << 2)
	c.PC++
}

func (c *CPU) SET_3_B() {
	// 0xD8 Set bit 3 of register B
	c.B |= (1 << 3)
	c.PC++
}

func (c *CPU) SET_3_C() {
	// 0xD9 Set bit 3 of register C
	c.C |= (1 << 3)
	c.PC++
}

func (c *CPU) SET_3_D() {
	// 0xDA Set bit 3 of register D
	c.D |= (1 << 3)
	c.PC++
}

func (c *CPU) SET_3_E() {
	// 0xDB Set bit 3 of register E
	c.E |= (1 << 3)
	c.PC++
}

func (c *CPU) SET_3_H() {
	// 0xDC Set bit 3 of register H
	c.H |= (1 << 3)
	c.PC++
}

func (c *CPU) SET_3_L() {
	// 0xDD Set bit 3 of register L
	c.L |= (1 << 3)
	c.PC++
}

func (c *CPU) SET_3_ADDR_HL() {
	// 0xDE Set bit 3 of register [HL]
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	b |= (1 << 3)
	c.Bus.Write(hl, b)
	c.PC++
}

func (c *CPU) SET_3_A() {
	// 0xDF Set bit 3 of register A
	c.A |= (1 << 3)
	c.PC++
}

func (c *CPU) SET_4_B() {
	// 0xE0 Set bit 4 of register B
	c.B |= (1 << 4)
	c.PC++
}

func (c *CPU) SET_4_C() {
	// 0xE1 Set bit 4 of register C
	c.C |= (1 << 4)
	c.PC++
}

func (c *CPU) SET_4_D() {
	// 0xE2 Set bit 4 of register D
	c.D |= (1 << 4)
	c.PC++
}

func (c *CPU) SET_4_E() {
	// 0xE3 Set bit 4 of register E
	c.E |= (1 << 4)
	c.PC++
}

func (c *CPU) SET_4_H() {
	// 0xE4 Set bit 4 of register H
	c.H |= (1 << 4)
	c.PC++
}

func (c *CPU) SET_4_L() {
	// 0xE5 Set bit 4 of register L
	c.L |= (1 << 4)
	c.PC++
}

func (c *CPU) SET_4_ADDR_HL() {
	// 0xE6 Set bit 4 of register [HL]
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	b |= (1 << 4)
	c.Bus.Write(hl, b)
	c.PC++
}

func (c *CPU) SET_4_A() {
	// 0xE7 Set bit 4 of register A
	c.A |= (1 << 4)
	c.PC++
}

func (c *CPU) SET_5_B() {
	// 0xE8 Set bit 5 of register B
	c.B |= (1 << 5)
	c.PC++
}

func (c *CPU) SET_5_C() {
	// 0xE9 Set bit 5 of register C
	c.C |= (1 << 5)
	c.PC++
}

func (c *CPU) SET_5_D() {
	// 0xEA Set bit 5 of register D
	c.D |= (1 << 5)
	c.PC++
}

func (c *CPU) SET_5_E() {
	// 0xEB Set bit 5 of register E
	c.E |= (1 << 5)
	c.PC++
}

func (c *CPU) SET_5_H() {
	// 0xEC Set bit 5 of register H
	c.H |= (1 << 5)
	c.PC++
}

func (c *CPU) SET_5_L() {
	// 0xED Set bit 5 of register L
	c.L |= (1 << 5)
	c.PC++
}

func (c *CPU) SET_5_ADDR_HL() {
	// 0xEE Set bit 5 of register [HL]
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	b |= (1 << 5)
	c.Bus.Write(hl, b)
	c.PC++
}

func (c *CPU) SET_5_A() {
	// 0xEF Set bit 5 of register A
	c.A |= (1 << 5)
	c.PC++
}

func (c *CPU) SET_6_B() {
	// 0xF0 Set bit 6 of register B
	c.B |= (1 << 6)
	c.PC++
}

func (c *CPU) SET_6_C() {
	// 0xF1 Set bit 6 of register C
	c.C |= (1 << 6)
	c.PC++
}

func (c *CPU) SET_6_D() {
	// 0xF2 Set bit 6 of register D
	c.D |= (1 << 6)
	c.PC++
}

func (c *CPU) SET_6_E() {
	// 0xF3 Set bit 6 of register E
	c.E |= (1 << 6)
	c.PC++
}

func (c *CPU) SET_6_H() {
	// 0xF4 Set bit 6 of register H
	c.H |= (1 << 6)
	c.PC++
}

func (c *CPU) SET_6_L() {
	// 0xF5 Set bit 6 of register L
	c.L |= (1 << 6)
	c.PC++
}

func (c *CPU) SET_6_ADDR_HL() {
	// 0xF6 Set bit 6 of register [HL]
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	b |= (1 << 6)
	c.Bus.Write(hl, b)
	c.PC++
}

func (c *CPU) SET_6_A() {
	// 0xF7 Set bit 6 of register A
	c.A |= (1 << 6)
	c.PC++
}

func (c *CPU) SET_7_B() {
	// 0xF8 Set bit 7 of register B
	c.B |= (1 << 7)
	c.PC++
}

func (c *CPU) SET_7_C() {
	// 0xF9 Set bit 7 of register C
	c.C |= (1 << 7)
	c.PC++
}

func (c *CPU) SET_7_D() {
	// 0xFA Set bit 7 of register D
	c.D |= (1 << 7)
	c.PC++
}

func (c *CPU) SET_7_E() {
	// 0xFB Set bit 7 of register E
	c.E |= (1 << 7)
	c.PC++
}

func (c *CPU) SET_7_H() {
	// 0xFC Set bit 7 of register H
	c.H |= (1 << 7)
	c.PC++
}

func (c *CPU) SET_7_L() {
	// 0xFD Set bit 7 of register L
	c.L |= (1 << 7)
	c.PC++
}

func (c *CPU) SET_7_ADDR_HL() {
	// 0xFE Set bit 7 of register [HL]
	hl := c.read_r16(HL)
	b := c.Bus.Read(hl)

	b |= (1 << 7)
	c.Bus.Write(hl, b)
	c.PC++
}

func (c *CPU) SET_7_A() {
	// 0xFF Set bit 7 of register A
	c.A |= (1 << 7)
	c.PC++
}
