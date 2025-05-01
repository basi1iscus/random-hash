package hash_writer

import (
	"errors"
	"fmt"
	"os"
	"random_hash/pkg/hasher"
)

type WriterTypeType string

const (
	Console WriterTypeType = "console"
	File WriterTypeType = "file"
)

type HashWriter struct {
	writers []Writer
}

func NewHashWriter() *HashWriter {
	return &HashWriter{
		writers: make([]Writer, 0),
	}
}

func (w *HashWriter) Close() error {
	var err error = nil
	for _, writer := range w.writers {
		e := writer.Close()
		if e != nil {
			err = e
		}
	}
	return err
}

type Writer interface {
	Open() error
	Write([]hasher.HashResult, int) error
	Close() error
}

type ConsoleWriter struct{
}

// Write implements Writer.

func (w *ConsoleWriter) Write(results []hasher.HashResult, start int) error {
	for i, result := range results {
		fmt.Printf("Value %d: %v\n", start + i + 1, result.Original)
		for algorithm, hash := range result.Hashes {
			fmt.Printf("  %s: %s\n", algorithm, hash)
		}
		fmt.Println()
	}
	return nil
}

func (w *ConsoleWriter) Open() error {
	fmt.Println("\nРезультаты:")
	return nil
}

func (w *ConsoleWriter) Close() error {
	return nil
}

type FileWriter struct {
	filename string
	file *os.File
}

func (w *FileWriter) Open() error {
	file, err := os.Create(w.filename);
	if err != nil {
		return err
	}

	w.file = file
	return nil
}

func (w *FileWriter) Close() error {
	if w.file == nil {
		return errors.New("file not open")
	}
	w.file.Close()
	return nil
}

func (w *FileWriter) Write(results []hasher.HashResult, start int) error {
	for i, result := range results {
		fmt.Fprintf(w.file, "Value %d: %v\n", start + i + 1, result.Original)
		for algorithm, hash := range result.Hashes {
			fmt.Fprintf(w.file, "  %s: %s\n", algorithm, hash)
		}
		fmt.Fprintln(w.file)
	}
	return nil
}

func (w *HashWriter) AddWriter(writerType WriterTypeType, path string) error {
	var writer Writer
	switch writerType {
	case Console:
		writer = &ConsoleWriter{}
	case File:
		writer = &FileWriter{filename: path}
	default:
		return fmt.Errorf("unsupported writer: %d", writer)
	}
	err := writer.Open()
	if err != nil {
		return err
	}
	w.writers = append(w.writers, writer)
	return nil
}

func (w *HashWriter) WriteHash(hashes []hasher.HashResult, start int) error {
	for _, writer := range w.writers {
		err := writer.Write(hashes, start)
		if err != nil {
			return err
		}
	}
	return nil
}
