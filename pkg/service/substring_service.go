package service

import (
	"errors"
	"fmt"
)

type serviceSub struct {
}

func NewSubService() *serviceSub {
	return &serviceSub{}
}

func (s *serviceSub) MaxLength(sub *string) (string, error) {

	arr := make(map[string]int)
	res := ""
	for j, i := range *sub {
		if i >= 48 && i <= 57 || i >= 65 && i <= 90 || i >= 97 && i <= 122 {
			_, prs := arr[string(i)]
			if prs != true {
				arr[string(i)] = j
				res += string(i)
			}
		} else {
			return "", errors.New("bad req")
		}
	}
	fmt.Println(res)
	return res, nil
}
