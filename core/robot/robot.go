package robot

import (
	"fmt"
	"github.com/yddeng/dutil/queue"
	"github.com/yddeng/webhook/conf"
)

var (
	makers     = map[string]RobotMaker{}
	robots     = map[string]*Robot{}
	eventQueue = queue.NewEventQueue(128, pcall)
)

type Event struct {
	Homepage string
	Cmd      string
	Args     []string
}

func PushEvent(e *Event) {
	_ = eventQueue.Push(e)
}

func pcall(i interface{}) {
	e := i.(*Event)

	fmt.Println("pcall", e)
	r, ok := robots[e.Homepage]
	if !ok {
		fmt.Println("no robot", e.Homepage)
		return
	}

	if r.checkCmd(e.Cmd) {
		r.instance.SendToClient(e.Cmd, e.Args...)
	}

}

type RobotMaker interface {
	Type() string
	Make(url string) RobotI
}

type RobotI interface {
	SendToClient(cmd string, args ...string)
}

func RegisterMaker(maker RobotMaker) {
	makers[maker.Type()] = maker
	fmt.Println("new maker", maker.Type())
}

type Robot struct {
	homepage string
	commands []string
	instance RobotI
}

func (this *Robot) checkCmd(cmd string) bool {
	for _, c := range this.commands {
		if c == cmd {
			return true
		}
	}
	return false
}

func InitRobots() {
	config := conf.GetConfig()

	for _, r := range config.Robot {
		m, ok := makers[r.RobotType]
		if !ok {
			fmt.Println("no maker", r.RobotType)
			continue
		}

		if r.RobotUrl == "" || len(r.NotifyCmd) == 0 {
			fmt.Println("failed url or notifyCmd")
			continue
		}

		robot := &Robot{
			homepage: r.Homepage,
			commands: r.NotifyCmd,
			instance: m.Make(r.RobotUrl),
		}

		robots[robot.homepage] = robot
		fmt.Println("new robot", robot)
	}

	eventQueue.Run(1)

}
