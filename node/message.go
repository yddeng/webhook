package node

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/dnet"
	"github.com/yddeng/webhook/codec"
	"github.com/yddeng/webhook/common"
	"github.com/yddeng/webhook/protocol"
	"os/exec"
)

func Command(path string) error {
	cmd := exec.Command("sh", "-c", path)

	//err := cmd.Run()
	//if err != nil {
	//	fmt.Println("Execute Command failed:" + err.Error())
	//}

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("execute shell:%s failed with error:%s \n\n", path, err.Error())
		return err
	}
	fmt.Printf("execute shell: %s finished with output:\n %s \n\n ", path, string(output))
	return nil
}

func onNotify(session dnet.Session, msg *codec.Message) {
	req := msg.GetData().(*protocol.Notify)

	for _, v := range myProxy.config.Hooks {
		if v.Homepage == req.GetHomepage() && v.Branch == req.GetBranch() {
			switch req.GetCmd() {
			case common.PushEvent:
			default:
				continue
			}
			err := Command(v.ShellPath)

			resp := &protocol.NotifyResp{
				Name:     proto.String(myProxy.config.Name),
				Homepage: proto.String(req.GetHomepage()),
				Message:  proto.String("脚本执行成功"),
			}

			if err != nil {
				resp.Message = proto.String("脚本执行失败")
			}
			err = myProxy.Send(resp)
		}
	}

}

func onLoginResp(session dnet.Session, msg *codec.Message) {
	req := msg.GetData().(*protocol.LoginResp)
	if !req.GetOk() {
		myProxy.CloseSignal <- req.GetMsg()
		return
	}
}

func onHeartbeat(session dnet.Session, msg *codec.Message) {}

func init() {
	dispatcher.Register(&protocol.Notify{}, onNotify)
	dispatcher.Register(&protocol.LoginResp{}, onLoginResp)
	dispatcher.Register(&protocol.Heartbeat{}, onHeartbeat)
}
