package node

type NodeOptions struct{}

var (
	nodeDescription = ``

	nodeExample = `
		# add node to cluster
		weaver node add
		
		`
)

func (o *NodeOptions) Validate() error {
	return nil
}

func NewNodeOptions() *NodeOptions {
	return &NodeOptions{}
}
