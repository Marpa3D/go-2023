// тип данных error
package main

import (
	"errors"
	"fmt"
	"os"
)

// check - пример создания пользовательского сообщения об ошибке
func check(a, b int) error {
	if a == 0 && b == 0 {
		return errors.New("Это пользовательское сообщение об ошибке")
	}
	return nil
}

// formattedError - сообщает об ошибке с помощью fmt.Errorf()
func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("a %d and b %d. UserID: %d", a, b, os.Getuid())
	}
}

func main() {

}
