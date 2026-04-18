package main

import (
	"fmt"
	"sort"
)

func commandPokemon(cfg *config, words ...string) error {
	pokemonResp, err := cfg.pokeapiClient.GetPokemon(words[0])
	if err != nil {
		return err
	}

	pokemonTypes := []string{}
	for _, pokemonType := range pokemonResp.Types {
		pokemonTypes = append(pokemonTypes, pokemonType.Type.Name)
	}

	// attacking type -> multiplier against this pokemon
	multipliers := map[string]float64{}

	// initialize all types to neutral
	allTypes := []string{
		"normal", "fire", "water", "electric", "grass", "ice",
		"fighting", "poison", "ground", "flying", "psychic",
		"bug", "rock", "ghost", "dragon", "dark", "steel", "fairy",
	}

	for _, t := range allTypes {
		multipliers[t] = 1.0
	}

	for _, pokemonType := range pokemonTypes {
		typeResp, err := cfg.pokeapiClient.GetType(pokemonType)
		if err != nil {
			return err
		}

		for _, t := range typeResp.DamageRelations.DoubleDamageFrom {
			multipliers[t.Name] *= 2
		}

		for _, t := range typeResp.DamageRelations.HalfDamageFrom {
			multipliers[t.Name] *= 0.5
		}

		for _, t := range typeResp.DamageRelations.NoDamageFrom {
			multipliers[t.Name] *= 0
		}
	}

	fmt.Printf("Name: %s\n", pokemonResp.Name)
	fmt.Printf("Types: %v\n", pokemonTypes)
	fmt.Println("Weaknesses:")

	keys := make([]string, 0, len(multipliers))
	for k, v := range multipliers {
		if v > 1 {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("- %s (x%.1f)\n", k, multipliers[k])
	}

	return nil
}
