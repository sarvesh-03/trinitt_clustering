package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/labstack/echo/v4"
	// "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/hamba/avro/v2"
	"github.com/trinitt/config"
	"github.com/trinitt/utils"
)




func Consume(c echo.Context) error {
	schema, _ := avro.Parse(Schema)
	
	go func(){
	run:=true
	for run {
		msg, err := config.GetConsumer().ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			out := Record{}
			err = avro.Unmarshal(schema, msg.Value, &out)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(out)
		} else  if err.(kafka.Error).Code().String()!=kafka.ErrMsgTimedOut.String(){
			fmt.Println(msg)
		}
	}
	}()
	return utils.SendResponse(c, http.StatusOK, "User created successfully")
}