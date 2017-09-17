// Author: Alangi Derick N
// LICENSE: MIT, all right reserved 2017
// Description: Entry point for the SQL analyzer tool

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	for {
		fmt.Printf("gosql-query-analyzer> ")

		reader := bufio.NewReader(os.Stdin)

		query, _ := reader.ReadString('\n')
		fmt.Println("query: ", query)
		// Enter Ctrl + C to quit.
	}

}
