package hardware

import (
	"log"
)

type Reg8 = byte
type Reg16 = uint16

type OPCODE = byte
type COMPOSITE_REGISTER_ID int

const (
	BC COMPOSITE_REGISTER_ID = iota
	DE
	HL
	AF
)

const (
	FLAG_Z = uint8(1 << 7) // Zero flag
	FLAG_N = uint8(1 << 6) // Subtraction flag
	FLAG_H = uint8(1 << 5) // Half carry flag
	FLAG_C = uint8(1 << 4) // Carry flag
)

type OPERATION struct {
	Mneumonic string
	Exec      func()
	Length    uint8
	T_States  uint8
	Flags     string
}

type CPU_STATUS struct {
	Stopped           bool
	Halted            bool
	Stepping          bool
	Interrupt_Enabled bool
}

type EXECUTION_INFO struct {
	Opcode      byte
	Address     uint16
	AddressMode uint8
	Instruction func()
	Operands    []interface{}
}

type CPU struct {
	A                   Reg8 // Accumulator
	F                   Reg8 // Flags
	B                   Reg8
	C                   Reg8
	D                   Reg8
	E                   Reg8
	H                   Reg8
	L                   Reg8
	SP                  Reg16 // Stack Pointer
	PC                  Reg16 // Program Counter
	Bus                 *Memory
	Operations          map[OPCODE]OPERATION
	Prefixed_Operations map[OPCODE]OPERATION

	Status   CPU_STATUS
	ExecInfo EXECUTION_INFO
}

var cpuInstance *CPU

var Operations map[OPCODE]OPERATION = map[OPCODE]OPERATION{}

func GetCPU() *CPU {
	if cpuInstance != nil {
		return cpuInstance
	}

	log.Println("Creating CPU Instance")
	cpuInstance = &CPU{Bus: GetBus()}
	cpuInstance.initFuncTable()
	cpuInstance.PC = 0x0100
	return cpuInstance
}

