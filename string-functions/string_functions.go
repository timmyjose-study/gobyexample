package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Contains:", strings.Contains("test", "es"))
	fmt.Println("Count:", strings.Count("test", "t"))
	fmt.Println("HasPrefix:", strings.HasPrefix("test", "te"))
	fmt.Println("HasSuffix:", strings.HasSuffix("test", "est"))
	fmt.Println("Index:", strings.Index("test", "st"))
	fmt.Println("Join:", strings.Join([]string{"hello", "hola", "Ni hao"}, ":"))
	fmt.Println("Repeat:", strings.Repeat("hello", 10))
	fmt.Println("Split:", strings.Split("helloxworldxagainx", "x"))
	fmt.Println("ToLower:", strings.ToLower("Hello, world"))
	fmt.Println("ToUpper:", strings.ToUpper("Hello, world"))

}