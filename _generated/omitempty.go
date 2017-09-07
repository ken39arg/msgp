package _generated

//go:generate msgp

type OmitemptyStrct struct {
	Foo string  `msg:",omitempty"`
	Bar float32 `msg:"b,omitempty"`
}

type Omitempty struct {
	Str    string         `msg:",omitempty"`
	Name   string         `msg:"nm,omitempty"`
	Num    int            `msg:",omitempty"`
	Boo    bool           `msg:",omitempty"`
	Nils   *string        `msg:",omitempty"`
	Arr    []string       `msg:",omitempty"`
	Strct  OmitemptyStrct `msg:"strct,omitempty"`
	Req    int            `msg:"req"`
	Ignore int            `msg:"-"`
}
