package control

import "jvm-book/ch07/instructions/base"
import "jvm-book/ch07/rtda"

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
