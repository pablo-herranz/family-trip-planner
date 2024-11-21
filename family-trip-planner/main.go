package main
import "fmt"

var HEADER string = "Welcome to Family Trip Planner!!\n"
var OPENTRIPMAP_APIKEY string = "5ae2e3f221c38a28845f05b63e95772af969fc4aa7e42b827ccb9e29"

func main() {
	var destination string
	var duration int

	fmt.Println(HEADER)
	fmt.Print("Please, let us know where you want to travel: ")
	fmt.Scan(&destination)
	fmt.Print("Nice! How many days would you like to stay there: ")
	fmt.Scan(&duration)
	fmt.Println("You would like to stay", duration, "days in", destination,".")
}