package mathutils

import "fmt"

func ArrOfInts() [5]int {
	var arr [5]int

	for i := range arr {
		arr[i] = i
	}

	return arr
}

func List() {
	var array [3]int

	array[0] = 1
	array[1] = 2
	array[2] = 3

	list := array[0:2]
	fmt.Println("Изначальный срез: ", list)

	list = append(list, 4)
	fmt.Println("Добавление в срез цифры 4: ", list)

	list = append(list[:1], list[2:]...)
	fmt.Println("После удаления 2го элемента: ", list)

}
