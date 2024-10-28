package hardware

func (c *CPU) NOP() {
	// 0x00, no operation
	c.PC++
}

func (c *CPU) LD_BC_n16() {
	// 0x01 Load word to BC register
	w := c.read_word()
	c.write_r16(BC, w)
	c.PC++
}

func (c *CPU) LD_ADDR_BC_A() {
	// 0x02 Store accumulator to location [BC]
	addr := c.read_r16(BC)
	c.Bus.Write(addr, c.A)
	c.PC++
}

func (c *CPU) INC_BC() {
	// 0x03 Increment BC register without carry
	bc := c.read_r16(BC) + 1
	c.write_r16(BC, bc)
	c.PC++
}

func (c *CPU) INC_B() {
	// 0x04 Increment B Register
	res, flags := c.add8(c.B, 0, 1)
	flags = (flags & ^FLAG_C) | (c.F & FLAG_C)

	c.B = res
	c.F = flags
	c.PC++
}

func (c *CPU) DEC_B() {
	// 0x05 Decrement B Register
	res, flags := c.sub8(c.B, 0, 1)
	flags = (flags & ^FLAG_C) | (c.F & FLAG_C)

	c.B = res
	c.F = flags
	c.PC++
}

func (c *CPU) LD_B_n8() {
	// 0x06 Load byte into register B
	c.B = c.read_byte()
	c.PC++
}

func (c *CPU) RLCA() {
	// 0x07 Rotate register A left with carry
	carry := c.A >> 7
	c.A = c.A<<1 | carry

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(carry << 4)
	c.PC++
}

func (c *CPU) LD_ADDR_a16_SP() {
	// 0x08 Load stack pointer into 16 bit address
	addr := c.read_word()
	s := uint8(c.SP >> 8)
	p := uint8(c.SP)
	c.Bus.Write(addr, p)
	c.Bus.Write(addr+1, s)
	c.PC++
}

func (c *CPU) ADD_HL_BC() {
	// 0x09 Add contents of register BC to HL
	bc := c.read_r16(BC)
	hl := c.read_r16(HL)

	res, flags := c.add16(hl, bc)
	c.F = (flags & ^FLAG_Z) | (c.F & FLAG_Z) // preserve zero flag
	c.write_r16(HL, res)
	c.PC++
}

func (c *CPU) LD_A_ADDR_BC() {
	// 0x0A Load contents of [BC] to A
	addr := c.read_r16(BC)
	c.A = c.Bus.Read(addr)
	c.PC++
}

func (c *CPU) DEC_BC() {
	// 0x0B Decrement BC register
	bc := c.read_r16(BC) - 1
	c.write_r16(BC, bc)
	c.PC++
}

func (c *CPU) INC_C() {
	// 0x0C Increment C Register
	res, flags := c.add8(c.C, 0, 1)
	flags = (flags & ^FLAG_C) | (c.F & FLAG_C)

	c.C = res
	c.F = flags
	c.PC++
}

func (c *CPU) DEC_C() {
	// 0x0D Decrement C Register
	res, flags := c.sub8(c.C, 0, 1)
	flags = (flags & ^FLAG_C) | (c.F & FLAG_C)

	c.C = res
	c.F = flags
	c.PC++
}

func (c *CPU) LD_C_n8() {
	// 0x0E Load byte into register C
	n := c.read_byte()
	c.C = n
	c.PC++
}

func (c *CPU) RRCA() {
	// 0x0F Rotate register A right with carry
	carry := c.A & 0x01
	c.A = c.A >> 1
	c.A |= carry << 7
	c.Flag_set(carry << 4)
	c.PC++
}

func (c *CPU) STOP() {
	// 0x10 Enter CPU low power mode
	//TODO
	c.PC++
}

func (c *CPU) LD_DE_n16() {
	// 0x11 Load word into DE register
	w := c.read_word()
	c.write_r16(DE, w)
	c.PC++
}

func (c *CPU) LD_ADDR_DE_A() {
	// 0x12 Store A into address [DE]
	addr := c.read_r16(DE)
	c.Bus.Write(addr, c.A)
	c.PC++
}

func (c *CPU) INC_DE() {
	// 0x13 Increment DE register
	de := c.read_r16(DE) + 1
	c.write_r16(DE, de)
	c.PC++
}

func (c *CPU) INC_D() {
	// 0x14 Increment D Register
	res, flags := c.add8(c.D, 0, 1)
	flags = (flags & ^FLAG_C) | (c.F & FLAG_C) // preserve carry flag

	c.D = res
	c.F = flags
	c.PC++
}

func (c *CPU) DEC_D() {
	// 0x15 Decrement D Register
	res, flags := c.sub8(c.D, 0, 1)
	flags = (flags & ^FLAG_C) | (c.F & FLAG_C) // preserve carry flag

	c.D = res
	c.F = flags
	c.PC++
}

func (c *CPU) LD_D_n8() {
	// 0x16 Load byte to register D
	c.D = c.read_byte()
	c.PC++
}

func (c *CPU) RLA() {
	// 0x17 Rotate left A with A[0] = C and C = A[7]
	bit_7 := (c.A & 0x80)
	carry := (c.F & FLAG_C) >> 4
	c.A = (c.A << 1) | carry

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(bit_7 >> 3)
	c.PC++
}

func (c *CPU) JR_e8() {
	// 0x18 Jump relative by [-128,127] offset
	e := c.read_byte()
	addr := int16(c.PC) + int16(int8(e))
	c.PC = uint16(addr)
}

