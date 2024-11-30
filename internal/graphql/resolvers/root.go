// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/cmd/apiserver/build"
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/logger"
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync"
)

const (
	// subscriptionQueueCapacity is the amount of subscriptions we buffer for processing.
	subscriptionQueueCapacity = 100

	// subscriptionInitialCapacity is the initial length of the subscription queue.
	subscriptionInitialCapacity = 100

	// listMaxEdgesPerRequest maximal number of edges end-client can request in one query.
	listMaxEdgesPerRequest uint32 = 100
)

// rootResolver represents the ApiResolver implementation.
type rootResolver struct {
	// service terminator
	wg      sync.WaitGroup
	cg      singleflight.Group
	sigStop chan bool
}

// log represents the logger to be used by the repository.
var log logger.Logger

// config represents the configuration setup used by the repository
// to establish and maintain required connectivity to external services
// as needed.
var cfg *config.Config

// SetLogger sets the repository logger to be used to collect logging info.
func SetLogger(l logger.Logger) {
	log = l
}

// SetConfig sets the repository configuration to be used to establish
// and maintain external repository connections.
func SetConfig(c *config.Config) {
	cfg = c
}

// New creates a new root resolver instance and initializes its internal structure.
func New() ApiResolver {
	if cfg == nil {
		panic(fmt.Errorf("missing configuration"))
	}
	if log == nil {
		panic(fmt.Errorf("missing logger"))
	}

	// create new resolver
	rs := rootResolver{
		sigStop: make(chan bool, 1),
	}

	// handle broadcast and subscriptions in a separate routine
	rs.wg.Add(1)
	go rs.run()

	return &rs
}

// Close terminates resolver's broadcast service.
func (rs *rootResolver) Close() {
	// log
	log.Notice("GraphQL resolver is closing")

	// send the signal
	rs.sigStop <- true
	rs.wg.Wait()
}

// run monitors and handles subscriptions and broadcasts incoming events to their subscribers.
func (rs *rootResolver) run() {
	// sign off on leaving
	defer func() {
		// terminate
		log.Notice("GraphQL resolver done")
		rs.wg.Done()
	}()

	// log action
	log.Notice("GraphQL resolver started")
	<-rs.sigStop
}

// listLimitCount enforces maximum size of a requested list to given limit
// amount of edges preserving the direction of the load. Note the count can
// be both positive and negative. It controls the direction in which the list
// of edges is built. Limit is always positive and adjusted to the count direction
// on return.
func listLimitCount(count int32, limit uint32) int32 {
	// requested count is zero?
	// this should not happen but let's return the max range than
	if count == 0 {
		return int32(limit)
	}

	// is the count already inside the limit range (e.g. correct input)?
	// return the valid original
	if (count > 0 && uint32(count) < limit) || (count < 0 && count > -int32(limit)) {
		return count
	}

	// the count is over the limit
	// so we return the limit being the max. value allowed
	// adjusted to the original direction
	if count < 0 {
		return -int32(limit)
	}

	return int32(limit)
}

// Version resolves the current version of the API server.
func (rs *rootResolver) Version() string {
	return build.Short(cfg)
}
