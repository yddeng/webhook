package robot

import (
	"fmt"
	"github.com/yddeng/webhook/conf"
)

var (
	makers map[string]RobotMaker
	robots map[string]*Robot
)

type RobotMaker interface {
	Type() string
	Make(url string) RobotI
}

type RobotI interface {
	SendToClient(args ...interface{})
}

func RegisterMaker(maker RobotMaker) {
	makers[maker.Type()] = maker
	fmt.Println("new maker", maker.Type())
}

type Robot struct {
	homepage string
	commands []string
	this     RobotI
}

func InitRobots() {
	config := conf.GetConfig()

	for _, r := range config.Robot {
		m, ok := makers[r.RobotType]
		if !ok {
			fmt.Println("no maker", r.RobotType)
			continue
		}

		robot := &Robot{
			homepage: r.Homepage,
			commands: r.NotifyCmd,
			this:     m.Make(r.RobotUrl),
		}

		robots[robot.homepage] = robot
		fmt.Println("new robot", robot)
	}

}
