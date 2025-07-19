package mqtt

type ClientOptions struct{}

func NewClientOptions() *ClientOptions                        { return &ClientOptions{} }
func (o *ClientOptions) AddBroker(b string) *ClientOptions    { return o }
func (o *ClientOptions) SetClientID(id string) *ClientOptions { return o }

type Token interface {
	Wait() bool
	Error() error
}
type noopToken struct{}

func (t noopToken) Wait() bool   { return true }
func (t noopToken) Error() error { return nil }

type Client interface {
	Connect() Token
	Publish(topic string, qos byte, retained bool, payload interface{}) Token
}

type mockClient struct{}

func NewClient(opts *ClientOptions) Client { return &mockClient{} }
func (c *mockClient) Connect() Token       { return noopToken{} }
func (c *mockClient) Publish(topic string, qos byte, retained bool, payload interface{}) Token {
	return noopToken{}
}
