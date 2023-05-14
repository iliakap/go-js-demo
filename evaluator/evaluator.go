package evaluator

import (
	_ "embed"
	"github.com/dop251/goja"
	"github.com/iliakap/go-js-demo/contextgetter"
	"regexp"
)

//go:embed setupFuncs.js
var setupFuncs string

//go:embed moment.js
var moment string

var contextGetRegex = regexp.MustCompile(`\$\.([0-9a-zA-Z\.]*)`)
var jsCommandRegex = regexp.MustCompile(`\{\{js (.*) \}\}`)

type Evaluator interface {
	Evaluate(string) (goja.Value, error)
}

type JSEvaluator struct {
	ContextGetter contextgetter.ContextGetter
	js            *goja.Runtime
}

func NewJSEvaluator(contextGetter contextgetter.ContextGetter) *JSEvaluator {
	runtime := goja.New()
	runtime.Set("$", func(j goja.FunctionCall) goja.Value {
		path := j.Arguments[0].String()
		val, _ := contextGetter.GetFromContext(path)
		return runtime.ToValue(val)
	})

	runtime.RunString(moment)
	runtime.RunString(setupFuncs)

	return &JSEvaluator{
		ContextGetter: contextGetter,
		js:            runtime,
	}
}

func (e *JSEvaluator) Evaluate(str string) (goja.Value, error) {
	str = jsCommandRegex.ReplaceAllString(str, "$1")
	return e.js.RunString(prepareContextString(str))
}

func prepareContextString(str string) string {
	return contextGetRegex.ReplaceAllString(str, "$('$1')")
}
