package main

import "os"

func main() {
	file, _ := os.Open("./x.txt")
	defer file.Close()
	var tmp = make([]byte, 128)
	n, err := file.Read(tmp)
	if err

}
