package day21

import (
	"fmt"
	"slices"
	"strings"

	"github.com/VBenny42/AoC/2020/golang/utils"
	"github.com/VBenny42/AoC/2020/golang/utils/set"
)

type food struct {
	ingredients []string
	allergens   []string
}

type day21 struct {
	foods []food
}

func (d *day21) Part1() (count int) {
	possibleAllergenicIngredients := set.NewSet[string]()

	allergens := set.NewSet[string]()
	for _, f := range d.foods {
		for _, allergen := range f.allergens {
			allergens.Add(allergen)
		}
	}

	for allergen := range allergens {
		candidateIngredients := set.NewSet[string]()
		first := true

		for _, f := range d.foods {
			hasAllergen := false
			for _, a := range f.allergens {
				if a == allergen {
					hasAllergen = true
					break
				}
			}

			if hasAllergen {
				if first {
					candidateIngredients = set.NewSet[string]()
					for _, ingredient := range f.ingredients {
						candidateIngredients.Add(ingredient)
					}
					first = false
				} else {
					newCandidates := set.NewSet[string]()
					for _, ingredient := range f.ingredients {
						if candidateIngredients.Contains(ingredient) {
							newCandidates.Add(ingredient)
						}
					}
					candidateIngredients = newCandidates
				}
			}
		}

		for ingredient := range candidateIngredients {
			if !possibleAllergenicIngredients.Contains(ingredient) {
				possibleAllergenicIngredients.Add(ingredient)
			}
		}
	}

	for _, f := range d.foods {
		for _, ingredient := range f.ingredients {
			if !possibleAllergenicIngredients.Contains(ingredient) {
				count++
			}
		}
	}

	return
}

// Claude did this :)
func (d *day21) Part2() string {
	// For each allergen, find ingredients that could contain it
	allergenToCandidates := make(map[string]map[string]bool)

	// Get all unique allergens
	allergens := make(map[string]bool)
	for _, f := range d.foods {
		for _, allergen := range f.allergens {
			allergens[allergen] = true
		}
	}

	// For each allergen, find ingredients that appear in ALL foods containing that allergen
	for allergen := range allergens {
		var candidateIngredients map[string]bool
		first := true

		// Find all foods that contain this allergen
		for _, f := range d.foods {
			hasAllergen := false
			for _, a := range f.allergens {
				if a == allergen {
					hasAllergen = true
					break
				}
			}

			if hasAllergen {
				if first {
					// Initialize with ingredients from first food containing this allergen
					candidateIngredients = make(map[string]bool)
					for _, ingredient := range f.ingredients {
						candidateIngredients[ingredient] = true
					}
					first = false
				} else {
					// Keep only ingredients that are also in this food (intersection)
					newCandidates := make(map[string]bool)
					for _, ingredient := range f.ingredients {
						if candidateIngredients[ingredient] {
							newCandidates[ingredient] = true
						}
					}
					candidateIngredients = newCandidates
				}
			}
		}

		allergenToCandidates[allergen] = candidateIngredients
	}

	// Process of elimination: repeatedly find allergens with only one candidate
	allergenToIngredient := make(map[string]string)
	usedIngredients := make(map[string]bool)

	for len(allergenToIngredient) < len(allergens) {
		for allergen, candidates := range allergenToCandidates {
			if _, solved := allergenToIngredient[allergen]; solved {
				continue
			}

			// Remove already used ingredients from candidates
			availableCandidates := make([]string, 0)
			for ingredient := range candidates {
				if !usedIngredients[ingredient] {
					availableCandidates = append(availableCandidates, ingredient)
				}
			}

			// If only one candidate remains, we found the match
			if len(availableCandidates) == 1 {
				ingredient := availableCandidates[0]
				allergenToIngredient[allergen] = ingredient
				usedIngredients[ingredient] = true
			}
		}
	}

	// Sort allergens alphabetically and build result string
	sortedAllergens := make([]string, 0, len(allergens))
	for allergen := range allergens {
		sortedAllergens = append(sortedAllergens, allergen)
	}

	slices.Sort(sortedAllergens) // Sort allergens alphabetically

	// Build comma-separated list of ingredients in alphabetical order of allergens
	result := make([]string, 0, len(sortedAllergens))
	for _, allergen := range sortedAllergens {
		result = append(result, allergenToIngredient[allergen])
	}

	return strings.Join(result, ",")
}

func Parse(filename string) *day21 {
	lines := utils.ReadLines(filename)
	foods := make([]food, 0, len(lines))

	for _, line := range lines {
		var f food
		left, right, ok := strings.Cut(line, " (contains ")
		f.ingredients = strings.Fields(left)
		if !ok {
			// food has no allergens
			fmt.Println("LOG: No allergens found in line:", line)
		} else {
			// food has allergens
			right = strings.TrimSuffix(right, ")")
			f.allergens = strings.Split(right, ", ")
		}
		foods = append(foods, f)
	}

	return &day21{foods: foods}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: number of ingredients that contain no allergens:", day.Part1())
	fmt.Println("ANSWER2: canonical dangerous ingredient list:", day.Part2())
}
