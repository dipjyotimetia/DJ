package stubs

type Config struct {
	Host        string
	Port        int
	Header      map[string]string
	Services    []Service
	ResponseDir string
}

type Service struct {
	Prefix    string
	Endpoints []Endpoint
}

type Endpoint struct {
	Method   string
	Name     string
	Response *Response
	Matches  []Match
}

type Response struct {
	Status int
	Header map[string]string
	Body   string
}

type Match struct {
	RequestBody string
	Response    *Response
}
