package classfile

import (
	"encoding/binary"
)

type ClassReader struct {
	data []byte
}

// u1
func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

// u2
func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

// u4 这是什么意思呢？func (bigEndian) Uint32(b []byte) uint32 {
//return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24 }
//终于理解了 这是为什么了，读取前几位，然后截取剩下的几位
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	//fmt.Printf("readUint32前: %v\n",len(self.data))
	self.data = self.data[4:]
	//fmt.Printf("readUint32后: %v\n",len(self.data))
	//fmt.Printf("readUint32值: %v\n",val)
	return val
}

func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

func (self *ClassReader) readBytes(n uint32) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}