func (c *CPU) initFuncTable() {
	c.Operations = map[OPCODE]OPERATION{
		0x00: {"nop", c.NOP, 1, 4, "----"}, 0x01: {"ld BC,n16", c.LD_BC_n16, 3, 12, "----"}, 0x02: {"ld [BC],A", c.LD_ADDR_BC_A, 1, 8, "----"}, 0x03: {"inc BC", c.INC_BC, 1, 8, "----"},
		0x04: {"inc B", c.INC_B, 1, 4, "Z0H-"}, 0x05: {"dec B", c.DEC_B, 1, 4, "Z1H-"}, 0x06: {"ld B,n8", c.LD_B_n8, 2, 8, "----"}, 0x07: {"rlca", c.RLCA, 1, 4, "000C"},
		0x08: {"ld [a16],SP", c.LD_ADDR_a16_SP, 3, 20, "----"}, 0x09: {"add HL,BC", c.ADD_HL_BC, 1, 8, "-0HC"}, 0x0A: {"ld A,[BC]", c.LD_A_ADDR_BC, 1, 8, "----"}, 0x0B: {"dec BC", c.DEC_BC, 1, 8, "----"},
		0x0C: {"inc C", c.INC_C, 1, 4, "Z0H-"}, 0x0D: {"dec C", c.DEC_C, 1, 8, "Z1H-"}, 0x0E: {"ld C,n8", c.LD_C_n8, 2, 8, "----"}, 0x0F: {"rrca", c.RRCA, 1, 4, "000C"},

		0x10: {"stop", c.STOP, 2, 4, "----"}, 0x11: {"ld DE,n16", c.LD_DE_n16, 3, 12, "----"}, 0x12: {"ld [DE],A", c.LD_ADDR_DE_A, 1, 8, "----"}, 0x13: {"inc DE", c.INC_DE, 1, 8, "----"},
		0x14: {"inc D", c.INC_D, 1, 4, "Z0H-"}, 0x15: {"dec D", c.DEC_D, 1, 4, "Z1H-"}, 0x16: {"ld D,n8", c.LD_D_n8, 2, 8, "----"}, 0x17: {"rla", c.RLA, 1, 4, "000C"},
		0x18: {"jr e8", c.JR_e8, 2, 12, "----"}, 0x19: {"add HL,DE", c.ADD_HL_DE, 1, 8, "-0HC"}, 0x1A: {"ld A,[DE]", c.LD_A_ADDR_DE, 1, 8, "----"}, 0x1B: {"dec DE", c.DEC_DE, 1, 8, "----"},
		0x1C: {"inc E", c.INC_E, 1, 4, "Z0H-"}, 0x1D: {"dec E", c.DEC_E, 1, 4, "Z1H-"}, 0x1E: {"ld E,n8", c.LD_E_n8, 2, 8, "----"}, 0x1F: {"rra", c.RRA, 1, 4, "000C"},

		0x20: {"jr nz,e8", c.JR_NZ_e8, 2, 12, "----"}, 0x21: {"ld HL,n16", c.LD_HL_n16, 3, 12, "----"}, 0x22: {"ld [HL+],A", c.LDI_ADDR_HL_A, 1, 8, "----"}, 0x23: {"inc HL", c.INC_HL, 1, 8, "----"},
		0x24: {"inc H", c.INC_H, 1, 4, "Z0H-"}, 0x25: {"dec H", c.DEC_H, 1, 4, "Z1H-"}, 0x26: {"ld H,n8", c.LD_H_n8, 2, 8, "----"}, 0x27: {"daa", c.DAA, 1, 4, "Z-0C"},
		0x28: {"jr Z,e8", c.JR_Z_e8, 2, 12, "----"}, 0x29: {"add HL,HL", c.ADD_HL_HL, 1, 8, "-0HC"}, 0x2A: {"ldi A,[HL]", c.LDI_A_ADDR_HL, 1, 8, "----"}, 0x2B: {"dec HL", c.DEC_HL, 1, 8, "----"},
		0x2C: {"inc L", c.INC_L, 1, 4, "Z0H-"}, 0x2D: {"dec L", c.DEC_L, 1, 4, "Z1H-"}, 0x2E: {"ld L,n8", c.LD_L_n8, 2, 8, "----"}, 0x2F: {"cpl", c.CPL, 1, 4, "-11-"},

		0x30: {"jr nc,e8", c.JR_NC_e8, 2, 12, "----"}, 0x31: {"ld sp,n16", c.LD_SP_n16, 3, 12, "----"}, 0x32: {"ldd [HL],A", c.LDD_ADDR_HL_A, 1, 8, "----"}, 0x33: {"inc SP", c.INC_SP, 1, 8, "----"},
		0x34: {"inc [HL]", c.INC_ADDR_HL, 1, 12, "Z0H-"}, 0x35: {"dec [HL]", c.DEC_ADDR_HL, 1, 12, "Z1H-"}, 0x36: {"ld [HL],n8", c.LD_ADDR_HL_n8, 2, 12, "----"}, 0x37: {"scf", c.SCF, 1, 4, "-001"},
		0x38: {"jr c,e8", c.JR_C_e8, 2, 12, "----"}, 0x39: {"add HL,HL", c.ADD_HL_SP, 1, 8, "-0HC"}, 0x3A: {"ldd A,[HL]", c.LDD_A_ADDR_HL, 1, 8, "----"}, 0x3B: {"dec SP", c.DEC_SP, 1, 8, "----"},
		0x3C: {"inc A", c.INC_A, 1, 4, "Z0H-"}, 0x3D: {"dec A", c.DEC_A, 1, 4, "Z1H-"}, 0x3E: {"ld A,n8", c.LD_A_n8, 2, 8, "----"}, 0x3F: {"cff", c.CCF, 1, 4, "-00C"},

		0x40: {"ld B,B", c.LD_B_B, 1, 4, "----"}, 0x41: {"ld B,C", c.LD_B_C, 1, 4, "----"}, 0x42: {"ld B,D", c.LD_B_D, 1, 4, "----"}, 0x43: {"ld B,E", c.LD_B_E, 1, 4, "----"},
		0x44: {"ld B,H", c.LD_B_H, 1, 4, "----"}, 0x45: {"ld B,L", c.LD_B_L, 1, 4, "----"}, 0x46: {"ld B,[HL]", c.LD_B_ADDR_HL, 1, 8, "----"}, 0x47: {"ld B,A", c.LD_B_A, 1, 4, "----"},
		0x48: {"ld C,B", c.LD_C_B, 1, 4, "----"}, 0x49: {"ld C,C", c.LD_C_C, 1, 4, "----"}, 0x4A: {"ld C,D", c.LD_C_D, 1, 4, "----"}, 0x4B: {"ld C,E", c.LD_C_E, 1, 4, "----"},
		0x4C: {"ld C,H", c.LD_C_H, 1, 4, "----"}, 0x4D: {"ld C,L", c.LD_C_L, 1, 4, "----"}, 0x4E: {"ld C,[HL]", c.LD_C_ADDR_HL, 1, 8, "----"}, 0x4F: {"ld C,A", c.LD_C_A, 1, 4, "----"},

		0x50: {"ld D,B", c.LD_D_B, 1, 4, "----"}, 0x51: {"ld D,C", c.LD_D_C, 1, 4, "----"}, 0x52: {"ld D,D", c.LD_D_D, 1, 4, "----"}, 0x53: {"ld D,E", c.LD_D_E, 1, 4, "----"},
		0x54: {"ld D,H", c.LD_D_H, 1, 4, "----"}, 0x55: {"ld D,L", c.LD_D_L, 1, 4, "----"}, 0x56: {"ld D,[HL]", c.LD_D_ADDR_HL, 1, 8, "----"}, 0x57: {"ld D,A", c.LD_D_A, 1, 4, "----"},
		0x58: {"ld E,B", c.LD_E_B, 1, 4, "----"}, 0x59: {"ld E,C", c.LD_E_C, 1, 4, "----"}, 0x5A: {"ld E,D", c.LD_E_D, 1, 4, "----"}, 0x5B: {"ld E,E", c.LD_E_E, 1, 4, "----"},
		0x5C: {"ld E,H", c.LD_E_H, 1, 4, "----"}, 0x5D: {"ld E,L", c.LD_E_L, 1, 4, "----"}, 0x5E: {"ld E,[HL]", c.LD_E_ADDR_HL, 1, 8, "----"}, 0x5F: {"ld E,A", c.LD_E_A, 1, 4, "----"},

		0x60: {"ld H,B", c.LD_H_B, 1, 4, "----"}, 0x61: {"ld H,C", c.LD_H_C, 1, 4, "----"}, 0x62: {"ld H,D", c.LD_H_D, 1, 4, "----"}, 0x63: {"ld H,E", c.LD_H_E, 1, 4, "----"},
		0x64: {"ld H,H", c.LD_H_H, 1, 4, "----"}, 0x65: {"ld H,L", c.LD_H_L, 1, 4, "----"}, 0x66: {"ld H,[HL]", c.LD_H_ADDR_HL, 1, 8, "----"}, 0x67: {"ld H,A", c.LD_H_A, 1, 4, "----"},
		0x68: {"ld L,B", c.LD_L_B, 1, 4, "----"}, 0x69: {"ld L,C", c.LD_L_C, 1, 4, "----"}, 0x6A: {"ld L,D", c.LD_L_D, 1, 4, "----"}, 0x6B: {"ld L,E", c.LD_L_E, 1, 4, "----"},
		0x6C: {"ld L,H", c.LD_L_H, 1, 4, "----"}, 0x6D: {"ld L,L", c.LD_L_L, 1, 4, "----"}, 0x6E: {"ld L,[HL]", c.LD_L_ADDR_HL, 1, 8, "----"}, 0x6F: {"ld L,A", c.LD_L_A, 1, 4, "----"},

		0x70: {"ld [HL],B", c.LD_ADDR_HL_B, 1, 8, "----"}, 0x71: {"ld [HL],C", c.LD_ADDR_HL_C, 1, 8, "----"}, 0x72: {"ld [HL],D", c.LD_ADDR_HL_D, 1, 8, "----"}, 0x73: {"ld [HL],E", c.LD_ADDR_HL_E, 1, 8, "----"},
		0x74: {"ld [HL],H", c.LD_ADDR_HL_H, 1, 8, "----"}, 0x75: {"ld [HL],L", c.LD_ADDR_HL_L, 1, 8, "----"}, 0x76: {"halt", c.HALT, 1, 4, "----"}, 0x77: {"ld [HL],A", c.LD_ADDR_HL_A, 1, 8, "----"},
		0x78: {"ld A,B", c.LD_A_B, 1, 4, "----"}, 0x79: {"ld A,C", c.LD_A_C, 1, 4, "----"}, 0x7A: {"ld A,D", c.LD_A_D, 1, 4, "----"}, 0x7B: {"ld A,E", c.LD_A_E, 1, 4, "----"},
		0x7C: {"ld A,H", c.LD_A_H, 1, 4, "----"}, 0x7D: {"ld A,L", c.LD_A_L, 1, 4, "----"}, 0x7E: {"ld A,[HL]", c.LD_A_ADDR_HL, 1, 8, "----"}, 0x7F: {"ld A,A", c.LD_A_A, 1, 4, "----"},

		0x80: {"add B", c.ADD_B, 1, 4, "Z0HC"}, 0x81: {"add C", c.ADD_C, 1, 4, "Z0HC"}, 0x82: {"add D", c.ADD_D, 1, 4, "Z0HC"}, 0x83: {"add E", c.ADD_E, 1, 4, "Z0HC"},
		0x84: {"add H", c.ADD_H, 1, 4, "Z0HC"}, 0x85: {"add L", c.ADD_L, 1, 4, "Z0HC"}, 0x86: {"add [HL]", c.ADD_ADDR_HL, 1, 8, "Z0HC"}, 0x87: {"add A", c.ADD_A, 1, 4, "Z0HC"},
		0x88: {"adc B", c.ADC_B, 1, 4, "Z0HC"}, 0x89: {"adc C", c.ADC_C, 1, 4, "Z0HC"}, 0x8A: {"adc D", c.ADC_D, 1, 4, "Z0HC"}, 0x8B: {"adc E", c.ADC_E, 1, 4, "Z0HC"},
		0x8C: {"adc H", c.ADC_H, 1, 4, "Z0HC"}, 0x8D: {"adc L", c.ADC_L, 1, 4, "Z0HC"}, 0x8E: {"adc [HL]", c.ADC_ADDR_HL, 1, 8, "Z0HC"}, 0x8F: {"adc A", c.ADC_A, 1, 4, "Z0HC"},

		0x90: {"sub B", c.SUB_B, 1, 4, "Z1HC"}, 0x91: {"sub C", c.SUB_C, 1, 4, "Z1HC"}, 0x92: {"sub D", c.SUB_D, 1, 4, "Z1HC"}, 0x93: {"sub E", c.SUB_E, 1, 4, "Z1HC"},
		0x94: {"sub H", c.SUB_H, 1, 4, "Z1HC"}, 0x95: {"sub L", c.SUB_L, 1, 4, "Z1HC"}, 0x96: {"sub [HL]", c.SUB_ADDR_HL, 1, 8, "Z1HC"}, 0x97: {"sub A", c.SUB_A, 1, 4, "1100"},
		0x98: {"sbc B", c.SBC_B, 1, 4, "Z1HC"}, 0x99: {"sbc C", c.SBC_C, 1, 4, "Z1HC"}, 0x9A: {"sbc D", c.SBC_D, 1, 4, "Z1HC"}, 0x9B: {"sbc E", c.SBC_E, 1, 4, "Z1HC"},
		0x9C: {"sbc H", c.SBC_H, 1, 4, "Z1HC"}, 0x9D: {"sbc L", c.SBC_L, 1, 4, "Z1HC"}, 0x9E: {"sbc [HL]", c.SBC_ADDR_HL, 1, 8, "Z1HC"}, 0x9F: {"sbc A", c.SBC_A, 1, 4, "Z1H-"},

		0xA0: {"and B", c.AND_B, 1, 4, "Z010"}, 0xA1: {"and C", c.AND_C, 1, 4, "Z010"}, 0xA2: {"and D", c.AND_D, 1, 4, "Z010"}, 0xA3: {"and E", c.AND_E, 1, 4, "Z010"},
		0xA4: {"and H", c.AND_H, 1, 4, "Z010"}, 0xA5: {"and L", c.AND_L, 1, 4, "Z010"}, 0xA6: {"and [HL]", c.AND_ADDR_HL, 1, 8, "Z010"}, 0xA7: {"and A", c.AND_A, 1, 4, "Z010"},
		0xA8: {"xor B", c.XOR_B, 1, 4, "Z000"}, 0xA9: {"xor C", c.XOR_C, 1, 4, "Z000"}, 0xAA: {"xor D", c.XOR_D, 1, 4, "Z000"}, 0xAB: {"xor E", c.XOR_E, 1, 4, "Z000"},
		0xAC: {"xor H", c.XOR_H, 1, 4, "Z000"}, 0xAD: {"xor L", c.XOR_L, 1, 4, "Z000"}, 0xAE: {"xor [HL]", c.XOR_ADDR_HL, 1, 8, "Z000"}, 0xAF: {"xor A", c.XOR_A, 1, 4, "1000"},

		0xB0: {"or B", c.OR_B, 1, 4, "Z000"}, 0xB1: {"or C", c.OR_C, 1, 4, "Z000"}, 0xB2: {"or D", c.OR_D, 1, 4, "Z000"}, 0xB3: {"or E", c.OR_E, 1, 4, "Z000"},
		0xB4: {"or H", c.OR_H, 1, 4, "Z000"}, 0xB5: {"or L", c.OR_L, 1, 4, "Z000"}, 0xB6: {"or [HL]", c.OR_ADDR_HL, 1, 8, "Z000"}, 0xB7: {"or A", c.OR_A, 1, 4, "Z000"},
		0xB8: {"cp B", c.CP_B, 1, 4, "Z1HC"}, 0xB9: {"cp C", c.CP_C, 1, 4, "Z1HC"}, 0xBA: {"cp D", c.CP_D, 1, 4, "Z1HC"}, 0xBB: {"cp E", c.CP_E, 1, 4, "Z1HC"},
		0xBC: {"cp H", c.CP_H, 1, 4, "Z1HC"}, 0xBD: {"cp L", c.CP_L, 1, 4, "Z1HC"}, 0xBE: {"cp [HL]", c.CP_ADDR_HL, 1, 8, "Z1HC"}, 0xBF: {"cp A", c.CP_A, 1, 4, "1100"},

		0xC0: {"ret NZ", c.RET_NZ, 1, 20, "----"}, 0xC1: {"pop BC", c.POP_BC, 1, 12, "----"}, 0xC2: {"jp NZ,a16", c.JP_NZ_a16, 3, 16, "----"}, 0xC3: {"jp a16", c.JP_a16, 3, 16, "----"},
		0xC4: {"call NZ,a16", c.CALL_NZ_a16, 3, 24, "----"}, 0xC5: {"push BC", c.PUSH_BC, 1, 16, "----"}, 0xC6: {"add n8", c.ADD_n8, 2, 8, "Z0HC"}, 0xC7: {"rst 00", c.RST_00, 1, 16, "----"},
		0xC8: {"ret Z", c.RET_Z, 1, 20, "----"}, 0xC9: {"ret", c.RET, 1, 16, "----"}, 0xCA: {"jp Z,a16", c.JP_Z_a16, 3, 16, "----"}, 0xCB: {"prefix", c.PREFIX, 1, 4, "----"},
		0xCC: {"call Z,a16", c.CALL_Z_a16, 3, 24, "----"}, 0xCD: {"call a16", c.CALL_a16, 3, 24, "----"}, 0xCE: {"adc n8", c.ADC_n8, 2, 8, "Z0HC"}, 0xCF: {"rst 08", c.RST_08, 1, 16, "----"},

		0xD0: {"ret NC", c.RET_NC, 1, 20, "----"}, 0xD1: {"pop DE", c.POP_DE, 1, 12, "----"}, 0xD2: {"jp NC,a16", c.JP_NC_a16, 1, 16, "----"}, 0xD3: {"???", c.NOP, 1, 4, "----"},
		0xD4: {"call NC,a16", c.CALL_NC_a16, 3, 24, "----"}, 0xD5: {"push DE", c.PUSH_DE, 1, 16, "----"}, 0xD6: {"sub n8", c.SUB_n8, 2, 8, "Z1HC"}, 0xD7: {"rst 10", c.RST_10, 1, 16, "----"},
		0xD8: {"ret C", c.RET_C, 1, 20, "----"}, 0xD9: {"reti", c.RETI, 1, 16, "----"}, 0xDA: {"jp C,a16", c.JP_C_a16, 1, 16, "----"}, 0xDB: {"???", c.NOP, 1, 4, "----"},
		0xDC: {"call C,a16", c.CALL_C_a16, 3, 24, "----"}, 0xDD: {"???", c.NOP, 1, 4, "----"}, 0xDE: {"sbc n8", c.SBC_n8, 2, 8, "Z1HC"}, 0xDF: {"rst 18", c.RST_18, 1, 16, "----"},

		0xE0: {"ldh [a8],A", c.LDH_ADDR_a8_A, 2, 12, "----"}, 0xE1: {"pop HL", c.POP_HL, 1, 12, "----"}, 0xE2: {"ld [C],A", c.LD_ADDR_C_A, 1, 8, "----"}, 0xE3: {"???", c.NOP, 1, 4, "----"},
		0xE4: {"???", c.NOP, 1, 4, "----"}, 0xE5: {"push HL", c.PUSH_HL, 1, 16, "----"}, 0xE6: {"and n8", c.AND_n8, 2, 8, "Z010"}, 0xE7: {"rst 20", c.RST_20, 1, 16, "----"},
		0xE8: {"add SP,e8", c.ADD_SP_e8, 2, 16, "00HC"}, 0xE9: {"jp HL", c.JP_HL, 1, 4, "----"}, 0xEA: {"ld [a16],A", c.LD_ADDR_a16_A, 3, 16, "----"}, 0xEB: {"???", c.NOP, 1, 4, "----"},
		0xEC: {"???", c.NOP, 1, 4, "----"}, 0xED: {"???", c.NOP, 1, 4, "----"}, 0xEE: {"xor n8", c.XOR_n8, 2, 8, "Z000"}, 0xEF: {"rst 28", c.RST_28, 1, 16, "----"},

		0xF0: {"ldh A,[a8]", c.LDH_A_ADDR_a8, 2, 12, "----"}, 0xF1: {"pop AF", c.POP_AF, 1, 12, "ZNHC"}, 0xF2: {"ld A,[C]", c.LD_A_ADDR_C, 1, 8, "----"}, 0xF3: {"di", c.DI, 1, 4, "----"},
		0xF4: {"???", c.NOP, 1, 4, "----"}, 0xF5: {"push AF", c.PUSH_AF, 1, 16, "----"}, 0xF6: {"or n8", c.OR_n8, 2, 8, "Z000"}, 0xF7: {"rst 30", c.RST_30, 1, 16, "----"},
		0xF8: {"ld HL,SP+e8", c.LD_HL_SP_PLUS_e8, 2, 12, "00HC"}, 0xF9: {"ld SP,HL", c.LD_SP_HL, 1, 8, "----"}, 0xFA: {"ld A,[a16]", c.LD_A_ADDR_a16, 3, 16, "----"}, 0xFB: {"ei", c.EI, 1, 4, "----"},
		0xFC: {"???", c.NOP, 1, 4, "----"}, 0xFD: {"???", c.NOP, 1, 4, "----"}, 0xFE: {"cp n8", c.CP_n8, 2, 8, "Z1HC"}, 0xFF: {"rst 38", c.RST_38, 1, 16, "----"},
	}

	c.Prefixed_Operations = map[OPCODE]OPERATION{
		0x00: {"rlc B", c.RLC_B, 2, 8, "Z00C"}, 0x01: {"rlc C", c.RLC_C, 2, 8, "Z00C"}, 0x02: {"rlc D", c.RLC_D, 2, 8, "Z00C"}, 0x03: {"rlc E", c.RLC_E, 2, 8, "Z00C"},
		0x04: {"rlc H", c.RLC_H, 2, 8, "Z00C"}, 0x05: {"rlc L", c.RLC_L, 2, 8, "Z00C"}, 0x06: {"rlc [HL]", c.RLC_ADDR_HL, 2, 16, "Z00C"}, 0x07: {"rlc A", c.RLC_A, 2, 8, "Z00C"},
		0x08: {"rrc B", c.RRC_B, 2, 8, "Z00C"}, 0x09: {"rrc C", c.RRC_C, 2, 8, "Z00C"}, 0x0A: {"rrc D", c.RRC_D, 2, 8, "Z00C"}, 0x0B: {"rrc E", c.RRC_E, 2, 8, "Z00C"},
		0x0C: {"rrc H", c.RRC_H, 2, 8, "Z00C"}, 0x0D: {"rrc L", c.RRC_L, 2, 8, "Z00C"}, 0x0E: {"rrc [HL]", c.RRC_ADDR_HL, 2, 16, "Z00C"}, 0x0F: {"rrc A", c.RRC_A, 2, 8, "Z00C"},

		0x10: {"rl B", c.RL_B, 2, 8, "Z00C"}, 0x11: {"rl C", c.RL_C, 2, 8, "Z00C"}, 0x12: {"rl D", c.RL_D, 2, 8, "Z00C"}, 0x13: {"rl E", c.RL_E, 2, 8, "Z00C"},
		0x14: {"rl H", c.RL_H, 2, 8, "Z00C"}, 0x15: {"rl L", c.RL_L, 2, 8, "Z00C"}, 0x16: {"rl [HL]", c.RL_ADDR_HL, 2, 16, "Z00C"}, 0x17: {"rl A", c.RL_A, 2, 8, "Z00C"},
		0x18: {"rr B", c.RR_B, 2, 8, "Z00C"}, 0x19: {"rr C", c.RR_C, 2, 8, "Z00C"}, 0x1A: {"rr D", c.RR_D, 2, 8, "Z00C"}, 0x1B: {"rr E", c.RR_E, 2, 8, "Z00C"},
		0x1C: {"rr H", c.RR_H, 2, 8, "Z00C"}, 0x1D: {"rr L", c.RR_L, 2, 8, "Z00C"}, 0x1E: {"rr [HL]", c.RR_ADDR_HL, 2, 16, "Z00C"}, 0x1F: {"rr A", c.RR_A, 2, 8, "Z00C"},

		0x20: {"sla B", c.SLA_B, 2, 8, "Z00C"}, 0x21: {"sla C", c.SLA_C, 2, 8, "Z00C"}, 0x22: {"sla D", c.SLA_D, 2, 8, "Z00C"}, 0x23: {"sla E", c.SLA_E, 2, 8, "Z00C"},
		0x24: {"sla H", c.SLA_H, 2, 8, "Z00C"}, 0x25: {"sla L", c.SLA_L, 2, 8, "Z00C"}, 0x26: {"sla [HL]", c.SLA_ADDR_HL, 2, 16, "Z00C"}, 0x27: {"sla A", c.SLA_A, 2, 8, "Z00C"},
		0x28: {"sra B", c.SRA_B, 2, 8, "Z00C"}, 0x29: {"sra C", c.SRA_C, 2, 8, "Z00C"}, 0x2A: {"sra D", c.SRA_D, 2, 8, "Z00C"}, 0x2B: {"sra E", c.SRA_E, 2, 8, "Z00C"},
		0x2C: {"sra H", c.SRA_H, 2, 8, "Z00C"}, 0x2D: {"sra L", c.SRA_L, 2, 8, "Z00C"}, 0x2E: {"sra [HL]", c.SRA_ADDR_HL, 2, 16, "Z00C"}, 0x2F: {"sra A", c.SRA_A, 2, 8, "Z00C"},

		0x30: {"swap B", c.SWAP_B, 2, 8, "Z00C"}, 0x31: {"swap C", c.SWAP_C, 2, 8, "Z00C"}, 0x32: {"swap D", c.SWAP_D, 2, 8, "Z00C"}, 0x33: {"swap E", c.SWAP_E, 2, 8, "Z00C"},
		0x34: {"swap H", c.SWAP_H, 2, 8, "Z00C"}, 0x35: {"swap L", c.SWAP_L, 2, 8, "Z00C"}, 0x36: {"swap [HL]", c.SWAP_ADDR_HL, 2, 16, "Z00C"}, 0x37: {"swap A", c.SWAP_A, 2, 8, "Z00C"},
		0x38: {"srl B", c.SRL_B, 2, 8, "Z00C"}, 0x39: {"srl C", c.SRL_C, 2, 8, "Z00C"}, 0x3A: {"srl D", c.SRL_D, 2, 8, "Z00C"}, 0x3B: {"srl E", c.SRL_E, 2, 8, "Z00C"},
		0x3C: {"srl H", c.SRL_H, 2, 8, "Z00C"}, 0x3D: {"srl L", c.SRL_L, 2, 8, "Z00C"}, 0x3E: {"srl [HL]", c.SRL_ADDR_HL, 2, 16, "Z00C"}, 0x3F: {"srl A", c.SRL_A, 2, 8, "Z00C"},

		0x40: {"bit 0,B", c.BIT_0_B, 2, 8, "Z01-"}, 0x41: {"bit 0,C", c.BIT_0_C, 2, 8, "Z01-"}, 0x42: {"bit 0,D", c.BIT_0_D, 2, 8, "Z01-"}, 0x43: {"bit 0,E", c.BIT_0_E, 2, 8, "Z01-"},
		0x44: {"bit 0,H", c.BIT_0_H, 2, 8, "Z01-"}, 0x45: {"bit 0,L", c.BIT_0_L, 2, 8, "Z01-"}, 0x46: {"bit 0,[HL]", c.BIT_0_ADDR_HL, 2, 16, "Z01-"}, 0x47: {"bit 0,A", c.BIT_0_A, 2, 8, "Z01-"},
		0x48: {"bit 1,B", c.BIT_1_B, 2, 8, "Z01-"}, 0x49: {"bit 1,C", c.BIT_1_C, 2, 8, "Z01-"}, 0x4A: {"bit 1,D", c.BIT_1_D, 2, 8, "Z01-"}, 0x4B: {"bit 1,E", c.BIT_1_E, 2, 8, "Z01-"},
		0x4C: {"bit 1,H", c.BIT_1_H, 2, 8, "Z01-"}, 0x4D: {"bit 1,L", c.BIT_1_L, 2, 8, "Z01-"}, 0x4E: {"bit 1,[HL]", c.BIT_1_ADDR_HL, 2, 16, "Z01-"}, 0x4F: {"bit 1,A", c.BIT_1_A, 2, 8, "Z01-"},

		0x50: {"bit 2,B", c.BIT_2_B, 2, 8, "Z01-"}, 0x51: {"bit 2,C", c.BIT_2_C, 2, 8, "Z01-"}, 0x52: {"bit 2,D", c.BIT_2_D, 2, 8, "Z01-"}, 0x53: {"bit 2,E", c.BIT_2_E, 2, 8, "Z01-"},
		0x54: {"bit 2,H", c.BIT_2_H, 2, 8, "Z01-"}, 0x55: {"bit 2,L", c.BIT_2_L, 2, 8, "Z01-"}, 0x56: {"bit 2,[HL]", c.BIT_2_ADDR_HL, 2, 16, "Z01-"}, 0x57: {"bit 2,A", c.BIT_2_A, 2, 8, "Z01-"},
		0x58: {"bit 3,B", c.BIT_3_B, 2, 8, "Z01-"}, 0x59: {"bit 3,C", c.BIT_3_C, 2, 8, "Z01-"}, 0x5A: {"bit 3,D", c.BIT_3_D, 2, 8, "Z01-"}, 0x5B: {"bit 3,E", c.BIT_3_E, 2, 8, "Z01-"},
		0x5C: {"bit 3,H", c.BIT_3_H, 2, 8, "Z01-"}, 0x5D: {"bit 3,L", c.BIT_3_L, 2, 8, "Z01-"}, 0x5E: {"bit 3,[HL]", c.BIT_3_ADDR_HL, 2, 16, "Z01-"}, 0x5F: {"bit 3,A", c.BIT_3_A, 2, 8, "Z01-"},

		0x60: {"bit 4,B", c.BIT_4_B, 2, 8, "Z01-"}, 0x61: {"bit 4,C", c.BIT_4_C, 2, 8, "Z01-"}, 0x62: {"bit 4,D", c.BIT_4_D, 2, 8, "Z01-"}, 0x63: {"bit 4,E", c.BIT_4_E, 2, 8, "Z01-"},
		0x64: {"bit 4,H", c.BIT_4_H, 2, 8, "Z01-"}, 0x65: {"bit 4,L", c.BIT_4_L, 2, 8, "Z01-"}, 0x66: {"bit 4,[HL]", c.BIT_4_ADDR_HL, 2, 16, "Z01-"}, 0x67: {"bit 4,A", c.BIT_4_A, 2, 8, "Z01-"},
		0x68: {"bit 5,B", c.BIT_5_B, 2, 8, "Z01-"}, 0x69: {"bit 5,C", c.BIT_5_C, 2, 8, "Z01-"}, 0x6A: {"bit 5,D", c.BIT_5_D, 2, 8, "Z01-"}, 0x6B: {"bit 5,E", c.BIT_5_E, 2, 8, "Z01-"},
		0x6C: {"bit 5,H", c.BIT_5_H, 2, 8, "Z01-"}, 0x6D: {"bit 5,L", c.BIT_5_L, 2, 8, "Z01-"}, 0x6E: {"bit 5,[HL]", c.BIT_5_ADDR_HL, 2, 16, "Z01-"}, 0x6F: {"bit 5,A", c.BIT_5_A, 2, 8, "Z01-"},

		0x70: {"bit 6,B", c.BIT_6_B, 2, 8, "Z01-"}, 0x71: {"bit 6,C", c.BIT_6_C, 2, 8, "Z01-"}, 0x72: {"bit 6,D", c.BIT_6_D, 2, 8, "Z01-"}, 0x73: {"bit 6,E", c.BIT_6_E, 2, 8, "Z01-"},
		0x74: {"bit 6,H", c.BIT_6_H, 2, 8, "Z01-"}, 0x75: {"bit 6,L", c.BIT_6_L, 2, 8, "Z01-"}, 0x76: {"bit 6,[HL]", c.BIT_6_ADDR_HL, 2, 16, "Z01-"}, 0x77: {"bit 6,A", c.BIT_6_A, 2, 8, "Z01-"},
		0x78: {"bit 7,B", c.BIT_7_B, 2, 8, "Z01-"}, 0x79: {"bit 7,C", c.BIT_7_C, 2, 8, "Z01-"}, 0x7A: {"bit 7,D", c.BIT_7_D, 2, 8, "Z01-"}, 0x7B: {"bit 7,E", c.BIT_7_E, 2, 8, "Z01-"},
		0x7C: {"bit 7,H", c.BIT_7_H, 2, 8, "Z01-"}, 0x7D: {"bit 7,L", c.BIT_7_L, 2, 8, "Z01-"}, 0x7E: {"bit 7,[HL]", c.BIT_7_ADDR_HL, 2, 16, "Z01-"}, 0x7F: {"bit 7,A", c.BIT_7_A, 2, 8, "Z01-"},

		0x80: {"res 0,B", c.RES_0_B, 2, 8, "----"}, 0x81: {"res 0,C", c.RES_0_C, 2, 8, "----"}, 0x82: {"res 0,D", c.RES_0_D, 2, 8, "----"}, 0x83: {"res 0,E", c.RES_0_E, 2, 8, "----"},
		0x84: {"res 0,H", c.RES_0_H, 2, 8, "----"}, 0x85: {"res 0,L", c.RES_0_L, 2, 8, "----"}, 0x86: {"res 0,[HL]", c.RES_0_ADDR_HL, 2, 16, "----"}, 0x87: {"res 0,A", c.RES_0_A, 2, 8, "----"},
		0x88: {"res 1,B", c.RES_1_B, 2, 8, "----"}, 0x89: {"res 1,C", c.RES_1_C, 2, 8, "----"}, 0x8A: {"res 1,D", c.RES_1_D, 2, 8, "----"}, 0x8B: {"res 1,E", c.RES_1_E, 2, 8, "----"},
		0x8C: {"res 1,H", c.RES_1_H, 2, 8, "----"}, 0x8D: {"res 1,L", c.RES_1_L, 2, 8, "----"}, 0x8E: {"res 1,[HL]", c.RES_1_ADDR_HL, 2, 16, "----"}, 0x8F: {"res 1,A", c.RES_1_A, 2, 8, "----"},

		0x90: {"res 2,B", c.RES_2_B, 2, 8, "----"}, 0x91: {"res 2,C", c.RES_2_C, 2, 8, "----"}, 0x92: {"res 2,D", c.RES_2_D, 2, 8, "----"}, 0x93: {"res 2,E", c.RES_2_E, 2, 8, "----"},
		0x94: {"res 2,H", c.RES_2_H, 2, 8, "----"}, 0x95: {"res 2,L", c.RES_2_L, 2, 8, "----"}, 0x96: {"res 2,[HL]", c.RES_2_ADDR_HL, 2, 16, "----"}, 0x97: {"res 2,A", c.RES_2_A, 2, 8, "----"},
		0x98: {"res 3,B", c.RES_3_B, 2, 8, "----"}, 0x99: {"res 3,C", c.RES_3_C, 2, 8, "----"}, 0x9A: {"res 3,D", c.RES_3_D, 2, 8, "----"}, 0x9B: {"res 3,E", c.RES_3_E, 2, 8, "----"},
		0x9C: {"res 3,H", c.RES_3_H, 2, 8, "----"}, 0x9D: {"res 3,L", c.RES_3_L, 2, 8, "----"}, 0x9E: {"res 3,[HL]", c.RES_3_ADDR_HL, 2, 16, "----"}, 0x9F: {"res 3,A", c.RES_3_A, 2, 8, "----"},

		0xA0: {"res 4,B", c.RES_4_B, 2, 8, "----"}, 0xA1: {"res 4,C", c.RES_4_C, 2, 8, "----"}, 0xA2: {"res 4,D", c.RES_4_D, 2, 8, "----"}, 0xA3: {"res 4,E", c.RES_4_E, 2, 8, "----"},
		0xA4: {"res 4,H", c.RES_4_H, 2, 8, "----"}, 0xA5: {"res 4,L", c.RES_4_L, 2, 8, "----"}, 0xA6: {"res 4,[HL]", c.RES_4_ADDR_HL, 2, 16, "----"}, 0xA7: {"res 4,A", c.RES_4_A, 2, 8, "----"},
		0xA8: {"res 5,B", c.RES_5_B, 2, 8, "----"}, 0xA9: {"res 5,C", c.RES_5_C, 2, 8, "----"}, 0xAA: {"res 5,D", c.RES_5_D, 2, 8, "----"}, 0xAB: {"res 5,E", c.RES_5_E, 2, 8, "----"},
		0xAC: {"res 5,H", c.RES_5_H, 2, 8, "----"}, 0xAD: {"res 5,L", c.RES_5_L, 2, 8, "----"}, 0xAE: {"res 5,[HL]", c.RES_5_ADDR_HL, 2, 16, "----"}, 0xAF: {"res 5,A", c.RES_5_A, 2, 8, "----"},

		0xB0: {"res 6,B", c.RES_6_B, 2, 8, "----"}, 0xB1: {"res 6,C", c.RES_6_C, 2, 8, "----"}, 0xB2: {"res 6,D", c.RES_6_D, 2, 8, "----"}, 0xB3: {"res 6,E", c.RES_6_E, 2, 8, "----"},
		0xB4: {"res 6,H", c.RES_6_H, 2, 8, "----"}, 0xB5: {"res 6,L", c.RES_6_L, 2, 8, "----"}, 0xB6: {"res 6,[HL]", c.RES_6_ADDR_HL, 2, 16, "----"}, 0xB7: {"res 6,A", c.RES_6_A, 2, 8, "----"},
		0xB8: {"res 7,B", c.RES_7_B, 2, 8, "----"}, 0xB9: {"res 7,C", c.RES_7_C, 2, 8, "----"}, 0xBA: {"res 7,D", c.RES_7_D, 2, 8, "----"}, 0xBB: {"res 7,E", c.RES_7_E, 2, 8, "----"},
		0xBC: {"res 7,H", c.RES_7_H, 2, 8, "----"}, 0xBD: {"res 7,L", c.RES_7_L, 2, 8, "----"}, 0xBE: {"res 7,[HL]", c.RES_7_ADDR_HL, 2, 16, "----"}, 0xBF: {"res 7,A", c.RES_7_A, 2, 8, "----"},

		0xC0: {"set 0,B", c.SET_0_B, 2, 8, "----"}, 0xC1: {"set 0,C", c.SET_0_C, 2, 8, "----"}, 0xC2: {"set 0,D", c.SET_0_D, 2, 8, "----"}, 0xC3: {"set 0,E", c.SET_0_E, 2, 8, "----"},
		0xC4: {"set 0,H", c.SET_0_H, 2, 8, "----"}, 0xC5: {"set 0,L", c.SET_0_L, 2, 8, "----"}, 0xC6: {"set 0,[HL]", c.SET_0_ADDR_HL, 2, 16, "----"}, 0xC7: {"set 0,A", c.SET_0_A, 2, 8, "----"},
		0xC8: {"set 1,B", c.SET_1_B, 2, 8, "----"}, 0xC9: {"set 1,C", c.SET_1_C, 2, 8, "----"}, 0xCA: {"set 1,D", c.SET_1_D, 2, 8, "----"}, 0xCB: {"set 1,E", c.SET_1_E, 2, 8, "----"},
		0xCC: {"set 1,H", c.SET_1_H, 2, 8, "----"}, 0xCD: {"set 1,L", c.SET_1_L, 2, 8, "----"}, 0xCE: {"set 1,[HL]", c.SET_1_ADDR_HL, 2, 16, "----"}, 0xCF: {"set 1,A", c.SET_1_A, 2, 8, "----"},

		0xD0: {"set 2,B", c.SET_2_B, 2, 8, "----"}, 0xD1: {"set 2,C", c.SET_2_C, 2, 8, "----"}, 0xD2: {"set 2,D", c.SET_2_D, 2, 8, "----"}, 0xD3: {"set 2,E", c.SET_2_E, 2, 8, "----"},
		0xD4: {"set 2,H", c.SET_2_H, 2, 8, "----"}, 0xD5: {"set 2,L", c.SET_2_L, 2, 8, "----"}, 0xD6: {"set 2,[HL]", c.SET_2_ADDR_HL, 2, 16, "----"}, 0xD7: {"set 2,A", c.SET_2_A, 2, 8, "----"},
		0xD8: {"set 3,B", c.SET_3_B, 2, 8, "----"}, 0xD9: {"set 3,C", c.SET_3_C, 2, 8, "----"}, 0xDA: {"set 3,D", c.SET_3_D, 2, 8, "----"}, 0xDB: {"set 3,E", c.SET_3_E, 2, 8, "----"},
		0xDC: {"set 3,H", c.SET_3_H, 2, 8, "----"}, 0xDD: {"set 3,L", c.SET_3_L, 2, 8, "----"}, 0xDE: {"set 3,[HL]", c.SET_3_ADDR_HL, 2, 16, "----"}, 0xDF: {"set 3,A", c.SET_3_A, 2, 8, "----"},

		0xE0: {"set 4,B", c.SET_4_B, 2, 8, "----"}, 0xE1: {"set 4,C", c.SET_4_C, 2, 8, "----"}, 0xE2: {"set 4,D", c.SET_4_D, 2, 8, "----"}, 0xE3: {"set 4,E", c.SET_4_E, 2, 8, "----"},
		0xE4: {"set 4,H", c.SET_4_H, 2, 8, "----"}, 0xE5: {"set 4,L", c.SET_4_L, 2, 8, "----"}, 0xE6: {"set 4,[HL]", c.SET_4_ADDR_HL, 2, 16, "----"}, 0xE7: {"set 4,A", c.SET_4_A, 2, 8, "----"},
		0xE8: {"set 5,B", c.SET_5_B, 2, 8, "----"}, 0xE9: {"set 5,C", c.SET_5_C, 2, 8, "----"}, 0xEA: {"set 5,D", c.SET_5_D, 2, 8, "----"}, 0xEB: {"set 5,E", c.SET_5_E, 2, 8, "----"},
		0xEC: {"set 5,H", c.SET_5_H, 2, 8, "----"}, 0xED: {"set 5,L", c.SET_5_L, 2, 8, "----"}, 0xEE: {"set 5,[HL]", c.SET_5_ADDR_HL, 2, 16, "----"}, 0xEF: {"set 5,A", c.SET_5_A, 2, 8, "----"},

		0xF0: {"set 6,B", c.SET_6_B, 2, 8, "----"}, 0xF1: {"set 6,C", c.SET_6_C, 2, 8, "----"}, 0xF2: {"set 6,D", c.SET_6_D, 2, 8, "----"}, 0xF3: {"set 6,E", c.SET_6_E, 2, 8, "----"},
		0xF4: {"set 6,H", c.SET_6_H, 2, 8, "----"}, 0xF5: {"set 6,L", c.SET_6_L, 2, 8, "----"}, 0xF6: {"set 6,[HL]", c.SET_6_ADDR_HL, 2, 16, "----"}, 0xF7: {"set 6,A", c.SET_6_A, 2, 8, "----"},
		0xF8: {"set 7,B", c.SET_7_B, 2, 8, "----"}, 0xF9: {"set 7,C", c.SET_7_C, 2, 8, "----"}, 0xFA: {"set 7,D", c.SET_7_D, 2, 8, "----"}, 0xFB: {"set 7,E", c.SET_7_E, 2, 8, "----"},
		0xFC: {"set 7,H", c.SET_7_H, 2, 8, "----"}, 0xFD: {"set 7,L", c.SET_7_L, 2, 8, "----"}, 0xFE: {"set 7,[HL]", c.SET_7_ADDR_HL, 2, 16, "----"}, 0xFF: {"set 7,A", c.SET_7_A, 2, 8, "----"},
	}
}

