package webhook

import (
	"github.com/aiteung/atapi"
	"github.com/aiteung/atmessage"
	"github.com/aiteung/module/model"
	"github.com/whatsauth/wa"
)

func HandlerIncomingMessage(msg model.IteungMessage) (resp atmessage.Response) {
	dt := &wa.TextMessage{
		To:       msg.Chat_number,
		IsGroup:  false,
		Messages: "Konnichiwa mazzeh" + msg.Alias_name,
	}
	if msg.Chat_server == "g.us" { //jika pesan datang dari group maka balas ke group
		dt.IsGroup = true
	}
	if msg.Phone_number != "628112000279" { //ignore pesan datang dari iteung
		resp, _ = atapi.PostStructWithToken[atmessage.Response]("Token", WAAPIToken, dt, "https://api.wa.my.id/api/send/message/text")
	}
	return
}
