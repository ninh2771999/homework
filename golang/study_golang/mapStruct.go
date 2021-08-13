package main
import (
	"fmt"
)
type Student struct{
	number 		int
	name 		string
	isMale	 	bool
	subjects 	[]string
}
func main(){
		// studentNameAgeMap := make (map[string]int) <----------map
		// studentNameAgeMap = map[string]int {
		// 	"ninh": 19,
		// 	"tien": 20,
		// 	"hoa": 25,
		// }
		// studentNameAgeMap["vietanh"] = 22
		// studentNameAgeMap["hoa"] = 23
		// delete(studentNameAgeMap,"tien")
		// fmt.Println(studentNameAgeMap)
		// // fmt.Println(studentNameAgeMap["ninh"])
		// copyMap := studentNameAgeMap
		// fmt.Println(copyMap)
			// studentNinh := Student{ ----------> struc
			// 	number: 	3,
			// 	name:		"ninh",
			// 	isMale: 	true,
			// 	subjects: 	[]string{
			// 		"math",
			// 		"english",
			// 		"computer",
			// 	},		
			// }
			// fmt.Println(studentNinh.name)
			// fmt.Println(studentNinh.number)

			studentNinh := Student {}
			studentNinh.number = 3
			studentNinh.name = "Ninh"
			studentNinh.isMale = true
			studentNinh.subjects = []string{"Math","English","Computer"}
			fmt.Println(studentNinh)
}