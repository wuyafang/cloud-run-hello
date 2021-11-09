// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"log"
)

func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)

	ip = strings.Split(localAddr.String(), ":")[0]
	log.Print(fmt.Sprintf("localIP is %s", ip))
	return
}

func GetRemoteIP(conn net.Conn) (ip string, err error) {
	remoteAddr := conn.RemoteAddr().(*net.TCPAddr)
	ip = strings.Split(remoteAddr.String(), ":")[0]
	log.Print(fmt.Sprintf("remoteIP is %s", remoteAddr.String()))
	log.Print(fmt.Sprintf("LocalIP is %s", conn.LocalAddr().String()))
	return
}

func main() {
	log.Print("HELLO WORLD!")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	ip, err := GetOutBoundIP()
	if err != nil {
		log.Print(fmt.Sprintf("Get OutBound IP error:%s",err))
	}
	log.Print(ip)


	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("net listen error:%s",err))
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go GetRemoteIP(conn)
	}

}


