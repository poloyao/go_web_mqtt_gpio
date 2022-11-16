package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/stianeikeland/go-rpio/v4"
)

var (
	pin = rpio.Pin(21)
)

// 根据订阅pi的内容操控GPIO引脚动作高低电平
func SubscriptionFunc(topic string, data []byte) {
	println(topic, string(data))

	if topic == "pi" {
		if err := rpio.Open(); err != nil {
			fmt.Println(err)
		} else {
			pin.Output()
			if string(data) == "1" {
				pin.High()
			} else {
				pin.Low()
			}
			rpio.Close()
		}
	}
}

func main() {

	r := gin.Default()

	r.GET("/mqpush", func(ctx *gin.Context) {
		//payload := time.Now().GoString()
		topic := ctx.Query("topic")
		payload := ctx.Query("payload")
		Push(topic, 1, false, payload)
		ctx.JSON(200, gin.H{
			"message": payload,
		})
	})

	r.GET("/mqsub", func(ctx *gin.Context) {
		name := ctx.Query("name")
		Subscription(name, 1, true, SubscriptionFunc)
		ctx.JSON(200, gin.H{
			"topic":      name,
			"subsucceed": true,
		})
	})

	r.Run()
}
