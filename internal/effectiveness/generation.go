package effectiveness

import "fmt"

type Generation string

const (
	Current        Generation = ""
	GenerationI    Generation = "generation-i"
	GenerationII   Generation = "generation-ii"
	GenerationIII  Generation = "generation-iii"
	GenerationIV   Generation = "generation-iv"
	GenerationV    Generation = "generation-v"
	GenerationVI   Generation = "generation-vi"
	GenerationVII  Generation = "generation-vii"
	GenerationVIII Generation = "generation-viii"
	GenerationIX   Generation = "generation-ix"
)

func generationNumber(gen Generation) int {
	switch gen {
	case GenerationI:
		return 1
	case GenerationII:
		return 2
	case GenerationIII:
		return 3
	case GenerationIV:
		return 4
	case GenerationV:
		return 5
	case GenerationVI:
		return 6
	case GenerationVII:
		return 7
	case GenerationVIII:
		return 8
	case GenerationIX, Current:
		return 9
	default:
		return 0
	}
}

func ParseGeneration(input string) (Generation, error) {
	switch input {
	case "1", "gen1", "generation-i":
		return GenerationI, nil
	case "2", "gen2", "generation-ii":
		return GenerationII, nil
	case "3", "gen3", "generation-iii":
		return GenerationIII, nil
	case "4", "gen4", "generation-iv":
		return GenerationIV, nil
	case "5", "gen5", "generation-v":
		return GenerationV, nil
	case "6", "gen6", "generation-vi":
		return GenerationVI, nil
	case "7", "gen7", "generation-vii":
		return GenerationVII, nil
	case "8", "gen8", "generation-viii":
		return GenerationVIII, nil
	case "9", "gen9", "generation-ix", "current":
		return GenerationIX, nil
	default:
		return "", fmt.Errorf("unknown generation: %s", input)
	}
}
