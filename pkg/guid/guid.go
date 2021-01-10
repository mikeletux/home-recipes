package guid

import "github.com/rs/xid"

type Guid interface {
	GetGUID() string
}

type GuidXid struct {
	id xid.ID
}

func NewGuidXid() Guid {
	guid := &GuidXid{xid.New()}
	return guid
}

func (x *GuidXid) GetGUID() string {
	return x.id.String()
}
