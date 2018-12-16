package config

type Client struct {
	Send chan string
	Stop chan bool
}
