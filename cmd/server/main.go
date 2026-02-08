package main

import "fmt"

//Run - is going to be responsibe for the instantiation and startup of the application
func Run() error {
	fmt.Println("starting up our application")
	return nil

}

func main() {
	fmt.Println("Go Rest API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
