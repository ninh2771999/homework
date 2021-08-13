package main
import (
	"fmt"
)
func main(){
	m := map[string]int{
		"ninh":20,
		"tien":22,
	}
	if age , isExist := m["ninh"];isExist{
		fmt.Println(age)
	}
}