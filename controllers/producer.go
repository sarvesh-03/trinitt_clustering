package controllers

import (
	"fmt"
	"log"

	"bitbucket.org/sjbog/go-dbscan"
	"github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/hamba/avro/v2"
	"github.com/trinitt/config"
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

var Schema1 = `{
	"type": "record",
	"name": "Record",
	"fields": [
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
			"name": "x",
			"type": "double"
			},
			{
			"name": "y",
			"type": "double"
			},
			{
			"name": "cluster",
			"type": "int"
			}
			]
		  }
		}
	  }
	]
  }`

type Node struct {
	X   float64      `avro:"x" json:"x"`
	Y 	float64      `avro:"y" json:"y"`
	Cluster     int  `avro:"cluster" json:"cluster"`
}

type Nodes struct{
	Param     []Node `avro:"param" json:"param"`
}

type ParamType struct {
	Data_type string `avro:"data_type" json:"data_type"`
	Value     string `avro:"value" json:"value"`
}



type Record struct {
	User_id   int      `avro:"user_id" json:"user_id"`
	Entity_id int      `avro:"entity_id" json:"entity_id"`
	Param     []ParamType `avro:"param" json:"param"`
}


func Produce(cluster [][]dbscan.ClusterablePoint){

	schema, err := avro.Parse(Schema1)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
	
	fmt.Println(schema)
	var node  []Node
	for ind, cluster := range cluster {
		fmt.Println("Cluster:")
		for _, point := range cluster {
			
			
			node= append(node, Node{
				X: point.(*dbscan.NamedPoint).Point[0],
				Y: point.(*dbscan.NamedPoint).Point[1],
				Cluster: ind,

			})
			fmt.Println(node)
		}
	}
	nodes:= Nodes{Param:node}

	data, err := avro.Marshal(schema, nodes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", data)
	topic := "yTopic"

	config.GetProducer().Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: 0},
		Value:          data,
	}, nil)

	out := Record{}
	err = avro.Unmarshal(schema, data, &out)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out.Entity_id)

}

