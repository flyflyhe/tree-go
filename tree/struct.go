package main

import (
	"fmt"
)

type S struct {
	Sli map[int]S
}

func  (s *S) init ()  {
	fmt.Println("init")
}

func main() {
	s := new(S)
	s.init()
	s.Sli = make(map[int]S)
	
	if _, ok := s.Sli[0]; ok {
		fmt.Println("gg")
	}

	str := "test"
	strRune := []rune(str)

	if (string(strRune[4:]) == "") {
		fmt.Println("可以")		
	}
}