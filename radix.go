package radix

import (
	"reflect"

	yaml "gopkg.in/yaml.v2"

	"github.com/TBoileau/go-micro-framework"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mediocregopher/radix.v2/redis"
)

type RadixMiddleware struct {
	framework.Middleware
	Parameters map[string]map[string]string
}

func Initialize(configFile string) *RadixMiddleware {
	middleware := RadixMiddleware{}
	middleware.ConfigFile = configFile
	return &middleware
}

func (middleware *RadixMiddleware) Get(connection string) *redis.Client {
	client, _ := redis.Dial("tcp", middleware.Parameters[connection]["hostname"]+":"+middleware.Parameters[connection]["port"])
	return client
}

func (middleware RadixMiddleware) Register() reflect.Value {
	middleware.Parameters = make(map[string]map[string]string)
	yaml.Unmarshal(middleware.LoadParameters(), &middleware.Parameters)
	return reflect.ValueOf(&middleware)
}

func (middleware RadixMiddleware) GetName() string {
	return "radix"
}
