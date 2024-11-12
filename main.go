package main

import (
	"context"
	"k8s.io/apiserver/pkg/storage/value/encrypt/envelope"
	"k8s.io/klog/v2"
	"time"
)

func main() {
	klog.ClearLogger()
	_, err := envelope.NewGRPCService(context.Background(), "", time.Minute)
	if err != nil {
		return
	}
}
