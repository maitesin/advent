package cpu

import (
	"fmt"
	"strconv"
	"strings"
)

type Instruction interface {
	Compute(*CPU)
}

func NewInstruction(code string) (Instruction, error) {
	parts := strings.Split(code, " ")

	switch len(parts) {
	case 1:
		return &Noop{}, nil
	case 2:
		value, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		return &AddX{x: value}, nil
	default:
		return nil, fmt.Errorf("instruction %q not recognized", code)
	}
}

type Noop struct{}

func (n *Noop) Compute(cpu *CPU) {
}

type AddX struct {
	x int
}

func (a *AddX) Compute(cpu *CPU) {
	cpu.wait = 1
	cpu.nextX = cpu.x + a.x
}

type CPU struct {
	x     int
	nextX int
	wait  int
}

func NewCPU() *CPU {
	return &CPU{nextX: 1}
}

func (c *CPU) Tick() bool {
	if c.wait == 0 {
		c.x = c.nextX
		return true
	}
	c.wait--
	return false
}

func (c *CPU) Compute(instruction Instruction) {
	instruction.Compute(c)
}

func (c *CPU) X() int {
	return c.x
}
