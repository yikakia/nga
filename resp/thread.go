package resp

type Thread struct {
	Data struct {
		T []struct {
			Tid        int    `json:"tid"` // 帖子ID
			Fid        int    `json:"fid"`
			QuoteFrom  int    `json:"quote_from"`
			TopicMisc  any    `json:"topic_misc"`
			Author     string `json:"author"`
			Authorid   any    `json:"authorid"` // 作者ID 匿名的时候是 string?? 这什么诡异的接口
			Subject    string `json:"subject"`
			Type       int    `json:"type"`
			Postdate   int    `json:"postdate"` // 帖子创建时间 Unix时间戳
			Lastpost   int    `json:"lastpost"` // 最后回复时间 Unix时间戳
			Lastposter string `json:"lastposter"`
			Replies    int    `json:"replies"` // 回复数
			Lastmodify int    `json:"lastmodify"`
			Recommend  int    `json:"recommend"`
			Tpcurl     string `json:"tpcurl"`
		} `json:"__T"`
		TROWS     int `json:"__T__ROWS"`
		TROWSPAGE int `json:"__T__ROWS_PAGE"`
		RROWSPAGE int `json:"__R__ROWS_PAGE"`
	} `json:"data"`
	Time int `json:"time"`
}
