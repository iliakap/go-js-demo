package contextgetter

import (
	"encoding/json"
	"github.com/yalp/jsonpath"
)

const dummyContext = `{"data": {"duration": "P10D"}}`

type ContextGetter interface {
	GetFromContext(string) (string, error)
}

type Ctx struct {
}

func New() *Ctx {
	return &Ctx{}
}

func (c *Ctx) GetFromContext(str string) (string, error) {
	p := "$." + str
	var con interface{}
	_ = json.Unmarshal([]byte(dummyContext), &con)
	read, err := jsonpath.Read(con, p)
	return read.(string), err
}
