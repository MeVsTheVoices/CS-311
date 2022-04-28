package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MAX_DRIVE_TIME    = 10
	MAX_CUSTOMER_TIME = 4
)

// Have driver pick up, then drop off passenger and make driver available again
func drive(customerName string, drivers chan string) {
	//read in a driver from the channel parameter, causes a pause if none are available
	currentDriver := <-drivers
	fmt.Printf("%s has picked up %s\n", currentDriver, customerName)
	//pause a few seconds
	time.Sleep(time.Second * time.Duration(rand.Intn(MAX_DRIVE_TIME)))
	fmt.Printf("%s has finished driving %s\n", currentDriver, customerName)
	//put that driver back in the queue to pick up someone else
	drivers <- currentDriver
}

func main() {
	fmt.Println("Ride-hailing Service Simulation")

	//create a couple buffered channels so we can query whether we still have customers in the queue
	waitingCustomers := make(chan string, 20)
	availableDrivers := make(chan string, 20)

	//populate drivers
	availableDrivers <- "Joel"
	availableDrivers <- "James"
	availableDrivers <- "Rebecca"
	availableDrivers <- "Francesca"

	//add a list of distinct people waiting to be picked up
	waitingCustomers <- "Frankie"
	waitingCustomers <- "Charles"
	waitingCustomers <- "DeLouise"
	waitingCustomers <- "Jessica"
	waitingCustomers <- "Bella"
	waitingCustomers <- "Justine"
	waitingCustomers <- "Bobby"
	waitingCustomers <- "Paul"

	//iterate through waiting customers, pause a few seconds between each showing up to wait
	//keep checking whether there's any that are still in the channel
	for a, _ := <-waitingCustomers; len(waitingCustomers) > 0; a, _ = <-waitingCustomers {
		//announce who's waiting, start a goroutine that will wait for an available driver
		fmt.Printf("%s is waiting to be picked up\n", a)
		go drive(a, availableDrivers)
		//pause between to give a little reality to things
		time.Sleep(time.Second * time.Duration(rand.Intn((MAX_CUSTOMER_TIME))))
	}

	fmt.Println("Simulation is over")
}
