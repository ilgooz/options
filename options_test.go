package options

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Branch struct {
	Id   int64  `option:"id"`
	Name string `option:"value"`
}

var branches = []Branch{{1, "İstanbul"}, {2, "Ankara"}}

func TestBasic(t *testing.T) {
	assert.Equal(t, []Option{
		{"id": "1", "value": "İstanbul"},
		{"id": "2", "value": "Ankara"},
	}, Encode(branches))
}

type Currency struct {
	Id       int64 `option:"id"`
	Name     string
	Code     string
	FullName string `option:"value"`
}

type OCurrency Currency

func (c Currency) EncodeOption() Option {
	c.FullName = fmt.Sprintf("%s - %s", c.Code, c.Name)
	return Encode(OCurrency(c))[0]
}

var currencies = []Currency{
	{Id: 1, Name: "Lira", Code: "TRY"},
	{Id: 2, Name: "US Dollar", Code: "USD"},
}

func TestCustomEncoder(t *testing.T) {
	assert.Equal(t, []Option{
		{"id": "1", "value": "TRY - Lira"},
		{"id": "2", "value": "USD - US Dollar"},
	}, Encode(currencies))
}
