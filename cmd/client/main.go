/// Client application to make requests via gRPC to a dockerised server
/// Author: Chris Clarkson
/// Date: 1st June 2017
/// Company: CJC Software Solutions ltd.
package main

import (
	"flag"
)

var address = flag.String("Address", "localhost", "Server address of the gRPC remote application host")
var port = flag.Int("Port", 30555, "Port number exposed by remote application host")

func main() {
	flag.Parse()

}
