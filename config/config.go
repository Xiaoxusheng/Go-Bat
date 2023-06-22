package config

type Messages struct {
	Post_type    string  `json:"post_type,omitempty"`
	Message_type string  `json:"message_type,omitempty"`
	Time         int64   `json:"time,omitempty"`
	Self_id      int64   `json:"self_id,omitempty"`
	Sub_type     string  `json:"sub_type,omitempty"`
	Message_id   int64   `json:"message_id,omitempty"`
	User_id      int64   `json:"user_id,omitempty"`
	Target_id    int64   `json:"target_id,omitempty"`
	Message      string  `json:"message,omitempty"`
	Sender       *Sender `json:"sender,omitempty"`
}

type Sender struct {
	Age     int64  `json:"age,omitempty"`
	Sex     string `json:"sex,omitempty"`
	User_id int64  `json:"user_id,omitempty"`
}
