package main

import (
	"fmt"
	"time"
	"sync/atomic"
	"runtime"
	"math"
)

func main() {
	fmt.Println("Hello,World")
	fmt.Print("The time is ", time.Now())
	var x, y int = 3, 4
	fmt.Println(x, y)

	var ops uint64 = 0
	for i := 0;i<50;i++ {
		go func() {
			atomic.AddUint64(&ops,1)
			runtime.Gosched()
		}()
	}
	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:",opsFinal)
	i := 0
	for i<100 {
		i ++
	}
	fmt.Println(sqrt(2),sqrt(-4))
	fmt.Println(
		power(3,2,10),
		power(3,3,20),
	)
	fmt.Print("Go runs on ")
	switch os:=runtime.GOOS;os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.",os)
	}
	fmt.Println("counting")

	for i:=0;i<10;i++ {
		//defer fmt.Println(i)
	}
	fmt.Println("done")

	i,j := 42,2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p /37
	fmt.Println(j)
	ap := Point{1,2}
	fmt.Printf("%+v\n",ap)

	var ap_i *Point
	ap_i = &ap
	ap_i.x = 1e9
	fmt.Printf("%+v\n",ap_i)
}
type Point struct {
	x,y int
}
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x)+"i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func power(x,n,lim float64)  float64 {
	if v:=math.Pow(x,n);v<lim {
		return v
	} else {
		fmt.Printf("%g>=%g\n",v,lim)
	}
	return lim
}