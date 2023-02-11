package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/labstack/echo/v4"

	"github.com/trinitt/config"
	"github.com/trinitt/utils"
	"github.com/hamba/avro/v2"
)

type SignupRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsSeller  bool   `json:"is_seller"`
}

var Schema = `{
	"type": "record",
	"name": "Record",
	"fields": [
	  {
		"name": "user_id",
		"type": "int"
	  },
	  {
		"name": "entity_id",
		"type": "int"
	  },
	  {
		"name": "param",
		"type": {
		  "type": "array",
		  "items": {
			"type": "record",
			"namespace": "Record",
			"name": "param",
			"fields": [
			  {
				"name": "data_type",
				"type": "string"
			  },
			  {
				"name": "value",
				"type": "string"
			  }
			]
		  }
		}
	  }
	]
  }`

type record struct{
	Data_type string `avro:"data_type" json:"data_type"`
	Value string `avro:"value" json:"value"`
}

type Record struct {
	User_id int  `avro:"user_id" json:"user_id"`
	Entity_id  int`avro:"entity_id" json:"entity_id"`
	Param []record `avro:"param" json:"param"`
}


func conv1(rec record) map[string]interface{}{
	m := make(map[string]interface{})
	m["data_type"] = rec.Data_type
	m["value"] = rec.Value
	return m
}
type M map[string]interface{}
func conv(rec Record) map[string]interface{}{
	m := make(map[string]interface{})
	m["user_id"] = rec.User_id
	m["entity_id"] = rec.Entity_id
	var h [] map[string]interface{}
	for _, value := range rec.Param {
        h =append(h, conv1(value))
    }
	m["param"] = h
	return m
}

func SignupUser(c echo.Context) error {

	schema, err := avro.Parse(Schema)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}





	in := Record{	
	
		User_id: 1,
		Entity_id: 1,
		Param: []record{
			{
				Data_type: "INT",
				Value: "5",
			},
			{
				Data_type: "string",
				Value: "hello",
			},
		},

	}

	data, err := avro.Marshal(schema, in)
	if err != nil {
		log.Fatal(err)
	}



	fmt.Printf("%+v\n", data)
	topic:="myTopic"

	config.GetProducer().Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: 0},
		Value:         data,
	}, nil)
    
	
	out := Record{}
	err = avro.Unmarshal(schema, data, &out)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out.Entity_id)
    
	return utils.SendResponse(c, http.StatusOK, "User created successfully")
}