func (c *CPU) ADD_HL_DE() {
	// 0x19 Add value of DE to HL
	de := c.read_r16(DE)
	hl := c.read_r16(HL)

	res, flags := c.add16(hl, de)
	c.F = (flags & ^FLAG_Z) | (c.F & FLAG_Z) // preserve zero flag
	c.write_r16(HL, res)
	c.PC++
}

func (c *CPU) LD_A_ADDR_DE() {
	// 0x1A Load A with contents of [DE]
	addr := c.read_r16(DE)
	c.A = c.Bus.Read(addr)
	c.PC++
}

func (c *CPU) DEC_DE() {
	// 0x1B Decrement DE by one
	de := c.read_r16(DE) - 1
	c.write_r16(DE, de)
	c.PC++
}

func (c *CPU) INC_E() {
	// 0x1C Increment E register with carry
	res, flags := c.add8(c.E, 0, 1)
	flags = (flags & ^FLAG_C) | (c.F & FLAG_C)

	c.E = res
	c.F = flags
	c.PC++
}

func (c *CPU) DEC_E() {
	// 0x1D Decrement E register with carry
	res, flags := c.sub8(c.E, 0, 1)
	flags = (flags & ^FLAG_C) | (c.F & FLAG_C)

	c.E = res
	c.F = flags
	c.PC++
}

func (c *CPU) LD_E_n8() {
	// 0x1E Load byte to E register
	c.E = c.read_byte()
	c.PC++
}

func (c *CPU) RRA() {
	// 0x1F Rotate Right A with A[0] -> C and C -> A[7]
	bit_0 := uint8(c.A&0x01) << 4
	carry := (c.F & FLAG_C) << 3

	c.A = (c.A >> 1) | carry
	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H)
	c.Flag_set(bit_0)
	c.PC++
}

func (c *CPU) JR_NZ_e8() {
	// 0x20 Jump by offset e [-128,127] if Z flag is not set
	e := c.read_byte()
	z := (c.F) & FLAG_Z

	if z == 0 {
		des := int16(c.PC) + int16(int8(e))
		c.PC = uint16(des)
		return
	}
	c.PC++
}

func (c *CPU) LD_HL_n16() {
	// 0x21 Load word to HL register
	n := c.read_word()
	c.write_r16(HL, n)
	c.PC++
}

func (c *CPU) LDI_ADDR_HL_A() {
	// 0x22 Load A to address [HL] and increment HL
	hl := c.read_r16(HL)
	c.Bus.Write(hl, c.A)
	c.write_r16(HL, hl+1)
	c.PC++
}

func (c *CPU) INC_HL() {
	// 0x23 Increment HL
	hl := c.read_r16(HL) + 1
	c.write_r16(HL, hl)
	c.PC++
}

func (c *CPU) INC_H() {
	// 0x24 Increment H register
	res, flags := c.add8(c.H, 0, 1)
	flags = (flags & ^FLAG_C) | (c.F & FLAG_C)

	c.H = res
	c.F = flags
	c.PC++
}

func (c *CPU) DEC_H() {
	// 0x25 Decrement H register
	res, flags := c.sub8(c.H, 0, 1)
	flags = (flags & ^FLAG_C) | (c.F & FLAG_C)

	c.H = res
	c.F = flags
	c.PC++
}

func (c *CPU) LD_H_n8() {
	// 0x26 Load byte to H
	c.H = c.read_byte()
	c.PC++
}

func (c *CPU) DAA() {
	// 0x27 Adjust Accumulator for BCD, if A[0:3] > 9 then A <- A + 0x06, if A[4:8] > 9, then A <- A + 0x60
	a := uint16(c.A)
	h_flag := (c.F & FLAG_H)
	c_flag := (c.F & FLAG_C)

	lo_nibble := a & 0x0f
	hi_nibble := (a & 0xf0) >> 4

	if lo_nibble > 9 || h_flag != 0 {
		a += 0x06
	}
	if hi_nibble > 9 || c_flag != 0 {
		a += 0x60
	}
	c.A = uint8(a)

	flags := (uint8(a>>8) << 4) // Get C -> 000C
	flags |= (c.F & FLAG_N)     // Preserve N -> 0-0C
	if c.A == 0 {
		flags |= FLAG_Z // Get Z -> Z-0C
	}
	c.F = flags
	c.PC++
}

func (c *CPU) JR_Z_e8() {
	// 0x28 Jump by offset e [-128,127] if Z flag is set
	e := c.read_byte()
	z := (c.F) & FLAG_Z

	if z != 0 {
		dest := int16(c.PC) + int16(int8(e))
		c.PC = uint16(dest)
		return
	}
	c.PC++
}

func (c *CPU) ADD_HL_HL() {
	// 0x29 Add the contents of HL to itself
	hl := c.read_r16(HL)

	res, flags := c.add16(hl, hl)
	c.F = (flags & ^FLAG_Z) | (c.F & FLAG_Z) // preserve zero flag
	c.write_r16(HL, res)
	c.PC++
}

func (c *CPU) LDI_A_ADDR_HL() {
	// 0x2A Load [HL] to A and increment HL
	hl := c.read_r16(HL)
	c.A = c.Bus.Read(hl)
	c.write_r16(HL, hl+1)
	c.PC++
}

func (c *CPU) DEC_HL() {
	// 0x2B Decrement HL by one
	hl := c.read_r16(HL) - 1
	c.write_r16(HL, hl)
	c.PC++
}

func (c *CPU) INC_L() {
	// 0x2C Increment L by one
	res, flags := c.add8(c.L, 0, 1)
	flags = (flags & ^FLAG_C) | (c.F & FLAG_C)

	c.L = res
	c.F = flags
	c.PC++
}

