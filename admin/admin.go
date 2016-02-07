package admin

import (
	"time"

	"github.com/uber/tchannel-go"
	"github.com/uber/tchannel-go/json"
)

type AdminClient struct {
	jsonClient      *json.Client
	tchanReqTimeout time.Duration
}

// NewAdminClient creates a new ringpop admin client for the node running on
// hostport, where hostport is a string with the form <host>:<port>.
// tchanTimeout is the timeout value used to cap all tchannel requests.
func NewAdminClient(hostport string, tchan *tchannel.Channel, tchanTimeout time.Duration) *AdminClient {
	client := json.NewClient(tchan, "ringpop", &json.ClientOptions{HostPort: hostport})
	return &AdminClient{
		jsonClient:      client,
		tchanReqTimeout: tchanTimeout,
	}
}

func (c *AdminClient) Stats() (*Stats, error) {
	response := make(map[string]interface{})

	ctx, cancel := json.NewContext(c.tchanReqTimeout)
	defer cancel()

	err := c.jsonClient.Call(ctx, "/admin/stats", &struct{}{}, &response)
	if err != nil {
		return nil, err
	}

	stats := &Stats{response}

	return stats, nil
}
