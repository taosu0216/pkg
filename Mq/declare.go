package Mq

import (
	"fmt"
	"github.com/streadway/amqp"
)

func NewMqObj(ch *amqp.Channel, exchangeType, exchangeName string,
	serverSendQueueName, agentSendQueueName, serverSendRoutingKey, agentSendRoutingKey string,
	options ...Option,
) (*RabbitMqObject, error) {
	var err error
	obj := &RabbitMqObject{
		Ch:                   ch,
		ServerSendQueueName:  serverSendQueueName,
		AgentSendQueueName:   agentSendQueueName,
		ExchangeName:         exchangeName,
		ServerSendRoutingKey: serverSendRoutingKey,
		AgentSendRoutingKey:  agentSendRoutingKey,
		ExchangeType:         exchangeType,
	}
	opts := &Options{
		Exchange:        NewDefaultExchangeInfo(),
		AgentSendQueue:  NewDefaultAgentSendQueueInfo(),
		ServerSendQueue: NewDefaultServerSendQueueInfo(),
		AgentSendBind:   NewDefaultAgentSendBindInfo(),
		ServerSendBind:  NewDefaultServerSendBindInfo(),
	}
	// 应用可选参数
	for _, opt := range options {
		if opt == nil {
			continue
		}
		opt(opts)
	}

	// 所有类型都需要声明两个队列
	// Declare AgentSend Queue
	obj.AgentSendQueue, err = ch.QueueDeclare(
		agentSendQueueName,
		opts.AgentSendQueue.Durable,
		opts.AgentSendQueue.AutoDelete,
		opts.AgentSendQueue.Exclusive,
		opts.AgentSendQueue.NoWait,
		opts.AgentSendQueue.Args,
	)
	if err != nil {
		return obj, fmt.Errorf("declare AgentSendQueue Failed: %w", err)
	}
	// Declare ServerSend Queue
	obj.ServerSendQueue, err = ch.QueueDeclare(
		serverSendQueueName,
		opts.ServerSendQueue.Durable,
		opts.ServerSendQueue.AutoDelete,
		opts.ServerSendQueue.Exclusive,
		opts.ServerSendQueue.NoWait,
		opts.ServerSendQueue.Args,
	)
	if err != nil {
		return obj, fmt.Errorf("declare ServerSendQueue Failed: %w", err)
	}

	// 如果是simple或者work类型就不需要声明exchange和绑定queue
	if exchangeType == "simple" || exchangeType == "work" {
		return obj, nil
	} else if exchangeType == "fanout" {
		// fanout类型需要声明交换机,但不需要queueName和routingKey
		agentSendQueueName = ""
		serverSendQueueName = ""
		agentSendRoutingKey = ""
		serverSendRoutingKey = ""
	}

	// 声明交换机
	err = ch.ExchangeDeclare(
		exchangeName,
		exchangeType,
		opts.Exchange.Durable,
		opts.Exchange.AutoDelete,
		opts.Exchange.Internal,
		opts.Exchange.NoWait,
		opts.Exchange.Args,
	)
	if err != nil {
		return obj, fmt.Errorf("declare Mq Exchange Failed: %w", err)
	}

	// 将agent端发送消息的queue绑定到交换机
	err = ch.QueueBind(
		agentSendQueueName,
		agentSendRoutingKey,
		exchangeName,
		opts.AgentSendBind.NoWait,
		opts.AgentSendBind.Args,
	)
	if err != nil {
		return obj, fmt.Errorf("bind AgentSendQueue Failed: %w", err)
	}

	// 将server端发送消息的queue绑定到交换机
	err = ch.QueueBind(
		serverSendQueueName,
		serverSendRoutingKey,
		exchangeName,
		opts.ServerSendBind.NoWait,
		opts.ServerSendBind.Args,
	)
	if err != nil {
		return obj, fmt.Errorf("bind ServerSendQueue Failed: %w", err)
	}

	return obj, nil
}

