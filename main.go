package main

import (
	_ "embed"
	"fmt"
	"github.com/iliakap/go-js-demo/contextgetter"
	"github.com/iliakap/go-js-demo/evaluator"
)

const dummyContext = `{"data": {"duration": "P10D"}}`

func main() {
	ctxGetter := contextgetter.New(dummyContext)
	jsEvaluator := evaluator.NewJSEvaluator(ctxGetter)

	fmt.Println(jsEvaluator.Evaluate(`moment().subtract(moment.duration($.data.duration)).format("YYYY-MM-DD")`))
}
