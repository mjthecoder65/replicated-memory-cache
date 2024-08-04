# Replicated Memory Cache

Distributed in-memory cache system using Go, Protocal Buffers(protobuf), and gRPC
to improve data retrieval times and ensure efficient inter-service communication.

### Technologies

Go, Protocal Buffers, gRPC, Consul, Docker, Gin Framework, Synchronization mechanism.

### Features

- Service discovery and node management using Consul
- Basic Auth for securing cache endpoints.
- Cache enviction policies and data synchronization across distributed nodes.
- Conducated Unit testing to maintain high code quality and reliability.

### Advanced Requirements.

- Allow immediate deletion of some data
- Add "mininimum replication count" functionality that will block (wait)
  until the date gets replicated numbers of peers. \* Add a "full replication" metadata flag which will replicate data
  all active peers. Block until all peers have the data.
- Add distributed lock. Lock should work per key (use full replication to implement)
- Add expiration metadata to concurrent map entries
  - Delete expired entries when appropriate
- Remove peers from the peer list if they stop receiving synchronization data.
- Add keepalive/heartbeart signals between the peers. If a peers stops sending signal, remove it from the peer list.