// ConsumeQueue 创建消费者通道
func (obj *RabbitMqObject) ConsumeQueue(queueName string) (<-chan amqp.Delivery, error) {
	return obj.Ch.Consume(
		queueName, // 队列名称
		// TODO: 消费者是否需要唯一标识符?
		"",    // 消费者标识
		false, // 自动应答
		false, // 非排他
		false, // 非本地
		false, // 不等待服务器确认
		nil,   // 其他参数
	)
}

// 废弃方法
//func GetNewMqObject(ch *amqp.Channel, serverSendQueueName, agentSendQueueName, exchangeName, serverSendRoutingKey, agentSendRoutingKey, exchangeKind string, args amqp.Table) (*RabbitMqObject, error) {
//	var err error
//
//	obj := &RabbitMqObject{
//		Ch:                   ch,
//		ServerSendQueueName:  serverSendQueueName,
//		AgentSendQueueName:   agentSendQueueName,
//		ExchangeName:         exchangeName,
//		ServerSendRoutingKey: serverSendRoutingKey,
//		AgentSendRoutingKey:  agentSendRoutingKey,
//		ExchangeType:         exchangeKind,
//	}
//
//	if exchangeKind == "simple" || exchangeKind == "work" {
//		err = obj.DeclareQueue(args)
//		if err != nil {
//			combinedErr := fmt.Errorf("declare Mq ServerSendQueue Failed: %w", err)
//			return nil, combinedErr
//		} else {
//			return obj, nil
//		}
//	} else {
//		// DeclareExchange 声明交换机
//		err = obj.DeclareExchange()
//		if err != nil {
//			combinedErr := fmt.Errorf("declare Mq Exchange Failed: %w", err)
//			return nil, combinedErr
//		}
//
//		// DeclareQueue 声明队列
//		err = obj.DeclareQueue(args)
//		if err != nil {
//			combinedErr := fmt.Errorf("declare Mq ServerSendQueue Failed: %w", err)
//			return nil, combinedErr
//		}
//
//		// BindQueue 声明交换机到队列的绑定
//		err = obj.BindQueue()
//		if err != nil {
//			combinedErr := fmt.Errorf("bind Mq ServerSendQueue Failed: %w", err)
//			return nil, combinedErr
//		}
//	}
//	return obj, nil
//}
//
//// DeclareExchange 声明交换机
//func (obj *RabbitMqObject) DeclareExchange() error {
//	return obj.Ch.ExchangeDeclare(
//		obj.ExchangeName, // 交换机名称
//		obj.ExchangeType, // 交换机类型
//		true,             // 是否持久化
//		false,            // 是否自动删除
//		false,            // 是否为排他
//		false,            // 是否阻塞
//		nil,              // 其他参数
//	)
//}
//
//// DeclareQueue 声明队列
//func (obj *RabbitMqObject) DeclareQueue(args amqp.Table) error {
//	var err error
//	obj.ServerSendQueue, err = obj.Ch.QueueDeclare(
//		obj.ServerSendQueueName, // 队列名称
//		true,  // 是否持久化
//		false, // 是否自动删除
//		false, // 是否为排他队列
//		false, // 是否等待服务器确认
//		args,  // 其他参数
//	)
//	if err != nil {
//		return err
//	}
//
//	obj.AgentSendQueue, err = obj.Ch.QueueDeclare(
//		obj.AgentSendQueueName, // 队列名称
//		true,  // 是否持久化
//		false, // 是否自动删除
//		false, // 是否为排他队列
//		false, // 是否等待服务器确认
//		args,  // 其他参数
//	)
//	return err
//}
//
//// BindQueue 声明交换机到队列的绑定
//func (obj *RabbitMqObject) BindQueue() error {
//	err := obj.Ch.QueueBind(
//		obj.ServerSendQueueName,  // 队列名称
//		obj.ServerSendRoutingKey, // 路由键
//		obj.ExchangeName,         // 交换机名称
//		false,                    // 是否阻塞
//		nil,                      // 其他参数
//	)
//	if err != nil {
//		return err
//	}
//	err = obj.Ch.QueueBind(
//		obj.AgentSendQueue.Name, // 队列名称
//		obj.AgentSendRoutingKey, // 路由键
//		obj.ExchangeName,        // 交换机名称
//		false,                   // 是否阻塞
//		nil,                     // 其他参数
//	)
//	return err
//}
