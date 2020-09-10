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

import (
	"testing"
)

type testType struct {
	a int
	s string
}

func TestGetChan(t *testing.T) {
	ch := GetChan()
	if ch == nil {
		t.Error("Failed to create ChanUp Channel")
	}
}

func TestChanUp_Put(t *testing.T) {
	ch := GetChan()
	status := ch.Put(testType{
		a: 10,
		s: "Sample",
	})
	if status == FAILED {
		t.Error("Failed to put a value in ChanUp Channel")
	}
}

func TestChanUp_Put2(t *testing.T) {
	ch := GetChan()
	status := ch.Put(testType{
		a: 10,
		s: "Sample",
	})
	if status == FAILED {
		t.Error("Failed to put a value in ChanUp Channel")
	}

	status = ch.Put(testType{
		a: 20,
		s: "Sample 2",
	})
	if status == PUT {
		t.Error("Put is not supposed to update ChanUp Channel")
	}
}

func TestChanUp_Get(t *testing.T) {
	ch := GetChan()
	status := ch.Put(testType{
		a: 10,
		s: "Sample",
	})
	if status == FAILED {
		t.Error("Failed to put a value in ChanUp Channel")
	}

	val := ch.Get()
	if val == nil {
		t.Error("Failed to Get Value from ChanUp Channel")
	}
	tv := val.(testType)

	if tv.a != 10 || tv.s != "Sample" {
		t.Error("Invalid value from ChanUp Channel")

	}
}

func TestGetChan2(t *testing.T) {
	ch := GetChan()

	val := ch.Get()
	if val != nil {
		t.Error("ChanUp Channel should be empty ang give 'nil' value")
	}
}

func TestChanUp_Update(t *testing.T) {
	ch := GetChan()

	testValue := testType{
		a: 10,
		s: "Sample",
	}
	status := ch.Update(testValue)

	if status != PUT {
		t.Error("Failed. Status must be PUT while Updating empty Channel")
	}

	val := ch.Get()
	if val == nil {
		t.Error("Failed to GET")
	}
	tv := val.(testType)

	if testValue != tv {
		t.Error("Got Wrong Value from ChanUp Channel after Updating")
	}
}

func TestChanUp_Update2(t *testing.T) {
	ch := GetChan()

	testValue := testType{
		a: 10,
		s: "Sample",
	}

	testValue2 := testType{
		a: 20,
		s: "Sample ",
	}
	_ = ch.Put(testValue)
	status := ch.Update(testValue2)

	if status != UPDATE {
		t.Error("Failed to Update ChanUp Channel")
	}

	val := ch.Get()
	if val == nil {
		t.Error("Failed to GET")
	}
	tv := val.(testType)

	if testValue2 != tv {
		t.Error("Got Wrong Value from ChanUp Channel after Updating")
	}
}
