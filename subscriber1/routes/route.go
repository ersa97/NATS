package routes

import (
	"fmt"
	"log"

	"github.com/ersa97/NATS/subscriber1/controllers"
)

func HandleRequest(route string, param interface{}, data interface{}) {
	log.Println("Accepted")
	log.Println(route)

	switch route {
	case "testing":
		controllers.Testing(data)
	}
	//create a switch case to handle the command
	//and direct it into a controller where it will be handled logically

	log.Println("Closing")
	fmt.Println()

}
