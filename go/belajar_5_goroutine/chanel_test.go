package belajar5goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateCHannel(t *testing.T)  {
	channel := make(chan string)
	defer close(channel)

	// channel <- "Lutfan"
	// message := <- channel
	// fmt.Println(<- channel)
	// fmt.Println(message)

	go func(){
		time.Sleep(2 * time.Second)
		channel <- "Lutfan Zainul Haq"
		fmt.Println("Success Mengirim Data ke Channel")
	}()

	data := <- channel
	fmt.Println(data)

	time.Sleep(3 * time.Second)
}

func GiveMeResponse(channel chan string)  {
	time.Sleep(2 * time.Second)
	channel <- "Lutfan Zainul Haq"
}

// func TestChannelASParameter()  {
// 	channel := make(chan string)
// 	defer close(channel)

// 	go GiveMeResponse(channel)

// 	data := <- channel
// 	fmt.Println(data)

// 	time.Sleep(3 * time.Second)
// }

func TestBufferedChannel(t *testing.T)  {
	channel := make(chan string, 3)
	defer close(channel)

	// channel <- "Lutfan"
	// channel <- "Zainul"
	// channel <- "Haq"

	// fmt.Println(<- channel)
	// fmt.Println(<- channel)
	// fmt.Println(<- channel)

	go func ()  {
		channel <- "Lutfan"
		channel <- "Zainul"
		channel <- "Haq"
	}()
	
	go func ()  {
		fmt.Println(<- channel)
		fmt.Println(<- channel)
		fmt.Println(<- channel)
	}()
	
	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T)  {
	channel := make(chan string)

	go func ()  {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()
	
	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	// select {
	// 	case data := <- channel1:
	// 		fmt.Println("Data dari channel 1 ",data)
	// 	case data := <- channel2:
	// 		fmt.Println("Data dari channel 2",data)
	// }

	counter := 0
	for  {
		select {
		case data := <- channel1:
			fmt.Println("Data dari channel 1 ",data)
			counter++
		case data := <- channel2:
			fmt.Println("Data dari channel 2",data)
			counter++
		}

		if counter == 2 {
			break
		}
	}

}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for  {
		select {
		case data := <- channel1:
			fmt.Println("Data dari channel 1 ",data)
			counter++
		case data := <- channel2:
			fmt.Println("Data dari channel 2",data)
			counter++
		default:
			fmt.Println("menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}

