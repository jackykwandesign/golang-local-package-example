package main

import (
	"fmt"

	"jacky.com/dddtest/patient"
	"jacky.com/dddtest/twoLayerFolder/subtwolayerfolder"
)

func main() {
	fmt.Println("HelloWorld")
	name, err := patient.NewName("JackyKwanasdasdasdasdasdasdasdasdasd")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(name.String())
	}

	newname, err := subtwolayerfolder.NewName("JackyKwanasdasdasdasdasdasdasdasdasd")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(newname.String())
	}
}
