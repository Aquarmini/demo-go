package main

var a = "w3cschool菜鸟教程";

var b string = "w3cschool.cc"
var c bool
var num int = 0;

var arr [10] int;

func main() {
	println(a, b, c)

	for j := 0; j <= 10000; j++ {
		for i := 0; i <= 10000; i++ {
			num++;
		}
	}
	println(num);

	arr1 := []int{1, 2, 3};
	for _, v := range arr1 {
		println(v);
	}

	arr1 = append(arr1, 5);
	for _, v := range arr1 {
		println(v);
	}
}