package main

import (
	"go-boy/hardware"
	op "go-boy/opcodes"
	"log"
	"os"
)

func main() {
	// program := []byte{
	// 	op.LD_HL_n16, 0x20, 0x00, // HL <- 2000
	// 	op.LD_A_n8, 0x01,
	// 	op.LD_B_A,           // B <- 1
	// 	op.LD_A_B,           // loop: A <- B
	// 	op.LDI_ADDR_HL_A,    // [HL++] <- A
	// 	op.PREFIX, op.RLC_B, // Rotate Left B
	// 	op.JR_NC_e8, 0xfb, // jump to loop
	// 	op.STOP,
	// }
	program := []byte{
		op.LD_A_n8, 0x11,
		op.LD_ADDR_a16_A, 0x0A, 0x00,
		op.STOP,
	}
	// program := []byte{
	// 	op.LD_HL_n16, 0x20, 0x00, // HL <- 2000
	// 	op.LD_A_n8, 0x00,
	// 	op.PREFIX, op.SET_0_A,
	// 	op.LDI_ADDR_HL_A,
	// 	op.PREFIX, op.SET_1_A,
	// 	op.LDI_ADDR_HL_A,
	// 	op.PREFIX, op.SET_2_A,
	// 	op.LDI_ADDR_HL_A,
	// 	op.PREFIX, op.SET_3_A,
	// 	op.LDI_ADDR_HL_A,
	// 	op.PREFIX, op.SET_4_A,
	// 	op.LDI_ADDR_HL_A,
	// 	op.PREFIX, op.SET_5_A,
	// 	op.LDI_ADDR_HL_A,
	// 	op.PREFIX, op.SET_6_A,
	// 	op.LDI_ADDR_HL_A,
	// 	op.PREFIX, op.SET_7_A,
	// 	op.LDI_ADDR_HL_A,
	// 	op.STOP,
	// }
	cpu := hardware.GetCPU()
	ram := hardware.GetBus()

	ram.WriteBytes(program, 0x0000)
	cpu.PC = 0x0000
	cpu.Run()

	ram_contents := ram.String()
	err := os.WriteFile("ram.txt", []byte(ram_contents), 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
}
