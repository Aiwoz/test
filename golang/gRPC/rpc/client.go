package main

import (
	rpc "Golang/important/gRPC/rpc/proto"
	"crypto/tls"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

const (
	networkType = "tcp"
	server      = "127.0.0.1"
	//port        = "8090"
	parallel = 50     //连接并行度
	times    = 100000 //每连接请求次数
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	currTime := time.Now()

	//并行请求
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			exe()
		}()
	}
	wg.Wait()

	log.Printf("time taken: %.2f ", time.Now().Sub(currTime).Seconds())
}

func exe() {
	//建立连接
	//conn, _ := grpc.Dial(server + ":" + port)
	tlsCreds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})

	conn, err := grpc.Dial("127.0.0.1:8090", grpc.WithTransportCredentials(tlsCreds))
	defer conn.Close()
	if err != nil {
		log.Fatalf("connect Error : %v\n", err)
	}

	client := rpc.NewDataClient(conn)

	for i := 0; i < int(times); i++ {
		getUser(client)
	}
}

func getUser(client rpc.DataClient) {
	var request rpc.UserRq
	r := rand.Intn(parallel)
	request.Id = int32(r)

	response, err := client.GetUser(context.Background(), &request) //调用远程方法
	if err != nil {
		log.Fatalf("GetUser's Error : %v\n", err)
	}

	//判断返回结果是否正确
	if id, _ := strconv.Atoi(strings.Split(response.Name, ":")[0]); id != r {
		log.Printf("response error  %v\n", response)
	}

}
