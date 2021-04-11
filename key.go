package code

// code common key
const (
	RemoteIP    = "remote_ip"   // Network
	RemotePort  = "remote_port" // Network
	ServerAddr  = "server_addr" // Network
	ClientAddr  = "client_addr" // Network
	Cluster     = "cluster"     // Router
	Color       = "color"       // Router
	Trace       = "trace"       // Trace
	Caller      = "caller"      // Trace
	CPUUsage    = "cpu_usage"   // Dispatch
	Errors      = "errors"      // Dispatch
	Requests    = "requests"    // Dispatch
	Timeout     = "timeout"     // Timeout
	Mirror      = "mirror"      // Mirror
	Mid         = "mid"         // Mid 外网账户用户id，NOTE: ！！！业务可重新修改key名！！！
	Device      = "device"      // Device 客户端信息
	Criticality = "criticality" // Criticality 重要性
)

var outgoingKey = map[string]struct{}{
	Color:       {},
	RemoteIP:    {},
	RemotePort:  {},
	Mirror:      {},
	Criticality: {},
}

var incomingKey = map[string]struct{}{
	Caller: {},
}

// IsOutgoingKey represent this key should propagate by rpc.
func IsOutgoingKey(key string) bool {
	_, ok := outgoingKey[key]
	return ok
}

// IsIncomingKey represent this key should extract from rpc metadata.
func IsIncomingKey(key string) (ok bool) {
	_, ok = outgoingKey[key]
	if ok {
		return
	}
	_, ok = incomingKey[key]
	return
}
