// Cisco DNA Center ISE Health Check Console Script.
//
// Copyright (c) 2019 Cisco and/or its affiliates.
//
// This software is licensed to you under the terms of the Cisco Sample
// Code License, Version 1.1 (the "License"). You may obtain a copy of the
// License at
//
//               https://developer.cisco.com/docs/licenses
//
// All use of the material herein must be in accordance with the terms of
// the License. All rights not expressly granted by the License are
// reserved. Unless required by applicable law or agreed to separately in
// writing, software distributed under the License is distributed on an "AS
// IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied.
//
// __author__ = "Robert Csapo"
// __email__ = "rcsapo@cisco.com"
// __version__ = "0.1"
// __copyright__ = "Copyright (c) 2019 Cisco and/or its affiliates."
// __license__ = "Cisco Sample Code License, Version 1.1"

package main

import (
  "log"
  "net"
  "flag"
  "fmt"
)

// Cisco ISE Settings Struct
type iseStruct struct {
  Host              string
  Proto             string
  sshPort           string
  webPort           string
  ersPort           string
  pxGridInterPort   string
  pxGridSubPort     string
}

// Create a var based of ISE Struct
var ise iseStruct

// Connection Test function
func connectTest(host,port,proto string) string {
  host = host+":"+port
  var status string
  conn, err := net.Dial(proto, host)
  if err != nil {
          log.Println("ERROR:\t\t", err)
          status = "Unreachable"
  } else {
          status = "Online"
          defer conn.Close()
  }
  return(status)
}

// Main function
func main() {
  flagHost := flag.String("host", "", "cisco ise hostname/ip-address")
  flag.Parse()


  if (*flagHost != "") {
    // set ISE host if flag is used
    ise.Host = string(*flagHost)
  } else {
    // get ISE host from input, as flag is missing
    fmt.Print("Enter host (FQDN): ")
    var host string
    fmt.Scanln(&host)
    ise.Host = host
  }

  ise.Proto = "tcp"
  ise.sshPort = "22"
  ise.webPort = "443"
  ise.ersPort = "9060"
  ise.pxGridInterPort = "5222"
  ise.pxGridSubPort = "8910"

  if (connectTest(ise.Host, ise.sshPort, ise.Proto) == "Online") {
    log.Println("SUCCESS:\t\tCisco ISE ("+ise.Host+") - SSH port ("+ise.sshPort+") is accessible")
  } else {
    log.Println("ERROR:\t\tCisco ISE ("+ise.Host+") - SSH port ("+ise.sshPort+") is NOT accessible")
  }
  if (connectTest(ise.Host, ise.webPort, ise.Proto) == "Online") {
    log.Println("SUCCESS:\t\tCisco ISE ("+ise.Host+") - Web port ("+ise.webPort+") is accessible")
  } else {
    log.Println("ERROR:\t\tCisco ISE ("+ise.Host+") - Web port ("+ise.webPort+") is NOT accessible")
  }
  if (connectTest(ise.Host, ise.ersPort, ise.Proto) == "Online") {
    log.Println("SUCCESS:\t\tCisco ISE ("+ise.Host+") - ERS API port ("+ise.ersPort+") is accessible")
  } else {
    log.Println("ERROR:\t\tCisco ISE ("+ise.Host+") - ERS API port ("+ise.ersPort+") is NOT accessible")
  }
  if (connectTest(ise.Host, ise.pxGridInterPort, ise.Proto) == "Online") {
    log.Println("SUCCESS:\t\tCisco ISE ("+ise.Host+") - pxGrid Inter-Node Communication port ("+ise.pxGridInterPort+") is accessible")
  } else {
    log.Println("ERROR:\t\tCisco ISE ("+ise.Host+") - pxGrid Inter-Node Communication port ("+ise.pxGridInterPort+") is NOT accessible")
  }
  if (connectTest(ise.Host, ise.pxGridSubPort, ise.Proto) == "Online") {
    log.Println("SUCCESS:\t\tCisco ISE ("+ise.Host+") - pxGrid Subscribers port ("+ise.pxGridSubPort+") is accessible")
  } else {
    log.Println("ERROR:\t\tCisco ISE ("+ise.Host+") - pxGrid Subscribers port ("+ise.pxGridSubPort+") is NOT accessible")
  }
}
