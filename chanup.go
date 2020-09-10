/*
   Copyright 2020 Subba Reddy Veeramreddy

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package chanup

// ChanUp is the Wrapper for a go channel
type ChanUp struct {
	// channel takes value of any type
	channel chan interface{}
}

// Status constants are used to express the outcome of an operation on channel with the Methods of ChanUp
type Status int

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
