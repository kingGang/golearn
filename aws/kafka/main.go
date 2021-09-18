package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kafka"
)

func main() {
	mySession := session.Must(session.NewSession())
	svc := kafka.New(mySession, aws.NewConfig().WithRegion("ap-northeast-1").WithEndpoint("b-3.kafka-demo.x3ik83.c4.kafka.ap-northeast-1.amazonaws.com:9092"))
	fmt.Println(svc.ClientInfo, svc.APIVersion)
	// fmt.Println(svc.ListClusters(&kafka.ListClustersInput{}))
	t := time.Now()
	fmt.Println(t.Hour(), t.Minute(), t.Second())
	time.Now().Weekday()
	fmt.Println(time.Since(t))
}
