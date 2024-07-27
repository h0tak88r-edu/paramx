package types


type TemplatePart int8


const (
	Subdomain TemplatePart = iota
	Endpoint
	Query
)

func (t TemplatePart) String() string {
	return [...]string{"subdomain", "endpoint", "query"}[t]
}
