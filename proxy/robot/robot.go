package robot

import (
	"fmt"
)

var (
	makers = map[string]RobotMaker{}
	robots = map[string]*Robot{}
)

type RobotMaker interface {
	Type() string
	Make(url string) RobotI
}

type RobotI interface {
	SendToClient(cmd string, args ...string) error
}

func RegisterMaker(maker RobotMaker) {
	makers[maker.Type()] = maker
}

func GetRobot(homepage string) *Robot {
	return robots[homepage]
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

func (this *Robot) Notify(cmd string, args ...string) error {
	if this.checkCmd(cmd) {
		return this.instance.SendToClient(cmd, args...)
	}
	return nil
}

func MakeRobot(tt, homepage, url string, cmds []string) error {
	m, ok := makers[tt]
	if !ok {
		return fmt.Errorf("no maker:%s", tt)
	}

	if url == "" || len(cmds) == 0 {
		return fmt.Errorf("failed url or notifyCmd")
	}

	robot := &Robot{
		homepage: homepage,
		commands: cmds,
		instance: m.Make(url),
	}

	robots[robot.homepage] = robot

	return nil
}
