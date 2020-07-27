package fsops

import (
	"github.com/kalafut/imohash"
)

func Compare(path1, path2 string) (bool, error) {
	s1, err := imohash.SumFile(path1)
	if err != nil {
		return false, err
	}
	s2, err := imohash.SumFile(path2)
	if err != nil {
		return false, err
	}
	return s1 == s2, nil
}
