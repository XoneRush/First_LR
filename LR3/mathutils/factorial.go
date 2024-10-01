package mathutils

func Factorial(num int) int{
	if num== 0 {return 1
	}else {
	res := 1
	for i := 1; i <= num; i++{
		res *= i
	}
	return res
	}
}