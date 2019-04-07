package main

import (
	"flag"
	"fmt"
	"time"
)

var Occupied = "Occupied"
var NotOccupied = "NotOccupied"

type Cashiers struct {
	Id     int
	Status string
}

func (c *Cashiers) Serve(cust *Customer, manqueue chan<- *Cashiers) {
	c.Status = Occupied
	fmt.Println(time.Now(), "----> cashier %d started serving customer %d", c.Id, cust.Id)
	time.Sleep(time.Second * cust.WillTakeTime)
	c.Status = NotOccupied
	fmt.Println(time.Now(), "----> cashier %d completed customer %d", c.Id, cust.Id)
	manqueue <- c
}

type Customer struct {
	Id           int
	WillTakeTime time.Duration
	Status       string
}

func QManager(manqueue chan *Cashiers) {

}

func main() {
	numCashiers := flag.Int("numCashiers", 1, "an int")
	numCustomers := flag.Int("numCustomers", 10, "an int")
	timePerCustomer := flag.Int("timePerCustomer", 3, "an int")

	time.Sleep(time.Second)
	fmt.Printf("hello world %d %d %d", numCashiers, numCustomers, timePerCustomer)
}
