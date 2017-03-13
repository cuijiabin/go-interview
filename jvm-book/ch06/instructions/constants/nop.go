package constants

import "jvm-book/ch06/instructions/base"
import "jvm-book/ch06/rtda"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
