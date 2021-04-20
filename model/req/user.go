package req

type Register struct {
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	NickName  string `json:"nick_name,omitempty"`
	HeaderImg string `json:"header_img,omitempty"`
}
