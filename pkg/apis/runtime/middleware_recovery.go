package runtime

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"

	"github.com/seal-io/seal/utils/log"
)

// Recovering is a gin middleware,
// which is the same as gin.Recovery,
// but friendly message information can be provided according to the request header.
func Recovering() Handle {
	logger := log.WithName("api")

	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				var e error
				if err, ok := r.(error); ok {
					e = err
				} else {
					e = fmt.Errorf("%v", r)
				}
				cs := callstack(3)
				logger.Errorf("panic observing: %v, callstack: \n%s", e, cs)

				if isStreamRequest(c) {
					// Stream request always send header at first,
					// so we don't need to rewrite.
					return
				}

				c.AbortWithStatusJSON(http.StatusInternalServerError, httpError{
					code: http.StatusInternalServerError,
				}) // TODO negotiate.
			}
		}()

		c.Next()
	}
}

func callstack(skip int) []byte {
	var buf bytes.Buffer

	for i := skip; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fn := "???"
		f := runtime.FuncForPC(pc)

		if f != nil {
			fn = f.Name()
		}
		_, _ = fmt.Fprintf(&buf, "%s\n\t%s:%d (0x%x)\n", fn, file, line, pc)
	}

	return buf.Bytes()
}
