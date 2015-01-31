package onlinelabs

const (
	AccountURL = "https://account.cloud.online.net"
	APIURL     = "https://api.cloud.online.net"
)

// ClientInterface as documented at https://doc.cloud.online.net/api/
type ClientInterface interface {
	CreateKey(string, string) (string, error)
	DestroyKey(string) error
	CreateServer(string, string, string, []*Volume, []string) (*Server, error)
	GetServer(string) (*Server, error)
	PowerOffServer(string) error
	DestroyServer(string) error
	CreateSnapshot(string, string, string) (*Snapshot, error)
	CreateImage(string, string, string, string) (*Image, error)
	DestroyImage(string) error
}

type Client struct {
	APIToken       string
	OrganizationID string
	AccountURL     string
	APIURL         string
}

func NewClient(apiToken, orgID string) ClientInterface {
	return &Client{
		APIToken:       apiToken,
		OrganizationID: orgID,
		AccountURL:     AccountURL,
		APIURL:         APIURL,
	}
}

func (c *Client) CreateKey(name, format string) (string, error) {
	return "", nil
}

func (c *Client) DestroyKey(keyID string) error {
	return nil
}

func (c *Client) CreateServer(name, org, image string, volumes []*Volume, tags []string) (*Server, error) {
	return nil, nil
}

func (c *Client) GetServer(id string) (*Server, error) {
	return nil, nil
}

func (c *Client) DestroyServer(id string) error {
	return nil
}

func (c *Client) PowerOffServer(id string) error {
	return nil
}

func (c *Client) CreateSnapshot(name, org, volumeID string) (*Snapshot, error) {
	return nil, nil
}

func (c *Client) CreateImage(org, name, arch, rootVolume string) (*Image, error) {
	return nil, nil
}

func (c *Client) DestroyImage(id string) error {
	return nil
}

func NewAPIRequest(c *Client, path string, body interface{}) ([]byte, error) {
	return nil, nil
}
