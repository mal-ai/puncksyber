package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	usb "github.com/google/gousb"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"

	"go.bug.st/serial/enumerator"
)

func main() {
	fmt.Println("*********************cyberpuk*************************")

	ports, err := enumerator.GetDetailedPortsList()
	var cyberpunk *enumerator.PortDetails
	if err != nil {
		fmt.Println("Erro of getting ports")
	}
	for _, port := range ports {
		if port.IsUSB {
			if strings.Contains(port.Name, "1420") {
				cyberpunk = port
			}
		}

	}
	fmt.Printf("   USB Name   %s\n", cyberpunk.Name)
	fmt.Printf("   USB ID     %s:%s\n", cyberpunk.VID, cyberpunk.PID)
	fmt.Printf("   USB serial %s\n", cyberpunk.SerialNumber)
	fmt.Println("******************************************************")

	ctx := usb.NewContext()
	defer ctx.Close()
	vid, pid := usb.ID(0x1A86), usb.ID(0x7523)
	devs, err := ctx.OpenDevices(func(desc *usb.DeviceDesc) bool {
		// this function is called for every device present.
		// Returning true means the device should be opened.
		return desc.Vendor == vid && desc.Product == pid
	})

	// All returned devices are now open and will need to be closed.
	for _, d := range devs {
		defer d.Close()
	}
	if err != nil {
		log.Fatalf("OpenDevices(): %v", err)
	}
	if len(devs) == 0 {
		log.Fatalf("no devices found matching VID %s and PID %s", vid, pid)
	}

	dev := devs[0]
	fmt.Println()
	fmt.Println("connected to device:  ", dev.Desc)

	cfg, err := dev.Config(1)
	if err != nil {
		log.Fatalf("%s.Config(0): %v", dev, err)
	}
	defer cfg.Close()

	fmt.Printf("avalibe config:        %s=%s:", cfg.Desc, cfg.String())
	fmt.Println()

	intf, err := cfg.Interface(0, 0)
	if err != nil {
		log.Fatalf("%s.Interface(0, 0): %v", cfg, err)
	}
	defer intf.Close()
	fmt.Println("avalibe interface:    ", intf.String())

	epOut, err := intf.OutEndpoint(2)
	if err != nil {
		log.Fatalf("%s.OutEndpoint: %v", intf, err)
	}

	fmt.Println("open endpont ep0out:  ", epOut.String())
	//here ven dor specific requests

	//req 1
	resp, err2 := dev.Control(64, 161, 924, 0, nil)
	if err2 != nil {
		fmt.Println("Error 1 request =", err2)
	}
	fmt.Println("req1 Resp==", resp)
	//	time.Sleep(600 * time.Millisecond)

	//req 2
	packet := "40a19cc38ad90000"
	dt, _ := hex.DecodeString(packet)
	fmt.Printf("setup % x  -  ", dt)

	resp, err2 = dev.Control(64, 161, 50076, 55690, dt)
	if err2 != nil {
		fmt.Println("Setup packet error:", err2)
	}
	fmt.Println("Setup packet 2 Ok ==", resp)
	//	time.Sleep(600 * time.Millisecond)

	//req 3
	resp, err2 = dev.Control(64, 154, 3884, 0, nil)
	if err2 != nil {
		fmt.Println("Error 3 request =", err2)
	}
	fmt.Println("req3 Resp ==", resp)
	//req 4
	//	time.Sleep(600 * time.Millisecond)

	packet4 := "409a2c0f07000000"
	dt4, _ := hex.DecodeString(packet4)
	fmt.Printf("setup % x  -  ", dt4)

	resp, err2 = dev.Control(64, 154, 3884, 7, dt4)
	if err2 != nil {
		fmt.Println("Setup packet error:", err2)
	}
	fmt.Println("Setup packet 4 Ok ==", resp)
	//	time.Sleep(600 * time.Millisecond)

	//req 5
	resp, err2 = dev.Control(64, 164, 223, 0, nil)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("req5 Resp ==", resp)
	// req 6
	//	time.Sleep(600 * time.Millisecond)

	packet6 := "40a4df0000000000"
	dt6, _ := hex.DecodeString(packet6)
	fmt.Printf("setup % x  -  ", dt6)

	resp, err2 = dev.Control(64, 164, 223, 0, dt6)
	if err2 != nil {
		fmt.Println("Setup packet error:", err2)
	}
	fmt.Println("Setup packet 6 Ok ==", resp)
	//	time.Sleep(600 * time.Millisecond)

	//req 7
	resp, err2 = dev.Control(64, 164, 159, 0, nil)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("req7 Resp==", resp)
	// req 8
	//	time.Sleep(600 * time.Millisecond)

	packet8 := "40a49f0000000000"
	dt8, _ := hex.DecodeString(packet8)
	fmt.Printf("setup % x  -  ", dt8)

	resp, err2 = dev.Control(64, 164, 159, 0, dt8)
	if err2 != nil {
		fmt.Println("Setup packet error:", err2)
	}
	fmt.Println("Setup packet 8 Ok ==", resp)
	//	time.Sleep(600 * time.Millisecond)

	//req 9
	resp, err2 = dev.Control(192, 149, 1798, 0, nil)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("req9 Resp==", resp)
	// req 10
	//	time.Sleep(600 * time.Millisecond)

	packet10 := "C095060700000200"
	dt10, _ := hex.DecodeString(packet10)
	fmt.Printf("setup % x  -  ", dt10)

	resp, err2 = dev.Control(192, 149, 1798, 0, dt10)
	if err2 != nil {
		fmt.Println("Setup packet error:", err2)
	}
	fmt.Println("Setup packet 10 Ok ==", resp)
	//	time.Sleep(600 * time.Millisecond)

	//req 11
	resp, err2 = dev.Control(64, 154, 10023, 0, nil)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("req11 Resp==", resp)
	//req 12
	//	time.Sleep(600 * time.Millisecond)

	packet12 := "409a272700000000"
	dt12, _ := hex.DecodeString(packet12)
	fmt.Printf("setup % x  -  ", dt12)
	resp, err2 = dev.Control(64, 154, 10023, 0, dt12)
	if err2 != nil {
		fmt.Println("Setup packet error:", err2)
	}
	fmt.Println("Setup packet 12 Ok ==", resp)
	//	time.Sleep(600 * time.Millisecond)

	//req 13
	resp, err2 = dev.Control(192, 149, 1798, 0, nil)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("req13 Resp==", resp)
	// потом послушать что ответит 0000014
	//req 14
	//	time.Sleep(600 * time.Millisecond)

	packet14 := "c095060700000200"
	dt14, _ := hex.DecodeString(packet14)
	fmt.Printf("setup % x  -  ", dt14)
	resp, err2 = dev.Control(192, 149, 1798, 0, dt14)
	if err2 != nil {
		fmt.Println("Setup packet error:", err2)
	}
	fmt.Println("Setup packet 14 Ok ==", resp)
	//	time.Sleep(600 * time.Millisecond)

	//req 15
	resp, err2 = dev.Control(64, 154, 10023, 0, nil)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("req15 Resp==", resp)
	//req16
	//	time.Sleep(600 * time.Millisecond)

	packet16 := "409a27270000"
	dt16, _ := hex.DecodeString(packet16)
	fmt.Printf("setup % x  -  ", dt16)
	resp, err2 = dev.Control(64, 154, 10023, 0, dt16)
	if err2 != nil {
		fmt.Println("Setup packet error:", err2)
	}
	fmt.Println("Setup packet 16 Ok ==", resp)
	//	time.Sleep(600 * time.Millisecond)

	//req 17
	resp, err2 = dev.Control(104, 154, 4882, 52355, nil)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("req17 Resp==", resp)
	//req 18
	//	time.Sleep(600 * time.Millisecond)

	packet18 := "409a121383cc0000"
	dt18, _ := hex.DecodeString(packet18)
	fmt.Printf("setup % x  -  ", dt18)
	resp, err2 = dev.Control(64, 154, 4882, 52355, dt18)
	if err2 != nil {
		fmt.Println("Setup packet error:", err2)
	}
	fmt.Println("Setup packet 18 Ok ==", resp)
	//	time.Sleep(600 * time.Millisecond)

	//req 19
	resp, err2 = dev.Control(64, 164, 191, 0, nil)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("req19 Resp==", resp)
	//req 20
	//	time.Sleep(600 * time.Millisecond)

	packet20 := "40a4bf0000000000"
	dt20, _ := hex.DecodeString(packet20)
	fmt.Printf("setup % x  -  ", dt20)
	resp, err2 = dev.Control(64, 164, 191, 0, dt20)
	if err2 != nil {
		fmt.Println("Setup packet error:", err2)
	}
	fmt.Println("Setup packet 20 Ok ==", resp)
	//	time.Sleep(600 * time.Millisecond)

	//req 21
	resp, err2 = dev.Control(192, 149, 1798, 0, nil)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("req21 Resp==", resp)
	//req 22
	//	time.Sleep(600 * time.Millisecond)

	packet22 := "c095060700000200"
	dt22, _ := hex.DecodeString(packet22)
	fmt.Printf("setup % x  -  ", dt22)
	resp, err2 = dev.Control(192, 149, 1798, 0, dt22)
	if err2 != nil {
		fmt.Println("Setup packet error:", err2)
	}
	fmt.Println("Setup packet 22 Ok ==", resp)
	//	time.Sleep(600 * time.Millisecond)

	//req 23
	resp, err2 = dev.Control(64, 154, 10023, 0, nil)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("req23 Resp==", resp)
	//req 24
	packet24 := "409a272700000000"
	dt24, _ := hex.DecodeString(packet24)
	fmt.Printf("setup % x  -  ", dt24)
	resp, err2 = dev.Control(64, 154, 10023, 0, dt24)
	if err2 != nil {
		fmt.Println("Setup packet error:", err2)
	}
	fmt.Println("Setup packet 24 Ok ==", resp)
	//	time.Sleep(600 * time.Millisecond)

	//req 25
	resp, err2 = dev.Control(192, 149, 1798, 0, nil)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("req25 Resp==", resp)
	//req 26
	packet26 := "c095060700000200"
	dt26, _ := hex.DecodeString(packet26)
	fmt.Printf("setup % x  -  ", dt26)
	resp, err2 = dev.Control(192, 149, 1798, 0, dt26)
	if err2 != nil {
		fmt.Println("Setup packet error:", err2)
	}
	fmt.Println("Setup packet 26 Ok ==", resp)
	//	time.Sleep(600 * time.Millisecond)

	//resp 27
	resp, err2 = dev.Control(64, 154, 10023, 0, nil)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("req27 Resp==", resp)
	//resp 28
	packet28 := "409a272700000000"
	dt28, _ := hex.DecodeString(packet28)
	fmt.Printf("setup % x  -  ", dt28)
	resp, err2 = dev.Control(64, 154, 10023, 0, dt28)
	if err2 != nil {
		fmt.Println("Setup packet error:", err2)
	}
	fmt.Println("Setup packet 28 Ok ==", resp)
	//	time.Sleep(600 * time.Millisecond)

	//req 29
	resp, err2 = dev.Control(64, 164, 255, 0, nil)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("req29 Resp==", resp)
	//req 30
	packet30 := "40a4ff0000000000"
	dt30, _ := hex.DecodeString(packet30)
	fmt.Printf("setup % x  -  ", dt30)
	resp, err2 = dev.Control(64, 164, 255, 0, dt30)
	if err2 != nil {
		fmt.Println("Setup packet error:", err2)
	}
	fmt.Println("Setup packet 30 Ok ==", resp)
	//	time.Sleep(600 * time.Millisecond)

	//req 31
	resp, err2 = dev.Control(192, 149, 1798, 0, nil)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("req31 Resp==", resp)
	//req 32
	packet32 := "c095060700000200"
	dt32, _ := hex.DecodeString(packet32)
	fmt.Printf("setup % x  -  ", dt32)
	resp, err2 = dev.Control(192, 149, 1798, 0, dt32)
	if err2 != nil {
		fmt.Println("Setup packet error:", err2)
	}
	fmt.Println("Setup packet 32 Ok ==", resp)
	//	time.Sleep(600 * time.Millisecond)

	//req 33
	resp, err2 = dev.Control(64, 154, 10023, 0, nil)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("req33 Resp==", resp)
	//req 34
	packet34 := "409a272700000000"
	dt34, _ := hex.DecodeString(packet34)
	fmt.Printf("setup % x  -  ", dt34)
	resp, err2 = dev.Control(64, 154, 10023, 0, dt34)
	if err2 != nil {
		fmt.Println("Setup packet error:", err2)
	}
	fmt.Println("Setup packet 34 Ok ==", resp)
	//	time.Sleep(600 * time.Millisecond)

	//--
	var totaCpulUsage int8
	var totaMemoryUsage int8
	var totaMemoryCacheUsage int8

	for {

		before, err := cpu.Get()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return
		}
		time.Sleep(time.Duration(1) * time.Second)
		after, err := cpu.Get()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return
		}
		total := float64(after.Total - before.Total)
		totaCpulUsage = int8((float64(after.User-before.User) / total * 100) + (float64(after.System-before.System) / total * 100))
		memory, err := memory.Get()
		if err != nil {
			fmt.Println(err)
		}
		totaMemoryUsage = int8(float64(memory.Used) / float64(memory.Total) * 100)
		totaMemoryCacheUsage = int8(float64(memory.Cached) / float64(memory.Total) * 100)

		buf := new(bytes.Buffer)
		setstr := "SETCPUGPURAM"
		setArr := strings.Split(setstr, "")
		for _, val := range setArr {
			err = binary.Write(buf, binary.LittleEndian, []byte(val))
		}
		if err != nil {
			fmt.Println(err)
		}
		setcommArray := buf.Bytes()
		fmt.Printf("*-* setcommArray= % x\n", setcommArray)

		_, err = epOut.Write(setcommArray)

		if err != nil {
			fmt.Println(err)
		}
		buff := new(bytes.Buffer)
		cpugpu := []int8{totaCpulUsage, totaMemoryCacheUsage, totaMemoryUsage}
		for _, val := range cpugpu {
			err = binary.Write(buff, binary.LittleEndian, val)
		}
		if err != nil {
			fmt.Println(err)
		}
		cpugpuArray := buff.Bytes()

		_, err = epOut.Write(cpugpuArray)
		if err != nil {
			fmt.Println(err)
		}

		buff2 := new(bytes.Buffer)
		endArr := []int8{13, 10}

		for _, val := range endArr {
			err = binary.Write(buff2, binary.LittleEndian, val)

		}
		if err != nil {
			fmt.Println(err)
		}
		ends := buff2.Bytes()
		_, err = epOut.Write(ends)
		if err != nil {
			fmt.Println(err)
		}
	}

}
