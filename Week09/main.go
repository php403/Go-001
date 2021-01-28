package main

import (
	"fmt"
	"net"
	"os"
	"syscall"
)

const (
	EPOLLET        = 1 << 31
	MaxEpollEvents = 32
)

func main() {
	var event syscall.EpollEvent
	var events [MaxEpollEvents]syscall.EpollEvent

	fd, err := syscall.Socket(syscall.AF_INET, syscall.O_NONBLOCK|syscall.SOCK_STREAM, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer syscall.Close(fd)

	if err = syscall.SetNonblock(fd, true); err != nil {
		fmt.Println("setnonblock1: ", err)
		os.Exit(1)
	}

	addr := syscall.SockaddrInet4{Port: 8888}
	copy(addr.Addr[:], net.ParseIP("0.0.0.0").To4())

	syscall.Bind(fd, &addr)
	syscall.Listen(fd, 10)

	epfd, e := syscall.EpollCreate1(0)
	if e != nil {
		fmt.Println("epoll_create1: ", e)
		os.Exit(1)
	}
	defer syscall.Close(epfd)

	event.Events = syscall.EPOLLIN
	event.Fd = int32(fd)
	if e = syscall.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, fd, &event); e != nil {
		fmt.Println("epoll_ctl: ", e)
		os.Exit(1)
	}


	for {
		nevents, e := syscall.EpollWait(epfd, events[:], -1)
		if e != nil {
			fmt.Println("epoll_wait: ", e)
			break
		}

		for ev := 0; ev < nevents; ev++ {
			if int(events[ev].Fd) == fd {
				connFd, _, err := syscall.Accept(fd)
				if err != nil {
					fmt.Println("accept: ", err)
					continue
				}
				syscall.SetNonblock(fd, true)
				event.Events = syscall.EPOLLIN | EPOLLET
				event.Fd = int32(connFd)
				if err := syscall.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, connFd, &event); err != nil {
					fmt.Print("epoll_ctl: ", connFd, err)
					os.Exit(1)
				}
			} else {
				go func() {
					fd1 := int(events[ev].Fd)
					defer syscall.Close(fd)
					var buf [32 * 1024]byte
					for {
						nbytes, e := syscall.Read(fd1, buf[:])
						if nbytes > 0 {
							fmt.Printf(">>> %s", buf)
							syscall.Write(fd1, buf[:nbytes])
							fmt.Printf("<<< %s", buf)
						}
						if e != nil {
							break
						}
					}
				}()
			}
		}

	}
	/*go func() {
		select {
		case msg := <- writeChan:
			syscall.Write(int(msg.fd), []byte(msg.msg))
			//syscall.Write(int(events[ev].Fd), buf[:nbytes])
		}
	}()

	go func() {
		select {
		case msg := <- readChan:
			nbytes, e := syscall.Read(int(events[ev].Fd), buf[:])
			if nbytes > 0 {
				fmt.Printf(">>> %s", buf)
			}

		}
	}()*/

}