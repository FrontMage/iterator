package iterator

import (
	"math"
)

// Chunk split a given slice into mutilple small chunks
func Chunk(s []interface{}, chunk int) [][]interface{} {
	if len(s) == 0 {
		return [][]interface{}{s}
	}
	chunkSize := int(math.Ceil(float64(len(s) / chunk)))
	if chunkSize <= 1 {
		return [][]interface{}{s}
	}
	result := [][]interface{}{}
	for idx := 0; idx < len(s); idx += chunkSize {
		end := idx + chunkSize
		if end > len(s) {
			result = append(result, s[idx:])
		} else {
			result = append(result, s[idx:end])
		}
	}
	return result
}

// Iter opens mutiple goroutine to iter a given slice
func Iter(s []interface{}, callback func(idx int, item interface{}), workerCount int) {
	if workerCount == 0 {
		workerCount = 10
	}
	tasks := Chunk(s, workerCount)
	doneSigs := make(chan int, workerCount)
	for _, t := range tasks {
		go func(t []interface{}) {
			for idx, item := range t {
				callback(idx, item)
			}
			doneSigs <- 1
		}(t)
	}
	counter := 0
	for {
		select {
		case sig := <-doneSigs:
			counter += sig
			if counter == len(tasks) {
				return
			}
		}
	}
}