func (c *CPU) Flag_set(mask uint8) {
	c.F |= mask
}

func (c *CPU) Flag_reset(mask uint8) {
	c.F &= ^mask
}

func (c *CPU) read_r16(r COMPOSITE_REGISTER_ID) uint16 {
	var hi, lo uint8
	if r == BC {
		hi = c.B
		lo = c.C
	}
	if r == DE {
		hi = c.D
		lo = c.E
	}
	if r == HL {
		hi = c.H
		lo = c.L
	}
	if r == AF {
		hi = c.A
		lo = c.F
	}
	return uint16(hi)<<8 | uint16(lo)
}

func (c *CPU) write_r16(r COMPOSITE_REGISTER_ID, d uint16) {
	var hi, lo uint8
	hi = uint8(d >> 8)
	lo = uint8(d)
	if r == BC {
		c.B = hi
		c.C = lo
	}
	if r == DE {
		c.D = hi
		c.E = lo
	}
	if r == HL {
		c.H = hi
		c.L = lo
	}
	if r == AF {
		c.A = hi
		c.F = lo
	}
}

func (c *CPU) read_byte() byte {
	// length is 1 byte
	c.PC++
	return c.Bus[c.PC]
}

func (c *CPU) read_word() uint16 {
	// length is 2 bytes
	c.PC++
	lo := uint16(c.Bus[c.PC])
	c.PC++
	hi := uint16(c.Bus[c.PC])
	return hi<<8 | lo
}

