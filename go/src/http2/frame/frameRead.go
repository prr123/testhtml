// frameRead
// author prr, azul software
// date 26 July 2022
// copyright 2022 prr, azul software
//
//
package main

import (
	"os"
	"fmt"
//	"golang.org/x/net/http2"
//	"github.com/prr123/http2"
//	"net/http"
	 http2 "http2/frame/http2Lib"
	"crypto/tls"
	"io"
)

func cvtFrameTyp(typ int) (string) {

	var framTyp [12]string

	if typ < 0 || typ > 9 { return "" }
	framTyp[0] = "Data"
	framTyp[1] = "Header"
	framTyp[2] = "Pri"
	framTyp[3] = "RST"
	framTyp[4] = "Settings"
	framTyp[5] = "Push"
	framTyp[6] = "Ping"
	framTyp[7] = "GoAway"
	framTyp[8] = "Wupd"
	framTyp[9] = "Cont"

	return framTyp[typ]
/*
	FrameData         FrameType = 0x0
	FrameHeaders      FrameType = 0x1
	FramePriority     FrameType = 0x2
	FrameRSTStream    FrameType = 0x3
	FrameSettings     FrameType = 0x4
	FramePushPromise  FrameType = 0x5
	FramePing         FrameType = 0x6
	FrameGoAway       FrameType = 0x7
	FrameWindowUpdate FrameType = 0x8
	FrameContinuation FrameType = 0x9
*/
}

func printErr(msg string, err error) {
	if err != nil {
		fmt.Printf("error %s: %v\n",msg, err)
		os.Exit(-1)
	}
	return
}

func dispFrame (frame http2.Frame) {

	fmt.Printf("fh type: %d: %s\n", frame.Header().Type, cvtFrameTyp(int(frame.Header().Type)))
	fmt.Printf("fh flag: %d\n", frame.Header().Flags)
	fmt.Printf("fh length: %d\n", frame.Header().Length)
	fmt.Printf("fh streamid: %d\n", frame.Header().StreamID)

//	headersframe := (frame1.(*http2.HeadersFrame))
//	fmt.Printf("stream ended? %v\n", headersframe.StreamEnded())
//	fmt.Printf("block fragment: %x\n", headersframe.HeaderBlockFragment())
}

func main () {
// generate with: openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.pem -days 365 -nodes
	cert, err := tls.LoadX509KeyPair("/home/peter/newca/server_crt.pem", "/home/peter/newca/server_key.pem")
	printErr("loadx509keypair", err)

	tlsCfg := &tls.Config{
    	Certificates: []tls.Certificate{cert},
    	NextProtos:   []string{"h2"},
	}

	l, err := tls.Listen("tcp", ":8787", tlsCfg)
	printErr("listen", err)
	defer l.Close()

	conn, err := l.Accept()
	printErr("conn accept", err)
	defer conn.Close()

//	fmt.Printf("conn: %v\n", conn)
	fmt.Printf("local: %v remote %v\n\n", conn.LocalAddr(), conn.RemoteAddr())
	const preface = "PRI * HTTP/2.0\r\n\r\nSM\r\n\r\n"
	b := make([]byte, len(preface))
	_, err = io.ReadFull(conn, b)
	printErr("Readfull", err)

	if string(b) != preface {
		printErr("preface", fmt.Errorf("string not preface!"))
	}

	framer := http2.NewFramer(conn, conn)
//	fmt.Printf("framer: %v\n", framer)
	fmt.Println("*** frame 0 ***")
	frame, err := framer.ReadFrame()

	printErr("read frame", err)
	fmt.Printf("frame: %v\n", frame)
	dispFrame(frame)
//	p := make([]byte, 24 + 9)
//    buf := make([]byte, 512)
/*
    for {
        n, err := conn.Read(buf)
        if err == io.EOF {
            break
        }

	_, err = io.ReadFull(conn,p)
	printErr("read frame byte", err)
	fmt.Printf("p [%d]: %x\n", len(p), p)
*/

	fmt.Println("\n*** frame 1 ***")
	frame1, err := framer.ReadFrame()
	printErr("read frame", err)
	fmt.Printf("frame 1: %v\n", frame1)
	dispFrame(frame1)

	fmt.Println("\n*** frame 2 ***")
	frame2, err := framer.ReadFrame()
	printErr("read frame", err)
	fmt.Printf("frame 2: %v\n", frame2)
	dispFrame(frame2)

}
