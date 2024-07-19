package Mq

import "github.com/streadway/amqp"

// RabbitMqObject 每个接口一个mq对象
type RabbitMqObject struct {
	Ch                              *amqp.Channel
	ServerSendQueue, AgentSendQueue amqp.Queue

	//队列名称
	ServerSendQueueName, AgentSendQueueName string
	//交换机名称
	ExchangeName string
	//bind Key 名称
	ServerSendRoutingKey, AgentSendRoutingKey string
	//交换机类型
	ExchangeType string
}

type Options struct {
	Exchange        *ExchangeInfo
	AgentSendQueue  *AgentSendQueueInfo
	ServerSendQueue *ServerSendQueueInfo
	AgentSendBind   *AgentSendBindInfo
	ServerSendBind  *ServerSendBindInfo
}

type ExchangeInfo struct {
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	Args       amqp.Table
}

type AgentSendQueueInfo struct {
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
}
type ServerSendQueueInfo struct {
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
}

type AgentSendBindInfo struct {
	NoWait bool
	Args   amqp.Table
}
type ServerSendBindInfo struct {
	NoWait bool
	Args   amqp.Table
}
