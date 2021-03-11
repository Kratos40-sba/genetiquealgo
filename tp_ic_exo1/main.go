package main

import (
	"fmt"
	"genetiquealgo/tp_ic_exo1/genetic"
	"math/rand"
	"time"
)

var mutationRate = 1.0

func main() {
	start := time.Now()
	rand.Seed(time.Now().UTC().UnixNano())
	population := genetic.CreatePopulation()
	found := false
	generations := 0
	for !found {
		generations++
		bestChromosome := genetic.GetBestChromosome(population)
		fmt.Printf("\r Generation : %d | %d | Fitness : %2f", generations, bestChromosome.Dna, bestChromosome.Fitness)
		if bestChromosome.Fitness == float64(1) {
			found = true
		} else {
			maxFitness := bestChromosome.Fitness
			pool := genetic.CreatePool(population, maxFitness)
			population = genetic.Selection(pool, population, mutationRate)
		}

	}
	end := time.Since(start)
	fmt.Printf("\n time taken : %s \n", end)
}

/*

	rand.Seed(time.Now().UTC().UnixNano())
	population := genetic.CreatePopulation()
	bestChromosome := genetic.GetBestChromosome(population)
    fmt.Println(population)
	fmt.Println("---------------------------------------")
	fmt.Println(bestChromosome)
	fmt.Println("---------------------------------------")
	pool := genetic.CreatePool(population,bestChromosome.Fitness)
	fmt.Println(pool)
	fmt.Println("---------------------------------------")
	population = genetic.Selection(pool,population, mutationRate)
	fmt.Println(population)
*/