func (c *CPU) DEC_L() {
	// 0x2D Decrement L by one
	res, flags := c.sub8(c.L, 0, 1)
	flags = (flags & ^FLAG_C) | (c.F & FLAG_C)

	c.L = res
	c.F = flags
	c.PC++
}

func (c *CPU) LD_L_n8() {
	// 0x2E Load n8 Into L
	c.L = c.read_byte()
	c.PC++
}

func (c *CPU) CPL() {
	// 0x2F Complement Accumulator
	c.A = ^c.A
	c.Flag_set(FLAG_N | FLAG_H)
	c.PC++
}

func (c *CPU) JR_NC_e8() {
	// 0x30 Jump relative if C flag not set
	e := c.read_byte()
	_c := (c.F) & FLAG_C

	if _c == 0 {
		dest := int16(c.PC) + int16(int8(e))
		c.PC = uint16(dest)
		return
	}
	c.PC++
}

func (c *CPU) LD_SP_n16() {
	// 0x31 Load word into Stack Pointer
	c.SP = c.read_word()
	c.PC++
}

func (c *CPU) LDD_ADDR_HL_A() {
	// 0x32 Store A to [HL] and decrement HL
	hl := c.read_r16(HL)
	c.Bus.Write(hl, c.A)
	c.write_r16(HL, hl-1)
	c.PC++
}

func (c *CPU) INC_SP() {
	// 0x33 Increment Stack Pointer
	c.SP++
	c.PC++
}

func (c *CPU) INC_ADDR_HL() {
	// 0x34 Increment value at address [HL]
	hl := c.read_r16(HL)
	res, flags := c.add8(c.Bus.Read(hl), 0, 1)
	flags = (flags & ^FLAG_C) | (c.F & FLAG_C)

	c.Bus.Write(hl, res)
	c.F = flags
	c.PC++
}

func (c *CPU) DEC_ADDR_HL() {
	// 0x35 Decrement value at address [HL]
	hl := c.read_r16(HL)
	res, flags := c.sub8(c.Bus.Read(hl), 0, 1)
	flags = (flags & ^FLAG_C) | (c.F & FLAG_C)

	c.Bus.Write(hl, res)
	c.F = flags
	c.PC++
}

func (c *CPU) LD_ADDR_HL_n8() {
	// 0x36 Load value to address [HL]
	hl := c.read_r16(HL)
	n := c.read_byte()
	c.Bus.Write(hl, n)
	c.PC++
}

func (c *CPU) SCF() {
	// 0x37 Set Carry Flag
	c.Flag_set(FLAG_C)
	c.Flag_reset(FLAG_N | FLAG_H)
	c.PC++
}

func (c *CPU) JR_C_e8() {
	// 0x38 Jump Relative if C flag is set
	e := c.read_byte()
	_c := (c.F) & FLAG_C

	if _c != 0 {
		dest := int16(c.PC) + int16(int8(e))
		c.PC = uint16(dest)
		return
	}
	c.PC++
}

func (c *CPU) ADD_HL_SP() {
	// 0x39 Add Stack Pointer to HL
	hl := c.read_r16(HL)

	res, flags := c.add16(hl, c.SP)
	c.F = (flags & ^FLAG_Z) | (c.F & FLAG_Z) // preserve zero flag
	c.write_r16(HL, res)
	c.PC++
}

func (c *CPU) LDD_A_ADDR_HL() {
	// 0x3A Load [HL] to A and decrement HL
	hl := c.read_r16(HL)
	c.A = c.Bus.Read(hl)
	c.write_r16(HL, hl-1)
	c.PC++
}

func (c *CPU) DEC_SP() {
	// 0x3B Decrement stack pointer
	c.SP--
	c.PC++
}

func (c *CPU) INC_A() {
	// 0x3C Increment register A
	res, flags := c.add8(c.A, 0, 1)
	flags = (flags & ^FLAG_C) | (c.F & FLAG_C)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) DEC_A() {
	// 0x3D Decrement Register A
	res, flags := c.sub8(c.A, 0, 1)
	flags = (flags & ^FLAG_C) | (c.F & FLAG_C)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) LD_A_n8() {
	// 0x3E Load byte into A
	n := c.read_byte()
	c.A = n
	c.PC++
}

func (c *CPU) CCF() {
	// 0x3F Complement Carry Flag
	c.Flag_reset(FLAG_N | FLAG_H)
	c.F = c.F ^ FLAG_C
	c.PC++
}

func (c *CPU) LD_B_B() {
	// 0x40 Load B with B register
	c.B = uint8(c.B)
	c.PC++
}

func (c *CPU) LD_B_C() {
	// 0x41 Load B with C register
	c.B = c.C
	c.PC++
}

func (c *CPU) LD_B_D() {
	// 0x42 Load B with D register
	c.B = c.D
	c.PC++
}

func (c *CPU) LD_B_E() {
	// 0x43 Load B with E register
	c.B = c.E
	c.PC++
}

func (c *CPU) LD_B_H() {
	// 0x44 Load B with H register
	c.B = c.H
	c.PC++
}

func (c *CPU) LD_B_L() {
	// 0x45 Load B with H register
	c.B = c.L
	c.PC++
}

func (c *CPU) LD_B_ADDR_HL() {
	// 0x46 Load B with value at address [HL]
	hl := c.read_r16(HL)
	c.B = c.Bus.Read(hl)
	c.PC++
}

func (c *CPU) LD_B_A() {
	// 0x47 Load B with A register
	c.B = c.A
	c.PC++
}

func (c *CPU) LD_C_B() {
	// 0x48 Load C with B register
	c.C = c.B
	c.PC++
}

