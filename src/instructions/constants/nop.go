package constants

import "awesomeGvm/src/instructions/base"
import "awesomeGvm/src/rtda"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
