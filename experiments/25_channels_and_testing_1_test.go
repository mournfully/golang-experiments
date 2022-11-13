/*
	go - How can I explicitly empty a channel? - Stack Overflow
	https://stackoverflow.com/questions/26143091/how-can-i-explicitly-empty-a-channel

	Building an Unbounded Channel in Go | by Jon Bodner | Capital One Tech | Medium
	https://medium.com/capital-one-tech/building-an-unbounded-channel-in-go-789e175cd2cd

	Go by Example: Testing
	https://gobyexample.com/testing

	go - Golang testing: "no test files" - Stack Overflow
	https://stackoverflow.com/questions/28240489/golang-testing-no-test-files

	! oh dear, I think I was supposed to write unit tests while doing my code
	! not after i'm done most of it - i'll do it for my next exercise then
	! and especially for that personal project i've been cooking up

	TODO: there's 3* types?
	unbuffered...
	buffered...
	unbounded?

	TODO: what i actually needed for my use case
	think of the channel as an array
	* i wish to have 1 index and overwrite it's value n times
	! though i did hear concenrs about how this is rather 'racy'
	so it's probably not proper, but for my first exercise, it will have to do
	and since I don't have MULTIPLE concurrent goroutines, its much simpler for me

	It is not possible to empty a channel without a loop.
	If you don't have any concurrent receivers, then you can use this simple loop:
	This seems to work though :D
	for len(ch) > 0 {
		<-ch
	}

	* there were also a few comments
	1. there's a risk if someone else is reading from the channel as well, and this blocks it
	2. if channel gets closed, this would cause an inifinite loop

	TODO: i dont need this at all, perhaps another time
	im not too sure what he's trying to achieve
	and without rereading it, my understanding of it so far is this
	think of the channel as an array
	* he's trying to write infinite arrays to infinite values
	this doesn't seem like an upgrade over the former, but an entirely different use-case

	TODO: `go test 25_channels_1_test.go`
	oh yeah, this might not actually need a main()
	"A test is created by writing a function with a name beginning with Test."
	"Files containing tests should be called name_test, with the _test suffix. They should be alongside the code that they are testing."
*/

package main

import (
	"fmt"
	"sync"
	"testing"
)

func MakeInfinite() (chan<- interface{}, <-chan interface{}) {
	in := make(chan interface{})
	out := make(chan interface{})
	go func() {
		var in_queue []interface{}
		out_chan := func() chan interface{} {
			if len(in_queue) == 0 {
				return nil
			}
			return out
		}

		cur_val := func() interface{} {
			if len(in_queue) == 0 {
				return nil
			}
			return in_queue[0]
		}

	loop:
		for {
			select {
			case v, ok := <-in:
				if !ok {
					break loop
				} else {
					in_queue = append(in_queue, v)
				}
			case out_chan() <- cur_val():
				in_queue = in_queue[1:]
			}
		}
		close(out)
	}()
	return in, out
}

func TestMakeInfiniteNotPause(t *testing.T) {
	in, out := MakeInfinite()
	last_value := -1
	var wg sync.WaitGroup
	wg.Add(1)

	// anonymous function for testing recieves
	go func() {
		for v := range out {
			vi := v.(int)
			fmt.Println(vi)
			if last_value+1 != vi {
				t.Errorf("unexpected value; expected %d, got %d,", last_value+1, vi)
			}
			last_value = vi
		}
		wg.Done()
		fmt.Println("finished reading")
	}()

	// loop for testing sends
	for i := 0; i < 100; i++ {
		fmt.Println("writing", i)
		in <- i
	}
	close(in)
	fmt.Println("finished writing")
	wg.Wait()

	if last_value != 99 {
		t.Errorf("didn't get all values, last one recieved was %d", last_value)
	}
}
