package quiz

import (
	"fmt"
)

func Quiz1() {
	slice := []interface{}{1, "2", 10, "11"}
	for _, v := range slice {
		switch v.(type) {
		case int:
			if v.(int) < 10 {
				fmt.Printf("0%v\n", v)
			} else {
				fmt.Printf("%v\n", v)
			}
		case string:
			if len(v.(string)) < 2 {
				fmt.Printf("0%v\n", v)
			} else {
				fmt.Printf("%v\n", v)
			}
		}
	}
}
