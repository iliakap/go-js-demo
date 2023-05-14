package main

import (
	_ "embed"
	"fmt"
	"github.com/iliakap/go-js-demo/contextgetter"
	"github.com/iliakap/go-js-demo/evaluator"
)

const dummyContext = `{"data": {"duration": "P10DT3H30M"}}`

const demo1 = `{{js now().sub($.data.duration) }}`
const demo2 = `{{js now().sub($.data.duration).format("hh:mm:ss DD/MM/yyyy") }}`

func main() {
	ctxGetter := contextgetter.New(dummyContext)
	jsEvaluator := evaluator.NewJSEvaluator(ctxGetter)

	fmt.Println(jsEvaluator.Evaluate(`{{js now().sub($.data.duration) }}`))
	fmt.Println(jsEvaluator.Evaluate(`{{js now().sub($.data.duration).format("hh:mm:ss DD/MM/yyyy") }}`))
}