func (c *CPU) LD_C_C() {
	// 0x49 Load C with C register
	c.C = uint8(c.C)
	c.PC++
}

func (c *CPU) LD_C_D() {
	// 0x4A Load C with D register
	c.C = c.D
	c.PC++
}

func (c *CPU) LD_C_E() {
	// 0x4B Load C with E register
	c.C = c.E
	c.PC++
}

func (c *CPU) LD_C_H() {
	// 0x4C Load C with H register
	c.C = c.H
	c.PC++
}

func (c *CPU) LD_C_L() {
	// 0x4D Load C with L register
	c.C = c.L
	c.PC++
}

func (c *CPU) LD_C_ADDR_HL() {
	// 0x4E Load C with value at address [HL]
	hl := c.read_r16(HL)
	c.C = c.Bus.Read(hl)
	c.PC++
}

func (c *CPU) LD_C_A() {
	// 0x4F Load C with A register
	c.C = c.A
	c.PC++
}

func (c *CPU) LD_D_B() {
	// 0x50 Load D with B register
	c.D = c.B
	c.PC++
}

func (c *CPU) LD_D_C() {
	// 0x51 Load D with C register
	c.D = c.C
	c.PC++
}

func (c *CPU) LD_D_D() {
	// 0x52 Load D with D register
	c.D = uint8(c.D)
	c.PC++
}

func (c *CPU) LD_D_E() {
	// 0x53 Load D with E register
	c.D = c.E
	c.PC++
}

func (c *CPU) LD_D_H() {
	// 0x54 Load D with H register
	c.D = c.H
	c.PC++
}

func (c *CPU) LD_D_L() {
	// 0x55 Load D with L register
	c.D = c.L
	c.PC++
}

func (c *CPU) LD_D_ADDR_HL() {
	// 0x56 Load D with value at address [HL]
	hl := c.read_r16(HL)
	c.D = c.Bus.Read(hl)
	c.PC++
}

func (c *CPU) LD_D_A() {
	// 0x57 Load D with A register
	c.D = c.A
	c.PC++
}

func (c *CPU) LD_E_B() {
	// 0x58 Load C with B register
	c.E = c.B
	c.PC++
}

func (c *CPU) LD_E_C() {
	// 0x59 Load C with C register
	c.E = c.C
	c.PC++
}

func (c *CPU) LD_E_D() {
	// 0x5A Load C with D register
	c.E = c.D
	c.PC++
}

func (c *CPU) LD_E_E() {
	// 0x5B Load C with E register
	c.E = uint8(c.E)
	c.PC++
}

func (c *CPU) LD_E_H() {
	// 0x5C Load C with H register
	c.E = c.H
	c.PC++
}

func (c *CPU) LD_E_L() {
	// 0x5D Load C with L register
	c.E = c.L
	c.PC++
}

func (c *CPU) LD_E_ADDR_HL() {
	// 0x5E Load C with value at address [HL]
	hl := c.read_r16(HL)
	c.E = c.Bus.Read(hl)
	c.PC++
}

func (c *CPU) LD_E_A() {
	// 0x5F Load C with A register
	c.E = c.A
	c.PC++
}

func (c *CPU) LD_H_B() {
	// 0x60 Load H with B register
	c.H = c.B
	c.PC++
}

func (c *CPU) LD_H_C() {
	// 0x61 Load H with C register
	c.H = c.C
	c.PC++
}

func (c *CPU) LD_H_D() {
	// 0x62 Load H with D register
	c.H = c.D
	c.PC++
}

func (c *CPU) LD_H_E() {
	// 0x63 Load H with E register
	c.H = c.E
	c.PC++
}

func (c *CPU) LD_H_H() {
	// 0x64 Load H with H register
	c.H = uint8(c.H)
	c.PC++
}

func (c *CPU) LD_H_L() {
	// 0x65 Load H with L register
	c.H = c.L
	c.PC++
}

func (c *CPU) LD_H_ADDR_HL() {
	// 0x66 Load H with value at address [HL]
	hl := c.read_r16(HL)
	c.H = c.Bus.Read(hl)
	c.PC++
}

func (c *CPU) LD_H_A() {
	// 0x67 Load H with A register
	c.H = c.A
	c.PC++
}

func (c *CPU) LD_L_B() {
	// 0x68 Load L with B register
	c.L = c.B
	c.PC++
}

func (c *CPU) LD_L_C() {
	// 0x69 Load L with C register
	c.L = c.C
	c.PC++
}

func (c *CPU) LD_L_D() {
	// 0x6A Load L with D register
	c.L = c.D
	c.PC++
}

func (c *CPU) LD_L_E() {
	// 0x6B Load L with E register
	c.L = c.E
	c.PC++
}

func (c *CPU) LD_L_H() {
	// 0x6C Load L with H register
	c.L = c.H
	c.PC++
}

func (c *CPU) LD_L_L() {
	// 0x6D Load L with L register
	c.L = uint8(c.L)
	c.PC++
}

func (c *CPU) LD_L_ADDR_HL() {
	// 0x6E Load L with value at address [HL]
	hl := c.read_r16(HL)
	c.L = c.Bus.Read(hl)
	c.PC++
}

func (c *CPU) LD_L_A() {
	// 0x6F Load L with A register
	c.L = c.A
	c.PC++
}

func (c *CPU) LD_ADDR_HL_B() {
	// 0x70 Load address [HL] with B register
	hl := c.read_r16(HL)
	c.Bus.Write(hl, c.B)
	c.PC++
}

