package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("https://192.168.81.82:6443", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	e, err := clientset.CoreV1().Pods("kingkang").Watch(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for {
		select {
		case ev := <-e.ResultChan():
			fmt.Println("消息", ev.Type, ev.Object)
			if string(ev.Type) == "" {
				time.Sleep(time.Second)
				fmt.Println("为空：", ev)
				e, err = clientset.CoreV1().Pods("kingkang").Watch(metav1.ListOptions{})
				fmt.Println(err)
				break
			}
		case <-time.After(time.Second * 3):
			fmt.Println("3秒了")
		default:
			time.Sleep(time.Second * 1)
		}
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
