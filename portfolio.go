package main

import (
	"fmt"
)

type Portfolio struct {
	valuePrevDay int
	positions    []Position
}

func (p *Portfolio) value() (float64, error) {
	if len(p.positions) == 0 {
		return 0, fmt.Errorf("No Positions in Portfolio")
	}
	value := 0.0
	for _, posistion := range p.positions {
		revenue, _, err := posistion.revenue()
		if err != nil {
			return 0, err
		}
		value += revenue
	}
	return value, nil
}

func (p *Portfolio) addPosition(pos Position) {
	p.positions = append(p.positions, pos)
}

func (p *Portfolio) rmPosition(positionInArray int) error {
	if positionInArray > len(p.positions)-1 {
		return fmt.Errorf("Position %d is out of range: Number of current Positions: %d", positionInArray, len(p.positions)-1)
	}
	p.positions = append(p.positions[:positionInArray], p.positions[positionInArray+1:]...)
	return nil
}
