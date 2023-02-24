package command

import (
	"mydocker/src/container"
	"os"

	log "github.com/sirupsen/logrus"
)

// run函数，执行具体run操作
func Run(tty bool, command string) {
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	parent.Wait()
	os.Exit(-1)

}
