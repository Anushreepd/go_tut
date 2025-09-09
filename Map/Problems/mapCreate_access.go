package Problems

import "fmt"

func MapCreation() {
	//create
	studentGrades := map[string]int{}

	studentGrades["Alice"] = 85
	studentGrades["Bob"] = 90
	studentGrades["Charlie"] = 78

	fmt.Println("student data", studentGrades)

	//isExist
	idx, isExist := studentGrades["David"]
	if isExist {
		fmt.Println("david exist in index value = ", idx)
	} else {
		fmt.Println("no exist")
	}

	//update -> Change Bob’s grade to 95.
	studentGrades["Bob"] = 95
	fmt.Println("student data", studentGrades)

	//delete --> Remove Charlie from the map.
	delete(studentGrades, "Charlie")
	fmt.Println("student data", studentGrades)

	//. Iterate Over a Map
	for key, value := range studentGrades {
		fmt.Printf("Student name %s and Grade is %d\n", key, value)
	}
}
