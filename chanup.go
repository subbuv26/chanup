package chanup

// ChanUp is the Wrapper for a go channel
type ChanUp struct {
	// channel takes value of any type
	channel chan interface{}
}

type Status int

// Status constants are used to express the outcome of an operation on channel with the Methods of ChanUp
const (
	FAILED Status = iota
	PUT
	UPDATE
)

// GetChan creates and returns a new ChanUp channel
func GetChan() *ChanUp {
	return &ChanUp{
		channel: make(chan interface{}, 1),
	}
}

// Put puts the value in the channel
func (cu *ChanUp) Put(value interface{}) Status {
	select {
	case cu.channel <- value:
		return PUT
	default:
		return FAILED
	}
}

// Update updates the channel with new value by replacing old value
// if the channel is empty, then Update acts like Put and puts the value in the channel
func (cu *ChanUp) Update(value interface{}) Status {
	select {
	case <-cu.channel:
		cu.channel <- value
		return UPDATE
	case cu.channel <- value:
		return PUT
	}
}

// Get removes value from channel and returns the same
// if the channel is empty Get returns a nil value
func (cu *ChanUp) Get() interface{} {
	select {
	case value := <-cu.channel:
		return value
	default:
		return nil
	}
}
