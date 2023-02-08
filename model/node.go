package model

type NodeInfo struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []NodeData `json:"data"`
}

type NodeData struct {
	NodeID   string `json:"nodeId"`
	NodeName string `json:"nodeName"`
	NodeIP   string `json:"nodeIp"`
	XmNodeID string `json:"xmazeNodeId"`
}

type DelVul struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
