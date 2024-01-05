package types

type Website struct {
	Name   string `xml:"name,attr"`
	Url    string
	Course []string
}

type Website2 struct {
	Url int32
}
