package grpc_health

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	pb_empty "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type GrpcHealthChecker interface {
	IsHealthy(serverUrl, serviceName string, timeout int, verbose bool) bool
}

func NewGrpcHealthChecker() GrpcHealthChecker {
	return grpcHealthChecker{}
}

type grpcHealthChecker struct {
}

func (g grpcHealthChecker) IsHealthy(serverUrl, serviceName string, timeout int, verbose bool) bool {
	if verbose {
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(ioutil.Discard)
	}

	log.Printf("Starting health check for %s running at %s (timeout=%ds)\n", serviceName, serverUrl, timeout)
	methodName := fmt.Sprintf("/%s/Health", serviceName)

	log.Println("Attempting connection...")
	conn, err := grpc.Dial(
		serverUrl,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(time.Duration(timeout)*time.Second),
	)
	if err != nil {
		if err == context.DeadlineExceeded {
			log.Println("Error: Timeout")
			return false
		} else {
			log.Printf("Error: Did not connect\n%v\n", err)
			return false
		}
	}
	defer conn.Close()
	log.Println("Connection successful")

	log.Printf("Invoking %s\n", methodName)
	in := &pb_empty.Empty{}
	out := new(pb_empty.Empty)
	err = grpc.Invoke(context.Background(), methodName, in, out, conn)
	if err != nil {
		grpcStatus, ok := status.FromError(err)
		if ok {
			log.Printf("Error: grpc_code=%s; msg=%s\n", grpcStatus.Code(), grpcStatus.Message())
			return false
		}
		log.Printf("Error: invoke failed\n%v\n", err)
		return false
	}
	log.Println("Invoke successful")
	log.Println("Healthy")
	return true
}
