package ejudge

type Error struct {
	Num     int64  `json:"num"`
	Symbol  string `json:"symbol,omitempty"`
	Message string `json:"message,omitempty"`
	LogID   string `json:"log_id,omitempty"`
}

type BaseReply struct {
	Ok         bool   `json:"ok"`
	ServerTime int64  `json:"server_time"`
	Action     string `json:"action"`
	RequestId  int64  `json:"request_id,omitempty"`
	ReplyId    int64  `json:"reply_id,omitempty"`
	Error      *Error `json:"error,omitempty"`
}

type Reply[T any] struct {
	Ok         bool   `json:"ok"`
	ServerTime int64  `json:"server_time"`
	Action     string `json:"action"`
	RequestId  int64  `json:"request_id,omitempty"`
	ReplyId    int64  `json:"reply_id,omitempty"`
	Error      *Error `json:"error,omitempty"`
	Result     *T     `json:"result,omitempty"`
}
