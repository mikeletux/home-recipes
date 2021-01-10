package guid

import "github.com/rs/xid"

type Guid interface {
	GetGUID() string
}

type GuidXid struct{}

func NewGuidXid() Guid {
	guid := &GuidXid{}
	return guid
}

func (x *GuidXid) GetGUID() string {
	guid := xid.New()
	return guid.String()
}
