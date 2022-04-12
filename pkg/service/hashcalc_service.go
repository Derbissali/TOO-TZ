package service

import (
	"hash/crc64"
	"math/rand"
	"strconv"
	"time"
)

type serviceHash struct {
}

func NewHashService() *serviceHash {
	return &serviceHash{}
}

func (h *serviceHash) GetID() (string, error) {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(200)), nil
}

func (h *serviceHash) GetHash(str string) (string, error) {

	s := int64(crc64.Checksum([]byte(str), crc64.MakeTable(crc64.ISO)))

	n := time.Now()
	t := n.UnixNano()
	for i := 0; i < 12; i++ {
		s = s & t
		n = n.Add(5 * time.Second)
		t = n.UnixNano()
	}
	res := strconv.Itoa(scoreOne(s))

	return res, nil

}

func scoreOne(i int64) int {
	var k int
	for i != 0 {
		if i%2 == 1 {
			k++
		}
		i = i / 2
	}
	return k
}
