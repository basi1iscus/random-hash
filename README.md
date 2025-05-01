# random_hash

A Go library and CLI for generating random values and computing their hashes using various algorithms. Includes utilities for writing hash results to different outputs.

## Features
- Generate random integers, floats, mixed(integers and floats) and strings
- Hash values using MD5, SHA-256, SHA-512, and SHA-3
- Write hash results to console or file
- Easily extensible and modular

## Project Structure
- `cmd/main.go` — CLI entry point
- `pkg/generator/` — Random value generator
- `pkg/hasher/` — Hashing utilities
- `pkg/hash_writer/` — Output writers for hash results

## Usage

### Run the CLI
```sh
go run ./cmd/main.go
```

The command-line arguments for this project:

- -count: Number of random values to generate (default: 1000)
- -min: Minimum value for generated numbers (default: 1) or minimun string length for generating strings 
- -max: Maximum value for generated numbers (default: 1000) or maximum string length for generating strings
- -type: Type of values to generate: int, float, mixed, or string (default: mixed)
- -out: Output type: console or file (default: console)
- -path: Output file path (optional if path sets write results into file, default: results.txt)
- -threads: Number of concurrent threads/goroutines to use (default: 2)
- -algs: Comma-separated list of hash algorithms to use (default: MD5,SHA-256,SHA-3)
- -complexity: String prefix to filter hashes by complexity (optional, complexity count for the first hash algorithms by the algs option)

Example usage:
```sh
go run ./cmd/main.go --count=5000 --min=10 --max=9999 --type=int --threads=4 --algs=SHA-256,SHA-512
```
This will generate 5000 random integers between 10 and 9999, hash them with SHA-256 and SHA-512, and print results to the console using 4 threads.

Example usage:
```sh
go run ./cmd/main.go --count=100000 --min=32 --max=32 --type=string --out=file --path=./results.txt --threads=4 --algs=SHA-256,SHA-512,MD5 --complexity=00
```
This will generate 100000 random 32 chars strings, hash them with SHA-256, SHA-512 and MD5 algorithms, and print results to the console and file, using 4 threads, count SHA-256 hashes started with 00.

### Use as a Library
Import and use the packages in your Go code:
```go
import "random_hash/pkg/generator"
import "random_hash/pkg/hasher"
import "random_hash/pkg/hash_writer"
```

## Testing
Run all tests:
```sh
go test ./pkg/...
```

## Requirements
- Go 1.18 or newer

## License
MIT
