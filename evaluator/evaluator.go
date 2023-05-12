package evaluator

import (
	_ "embed"
	"github.com/iliakap/go-js-demo/contextgetter"
	"regexp"
	v8 "rogchap.com/v8go"
)

//go:embed setupContextGetter.js
var setupContextGetter string

//go:embed setupFuncs.js
var setupFuncs string

//go:embed moment.js
var moment string

var contextGetRegex = regexp.MustCompile(`\$\.([0-9a-zA-Z\.]*)`)

type Evaluator interface {
	Evaluate(string) (*v8.Value, error)
}

type JSEvaluator struct {
	ContextGetter contextgetter.ContextGetter
	V8Context     *v8.Context
}

func NewJSEvaluator(contextGetter contextgetter.ContextGetter) *JSEvaluator {
	iso := v8.NewIsolate()
	ctx := v8.NewContext(iso)

	global := ctx.Global()

	// get the values from the context
	getCtx := v8.NewFunctionTemplate(iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		path := info.Args()[0].String()
		val, _ := contextGetter.GetFromContext(path)
		value, _ := v8.NewValue(iso, val)
		return value
	})
	global.Set("getCtx", getCtx.GetFunction(ctx))

	// setup js environment
	ctx.RunScript(setupContextGetter, "setupContextGetter.js")
	ctx.RunScript(moment, "moment.js")
	ctx.RunScript(setupFuncs, "setupFuncs.js")

	return &JSEvaluator{
		ContextGetter: contextGetter,
		V8Context:     ctx,
	}
}

func (e *JSEvaluator) Evaluate(str string) (*v8.Value, error) {
	return e.V8Context.RunScript(prepareContextString(str), "eval.js")
}

func prepareContextString(str string) string {
	return contextGetRegex.ReplaceAllString(str, "$['$1']")
}
