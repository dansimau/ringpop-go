package admin

import (
	"encoding/json"
	"time"

	"github.com/uber/tchannel-go"
	tchannelJson "github.com/uber/tchannel-go/json"
)

type AdminClient struct {
	jsonClient      *tchannelJson.Client
	tchanReqTimeout time.Duration
}

// NewAdminClient creates a new ringpop admin client for the node running on
// hostport, where hostport is a string with the form <host>:<port>.
// tchanTimeout is the timeout value used to cap all tchannel requests.
func NewAdminClient(hostport string, tchan *tchannel.Channel, tchanTimeout time.Duration) *AdminClient {
	client := tchannelJson.NewClient(tchan, "ringpop", &tchannelJson.ClientOptions{HostPort: hostport})
	return &AdminClient{
		jsonClient:      client,
		tchanReqTimeout: tchanTimeout,
	}
}

func (c *AdminClient) call(endpoint string) (interface{}, error) {
	response := make(map[string]interface{})

	ctx, cancel := tchannelJson.NewContext(c.tchanReqTimeout)
	defer cancel()

	err := c.jsonClient.Call(ctx, endpoint, &struct{}{}, &response)

	return response, err
}

// Stats returns stats from the node.
func (c *AdminClient) Stats() (*Stats, error) {
	response, err := c.call("/admin/stats")
	if err != nil {
		return nil, err
	}

	// Marshal and then unmarshal again. This is stupid, but it's just a
	// temporary hack until we figure out how to use TChannel/raw.
	resJson, err := json.Marshal(response)
	if resJson == nil {
		return nil, err
	}

	stats := &Stats{}
	err = json.Unmarshal(resJson, stats)
	if stats == nil {
		return nil, err
	}

	return stats, err
}
