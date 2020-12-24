package servicebus

import (
	"context"
	"sync"
	"testing"

	"github.com/Azure/go-amqp"
)

type mockAMQPSender struct{}

func (*mockAMQPSender) send(ctx context.Context, msg *Message) (chan interface{}, error) {
	return nil, nil
}

var sessionID = "sessionId"

var namespace = Namespace{}
var message = Message{
	SessionID: &sessionID,
	ID:        "id",
}

func TestSender_Send(t *testing.T) {
	type fields struct {
		namespace          *Namespace
		client             *amqp.Client
		clientMu           sync.RWMutex
		session            *session
		sender             *amqp.Sender
		entityPath         string
		Name               string
		sessionID          *string
		doneRefreshingAuth func()
	}
	type args struct {
		ctx  context.Context
		msg  *Message
		opts []SendOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				namespace: &namespace,
				sender: &amqp.Sender{

				},
			},
			args: args{
				msg: &message,
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sender{
				namespace:          tt.fields.namespace,
				client:             tt.fields.client,
				clientMu:           tt.fields.clientMu,
				session:            tt.fields.session,
				sender:             tt.fields.sender,
				entityPath:         tt.fields.entityPath,
				Name:               tt.fields.Name,
				sessionID:          tt.fields.sessionID,
				doneRefreshingAuth: tt.fields.doneRefreshingAuth,
			}
			if err := s.Send(tt.args.ctx, tt.args.msg, tt.args.opts...); (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
