package models

type Payload struct{
	Method string
	Addr string
	URL string
	Proto string
	Host string
	StatusCode int
	ContentLength int
	ContentType string
}
