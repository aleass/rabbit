package main

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	//rabbitmq := NewRabbitMQSimple("queue")
	//go rabbitmq.ConsumeSimple(1)
	//go rabbitmq.ConsumeSimple(2)
	//time.Sleep(100*time.Second)

	go func() {
		r:= NewRabbitMQTopic("hxbExc","#")
		r.RecieveTopic(1)

	}()
	go func() {
		r:= NewRabbitMQTopic("hxbExc","huxiaobai.*.cs")
		r.RecieveTopic(2)
	}()

	go func() {
		r:= NewRabbitMQTopic("hxbExc","huxiaobai.o#")
		r.RecieveTopic(3)
	}()


	for {
		time.Sleep(1 * time.Second)
	}
}