package hasher

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"hash"
	"math"
)

type HashAlgorithm string

const (
	MD5    HashAlgorithm = "MD5"
	SHA256 HashAlgorithm = "SHA-256"
	SHA512 HashAlgorithm = "SHA-512"
	SHA3   HashAlgorithm = "SHA-3"
)

type Hash map[HashAlgorithm]string

type HashResult struct {
	Hashes    Hash
	Original  interface{}
}

type Hasher struct {
	algorithms []HashAlgorithm
}

var DEFAULT_ALGORITHM = []HashAlgorithm{SHA256}

func NewHasher(algorithms ...HashAlgorithm) *Hasher {
	if len(algorithms) == 0 {
		algorithms = DEFAULT_ALGORITHM
	}

	return &Hasher{
		algorithms: algorithms,
	}
}

func (h *Hasher) HashValue(value interface{}) (Hash, error) {
	results := make(Hash, len(h.algorithms))

	for _, alg := range h.algorithms {
		var hasher hash.Hash
		switch alg {
		case MD5:
			hasher = md5.New()
		case SHA256:
			hasher = sha256.New()
		case SHA512:
			hasher = sha512.New()
		case SHA3:
			hasher = sha3.New224()
		default:
			return nil, fmt.Errorf("unsupported hash algorithm: %s", alg)
		}

		var b []byte
		switch v := value.(type) {
		case string:
			b := []byte(v)
			hasher.Write(b)
		case int:
			b := make([]byte, 8)
			binary.LittleEndian.PutUint64(b, uint64(v))
		case float64:
			b := make([]byte, 8)
			binary.LittleEndian.PutUint64(b[:], math.Float64bits(v))
		default:
			return nil, fmt.Errorf("unsupported value type: %T", value)
		}

		hasher.Write(b)
		hashBytes := hasher.Sum(nil)
		hashHex := hex.EncodeToString(hashBytes)

		results[alg] = hashHex
	}
	return results, nil
}

func (h *Hasher) HashValues(values []interface{}) ([]HashResult, error) {
	results := make([]HashResult, len(values))
	for i, value := range values {
		hashes, err := h.HashValue(value)
		if err != nil {
			return nil, err
		}
		results[i] = HashResult {
			Original: value,
			Hashes: hashes,
		}
	}

	return results, nil
}
