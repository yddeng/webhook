package proxy

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/webhook/common"
	"github.com/yddeng/webhook/protocol"
	"github.com/yddeng/webhook/proxy/robot"
)

type Event struct {
	Homepage string
	Cmd      string
	Branch   string
	Args     []string
}

func doEvent(e *Event) {

	fmt.Println("doEvent:", *e)
	r := robot.GetRobot(e.Homepage)
	if r != nil {
		err := r.Notify(e.Cmd, e.Args...)
		if err != nil {
			fmt.Printf("notify homepage:%s err:%s\n", e.Homepage, err)
		} else {
			fmt.Printf("notify homepage:%s ok\n", e.Homepage)
		}
	}

	if tcpStarted && e.Cmd == common.PushEvent {
		clients, ok := homeNodes[e.Homepage]
		if ok {
			notify := &protocol.Notify{
				Cmd:      proto.String(e.Cmd),
				Homepage: proto.String(e.Homepage),
				Branch:   proto.String(e.Branch),
			}

			for _, c := range clients {
				err := c.Send(notify)
				if err != nil {
					doEvent(&Event{
						Homepage: e.Homepage,
						Cmd:      common.Message,
						Args:     []string{c.Name, err.Error()},
					})
				}
			}
		}
	}
}
