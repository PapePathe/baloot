package dumb_broker

type segmentConfig struct {
	MaxStoreBytes uint64
	MaxIndexBytes uint64
	InitialOffset uint64
}

type Config struct {
	Segment segmentConfig
}
