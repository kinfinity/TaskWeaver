package node

type Node struct {
	Name            string
	Ip              string
	Cores           int
	Memory          int
	MemoryAllocated int
	Disk            int
	DiskAllocated   int
}

// node information is sourced when node joins the cluster and sent to master
// how does join happen?

// Look into virtualizing - Nodes for testing ?
// https://en.wikipedia.org/wiki/Kernel-based_Virtual_Machine

func (n *Node) GetMetrics(name string, ip string) (*Node, error) {
	return n, nil
}

// Node Agent
// gathers information about the node
//
