package main

import (
	"syscall"
    "fmt"
)

func main() {
	
    disk := "\\\\.\\PHYSICALDRIVE3"

	fmt.Printf("Reading data[1024] from Disk: %v \n",disk)
    fd, err := syscall.Open(disk, syscall.O_RDONLY, 0)

    if err != nil {
        fmt.Print("Open Failed:%v ",err.Error(), "\n")
        return
    } else {
		fmt.Printf("Disk Opened : Handle [%v] \n",fd)
	}

    buffer := make([]byte, 1024*1024)
	
	var numOfBytesRead int
	
    for { 
		numOfBytesRead, err = syscall.Read(fd, buffer)
		if err != nil {
			fmt.Printf("Read Failed: %v \n", err)
		}
		if numOfBytesRead <= 0 {
			fmt.Println("Found End of disk. exiting...")
			break			
		}
		
		fmt.Printf("Number of bytes read: %d\n", numOfBytesRead)
		//fmt.Printf("Buffer: %v\n", buffer)
	}
	
	err = syscall.Close(fd)

    if err != nil {
        fmt.Printf("Close Failed:%v", err)
    }

	/*
			if buffer[510] != 0x55 || buffer[511] != 0xaa {
				fmt.Println("MBR not detected on physical drive")
			} else {
				fmt.Println("MBR detected on physical drive")
			}
	*/
	
	/*
	//Seek functionality
	fmt.Println("\nSeek to offset 57344")
	offset, err := syscall.Seek(fd, 57344, 0)
	fmt.Printf("Offset after seek(): %v\n", offset)

    numread, err = syscall.Read(fd, buffer)
    if err != nil {
		fmt.Printf("Read Failed: %v \n", err)
    }
	
	fmt.Printf("Numer of bytes read: %d\n", numread)
    fmt.Printf("Buffer: %v\n", buffer)
	*/

}