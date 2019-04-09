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
	fmt.Printf(time.Now().Format("2006-01-02 03:04:05")+" --> Cashier %d: Customer %d Started\n", c.Id, cust.Id)
	time.Sleep(time.Second * cust.WillTakeTime)
	c.Status = NotOccupied
	fmt.Printf(time.Now().Format("2006-01-02 03:04:05")+" --> Cashier %d: Customer %d Completed\n", c.Id, cust.Id)
	manqueue <- c
}

type Customer struct {
	Id           int
	WillTakeTime time.Duration
}

func QManager(manqueue chan *Cashiers, custqueue chan *Customer, numCashiers int) {
	fmt.Printf(time.Now().Format("2006-01-02 03:04:05") + " --> Bank simulations started\n")
	for {
		cashier := <-manqueue
		cust, more := <-custqueue
		if more {
			go cashier.Serve(cust, manqueue)
		} else {
			for i := 1; i < numCashiers; i++ {
				_ = <-manqueue
			}
			fmt.Printf(time.Now().Format("2006-01-02 03:04:05") + " --> Bank Simulated Ended\n")
			return
		}
	}

}

func main() {
	numCashiers := flag.Int("numCashiers", 2, "an int")
	numCustomers := flag.Int("numCustomers", 20, "an int")
	timePerCustomer := flag.Int("timePerCustomer", 3, "an int")

	flag.Parse()

	manqueue := make(chan *Cashiers, *numCashiers)
	custqueue := make(chan *Customer, *numCustomers)
	for i := 1; i <= *numCashiers; i++ {
		manqueue <- &Cashiers{
			Id:     i,
			Status: NotOccupied,
		}
	}
	for i := 1; i <= *numCustomers; i++ {
		custqueue <- &Customer{
			Id:           i,
			WillTakeTime: time.Duration(*timePerCustomer),
		}
	}
	close(custqueue)
	QManager(manqueue, custqueue, *numCashiers)
}
