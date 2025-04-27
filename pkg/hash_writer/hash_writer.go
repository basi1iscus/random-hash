package hash_writer

import (
	"crypto/sha256"
	"fmt"

	"github.com/basi1iscus/random-hash/pkg/hasher"
)

type WriterTypeType int

const (
	Console WriterTypeType = iota
	File
)

type HashWriter struct {
	writers []WriterTypeType
}

func NewHashWriter() *HashWriter {
	return &HashWriter{
		writers: make([]WriterTypeType, 0),
	}
}

func (w *HashWriter) AddWriter(WriterTypeType) {
	w.writers = append(w.writers)
}

func (w *HashWriter) WriteHash(hashes []hasher.HashResult) error {
	for _, writer := range w.writers {
		switch writer {
		case Console:
			consoleWriter(hashes)
		case File:
			hasher = sha256.New()
		default:
			return fmt.Errorf("unsupported writer: %s", writer)
		}
		return nil
	}

}

func consoleWriter(hashes [][]hasher.HashResult) {
	// Вывод результатов в консоль
	fmt.Println("\nРезультаты:")
	for i, results := range hashes {
		fmt.Printf("Value %d: %v\n", i+1, numbers[i])
		for _, result := range results {
			fmt.Printf("  %s: %s\n", result.Algorithm, result.Hash)
		}
		fmt.Println()
	}
}
