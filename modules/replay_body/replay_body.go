package replay_body

import (
	"github.com/golang/glog"
	"github.com/mailgun/vulcan/netutils"
	"github.com/mailgun/vulcan/vmod"
	"net/http"
)

func init() {
	vmod.Register(NewBodyBuffer())
}

type BodyBuffer struct {
	sizeInMemory int64
}

const DEFAULT_MEMORY_BUFFER_LIMIT = 1048576

func NewBodyBuffer() *BodyBuffer {
	return &BodyBuffer{sizeInMemory: DEFAULT_MEMORY_BUFFER_LIMIT}
}

func (bb *BodyBuffer) Name() string {
	return "replay_body"
}

func (bb *BodyBuffer) Configure() (vmod.HandlerFunc, error) {
	return func(w *vmod.ResponseWriter, req *vmod.Request) error {
		// We are allowed to fallback in case of upstream failure,
		// record the request body so we can replay it on errors.

		body, err := netutils.NewBodyBuffer(bb.sizeInMemory, req.Body)
		if err != nil {
			glog.Errorf("Request read error %s", err)
			return netutils.NewHttpError(http.StatusBadRequest)
		}

		req.Body = body
		/*		req.OnFinish(func() {
					defer body.Close()
				})
		*/
		return nil
	}, nil
}
