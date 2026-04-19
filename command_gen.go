package main

import (
	"fmt"

	"github.com/isaacwilkinsonlongden/pokemon-weakness-calculator/internal/effectiveness"
)

func commandGen(cfg *config, words ...string) error {
	if len(words) == 1 {
		fmt.Printf("Current generation: %s\n", cfg.generation)
		return nil
	}

	gen, err := effectiveness.ParseGeneration(words[1])
	if err != nil {
		return err
	}

	cfg.generation = gen
	fmt.Printf("Generation set to %s\n", cfg.generation)

	return nil
}
