# TimeMon

Go library to monitor clock synchronization across a cluster.

## Inspiration

Inspired by the book [Designing Data-Intensive Applications](https://dataintensive.net/) by Martin Kleppmann. On Pg. 291, Relying on Synchronized Clocks.

*“Any node whose clock drifts too far from the others should be declared dead and removed from the cluster”.*

I aim to implement this functionality by monitoring wall-clock time across a cluster and alerting of drifting nodes to the cluster group.

## Objectives

- Implement a transport-agnostic library to monitor wall-clock time across a process group using a leader-follower algorithm. This would allow users to implement applications dependent upon synchronized clocks (like linearizable databases).
- Design the library to enable applications to implement Spanner's TrueTime-like functionality (Commit Windows?) (Checkout [CockroachDB](https://www.cockroachlabs.com/blog/living-without-atomic-clocks)) i.e communicate upper bound on wall-clock time synchronization. 

## References

- My blog post on clocks and timers: [Clocks, Timers and Virtualization](https://arush15june.github.io/posts/2020-07-12-clocks-timers-virtualization/)
- [Practical Uses of Synchronized Clocks in Distributed Systems](http://www.dainf.cefetpr.br/~tacla/SDII/PracticalUseOfClocks.pdf) by Barbara Liskov
- CockroachDB: [Living without Atomic Clocks](https://www.cockroachlabs.com/blog/living-without-atomic-clocks)
- CockroachDB: [CockroachDB's Consistency Model](https://www.cockroachlabs.com/blog/consistency-model/)

## TODO

- Actually write code
- Design and Specification Document
- Benchmark Public Clouds: AWS, Azure, GCP
