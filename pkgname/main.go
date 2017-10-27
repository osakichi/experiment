package main

import "bufio"
import "encoding/json"
import "os"

func main() {
	var data interface{}
	reader := bufio.NewReader(os.Stdin)
	json, _, _ := reader.ReadLine()
	json.Unmarshal(json, &data)
}

// Assigning a variable with the same name as the package name
// results in a compile error.
