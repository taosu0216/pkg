package Mq

import "github.com/streadway/amqp"

type Option func(*Options)

func WithExchangeOptions(opts ...ExchangeOption) Option {
	return func(o *Options) {
		for _, opt := range opts {
			opt(o.Exchange)
		}
	}
}

func WithAgentSendQueueOptions(opts ...AgentSendQueueOption) Option {
	return func(o *Options) {
		for _, opt := range opts {
			opt(o.AgentSendQueue)
		}
	}
}

func WithServerSendQueueOptions(opts ...ServerSendQueueOption) Option {
	return func(o *Options) {
		for _, opt := range opts {
			opt(o.ServerSendQueue)
		}
	}
}

func WithAgentSendBindOptions(opts ...AgentSendBindOption) Option {
	return func(o *Options) {
		for _, opt := range opts {
			opt(o.AgentSendBind)
		}
	}
}

func WithServerSendBindOptions(opts ...ServerSendBindOption) Option {
	return func(o *Options) {
		for _, opt := range opts {
			opt(o.ServerSendBind)
		}
	}
}

type ExchangeOption func(*ExchangeInfo)

func WithExchangeDurable(durable bool) ExchangeOption {
	return func(e *ExchangeInfo) {
		e.Durable = durable
	}
}

func WithExchangeAutoDelete(autoDelete bool) ExchangeOption {
	return func(e *ExchangeInfo) {
		e.AutoDelete = autoDelete
	}
}

func WithExchangeInternal(internal bool) ExchangeOption {
	return func(e *ExchangeInfo) {
		e.Internal = internal
	}
}

func WithExchangeNoWait(noWait bool) ExchangeOption {
	return func(e *ExchangeInfo) {
		e.NoWait = noWait
	}
}

func WithExchangeArgs(args amqp.Table) ExchangeOption {
	return func(e *ExchangeInfo) {
		e.Args = args
	}
}

type AgentSendQueueOption func(*AgentSendQueueInfo)

func WithAgentSendQueueDurable(durable bool) AgentSendQueueOption {
	return func(q *AgentSendQueueInfo) {
		q.Durable = durable
	}
}

func WithAgentSendQueueAutoDelete(autoDelete bool) AgentSendQueueOption {
	return func(q *AgentSendQueueInfo) {
		q.AutoDelete = autoDelete
	}
}

func WithAgentSendQueueExclusive(exclusive bool) AgentSendQueueOption {
	return func(q *AgentSendQueueInfo) {
		q.Exclusive = exclusive
	}
}

func WithAgentSendQueueNoWait(noWait bool) AgentSendQueueOption {
	return func(q *AgentSendQueueInfo) {
		q.NoWait = noWait
	}
}

func WithAgentSendQueueArgs(args amqp.Table) AgentSendQueueOption {
	return func(q *AgentSendQueueInfo) {
		q.Args = args
	}
}

type ServerSendQueueOption func(*ServerSendQueueInfo)

func WithServerSendQueueDurable(durable bool) ServerSendQueueOption {
	return func(q *ServerSendQueueInfo) {
		q.Durable = durable
	}
}

func WithServerSendQueueAutoDelete(autoDelete bool) ServerSendQueueOption {
	return func(q *ServerSendQueueInfo) {
		q.AutoDelete = autoDelete
	}
}

func WithServerSendQueueExclusive(exclusive bool) ServerSendQueueOption {
	return func(q *ServerSendQueueInfo) {
		q.Exclusive = exclusive
	}
}

func WithServerSendQueueNoWait(noWait bool) ServerSendQueueOption {
	return func(q *ServerSendQueueInfo) {
		q.NoWait = noWait
	}
}

func WithServerSendQueueArgs(args amqp.Table) ServerSendQueueOption {
	return func(q *ServerSendQueueInfo) {
		q.Args = args
	}
}

type AgentSendBindOption func(*AgentSendBindInfo)

func WithAgentSendBindNoWait(noWait bool) AgentSendBindOption {
	return func(b *AgentSendBindInfo) {
		b.NoWait = noWait
	}
}

func WithAgentSendBindArgs(args amqp.Table) AgentSendBindOption {
	return func(b *AgentSendBindInfo) {
		b.Args = args
	}
}

type ServerSendBindOption func(*ServerSendBindInfo)

func WithServerSendBindNoWait(noWait bool) ServerSendBindOption {
	return func(b *ServerSendBindInfo) {
		b.NoWait = noWait
	}
}

func WithServerSendBindArgs(args amqp.Table) ServerSendBindOption {
	return func(b *ServerSendBindInfo) {
		b.Args = args
	}
}

func NewDefaultExchangeInfo(options ...ExchangeOption) *ExchangeInfo {
	exchange := ExchangeInfo{
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Args:       nil,
	}
	if options == nil {
		return &exchange
	}
	for _, opt := range options {
		opt(&exchange)
	}
	return &exchange
}

func NewDefaultAgentSendQueueInfo(options ...AgentSendQueueOption) *AgentSendQueueInfo {
	queue := AgentSendQueueInfo{
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
		Args:       nil,
	}
	if options == nil {
		return &queue
	}
	for _, opt := range options {
		opt(&queue)
	}
	return &queue
}

func NewDefaultServerSendQueueInfo(options ...ServerSendQueueOption) *ServerSendQueueInfo {
	queue := ServerSendQueueInfo{
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
		Args:       nil,
	}
	if options == nil {
		return &queue
	}
	for _, opt := range options {
		opt(&queue)
	}
	return &queue
}

func NewDefaultAgentSendBindInfo(options ...AgentSendBindOption) *AgentSendBindInfo {
	bind := AgentSendBindInfo{
		NoWait: false,
		Args:   nil,
	}
	if options == nil {
		return &bind
	}
	for _, opt := range options {
		opt(&bind)
	}
	return &bind
}

func NewDefaultServerSendBindInfo(options ...ServerSendBindOption) *ServerSendBindInfo {
	bind := ServerSendBindInfo{
		NoWait: false,
		Args:   nil,
	}
	if options == nil {
		return &bind
	}
	for _, opt := range options {
		opt(&bind)
	}
	return &bind
}
