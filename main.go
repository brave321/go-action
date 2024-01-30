
package main

import (
	"fmt"
	"log"
	"os"
)

const  (
	Ldate =1<<iota
	Ltime
	Lmicroseconds
	Llongfile
	Lshortfile
	LUTC
	LstdFlags= Ldate | Ltime
)

//func SetOutput(w io.Writer)

func main()  {
	fmt.Println("logger")
	log.Println("这是一条很普通的日志")
	v :="很普通的"
	log.Printf("这是一条 %s日志\n",v)
	//log.Fatal("这是一条触发fatal 的日志")

	fmt.Println("=======================")
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	//log.Panicln("这是一条会触发panic的日志")
	//fmt.Println("这是不会被执行的代码")
	//log.Println("这是一条很普通的日志。")
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[pprof]")
	log.Println("这是一条很普通的日志。")

    logFile,err :=os.OpenFile("./xx.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
    if err !=nil{
    	fmt.Println("\"open log file failed, err:\", err")
		return
	}
	log.SetOutput(logFile)

    log.Println("这是一条很普通的日志")
    log.Println("这是一条很普通的日志")
    log.Println("这是一条很普通的日志")
	log.Println("这是一条很普通的日志")

}

