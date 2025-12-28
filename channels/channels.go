package main

import (
	"fmt"
	"time"
)

// ====================================================================> Example for Buffered Channels
// Queue example where need to send multiple emails
// Single Channel example
// func emailSender(emailChan chan string, done chan bool) { // Without restriction signature
func emailSender(emailChan <-chan string, done chan<- bool) { // With restriction signature
	defer func() { done <- true }()

	// Before Restriction
	// <-done
	// emailChan <- "something@gmail.com"
	// Channel can send and receive as well. So we need to make sure type safety is handled properly so that we can limitize the scope of the channel's responsibility.
	// So for example here we are receiving emailChan so we need to restrict it to only receive it.
	// Similarly we do it for done channel to restrict it send only after receiving it as its purpose is to send completion flag to main function once the operation is done.

	// After Restriction
	// emailChan <- "something@gmail.com" // Giving error because email channel is now restricted to receive only
	// <-done // Giving error because done channel is now restricted to send only

	for email := range emailChan {
		fmt.Println("sending email to: ", email)
		time.Sleep(time.Second)
	}
}

//======================================================================> Example for Unbuffered Channels
// Fetching from second gorotuine to first goroutine
// func sum(result chan int, num1, num2 int) {
// 	result <- (num1 + num2)
// }

// Used for sending value from one goroutine to another one
// func processNum(numChan chan int) {
// 	for i := range numChan {
// 		fmt.Println("Processing number", i)
// 		time.Sleep(time.Second)
// 	}
// }

// Using channels to get similar behaviour we did with Wait Group
// goroutine synchronizer
// func task(done chan bool) {
// 	defer func() { done <- true }()
// 	fmt.Println("processing...")
// 	// done <- true // This will work as it will identify that the task is completed to the channel but sometime it can fail like before execution of this line any error comes up this line won't execute. So better to use it with defer keyword.
// }

//=====================================================================================================================================================

// Channels are a way by which multiple go-routines communicate with each other. Always use channels when dealing with small datasets if larger dataset then sync.WaitGroup as it gives controll on Add, Wait and Done functionality
func main() {
	// ========================================================================> Buffered Channels
	//-------------------> Single Channel Example

	// emailChan := make(chan string, 100) // This is buffered channel which means that until 100 size the channel won't be blocking but if value exceeds then those channel values will become blocking.
	// doneChan := make(chan bool)

	// go emailSender(emailChan, doneChan)

	// for i := 0; i < 10; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i) // 1. Non Blocking because of buffered size. 2. In order to add integer value in between string using "Sprintf"
	// }

	//fmt.Println("Done Sending")
	//close(emailChan) // Need to use "close" function on channel because if we don't close due to for range loop in email Sender will create infinite loop because of buffered channel, the done channel won't execute hence will create a deadlock.

	//<-doneChan // this will keep our main channel wait until our all emails are sent

	//-------------------> Multiple Channels Example

	channelOne := make(chan int)
	channelTwo := make(chan string)

	go func() {
		channelOne <- 10
	}()

	go func() {
		channelTwo <- "John"
	}()

	for i := 0; i < 2; i++ {
		select {
		case channelOneVal := <-channelOne:
			fmt.Println("Recieved Data from channel One: ", channelOneVal)
		case channelTwoVal := <-channelTwo:
			fmt.Println("Recieved Data from channel Two: ", channelTwoVal)
		}
	}

	// ==========================================================================> UnBuffered Channels
	// We can use it we want one by one functionality not any parallel activity because of the blocking purpose. At a time, it will operate only one job or one value. Until its received another task won't start working.

	//--------------------------------------------------------------------------------------------------
	// Using channels to get similar behaviour like we did with Wait Group
	// done := make(chan bool)
	// go task(done)

	//Either this
	// result := <-done // blocking here
	// fmt.Println(result)
	//Or
	// <-done // blocking here use it if not printing or using it

	//--------------------------------------------------------------------------------------------------
	// Receiving values from another goroutine
	// result := make(chan int)
	// go sum(result, 10, 20)

	// res := <-result // The reason we didn't use time.sleep is because sending and receiving line in channel are basically blocking which is partially true because in exceptional cases like buffered channel this is not the case.
	// fmt.Println(res)

	//--------------------------------------------------------------------------------------------------
	// Sending value from one goroutine to another go routine
	// By this way from one goroutine(main) we were able to send value to another goroutine(processNum) using channel
	// numChan := make(chan int)

	// go processNum(numChan)

	// numChan <- 5

	// for {
	// 	numChan <- rand.Intn(100)
	// }
	// time.Sleep(time.Second * 2)

	// ----------------------------------------------------------------------------------------------
	// messageChannel := make(chan string) // using "chan" keyword, we can make sure channel is getting used.

	// // Setting values inside channel
	// messageChannel <- "ping" // This is currently under blocking state and it can't be released until some receiver receives it from other end.

	// newData := <-messageChannel // Retrieve data from channel and assigned to variable.

	// fmt.Println(newData)
	// Deadlock is caused when multiple process are operating and our current task also captured by any process is holded due to another processor not releasing its exist captured task because of the completion not recieved.
	//Output: fatal error: all goroutines are asleep - deadlock!
}