func (c *CPU) LD_ADDR_HL_C() {
	// 0x71 Load address [HL] with C register
	hl := c.read_r16(HL)
	c.Bus.Write(hl, c.C)
	c.PC++
}

func (c *CPU) LD_ADDR_HL_D() {
	// 0x72 Load address [HL] with D register
	hl := c.read_r16(HL)
	c.Bus.Write(hl, c.D)
	c.PC++
}

func (c *CPU) LD_ADDR_HL_E() {
	// 0x73 Load address [HL] with E register
	hl := c.read_r16(HL)
	c.Bus.Write(hl, c.E)
	c.PC++
}

func (c *CPU) LD_ADDR_HL_H() {
	// 0x74 Load address [HL] with H register
	hl := c.read_r16(HL)
	c.Bus.Write(hl, c.H)
	c.PC++
}

func (c *CPU) LD_ADDR_HL_L() {
	// 0x75 Load address [HL] with L register
	hl := c.read_r16(HL)
	c.Bus.Write(hl, c.L)
	c.PC++
}

func (c *CPU) HALT() {
	// 0x76 Load address [HL] with value at address [HL]
	//TODO behavior depends on interrupts
	c.PC++
}

func (c *CPU) LD_ADDR_HL_A() {
	// 0x77 Load address [HL] with A register
	hl := c.read_r16(HL)
	c.Bus.Write(hl, c.A)
	c.PC++
}

func (c *CPU) LD_A_B() {
	// 0x78 Load A with B register
	c.A = c.B
	c.PC++
}

func (c *CPU) LD_A_C() {
	// 0x79 Load A with C register
	c.A = c.C
	c.PC++
}

func (c *CPU) LD_A_D() {
	// 0x7A Load A with D register
	c.A = c.D
	c.PC++
}

func (c *CPU) LD_A_E() {
	// 0x7B Load A with E register
	c.A = c.E
	c.PC++
}

func (c *CPU) LD_A_H() {
	// 0x7C Load A with H register
	c.A = c.H
	c.PC++
}

func (c *CPU) LD_A_L() {
	// 0x7D Load A with L register
	c.A = c.L
	c.PC++
}

func (c *CPU) LD_A_ADDR_HL() {
	// 0x7E Load A with value at address [HL]
	hl := c.read_r16(HL)
	c.A = c.Bus.Read(hl)
	c.PC++
}

func (c *CPU) LD_A_A() {
	// 0x7F Load A with A register
	c.A = uint8(c.A)
	c.PC++
}

