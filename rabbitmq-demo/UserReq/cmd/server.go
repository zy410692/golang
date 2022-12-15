package main

import (
	"log"
	"main/rabbitmq-demo/Lib"
	"main/rabbitmq-demo/UserReq/Models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Handle("POST", "/user", func(context *gin.Context) {
		userModel := Models.NewUserModel()
		err := context.BindJSON(&userModel)
		if err != nil {
			context.JSON(400, gin.H{"result": "params error"})
		} else {
			userModel.UserId = int(time.Now().Unix())
			if userModel.UserId > 0 {
				mq := Lib.NewMq()
				err = mq.SendMessage(Lib.ROUTER_KEY_USERREG, Lib.EXCHANGE_USER, strconv.Itoa(userModel.UserId))
				defer mq.Channel.Close()
				if err != nil {
					log.Println(err)
				}

			}
			context.JSON(200, gin.H{"result": userModel})
		}
	})

	c := make(chan error)

	go func() {
		err := router.Run(":8081")
		if err != nil {
			c <- err
		}
	}()

	go func() {
		err := Lib.UserInit()
		if err != nil {
			c <- err
		}

	}()

	err := <-c
	log.Fatal(err)

}
