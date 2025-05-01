package main

import (
	"flag"
	"fmt"
	"random_hash/pkg/generator"
	"random_hash/pkg/hash_writer"
	"random_hash/pkg/hasher"
	"strings"
	"sync"
)

const BufferSize = 100
type RandomHashOptions struct {
	count     int
	min       int
	max       int
	valueType generator.RandomValueType
	hashTypes    []hasher.HashAlgorithm
	threads   int
	out       hash_writer.WriterTypeType
	filePath  string
	complexity  string
}

func getCommandlineOptions() RandomHashOptions {
	// Парсинг аргументов командной строки
	count := flag.Int("count", 1000, "количество генерируемых чисел")
	min := flag.Int("min", 1, "минимальное значение")
	max := flag.Int("max", 1000, "максимальное значение")
	typeFlag := flag.String("type", "mixed", "тип данный (int, float, mixed, string)")
	out := flag.String("out", "console", "тип вывода (console, file)")
	filePath := flag.String("path", "results.txt", "выходной файл")
	threads := flag.Int("threads", 2, "число потоков")
	complexity := flag.String("complexity", "", "сложность")
	algString := flag.String("algs", "MD5,SHA-256,SHA-3", "сложность")
	flag.Parse()

	var algs = strings.Split(*algString, ",")
	var hashes = []hasher.HashAlgorithm{} 
	for _, alg := range algs {
		hashes = append(hashes, hasher.HashAlgorithm(alg))
	}

	return RandomHashOptions{
		count:     int(*count),
		min:       int(*min),
		max:       int(*max),
		valueType: generator.RandomValueType(*typeFlag),
		filePath:  *filePath,
		out:       hash_writer.WriterTypeType(*out),
		threads:   int(*threads),
		hashTypes: hashes,
		complexity: *complexity,
	}
}

func complexityCounter(results []hasher.HashResult, alg hasher.HashAlgorithm, complexity string) int {
	var count = 0
	for _, result := range results {
		if strings.HasPrefix(result.Hashes[alg], complexity) {
			count++
		}
	}
	return count
}

func main() {
	var options = getCommandlineOptions()

	// Генерация случайных чисел
	fmt.Println("Генерация", options.count, "случайных данных...")
	channel := make(chan []hasher.HashResult, BufferSize)
	restCount := options.count
	var wg sync.WaitGroup
	for i := 0; i < options.threads; i++ {
		var threadCount int
		if i == options.threads-1 {
			threadCount = restCount
		} else {
			threadCount = options.count / options.threads
			restCount -= threadCount
		}

		wg.Add(1)
		// Запускаем в отдельных горутинах
		go func(count int) {
			defer wg.Done()
			gen := generator.NewGenerator()
			hasher := hasher.NewHasher(options.hashTypes...)
			for count > 0 {
				var batchSize int
				if count < BufferSize {
					batchSize = count
				} else {
					batchSize = BufferSize
				}
				count -= batchSize
				values := gen.Generate(batchSize, options.valueType, options.min, options.max)
				hashResult, err := hasher.HashValues(values)
				if err != nil {
					fmt.Println("Ошибка при хэшировании:", err)
					continue
				}
				channel <- hashResult
			}
		}(threadCount)
	}

	go func() {
		wg.Wait()
		close(channel)
	}()

	writer := hash_writer.NewHashWriter()
	writer.AddWriter(options.out, options.filePath)
	counter := 0
	compCounter := 0
	for hashResult := range channel {
		writer.WriteHash(hashResult, counter)
		counter += len(hashResult)
		if len(options.complexity) > 0 {
			compCounter += complexityCounter(hashResult, options.hashTypes[0], options.complexity);
		}
	}
	writer.Close()
	if len(options.complexity) > 0 {
		fmt.Println("Удовлетворяют сложности ", compCounter, " хешей...")
	}
}
