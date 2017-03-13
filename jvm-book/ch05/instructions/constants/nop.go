package constants

import "jvm-book/ch05/instructions/base"
import "jvm-book/ch05/rtda"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
