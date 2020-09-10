# ChanUp Channels
A Wrapper on top of go channels which supports one complex use case. 
A ChanUp channel buffer has a length of One.
ChanUp channel can be used to make producer and consumer of channel independent of each other.
The execution flow of no process gets blocked. 

## Functionality
- ChanUp channel never blocks producer or consumer.
- A producer Puts a value if channel is empty
- A producer Updates channel with new value if it is holding a stale value.
- A consumer gets a value if available.
