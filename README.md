# goBroadcastUDP
Sample Setup to Show multi-NIC broadcast server and client. 

There is some good documents out there that give a bare minimum example on how to send Broadcast UDP with golang, but i found it quite limiting trying to use it on a multi-network interface server. Being able to sort out what interface the messaging comes in on, or even more important, gets sent from on a multi-network interface system, was not very clear.

I posted on google groups [golang-nuts post](https://groups.google.com/forum/#!msg/golang-nuts/nbmYWwHCgPc/ZBw2uH6Bdi4J) golang-nuts over a year ago, and circling back around with a resolution.


## Server
I have written a server.go and a client.go that demonstrate usage of a utility function that shows how this can work. The utility server.go will bind to 0.0.0.0:<port> and query all NIC at time of startup. When the server receives any broadcast packets from the network, it will query the previously cached IPNet collection, and identify what net.Interface it was originally transmitted from. It returns stdout a 

+ hex.Dump() of the message
+ the remote address info of the client
message => 
`the client sent a broadcast message on interface: en0`  
+ identifying what interface the message was received on.  

NOTE: The server will only bind to NON-localhost, IPv4 , broadcast enabled net.Interfaces

## Client
The client.go has a quick way to send a broadcast on any particular network, and bind the local address to the DialUDP function. This is only necessary if the client has more than 1 NIC. 


checkout the sample test here and run it: 

```
$ git clone https://github.com/dfense/goBroadcastUDP
$ cd goBroadcastUDP; export GOPATH=$PWD;
$ go run src/github.com/hupla/main/server.go

(change the src in client.go variables ipLocal and ipBcast to meet your local server/client test needs)
$ go run src/github.com/hupla/main/client.go
```
