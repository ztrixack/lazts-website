package utils

import (
	"math"
	"math/rand/v2"
	"sync"
)

type Blackhole struct {
	Size    int
	Rotate  int
	Opacity int
	Width   int
}

func RandomizeBlackholes(count int) []Blackhole {
	if count <= 0 {
		return []Blackhole{}
	}

	var wg sync.WaitGroup
	blackholes := make([]Blackhole, count)
	ch := make(chan Blackhole, count)

	worker := func() {
		defer wg.Done()
		size := rand.IntN(360) + 180
		rotate := rand.IntN(360)
		opacity := int(math.Max(110-float64(size*100/450), 5))
		width := (size - 90) / 6

		ch <- Blackhole{Size: size, Rotate: rotate, Opacity: opacity, Width: width}
	}

	for i := 0; i < count; i++ {
		wg.Add(1)
		go worker()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	index := 0
	for bh := range ch {
		blackholes[index] = bh
		index++
	}

	return blackholes
}

type Cloud struct {
	Top    int
	Left   int
	Rotate int
}

func RandomizeClouds(count int) []Cloud {
	if count <= 0 {
		return []Cloud{}
	}

	var wg sync.WaitGroup
	clouds := make([]Cloud, count)
	ch := make(chan Cloud, count)

	worker := func() {
		defer wg.Done()
		top := rand.IntN(150) - 50
		left := rand.IntN(100)
		rotate := rand.IntN(360)
		ch <- Cloud{Top: top, Left: left, Rotate: rotate}
	}

	for i := 0; i < count; i++ {
		wg.Add(1)
		go worker()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	index := 0
	for c := range ch {
		clouds[index] = c
		index++
	}

	return clouds
}
