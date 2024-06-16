package quiz

import (
	"errors"
)

func FindKeyByValue(m map[int]string, value string) (int, error) {
	for k, v := range m {
		if v == value {
			return k, nil
		}
	}
	return 0, errors.New("そのvalueは存在しません")
}
