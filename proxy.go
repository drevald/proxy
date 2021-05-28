package main

import (
	"fmt"
	"net"
	"os"
	"encoding/binary"
    "strconv"
)

const (
    connHost = "localhost"
    connPort = "8001"
    connType = "tcp"
)

type address struct {
    length []byte
    host []byte
    port []byte
}


// o  X'00' "NO AUTHENTICATION REQUIRED" method

// 4.  Requests

//    Once the method-dependent subnegotiation has completed, the client
//    sends the request details.  If the negotiated method includes
//    encapsulation for purposes of integrity checking and/or
//    confidentiality, these requests MUST be encapsulated in the method-
//    dependent encapsulation.

//    The SOCKS request is formed as follows:

//         +----+-----+-------+------+----------+----------+
//         |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
//         +----+-----+-------+------+----------+----------+
//         | 1  |  1  | X'00' |  1   | Variable |    2     |
//         +----+-----+-------+------+----------+----------+

//      Where:

//           o  VER    protocol version: X'05'
//           o  CMD
//              o  CONNECT X'01'
//              o  BIND X'02'
//              o  UDP ASSOCIATE X'03'
//           o  RSV    RESERVED
//           o  ATYP   address type of following address
//              o  IP V4 address: X'01'
//              o  DOMAINNAME: X'03'
//              o  IP V6 address: X'04'
//           o  DST.ADDR       desired destination address
//           o  DST.PORT desired destination port in network octet
//              order

func method_0(conn net.Conn) {
    req := []byte{5, 0}                 //  The server selects from one of the methods given in METHODS, and    
    conn.Write(req)                     //  sends a METHOD selection message
    res := make([]byte, 4)
    conn.Read(res)
    fmt.Println("Response", res)
    
    if (res[3]==3) {
        addr := address {
            length:make([]byte, 1),
            port:make([]byte, 2),
        }
        conn.Read(addr.length)
        addr.host = make([]byte, int(addr.length[0]))
        conn.Read(addr.host)            
        conn.Read(addr.port)
        fmt.Println("Host", addr.host)
        fmt.Println("Port", addr.port)
        port := binary.BigEndian.Uint16(addr.port)
        fmt.Println("Address", string(addr.host) + ":" + strconv.Itoa(int(port)));
        
        //_, err := net.Listen(connType, string(addr.host) + ":" + strconv.Itoa(int(port)))
        _, err := net.Listen(connType, "www.vk.com:80")
        if err != nil {
            fmt.Println("Error", err.Error())
        }

    }   

}


// o  X'02' "USERNAME/PASSWORD" method
func method_2(conn net.Conn) {
    fmt.Printf("Method 2 not implemented")
}

func main() {
    fmt.Println("Starting " + connType + " server on " + connHost + ":" + connPort)
    l, err := net.Listen(connType, connHost+":"+connPort)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    defer l.Close()


    // The client connects to the server, and sends a version
    // identifier/method selection message:
    // +----+----------+----------+
    // |VER | NMETHODS | METHODS  |
    // +----+----------+----------+
    // | 1  |    1     | 1 to 255 |
    // +----+----------+----------+    
    for {
        c, e := l.Accept()
        if e != nil {
            fmt.Println("Error accepting", e.Error())
            os.Exit(1)
        }
        // Reading 1 + 1 + 255 bytes
        buf := make([]byte, 257)
        _, e1 := c.Read(buf)
        if e1 != nil {
            fmt.Println("Error reading request", e.Error())
        }
        //Ignoring version and reading number of methods
        fmt.Println("Num methods", int(buf[1]))        
        for i := 0; i < int(buf[1]); i++ {

            // The values currently defined for METHOD are:
            // o  X'00' NO AUTHENTICATION REQUIRED
            // o  X'01' GSSAPI
            // o  X'02' USERNAME/PASSWORD
            // o  X'03' to X'7F' IANA ASSIGNED
            // o  X'80' to X'FE' RESERVED FOR PRIVATE METHODS
            // o  X'FF' NO ACCEPTABLE METHODS
            fmt.Printf("Method: %d\n", buf[2+i])
            if buf[2+i] == 0 {
                method_0(c)
            }
            if buf[2+i] == 2 {
                method_2(c)
            }

        }
        fmt.Println("End")
    }

}