func (c *CPU) add8(a, b, carry uint8) (res, flags uint8) {
	if carry > 1 {
		carry = 1
	}
	buffer := uint16(a) + uint16(b) + uint16(carry)
	half_buffer := (a & 0x0f) + (b & 0x0f) + carry

	res = uint8(buffer)
	flags = (uint8(buffer>>8) | (half_buffer&0x10)>>3) << 4
	if res == 0 {
		flags |= FLAG_Z
	}
	return
}

func (c *CPU) add16(a, b uint16) (res uint16, flags uint8) {
	lo, flags := c.add8(uint8(a), uint8(b), 0)
	carry := flags & FLAG_C
	hi, flags := c.add8(uint8(a>>8), uint8(b>>8), carry)
	res = uint16(hi)<<8 | uint16(lo)
	return
}

func (c *CPU) sub8(a, b, carry uint8) (res, flags uint8) {
	if carry > 1 {
		carry = 1
	}
	buffer := (uint16(a) - uint16(b) - uint16(carry)) & 0x1ff     // keep lower 9 bits only
	half_buffer := (uint8(a&0x0f) - uint8(b&0x0f) - carry) & 0x1f // keep lower 5 bits only

	res = uint8(buffer)
	flags = (uint8(buffer>>8) | (half_buffer&0x10)>>3) << 4
	flags |= FLAG_N
	if res == 0 {
		flags |= FLAG_Z
	}
	return
}

func (c *CPU) push(w uint16) {
	lo := uint8(w)
	hi := uint8(w >> 8)

	c.SP--
	c.Bus[c.SP] = hi
	c.SP--
	c.Bus[c.SP] = lo
}

func (c *CPU) pop() (w uint16) {
	lo := uint16(c.Bus[c.SP])
	c.SP++
	hi := uint16(c.Bus[c.SP])
	c.SP++

	return hi<<8 | lo
}

func (c *CPU) fetch() {
	op := c.Bus.Read(c.PC)
	c.ExecInfo.Opcode = op
	c.ExecInfo.Instruction = c.Operations[op].Exec
}

func (c *CPU) execute() {
	c.ExecInfo.Instruction()
}

func (c *CPU) Run() {
	for c.ExecInfo.Opcode != 0x10 {
		c.fetch()
		c.execute()
	}
}
