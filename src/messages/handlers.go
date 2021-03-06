package messageengine

import (
	"encoding/json"

	"github.com/alxarno/swap/logger"
	"github.com/alxarno/swap/src/api"
)

func authHandler(msg SystemMessage, connect *userConnection) {
	var ans = answer{}
	token := msg.Payload
	user, err := api.TestUserToken(token)
	if err != nil {
		ans = answer{MessageType: messageTypeSystem, Result: messageFailed,
			Action: messageActionAuth, Error: err.Error()}
	} else {
		connect.UserID = user.ID
		connect.Auth = true
		ans = answer{MessageType: messageTypeSystem, Result: messageSuccess, Action: messageActionAuth}
		userMove(connect.UserID, onlineUserInc)
		logger.Logger.Printf("WS user %d was authed from %s", connect.UserID, connect.RemoteAddr)
	}
	finish, _ := json.Marshal(ans)
	connect.SystemMessageChan <- string(finish)
}