func (c *CPU) ADD_B() {
	// 0x80 Add B
	res, flags := c.add8(c.A, c.B, 0)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) ADD_C() {
	// 0x81 Add C
	res, flags := c.add8(c.A, c.C, 0)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) ADD_D() {
	// 0x82 Add D
	res, flags := c.add8(c.A, c.D, 0)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) ADD_E() {
	// 0x83 Add E
	res, flags := c.add8(c.A, c.E, 0)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) ADD_H() {
	// 0x84 Add H
	res, flags := c.add8(c.A, c.H, 0)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) ADD_L() {
	// 0x85 Add L
	res, flags := c.add8(c.A, c.L, 0)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) ADD_ADDR_HL() {
	// 0x86 Add [HL]
	hl := c.read_r16(HL)
	res, flags := c.add8(c.A, c.Bus.Read(hl), 0)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) ADD_A() {
	// 0x87 Add A
	res, flags := c.add8(c.A, c.A, 0)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) ADC_B() {
	// 0x88 Add B With carry
	carry := c.F & FLAG_C
	res, flags := c.add8(c.A, c.B, carry)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) ADC_C() {
	// 0x89 Add C With carry
	carry := c.F & FLAG_C
	res, flags := c.add8(c.A, c.C, carry)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) ADC_D() {
	// 0x8A Add D With carry
	carry := c.F & FLAG_C
	res, flags := c.add8(c.A, c.D, carry)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) ADC_E() {
	// 0x8B Add E With carry
	carry := c.F & FLAG_C
	res, flags := c.add8(c.A, c.E, carry)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) ADC_H() {
	// 0x8C Add H With carry
	carry := c.F & FLAG_C
	res, flags := c.add8(c.A, c.H, carry)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) ADC_L() {
	// 0x8D Add L With carry
	carry := c.F & FLAG_C
	res, flags := c.add8(c.A, c.L, carry)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) ADC_ADDR_HL() {
	// 0x8E Add [HL] With carry
	carry := c.F & FLAG_C
	hl := c.read_r16(HL)
	res, flags := c.add8(c.A, c.Bus.Read(hl), carry)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) ADC_A() {
	// 0x8F Add A with carry
	carry := c.F & FLAG_C
	res, flags := c.add8(c.A, c.A, carry)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) SUB_B() {
	// 0x90 Subtract B
	res, flags := c.sub8(c.A, c.B, 0)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) SUB_C() {
	// 0x91 Subtract C
	res, flags := c.sub8(c.A, c.C, 0)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) SUB_D() {
	// 0x92 Subtract D
	res, flags := c.sub8(c.A, c.D, 0)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) SUB_E() {
	// 0x93 Subtract E
	res, flags := c.sub8(c.A, c.E, 0)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) SUB_H() {
	// 0x94 Subtract H
	res, flags := c.sub8(c.A, c.H, 0)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) SUB_L() {
	// 0x95 Subtract L
	res, flags := c.sub8(c.A, c.L, 0)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) SUB_ADDR_HL() {
	// 0x96 Subtract [HL]
	hl := c.read_r16(HL)
	res, flags := c.sub8(c.A, c.Bus.Read(hl), 0)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) SUB_A() {
	// 0x97 Subtract A
	res, flags := c.sub8(c.A, c.A, 0)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) SBC_B() {
	// 0x98 Subtract B with carry
	carry := c.F & FLAG_C
	res, flags := c.sub8(c.A, c.B, carry)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) SBC_C() {
	// 0x99 Subtract C with carry
	carry := c.F & FLAG_C
	res, flags := c.sub8(c.A, c.C, carry)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) SBC_D() {
	// 0x9A Subtract D with carry
	carry := c.F & FLAG_C
	res, flags := c.sub8(c.A, c.D, carry)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) SBC_E() {
	// 0x9B Subtract E with carry
	carry := c.F & FLAG_C
	res, flags := c.sub8(c.A, c.E, carry)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) SBC_H() {
	// 0x9C Subtract H with carry
	carry := c.F & FLAG_C
	res, flags := c.sub8(c.A, c.H, carry)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) SBC_L() {
	// 0x9D Subtract L with carry
	carry := c.F & FLAG_C
	res, flags := c.sub8(c.A, c.L, carry)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) SBC_ADDR_HL() {
	// 0x9E Subtract [HL] with carry
	carry := c.F & FLAG_C
	hl := c.read_r16(HL)
	res, flags := c.sub8(c.A, c.Bus.Read(hl), carry)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) SBC_A() {
	// 0x9F Subtract A with carry
	carry := c.F & FLAG_C
	res, flags := c.sub8(c.A, c.A, carry)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) AND_B() {
	// 0xA0 And B
	c.A &= c.B
	c.Flag_set(FLAG_H)
	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_C)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) AND_C() {
	// 0xA1 And C
	c.A &= c.C
	c.Flag_set(FLAG_H)
	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_C)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) AND_D() {
	// 0xA2 And D
	c.A &= c.D
	c.Flag_set(FLAG_H)
	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_C)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) AND_E() {
	// 0xA3 And E
	c.A &= c.E
	c.Flag_set(FLAG_H)
	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_C)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) AND_H() {
	// 0xA4 And H
	c.A &= c.H
	c.Flag_set(FLAG_H)
	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_C)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) AND_L() {
	// 0xA5 And L
	c.A &= c.L
	c.Flag_set(FLAG_H)
	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_C)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) AND_ADDR_HL() {
	// 0xA6 And [HL]
	hl := c.read_r16(HL)
	c.A &= c.Bus.Read(hl)
	c.Flag_set(FLAG_H)
	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_C)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) AND_A() {
	// 0xA7 And A
	c.A &= c.A
	c.Flag_set(FLAG_H)
	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_C)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) XOR_B() {
	// 0xA8 Xor B
	c.A ^= c.B

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H | FLAG_C)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) XOR_C() {
	// 0xA9 Xor C
	c.A ^= c.C

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H | FLAG_C)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) XOR_D() {
	// 0xAA Xor D
	c.A ^= c.D

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H | FLAG_C)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) XOR_E() {
	// 0xAB Xor E
	c.A ^= c.E

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H | FLAG_C)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) XOR_H() {
	// 0xAC Xor H
	c.A ^= c.H

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H | FLAG_C)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) XOR_L() {
	// 0xAD Xor L
	c.A ^= c.L

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H | FLAG_C)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) XOR_ADDR_HL() {
	// 0xAE Xor [HL]
	hl := c.read_r16(HL)
	c.A ^= c.Bus.Read(hl)

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H | FLAG_C)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) XOR_A() {
	// 0xAF Xor A
	c.A ^= c.A

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H | FLAG_C)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) OR_B() {
	// 0xB0 Or B
	c.A |= c.B

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H | FLAG_C)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) OR_C() {
	// 0xB1 Or C
	c.A |= c.C

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H | FLAG_C)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) OR_D() {
	// 0xB2 Or D
	c.A |= c.D

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H | FLAG_C)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) OR_E() {
	// 0xB3 Or E
	c.A |= c.E

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H | FLAG_C)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) OR_H() {
	// 0xB4 Or H
	c.A |= c.H

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H | FLAG_C)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) OR_L() {
	// 0xB5 Or L
	c.A |= c.L

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H | FLAG_C)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) OR_ADDR_HL() {
	// 0xB6 Or [HL]
	hl := c.read_r16(HL)
	c.A |= c.Bus.Read(hl)

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H | FLAG_C)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) OR_A() {
	// 0xB7 Or A
	c.A |= c.A

	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_H | FLAG_C)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) CP_B() {
	// 0xB8 Compare B
	_, flags := c.sub8(c.A, c.B, 0)
	c.F = flags
	c.PC++
}

func (c *CPU) CP_C() {
	// 0xB9 Compare C
	_, flags := c.sub8(c.A, c.C, 0)
	c.F = flags
	c.PC++
}

func (c *CPU) CP_D() {
	// 0xBA Compare D
	_, flags := c.sub8(c.A, c.D, 0)
	c.F = flags
	c.PC++
}

func (c *CPU) CP_E() {
	// 0xBB Compare E
	_, flags := c.sub8(c.A, c.E, 0)
	c.F = flags
	c.PC++
}

func (c *CPU) CP_H() {
	// 0xBC Compare H
	_, flags := c.sub8(c.A, c.H, 0)
	c.F = flags
	c.PC++
}

func (c *CPU) CP_L() {
	// 0xBD Compare L
	_, flags := c.sub8(c.A, c.L, 0)
	c.F = flags
	c.PC++
}

func (c *CPU) CP_ADDR_HL() {
	// 0xBE Compare [HL]
	hl := c.read_r16(HL)
	_, flags := c.sub8(c.A, c.Bus.Read(hl), 0)
	c.F = flags
	c.PC++
}

