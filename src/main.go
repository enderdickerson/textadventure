package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type choices struct {
	cmd         string
	description string
	node        *storyNode
	nextChoice  *choices
}

type storyNode struct {
	text    string
	choices *choices
}

func (sn *storyNode) addChoice(c string, d string, nn *storyNode) {
	choice := &choices{cmd: c, description: d, node: nn, nextChoice: nil}

	if sn.choices == nil {
		sn.choices = choice
	} else {
		cc := sn.choices
		for cc.nextChoice != nil {
			cc = cc.nextChoice
		}
		cc.nextChoice = choice
	}
}

func (sn *storyNode) render() {
	fmt.Println(sn.text)
	cc := sn.choices
	for cc != nil {
		fmt.Println(cc.cmd, ":", cc.description)
		cc = cc.nextChoice
	}
}

func (sn *storyNode) executeCmd(cmd string) *storyNode {
	cc := sn.choices
	for cc != nil {
		if strings.ToLower(cc.cmd) == strings.ToLower(cmd) {
			return cc.node
		}
		cc = cc.nextChoice
	}
	fmt.Println("Sorry, I didn't understand that.")
	return sn
}

func (sn *storyNode) play(s *bufio.Scanner) {
	sn.render()
	if sn.choices != nil {
		s.Scan()
		sn.executeCmd(s.Text()).play(s)
	}
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	start := storyNode{text: `
		You are in a large chamber, deep underground.
		You see three passages leading out. A north passage leads into darkness.
		To the south, a passage appears to head upward. The eastern passage appears
		flat and well traveled.
	`}

	darkRoom := storyNode{text: `
		It is pitch black. You cannot see a thing.
	`}

	darkRoomLit := storyNode{text: `
		The dark passage is now lit by your latern. You can continue north or head back south.
	`}

	grue := storyNode{text: `
		While stumbling around in the darkness, you are eaten by a grue.
	`}

	trap := storyNode{text: `
		You head down the well traveled path when suddenly a trap door opens and you fall into a pit.
	`}

	treasure := storyNode{text: `
		You arrive at a small chamber, filled with treasure!
	`}

	start.addChoice("N", "Go North", &darkRoom)
	start.addChoice("S", "Go South", &darkRoom)
	start.addChoice("E", "Go East", &trap)

	darkRoom.addChoice("S", "Try to go back south", &grue)
	darkRoom.addChoice("O", "Turn on latern", &darkRoomLit)

	darkRoomLit.addChoice("N", "Go North", &treasure)
	darkRoomLit.addChoice("S", "Go South", &start)

	start.play(s)
	fmt.Println()
	fmt.Println("The End.")
}
