package main

import (
	"fmt"
	"net/http"
	"strconv"
	//"reflect"
	"runtime"
	"strings"
)

var fin = make(chan string)
var limits = make(chan int, 1000)

func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func doit(url string) {
	limits <- 1
	defer func() {
		<-limits
	}()
	httpget(url)
}

func httpget(url string) {
	resp, err := http.Get(url)
	defer func() {
		//fmt.Printf("url is %s, goroutine id is %d\n", url, GoID())
		//v := reflect.ValueOf(resp)
		//fmt.Printf("pv, %p,r, %v\n",v,resp)
		//count := v.NumField()
		//for i := 0; i < count; i++ {
		//    f := v.Field(i)
		//    switch f.Kind() {
		//        case reflect.String:
		//            fmt.Println(f.String())
		//        case reflect.Int:
		//            fmt.Println(f.Int())
		//    }
		//    //fmt.Printf("Field %d: %v\n", i, value.Field(i))
		//}
		if resp != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		fmt.Println(url, err)
	} else {
		fmt.Println(url, resp.StatusCode)
	}
	fin <- url
}

func main() {

    var MULTICORE int = runtime.NumCPU() //number of core
	runtime.GOMAXPROCS(1) //running in multicore
	fmt.Printf("with %d core\n", MULTICORE)
	
	var httplist [50000]string
	for i := 0; i < len(httplist); i++ {
		//httplist[i] = "http://www.baidu.com/s?wd=search" + strconv.Itoa(i+1)
		//httplist[i] = "http://192.168.6.151:2003/"
		httplist[i] = "http://192.168.6.151:80/"
		go doit(httplist[i])
	}

	//等待结束
	for i := 0; i < len(httplist); i++ {
		<-fin
	}
}
