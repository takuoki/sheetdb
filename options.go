package sheetdb

// ClientOption is an option to change the behavior of the client.
type ClientOption func(client *Client) *Client

// ModelSetName is an option to change model set name using by client.
// If not used, the "default" model set is used.
func ModelSetName(modelSetName string) func(client *Client) *Client {
	return func(client *Client) *Client {
		if client != nil {
			client.modelSetName = modelSetName
		}
		return client
	}
}
