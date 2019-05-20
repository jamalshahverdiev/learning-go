package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	os.Setenv("ANSVAULTPASS", "2NsV2ultP244")
	ansVaultPass := os.Getenv("ANSVAULTPASS")

	//	f, err := os.OpenFile("./vault_pass.txt", os.O_APPEND, 0644)
	//
	//	if err != nil {
	//		fmt.Println("Error: Unable to open ./vault_pass.txt")
	//		os.Exit(1)
	//	}
	//
	//	defer f.Close()

	//	_, err = f.Write([]byte(ansVaultPass))
	err := ioutil.WriteFile("./vault_pass.txt", []byte(ansVaultPass), 0666)
	if err != nil {
		fmt.Println("Error: Failed to write to ./vault_pass.txt")
		os.Exit(1)
	}

	//cmd := exec.Command("bash", "getVar.sh")
	//_, err := cmd.Output()

	//if err != nil {
	//	println(err.Error())
	//	return
	//}
}
