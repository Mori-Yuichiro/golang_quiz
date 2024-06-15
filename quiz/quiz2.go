package quiz

import (
	"errors"
)

func FindKeyByValue(m map[int]string, value string) (int, error) {
	var err error
	var key int

	for k, v := range m {
		if v == value {
			key = k
			err = nil
			break
		} else {
			err = errors.New("そのvalueは存在しません")
		}
	}
	return key, err
}
