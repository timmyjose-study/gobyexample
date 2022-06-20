package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["hello"] = len("hello")
	m["world"] = len("world")
	m["nice"] = len("nice")
	m["to"] = len("to")
	m["meet"] = len("meet")
	m["you"] = len("you")

	fmt.Println("m:", m)

	for key, val := range m {
		fmt.Printf("%v => %v\n", key, val)
	}

	if v, ok := m["foo"]; !ok {
		fmt.Printf("%s not found in map\n", "foo")
	} else {
		fmt.Println(v)
	}

	if v, ok := m["world"]; !ok {
		fmt.Printf("%s not found in map\n", "world")
	} else {
		fmt.Println(v)
	}

	delete(m, "world")

	if v, ok := m["world"]; !ok {
		fmt.Printf("%s not found in map\n", "world")
	} else {
		fmt.Println(v)
	}

	foo := map[int]string{1: "hola", 2: "mundo", 3: "como", 4: "estas", 5: "ese?"}
	for k, v := range foo {
		fmt.Printf("%v => %v\n", k, v)
	}
}