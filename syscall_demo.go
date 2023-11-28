package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func openFile() {
	filename := "bytes.html"
	// open in read-only mode
	flags := syscall.O_RDONLY
	mode := uint32(0666) // file permission --> rw-rw-rw-

	fd, err := syscall.Open(filename, flags, mode)
	if err != nil {
		fmt.Printf("Error !!!: %s\n", err)
		return
	}

	defer syscall.Close(fd)
	fmt.Printf("Opened file: %s\n", filename)
}

func getSysinfo() {
	/*
		var info syscall.Sysinfo_t
		if err := syscall.Sysinfo(&info); err != nil {
			panic(err)
		}

		fmt.Printf("Total RAM: %v bytes\n", info.Totalram)
		fmt.Printf("Free RAM: %v bytes\n", info.Freeram)
		fmt.Printf("Uptime: %v seconds\n", info.Uptime)
	*/
}

func createProcess() {
	// Run ls via syscall
	/*
		err := syscall.Exec("/bin/dir", []string{"dir"}, syscall.Environ())
		if err != nil {
			fmt.Println("Error !!!: ", err)
			return
		}

		fmt.Println("Process created successfully: ")
	*/
}

func killProcess() {
	pid := os.Getpid()
	fmt.Println("Current process PID: ", pid)

	handle, err := syscall.GetCurrentProcess()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	var exitcode uint32 = 0
	err = syscall.TerminateProcess(handle, exitcode) // Ctrl + C

	time.Sleep(3 * time.Second)
	fmt.Println("Process terminated successfully")
}

func main() {
	// openFile()
	// getSysinfo()
	// createProcess()
	killProcess()
}
