/*
Copyright © 2023 EGBEWATT M. KOKOU	 kokou.egbewatt@gmail.com
*/
package main

func main() {

}

// Need threads for each system?
// figure out how to setup etcd & kafka for communication
// use GRPC or GraphQL for master - worker communication
// how are worker nodes synced to  master? install worker client on node
//

// Do we decide the type of machine our Master Node Runs? Need it for ETCD setup
// single ETCD or cluster in case of Multiple master Nodes with a leader
// control plane components and etcd communicate over secure channels using TLS certificates.
// backup etcd data stored on master node directories for data durability and consistency
// operators to automate etcd backup and restore procedures
// scale, upgrade & monitor health and performance of etcd

//
