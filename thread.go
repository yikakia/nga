package nga

import (
	"encoding/json"

	"github.com/yikakia/nga/resp"
)

func (c *Client) Thread(fid string) (*resp.Thread, error) {
	get, err := c.getDefaultClient().
		SetQueryParam("fid", fid).
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
