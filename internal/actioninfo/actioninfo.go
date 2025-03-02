package actioninfo

import (
	"fmt"
)

// создайте интерфейс DataParser
type DataParser interface {
	ActionInfo() string
	Parse(datastring string) (err error)
}

// создайте функцию Info()
func Info(dataset []string, dp DataParser) {
	for _, datastring := range dataset {

		err := dp.Parse(datastring)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(dp.ActionInfo())
	}
}
