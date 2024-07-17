package messages

import (
	pb "Content-Service/genproto/content-service"
)

func (m *MessageRepo) CreateMessage(in *pb.MessageReq) (*pb.MessageRes, error) {
	info := pb.MessageRes{}
	err := m.DB.QueryRow("insert into messages (sender_id, recipient_id, content) values ($1, $2, $3) returning id, created_at",
		in.SenderId, in.RecipientId, in.Content).Scan(&info.Id, &info.CreatedAt)
	if err != nil {
		m.Logger.Error("error in create message", "err", err)
		return nil, err
	}
	info.SenderId = in.SenderId
	info.RecipientId = in.RecipientId
	info.Content = in.Content

	return &info, nil
}

func (m *MessageRepo) GetMessages(in *pb.GetMessage) (*pb.Messages, error) {
	rows, err := m.DB.Query("select sender_id, recipient_id, content, created_at from messages order by created_at desc")
	if err != nil {
		m.Logger.Error("error in get messages", "err", err)
		return nil, err
	}
	defer rows.Close()
	messages := []*pb.Message{}
	for rows.Next() {
		message := pb.Message{}
		err = rows.Scan(&message.Sender, &message.Recipient.Id, &message.Content, &message.CreatedAt)
		if err != nil {
			m.Logger.Error("error in scan", "err", err)
			return nil, err
		}
		messages = append(messages, &message)
	}
	return &pb.Messages{Messages: messages}, nil
}
