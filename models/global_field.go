package models

import (
	"github.com/go-redis/redis"
	"github.com/jaimejorge/go-cobinhood/pkg/cobinhood"
	"github.com/patrickmn/go-cache"
	"github.com/preichenberger/go-gdax"
	pusher "github.com/pusher/pusher-http-go"
	"github.com/qor/transition"
)

var (
	Cred       *Credential
	AMQPQueue_ *AMQPQueue

	TransactionStateMachine *transition.StateMachine
	DepositStateMachine     *transition.StateMachine
	coinbaseClient          *gdax.Client
	PusherClient            *pusher.Client
	cobinhoodClient         *cobinhood.CobinhoodClient
	Cache                   *cache.Cache
	redisClient             *redis.Client
)
