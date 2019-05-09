package fileio

import "io/ioutil"

//WriteAFile
//write a line to a file
func WriteAFile() {
	name := "test.txt"
	contStr := "hello world"
	contByte := []byte(contStr)
	ioutil.WriteFile(name, contByte, 0644)
}
