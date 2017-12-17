package models

type Chat struct{
	ID float64
	Name string
	Addr_users []string
	MessageBlockId float64
	LastSender, LastMessage string
}
type MessageBlock struct {
	chatId   int64
	Messages []Message
}
type Message struct {
	Addr_author string
	Content string
	Type string
	Chat_Id float64
}
type UserChatInfo struct{
	ID int64 `json:"id"`
	Name string `json:"name"`
	Type int `json:"type"`
	//Addr_users []string
	LastSender string `json:"last_sender"`
	Admin_id int64 `json:"admin_id"`
	//Moders_ids []float64 `json:"moderators_ids"`
	LastMessage *MessageContent `json:"last_message"`
	LastMessageTime int64 `json:"last_message_time"`
	View int `json:"view"`
	Delete bool `json:"delete"`
	Online int64 `json:"online"`
}
type MessageContent struct{
	Message string `json:"content"`
	Documents []int64 `json:"documents"`
	Type string `json:"type"`
}
type User struct {
	ID float64
	Name string
	Login string
	Pass string
}

type NewMessageToUser struct{
	ID int64 `json:"id"`
	ChatId int64 `json:"chat_id"`
	Content *MessageContentToUser `json:"message"`
	AuthorId int64 `json:"author_id"`
	AuthorName string `json:"author_name"`
	AuthorLogin string `json:"author_login"`
	Time int64 `json:"time"`
}

type CreateDHData struct{
	CommonName string
	Organization string
	DNSNames1 string //wiki
	DNSNames2 string //192.168.0.2
	Type string
}

type MessageContentToUser struct{
	Message string `json:"content"`
	Documents []map[string]interface{} `json:"documents"`
	Type string `json:"type"`
}

//type MessageContent struct{
//	Message string `json:"content"`
//	Documents []int64 `json:"documents"`
//	Type string `json:"type"`
//}



type ForceMsgToUser struct{
	UserId int64
	Msg NewMessageToUser}


func GetModels() string{
	return "Info"
}
