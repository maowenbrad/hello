package main

import (
	"fmt"
	"github.com/cbroglie/mustache"
)

func main() {

	fmt.Printf("Hello, Matthew!!!\n")

	output, err := mustache.Render("Hello, {{c}}!!\n", map[string]string{"c": "Matthew"})

	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf(output)
	}

    // folderPath, err := osext.ExecutableFolder()
    // if err != nil {
    //     log.Fatal(err)
    // }
    // fmt.Println(folderPath)


}