package main

import "fmt"

func main() {
	fmt.Println("Array----------------------------------------------")
	//Array
	var a [5]int
	fmt.Println("emp:", a)
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	fmt.Println("Slice----------------------------------------------")
	//Slice
	s := make([]string, 3)
	fmt.Println("emp:", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	fmt.Println("len:", len(s))

	fmt.Println("Map----------------------------------------------")
	//Map
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)
	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len:", len(m))
	delete(m, "k2")
	fmt.Println("map:", m)
	//if the key was present in the map
	_, prs := m["k2"]
	fmt.Println("is the key present :", prs)

	fmt.Println("Range----------------------------------------------")
	//Range
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
