package models

type MessageCommand int

const (
	MessageCommandUserInsertedToChat MessageCommand = iota + 1
	MessageCommandUserCreatedChat
	MessageCommandUserInsertedToChannel
	MessageCommandUserCreatedChannel
	MessageCommandUserInsertedToDialog
	MessageCommandUserCreatedDialog
)

const (
	MessageActionOnlineUser       = "online_sser"
	MessageActionUserChatInserted = "sser_inserted"
	MessageActionChatCreated      = "chat_created"
	MessageActionTypeSystem       = "system"
)