func (c *CPU) CP_A() {
	// 0xBF Compare A
	_, flags := c.sub8(c.A, c.A, 0)
	c.F = flags
	c.PC++
}

func (c *CPU) RET_NZ() {
	// 0xC0 Return from subroutine if flag Z not set
	if c.F&FLAG_Z != 0 {
		c.PC++
		return
	}
	c.PC = c.pop()
}

func (c *CPU) POP_BC() {
	// 0xC1 Pop stack into BC
	w := c.pop()
	c.write_r16(BC, w)
	c.PC++
}

func (c *CPU) JP_NZ_a16() {
	// 0xC2 Jump to a16 if flag Z not set
	des := c.read_word()

	if c.F&FLAG_Z != 0 {
		c.PC++
		return
	}
	c.PC = des
}

func (c *CPU) JP_a16() {
	// 0xC3 Jump to a16
	des := c.read_word()
	c.PC = des
}

func (c *CPU) CALL_NZ_a16() {
	// 0xC4 Call subroutine at address a16 if Z flag not set
	subr := c.read_word()
	c.PC++

	if c.F&FLAG_Z != 0 {
		return
	}

	c.push(c.PC)
	c.PC = subr
}

func (c *CPU) PUSH_BC() {
	// 0xC5 Push BC register onto the stack
	bc := c.read_r16(BC)
	c.push(bc)
	c.PC++
}

func (c *CPU) ADD_n8() {
	// 0xC6 Add byte
	n := c.read_byte()
	res, flags := c.add8(c.A, n, 0)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) RST_00() {
	// 0xC7 Jump to vector 0x0000
	c.PC++
	c.push(c.PC)
	c.PC = 0x0000
}

func (c *CPU) RET_Z() {
	// 0xC8 Return from subroutine if flag Z is set
	if c.F&FLAG_Z == 0 {
		c.PC++
		return
	}
	c.PC = c.pop()
}

func (c *CPU) RET() {
	// 0xC9 Return from subroutine
	c.PC = c.pop()
}

func (c *CPU) JP_Z_a16() {
	// 0xCA Jump to address a16 if Z flag is set
	des := c.read_word()

	if c.F&FLAG_Z == 0 { // if Z is not set, continue
		c.PC++
		return
	}
	c.PC = des
}

func (c *CPU) PREFIX() {
	// 0xCB prefixed instructions, read next byte and perform operation from prefixed table
	op := c.read_byte()
	c.Prefixed_Operations[op].Exec()
}

func (c *CPU) CALL_Z_a16() {
	// 0xCC Call subroutine at address a16 if flag Z is set
	subr := c.read_word()
	c.PC++

	if c.F&FLAG_Z == 0 {
		return
	}

	c.push(c.PC)
	c.PC = subr
}

func (c *CPU) CALL_a16() {
	// 0xCD Call subroutine at address a16
	subr := c.read_word()
	c.PC++
	c.push(c.PC)
	c.PC = subr
}

func (c *CPU) ADC_n8() {
	// 0xCE Add byte with carry
	n := c.read_byte()
	carry := c.F & FLAG_C
	res, flags := c.add8(c.A, n, carry)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) RST_08() {
	// 0xCF Jump to vector 0x0008
	c.PC++
	c.push(c.PC)
	c.PC = 0x0008
}

func (c *CPU) RET_NC() {
	// 0xD0 Return from subroutine if flag C not set
	if c.F&FLAG_C != 0 {
		c.PC++
		return
	}
	c.PC = c.pop()
}

func (c *CPU) POP_DE() {
	// 0xD1 Pop stack into DE
	w := c.pop()
	c.write_r16(DE, w)
	c.PC++
}

func (c *CPU) JP_NC_a16() {
	// 0xD2 Jump to a16 if flag C not set
	des := c.read_word()

	if c.F&FLAG_C != 0 {
		c.PC++
		return
	}
	c.PC = des
}

func (c *CPU) CALL_NC_a16() {
	// 0xD4 Call subroutine at address a16 if C flag not set
	subr := c.read_word()
	c.PC++

	if c.F&FLAG_C != 0 {
		return
	}

	c.push(c.PC)
	c.PC = subr
}

func (c *CPU) PUSH_DE() {
	// 0xD5 Push DE register onto the stack
	de := c.read_r16(DE)
	c.push(de)
	c.PC++
}

func (c *CPU) SUB_n8() {
	// 0xD6 Subtract byte from A
	n := c.read_byte()
	res, flags := c.sub8(c.A, n, 0)

	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) RST_10() {
	// 0xD7 Jump to vector 0x0010
	c.PC++
	c.push(c.PC)
	c.PC = 0x0000
}

func (c *CPU) RET_C() {
	// 0xD8 Return from subroutine if flag C is set
	if c.F&FLAG_C == 0 {
		c.PC++
		return
	}
	c.PC = c.pop()
}

func (c *CPU) RETI() {
	// 0xD9 Return from subroutine and enable interrupts
	//TODO Add interrupt handling
	c.PC = c.pop()
}

func (c *CPU) JP_C_a16() {
	// 0xDA Jump to address a16 if C flag is set
	des := c.read_word()

	if c.F&FLAG_C == 0 { // if C is not set, continue
		c.PC++
		return
	}
	c.PC = des
}

func (c *CPU) CALL_C_a16() {
	// 0xDC Call subroutine if C flag is set
	subr := c.read_word()
	c.PC++

	if c.F&FLAG_C == 0 {
		return
	}

	c.push(c.PC)
	c.PC = subr
}

