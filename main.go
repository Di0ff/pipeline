package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	bufferSize    = 5
	flushInterval = 3 * time.Second
)

func main() {
	done := make(chan struct{})
	defer close(done)

	source := generateNumbers(done)

	positiveNumbers := removeNegative(done, source)
	multiplesOfThree := filterMultiplesOfThree(done, positiveNumbers)
	buffered := buffer(done, multiplesOfThree)

	consume(buffered)
}

// Источник данных (чтение с консоли)
func generateNumbers(done <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			select {
			case <-done:
				return
			default:
				text := scanner.Text()
				num, err := strconv.Atoi(text)
				if err != nil {
					fmt.Println("Введите корректное число")
					continue
				}
				out <- num
			}
		}
	}()
	return out
}

// Фильтрация отрицательных чисел
func removeNegative(done <-chan struct{}, in <-chan int) <-chan int {
	positive := make(chan int)
	go func() {
		defer close(positive)
		for i := range in {
			select {
			case <-done:
				return
			default:
				if i >= 0 {
					positive <- i
				}
			}
		}
	}()
	return positive
}

// Фильтрация чисел, не кратных 3 и исключение 0
func filterMultiplesOfThree(done <-chan struct{}, in <-chan int) <-chan int {
	multiplesOfThree := make(chan int)
	go func() {
		defer close(multiplesOfThree)
		for i := range in {
			select {
			case <-done:
				return
			default:
				if i%3 == 0 && i != 0 {
					multiplesOfThree <- i
				}
			}
		}
	}()
	return multiplesOfThree
}

// Буферизация данных в кольцевом буфере
func buffer(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	buffer := make([]int, 0, bufferSize)

	go func() {
		defer close(out)
		ticker := time.NewTicker(flushInterval)
		defer ticker.Stop()

		for {
			select {
			case <-done:
				return
			case i, ok := <-in:
				if !ok {
					flushBuffer(out, buffer)
					return
				}
				buffer = append(buffer, i)
				if len(buffer) >= bufferSize {
					flushBuffer(out, buffer)
					buffer = buffer[:0]
				}
			case <-ticker.C:
				flushBuffer(out, buffer)
				buffer = buffer[:0]
			}
		}
	}()
	return out
}

// Опустошение буфера
func flushBuffer(out chan<- int, buffer []int) {
	for _, i := range buffer {
		out <- i
	}
}

// Потребитель данных
func consume(in <-chan int) {
	for i := range in {
		fmt.Printf("Получены данные: %d\n", i)
	}
}
