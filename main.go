package main

import (
	_ "embed"
	"fmt"
	"github.com/iliakap/go-js-demo/contextgetter"
	"github.com/iliakap/go-js-demo/evaluator"
)

const dummyContext = `{"data": {"duration": "P10DT3H30M"}}`

func main() {
	ctxGetter := contextgetter.New(dummyContext)
	jsEvaluator := evaluator.NewJSEvaluator(ctxGetter)

	fmt.Println(jsEvaluator.Evaluate(`now().sub($.data.duration)`)) // simple
	fmt.Println(jsEvaluator.Evaluate(`now().sub($.data.duration)`))
}
