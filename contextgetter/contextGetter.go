package contextgetter

import (
	"encoding/json"
	"github.com/yalp/jsonpath"
)

type ContextGetter interface {
	GetFromContext(string) (string, error)
}

type Ctx struct {
	con string
}

func New(con string) *Ctx {
	return &Ctx{
		con: con,
	}
}

func (c *Ctx) GetFromContext(str string) (string, error) {
	p := "$." + str
	var con interface{}
	_ = json.Unmarshal([]byte(c.con), &con)
	read, err := jsonpath.Read(con, p)
	return read.(string), err
}
