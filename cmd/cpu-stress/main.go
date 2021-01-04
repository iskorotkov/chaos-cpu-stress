package main

import (
	"fmt"
	"github.com/iskorotkov/chaos-io-stress/pkg/bench"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	out := make(chan os.Signal, 1)
	signal.Notify(out, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan struct{})
	go func() {
		<-out
		done <- struct{}{}
	}()

	rng := rand.New(rand.NewSource(0))

	dimStr := os.Getenv("DIM")
	dim := 64
	if dimStr != "" {
		value, err := strconv.ParseInt(dimStr, 10, 32)
		if err != nil {
			panic(err)
		}

		dim = int(value)
	}

	a := createRandomMatrix(dim, rng)

	bench.Benchmark(func() {
		a.multiply(a)
	}, func(count int64) {
		fmt.Printf("executed %d times\n", count)
	}, time.Second, done)
}

type Matrix [][]int

func (m Matrix) multiply(other Matrix) Matrix {
	res := createMatrix(len(m))
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			for k := 0; k < len(m); k++ {
				res[i][k] += m[i][j] * other[j][k]
			}
		}
	}

	return res
}

func createMatrix(dim int) Matrix {
	m := make([][]int, 0)
	for i := 0; i < dim; i++ {
		r := make([]int, 0)
		for j := 0; j < dim; j++ {
			r = append(r, 0)
		}

		m = append(m, r)
	}

	return m
}

func createRandomMatrix(dim int, rng *rand.Rand) Matrix {
	m := createMatrix(dim)
	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			m[i][j] = rng.Int()
		}
	}

	return m
}
