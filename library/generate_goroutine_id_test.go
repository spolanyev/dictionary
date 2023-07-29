//@author Stanislav Polaniev <spolanyev@gmail.com>

package library

import (
	"sync"
	"testing"
)

func TestGoroutineIdGeneration(t *testing.T) {
	cases := []struct {
		testName          string
		placeIdentifier   []string
		goroutineQuantity int
		want              string
	}{
		{
			"1",
			[]string{},
			0,
			"[G-1] ",
		},
		{
			"2",
			[]string{"MethodName"},
			3,
			"[MethodName][G-5] ",
		},
		{
			"3",
			[]string{"ModuleName", "ClassName", "MethodName"},
			0,
			"[ModuleName][ClassName][MethodName][G-6] ",
		},
	}

	for _, theCase := range cases {
		if theCase.goroutineQuantity > 0 {
			done := make(chan struct{})
			group := sync.WaitGroup{}
			for i := 0; i < theCase.goroutineQuantity; i++ {
				group.Add(1)
				go func() {
					defer group.Done()
					GenerateGoroutineId( /*without place identifier*/ )
				}()
			}
			go func() {
				group.Wait()
				close(done)
			}()
			<-done
		}

		got := GenerateGoroutineId(theCase.placeIdentifier...)

		if got != theCase.want {
			t.Errorf("Test %v: output == %v, want %v", theCase.testName, got, theCase.want)
		}
	}
}
