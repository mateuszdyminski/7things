package main

import (
	"context"
	"fmt"
	"github.com/mateuszdyminski/7things/logs/logsctx"
	"github.com/pborman/uuid"
	"github.com/uber-go/zap"
	"math"
	"net/http"
	"strconv"
)

var httpContext = context.Background()

func main() {
	http.HandleFunc("/square", squareHandler)
	http.ListenAndServe(":9001", nil)
}

func squareHandler(resp http.ResponseWriter, req *http.Request) {
	rqCtx := logsctx.WithRqId(httpContext, uuid.NewRandom().String())

	arg := req.URL.Query().Get("arg")

	result, err := square(rqCtx, arg)
	if err != nil {
		http.Error(resp, fmt.Sprintf("can't square value: %s! err: %v", arg, err), http.StatusBadRequest)
		return
	}

	fmt.Fprint(resp, result)
}

func square(ctx context.Context, arg string) (float64, error) {
	logger := logsctx.Logger(ctx)

	if arg == "" {
		return 0, fmt.Errorf("arg should be set!")
	}

	val, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		return 0, err
	}
	logger.Info("got arg", zap.String("val", arg))

	result := math.Pow(val, 2)
	logger.Info("finished", zap.Float64("result", result))

	return result, nil
}