func (c *CPU) SBC_n8() {
	// 0xDE Subtract byte with carry
	n := c.read_byte()
	carry := c.F & FLAG_C

	res, flags := c.sub8(c.A, n, carry)
	c.A = res
	c.F = flags
	c.PC++
}

func (c *CPU) RST_18() {
	// 0xDF Jump to vector 0x0018
	c.PC++
	c.push(c.PC)
	c.PC = 0x0018
}

func (c *CPU) LDH_ADDR_a8_A() {
	// 0xE0 Store A register to address in High Ram offset by a8
	n := c.read_byte()
	addr := 0xff00 | uint16(n)
	c.Bus.Write(addr, c.A)
	c.PC++
}

func (c *CPU) POP_HL() {
	// 0xE1 Pop stack into HL register
	s := c.pop()
	c.write_r16(HL, s)
	c.PC++
}

func (c *CPU) LD_ADDR_C_A() {
	// 0xE2 Store A register to address in High Ram offset by C
	addr := 0xff00 | uint16(c.C)
	c.Bus.Write(addr, c.A)
	c.PC++
}

func (c *CPU) PUSH_HL() {
	// 0xE5 Push contents of HL on stack
	hl := c.read_r16(HL)
	c.push(hl)
	c.PC++
}

func (c *CPU) AND_n8() {
	// 0xE6 And byte
	n := c.read_byte()
	c.A &= n
	c.Flag_set(FLAG_H)
	c.Flag_reset(FLAG_Z | FLAG_N | FLAG_C)
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RST_20() {
	// 0xE7 Jump to vector 0x0020
	c.PC++
	c.push(c.PC)
	c.PC = 0x0020
}

func (c *CPU) ADD_SP_e8() {
	// 0xE8 Add the signed value e8 to SP
	e := c.read_byte()
	_, flags := c.add8(uint8(c.SP), e, 0)
	res, _ := c.add16(c.SP, uint16(int8(e))) // uint16(int8(e)): convert byte to signed 8bit to unsigned 16bit to preserve 2s complement
	c.SP = res
	c.F = (FLAG_H | FLAG_C) & flags // 0011 0000 -> keep H and C flags only
	c.PC++
}

func (c *CPU) JP_HL() {
	// 0xE9 Jump to address HL
	hl := c.read_r16(HL)
	c.PC = hl
}

func (c *CPU) LD_ADDR_a16_A() {
	// 0xEA Load contents of A to address a16
	addr := c.read_word()
	c.Bus.Write(addr, c.A)
	c.PC++
}

func (c *CPU) XOR_n8() {
	// 0xEE Xor byte
	n := c.read_byte()
	c.A ^= n
	c.F = 0
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RST_28() {
	// 0xEF Jump to vector 0x0028
	c.PC++
	c.push(c.PC)
	c.PC = 0x0028
}

func (c *CPU) LDH_A_ADDR_a8() {
	// 0xF0 Load into A contents of high ram byte a8
	a := c.read_byte()
	addr := 0xff00 | uint16(a)
	c.A = c.Bus.Read(addr)
	c.PC++
}

func (c *CPU) POP_AF() {
	// 0xF1 Pop stack into AF register
	s := c.pop()
	c.write_r16(AF, s)
	c.PC++
}

func (c *CPU) LD_A_ADDR_C() {
	// 0xF2 Load into A contents of high ram byte C
	addr := 0xff00 | uint16(c.C)
	c.A = c.Bus.Read(addr)
	c.PC++
}

func (c *CPU) DI() {
	// 0xF3 Disable Interrupt
	//TODO Add interrupt handling
	c.PC++
}

func (c *CPU) PUSH_AF() {
	// 0xF5 Push contents of AF onto stack
	af := c.read_r16(AF)
	c.push(af)
	c.PC++
}

func (c *CPU) OR_n8() {
	// 0xF6 Or byte
	n := c.read_byte()
	c.A |= n
	c.F = 0
	if c.A == 0 {
		c.Flag_set(FLAG_Z)
	}
	c.PC++
}

func (c *CPU) RST_30() {
	// 0xF7 Jump to vector 0x0030
	c.PC++
	c.push(c.PC)
	c.PC = 0x0030
}

func (c *CPU) LD_HL_SP_PLUS_e8() {
	// 0xF8 Add the signed value e8 to SP and store the result in HL
	e := c.read_byte()
	_, flags := c.add8(uint8(c.SP), e, 0)
	res, _ := c.add16(c.SP, uint16(int8(e))) // uint16(int8(e)): convert byte to signed 8bit to unsigned 16bit to preserve 2s complement
	c.write_r16(HL, res)
	c.F = (FLAG_H | FLAG_C) & flags // 0011 0000 -> keep H and C flags only
	c.PC++
}

func (c *CPU) LD_SP_HL() {
	// 0xF9 Load HL to SP
	hl := c.read_r16(HL)
	c.SP = hl
	c.PC++
}

func (c *CPU) LD_A_ADDR_a16() {
	// 0xFA Load contents of address a16 to A
	addr := c.read_word()
	c.A = c.Bus.Read(addr)
	c.PC++
}

func (c *CPU) EI() {
	// 0xFB Enable interrupt
	//TODO Add interrupt handling
	c.PC++
}

func (c *CPU) CP_n8() {
	// 0xFE Compare A with byte
	n := c.read_byte()
	_, flags := c.sub8(c.A, n, 0)
	c.F = flags
	c.PC++
}

func (c *CPU) RST_38() {
	// 0xF7 Jump to vector 0x0038
	c.PC++
	c.push(c.PC)
	c.PC = 0x0038
}
