package main

import (
	_ "embed"
	"fmt"
	"github.com/iliakap/go-js-demo/contextgetter"
	"github.com/iliakap/go-js-demo/evaluator"
)

func main() {
	ctx := contextgetter.New()
	jsEvaluator := evaluator.NewJSEvaluator(ctx)

	fmt.Println(jsEvaluator.Evaluate(`moment().subtract(moment.duration($.data.duration)).format("YYYY-MM-DD")`))
}
