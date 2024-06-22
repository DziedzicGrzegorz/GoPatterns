package samplehash

import (
	"bytes"
	"encoding/gob"
	"hash/fnv"
)

func MakeHash[T any](value T) uint64 {

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(value)
	if err != nil {
		panic(err)
	}

	h := fnv.New64a()
	_, err = h.Write(buf.Bytes())
	if err != nil {
		panic(err)
	}
	return h.Sum64()
}
