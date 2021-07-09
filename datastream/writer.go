package datastream

import (
	"encoding/binary"
	"math"
)

func NewDataStreamWriter() *DataStream {
	r := DataStream{head:0}
	return &r
}

func (dsr *DataStream) GetBuffer() []byte {
	return dsr.buffer
}

func (dsr *DataStream) WriteInt(data int) {
	byteCount := 4
	for i:=0;byteCount>0;i++{
		dsr.buffer = append(dsr.buffer,byte(data >> ((byteCount-1) *8)))
		byteCount--
	}
}

func (dsr *DataStream) WriteByte(data byte) {
	dsr.buffer = append(dsr.buffer,data)
}

func (dsr *DataStream) WriteString(data string) {
	bytes := []byte(data)
	byteLength := len(bytes)
	dsr.WriteByte(byte(byteLength))
	for i := 0 ;i <len(bytes);i++{
		dsr.WriteByte(bytes[i])
	}
}

func (dsr *DataStream) WriteFloat(data float32) {
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:],math.Float32bits(data))
	for i:=0;i<len(buf);i++ {
		dsr.buffer = append(dsr.buffer,buf[i])
	}
}