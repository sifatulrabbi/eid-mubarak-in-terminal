package canvas

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

type Canvas struct {
	Board [20][100]string
	Stop  bool
	Done  chan bool
}

func NewCanvas() *Canvas {
	c := Canvas{}
	c.Stop = false
	c.Done = make(chan bool)
	for i := 0; i < 20; i++ {
		for j := 0; j < 100; j++ {
			c.Board[i][j] = " "
		}
	}
	return &c
}

func (c *Canvas) ApplyPaint(addr [2]int, s string) {
	c.Board[addr[0]][addr[1]] = s
}

func (c *Canvas) RenderLoop() {
	for {
		if c.Stop {
			c.Done <- true
			break
		}

		out := ""
		for _, l := range c.Board {
			for _, s := range l {
				if s != "" {
					out += s
					out += " "
				}
			}
			out += "\n"
		}

		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			log.Fatalln(err)
		}
		fmt.Println(out)

		time.Sleep(time.Millisecond * 1000 / 60)
	}
}
