package main

type WordSize int

const (
	ZERO WordSize = 6 * iota
	SMALL
	MEDIUM
	LARGE
	XLARGE
	XXLARGE   WordSize = 50
	SEPARATOR          = ", "
)

type stringFunc func(s string) (result string)

type filterFunc func(ws string) bool

type ChainLink struct {
	Data []string
}

func (c *ChainLink) Value() []string {
	return c.Data
}

func (c *ChainLink) Map(fn stringFunc) *ChainLink {
	mapped := make([]string, len(c.Data))
	for i, v := range c.Data {
		mapped[i] = fn(v)
	}

	c.Data = mapped
	return c
}

func (c *ChainLink) Filter(fn filterFunc) *ChainLink {
	mapped := make([]string, 0)
	for _, v := range c.Data {
		if fn(v) {
			mapped = append(mapped, v)
		}
	}
	c.Data = mapped
	return c
}
