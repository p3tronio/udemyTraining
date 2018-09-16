package main

import "fmt"

func main() {

	fmt.Println(project_euler_problem_01()) // 233168
	fmt.Println(project_euler_problem_02())


}

func project_euler_problem_01() int{

	var total int
	for i:= 999; i > 0; i--{

	if(i%3 == 0){ total += i } // add if multiple of 3
	if(i%5 == 0 && i%3 != 0){ total += i } // add if multiple of 5 but not if multiple of 3

	}

	return total
}

func project_euler_problem_02() int {

	var a int
	var b int
	var c int
	var total int
	a = 1
	b = 1
	c = 1

	for c < 4000000	{

		c = a + b

		if c%2 == 0 {

			total += c

		}
		b = a
		a = c

	}

	return total

}







