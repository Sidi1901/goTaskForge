module Sidi1901/goTaskForge/worker

go 1.25.0

replace Sidi1901/goTaskForge/shared => ../shared

require (
	Sidi1901/goTaskForge/shared v0.0.0-00010101000000-000000000000
	github.com/redis/go-redis/v9 v9.19.0
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
)
