package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Portfolio struct {
	ValuePrevDay int
	Positions    []Position
}

func (p *Portfolio) value() (float64, error) {
	if len(p.Positions) == 0 {
		return 0, fmt.Errorf("No Positions in Portfolio")
	}
	value := 0.0
	for _, posistion := range p.Positions {
		revenue, _, err := posistion.revenue()
		if err != nil {
			return 0, err
		}
		value += revenue
	}
	return value, nil
}

func (p *Portfolio) addPosition(pos Position) {
	p.Positions = append(p.Positions, pos)
}

func (p *Portfolio) rmPosition(positionInArray int) error {
	if positionInArray > len(p.Positions)-1 {
		return fmt.Errorf("Position %d is out of range: Number of current Positions: %d", positionInArray, len(p.Positions)-1)
	}
	p.Positions = append(p.Positions[:positionInArray], p.Positions[positionInArray+1:]...)
	return nil
}

func (p *Portfolio) importJSON(location string) error {
	data, err := ioutil.ReadFile(location)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &p.Positions)
	if err != nil {
		return err
	}
	fmt.Printf("imported Portfolio from %s\n", location)
	return nil
}

func (p Portfolio) exportJSON() error {
	jsonExport, err := json.Marshal(&p.Positions)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("portfolio.json", jsonExport, 0644)
	if err != nil {
		return err
	}
	fmt.Println("Exported portfolio to portfolio.json")
	return nil
}
