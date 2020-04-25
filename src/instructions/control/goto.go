package control

import "awesomeGvm/src/instructions/base"
import "awesomeGvm/src/rtda"

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
