package zrpc

import (
	"fmt"

	"google.golang.org/grpc/resolver"
)

var addrs = []string{"localhost:50052", "localhost:50053", "localhost:50054", "localhost:50055"}

const scheme = "gametake"
const serviceName = "baloot.GameTakeLearning"

type gametakeResolverBuilder struct{}

func (gr *gametakeResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &gametakeResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{
			serviceName: addrs,
		},
	}

	r.start()
	return r, nil
}
func (gr *gametakeResolverBuilder) Scheme() string { return scheme }

type gametakeResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

func (r *gametakeResolver) start() {
	addrStrs := r.addrsStore[serviceName]
	addrs := make([]resolver.Address, len(addrStrs))
	for i, s := range addrStrs {
		addrs[i] = resolver.Address{Addr: s}
		fmt.Println(addrs[i])
	}
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}
func (r *gametakeResolver) ResolveNow(o resolver.ResolveNowOptions) {}
func (r *gametakeResolver) Close()                                  {}

func init() {
	resolver.Register(&gametakeResolverBuilder{})
}
