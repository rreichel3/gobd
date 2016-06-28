// GOBD initial round
package main
import (
	"io"
	"log"
	"os"
	"os/signal"
	"github.com/tarm/serial"
)


func readData(obdConn io.ReadWriteCloser, data chan byte) {
	writeData("atma")
	buffer := make([]byte, 1)	
	for {
		n, err := obdConn.read(buffer)
		data <- buffer[0]
	}
}
func writeData(obdConn io.ReadWriteCloser, cmd string) {
	cmd += "\n"
	_, err := obdConn.write(cmd)
	if err != nil {
		log.Fatal(err)
	}
}


func main() {
	// Initialize serial reader
	// Initialize data structures for the OBD Data
	// Spin off data reading on goroutine
	// Write data across channel
	// Process/Print data
	conf := &serial.Config{Name:"/dev/ttyUSBart0", Baud: 4800}
	obdConn, err := serial.OpenPort(conf)
	if err != nil {
		log.Fatal(err)
	}
	// This is a large buffer. Hopefully its temporarily like this and can be made smaller
	data := make(chan byte, 1000)
	defer close(data)
	////////////////////////////////////
	// Init ELM Device here
	////////////////////////////////////
	writeData("atz")
	writeData("ath1")
	writeData("ate0")
	writeData("atal")
	////////////////////////////////////
	// May also need to change baud rate
	////////////////////////////////////
	go readData(obdConn, data)
	i := 0
	for/*ever*/ {
		i++
		if (i%2 == 0) {
			fmt.Print(" ")
		}
		fmt.Print(<-data)
	}
}