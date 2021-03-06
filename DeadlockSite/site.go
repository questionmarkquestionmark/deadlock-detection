package DeadlockSite

//creates nodes

type Site struct {
	siteNodeCount  int
	TotalNodeCount int
	NodeList       []*DeadlockNode
	SiteList       []*Site
}

func NewSite(bossNodeAddr string, numNode int, totalNode int, siteList []*Site) *Site {
	site := &Site{
		siteNodeCount:  numNode,
		TotalNodeCount: totalNode,
		NodeList:       make([]*DeadlockNode, numNode),
		SiteList:       siteList,
	}
	for i := 0; i < numNode; i++ { //builds the nodes and gives them their init parameters
		var dependence string
		if i > 0 {
			dependence = site.NodeList[i-1].node.ID().Address
		} else {
			dependence = ""
		}
		dn := NewDeadlockNode(bossNodeAddr, numNode, site, dependence)
		go func(node *DeadlockNode) {
			node.Start()
		}(dn)
		site.NodeList[i] = dn
	}
	return site
}
