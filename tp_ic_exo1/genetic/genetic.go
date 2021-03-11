package genetic

import (
	"math"
	"math/rand"
)

type Chromosome struct {
	Dna     []int
	Fitness float64
}

func generateDna() (ch Chromosome) {
	dna := make([]int, 4)
	for i := 0; i < 4; i++ {
		dna[i] = rand.Intn(30)
	}
	ch = Chromosome{
		Dna:     dna,
		Fitness: 0.0,
	}

	return ch
}
func (ch *Chromosome) calculateFitness() float64 {
	var fitnessRate float64
	fitnessRate = float64(ch.Dna[0] + 2*ch.Dna[1] + 3*ch.Dna[2] + 4*ch.Dna[3] - 30)
	fitnessRate = math.Abs(fitnessRate)
	return 1 / (fitnessRate + 1)
}
func CreatePopulation() []Chromosome {
	ch := make([]Chromosome, 6)
	for i := 0; i < 6; i++ {
		ch[i] = generateDna()
		ch[i].Fitness = ch[i].calculateFitness()
	}
	return ch
}
func CreatePool(population []Chromosome, bestFit float64) (pool []Chromosome) {
	pool = make([]Chromosome, 0)
	for i := 0; i < 6; i++ {
		population[i].calculateFitness()
		iterations := int(population[i].Fitness/bestFit) * 100
		for j := 0; j < iterations; j++ {
			pool = append(pool, population[i])
		}
	}
	return pool
}
func crossover(ch1, ch2 Chromosome) (ch *Chromosome) {
	ch = &Chromosome{
		Dna:     make([]int, 4),
		Fitness: 0.0,
	}
	x := rand.Intn(4)
	for i := 0; i < 4; i++ {
		if i > x {
			ch.Dna[i] = ch1.Dna[i]
		} else {
			ch.Dna[i] = ch2.Dna[i]
		}
	}
	return ch
}
func (ch *Chromosome) mutate(mutationRate float64) *Chromosome {
	for i := 0; i < 4; i++ {
		if rand.Float64() < mutationRate {
			ch.Dna[i] = rand.Intn(30)
		}
	}
	return ch
}
func GetBestChromosome(population []Chromosome) Chromosome {
	bestFitness := 0.0
	index := 0
	for i := 0; i < len(population); i++ {
		if population[i].Fitness > bestFitness {
			bestFitness = population[i].Fitness
			index = i
		}
	}
	return population[index]
}
func Selection(pool, population []Chromosome, mutationRate float64) []Chromosome {
	nextGen := make([]Chromosome, len(population))
	for i := 0; i < len(population); i++ {
		r1, r2 := rand.Intn(len(pool)), rand.Intn(len(pool))
		a := pool[r1]
		b := pool[r2]
		child := crossover(a, b).
			mutate(mutationRate)
		ch := &Chromosome{
			Dna:     child.Dna,
			Fitness: child.calculateFitness(),
		}
		nextGen[i] = *ch

	}
	return nextGen
}
