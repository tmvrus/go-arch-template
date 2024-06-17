package pkg

type Client struct {
}

func NewClient(url string) Client {
	return Client{}
}

func (Client) DoWork() error {
	return nil
}
