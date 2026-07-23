package nga

import (
	"encoding/json"
	"slices"

	"github.com/yikakia/nga/resp"
)

func (c *Client) Thread(fid string, opts ...Option) (*resp.Thread, error) {
	req := c.getDefaultClient().
		SetQueryParam("fid", fid)

	_ctx := ApplyOptions(slices.Insert(opts, 0, WithReq(req))...)

	get, err := _ctx.Req.
		Get(c.BaseUrl + "/thread.php")
	if err != nil {
		return nil, err
	}

	utf8, err := gbkToUtf8(get.String())
	if err != nil {
		return nil, err
	}
	thread := new(resp.Thread)
	err = json.Unmarshal([]byte(utf8), thread)
	if err != nil {
		return nil, err
	}

	return thread, nil
}
