package main
import (
	"fmt"
)
func main() {
	language := []string{"java","python","javascript"};
	for index,item := range language {
		fmt.Printf("%d. %s\n",index+1,item)
	}
	for _,item := range language {
		fmt.Printf("%s, ",item)
	}
}
