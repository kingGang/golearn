package main

import (
	pb "golearn/istio-grpc/api"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"gitee.com/KaiErDeLe/higo/pkg/naming"
	"github.com/gin-gonic/gin"
)

var (
	r   = gin.Default()
	cli pb.HelloWorldClient
)

func main() {
	var err error
	cli, err = pb.NewClient(nil)
	if err != nil {
		panic(err)
	}
	r = gin.Default()
	setupRouter(r)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    time.Duration(10) * time.Second,
		WriteTimeout:   time.Duration(10) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func test(target string) {
	if u, e := url.Parse(target); e == nil {
		v := u.Query()
		for _, c := range c.conf.Clusters {
			v.Add(naming.MetaCluster, c)
		}
		if c.conf.Zone != "" {
			v.Add(naming.MetaZone, c.conf.Zone)
		}
		if v.Get("subset") == "" && c.conf.Subset > 0 {
			v.Add("subset", strconv.FormatInt(int64(c.conf.Subset), 10))
		}
		u.RawQuery = v.Encode()
		// 比较_grpcTarget中的appid是否等于u.path中的appid，并替换成mock的地址
		for _, t := range _grpcTarget {
			strs := strings.SplitN(t, "=", 2)
			if len(strs) == 2 && ("/"+strs[0]) == u.Path {
				u.Path = "/" + strs[1]
				u.Scheme = "passthrough"
				u.RawQuery = ""
				break
			}
		}
		target = u.String()
	}
}

func setupRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("grpc", func(c *gin.Context) {
		resp, err := cli.SayHelloURL(c, &pb.HelloReq{Name: "grpc test SayHelloURL"})
		if err != nil {
			c.JSON(200, err.Error())
			return
		}
		c.JSON(200, resp)
	})
	r.GET("chaobin", func(c *gin.Context) {
		resp, err := cli.SayHello(c, &pb.HelloRequest{Name: "grpc test SayHelloURL"})
		if err != nil {
			c.JSON(200, err.Error())
			return
		}
		c.JSON(200, resp)
	})
}
