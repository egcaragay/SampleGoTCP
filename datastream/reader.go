package datastream

import "math"

type DataStream struct {
	capacity int
	head int
	buffer []byte
}

func NewDataStreamReader(buffer []byte) *DataStream {
	length := len(buffer)
	r := DataStream{
		capacity: length,
		head: 0,
		buffer: buffer,
	}
	return &r
}
func (dsr *DataStream) getHeadBytes(byteCount int) []byte {
	resByte := dsr.buffer[dsr.head:dsr.head+byteCount]
	dsr.head += byteCount
	return resByte
}

func bytesToInt(num []byte) int32 {
	var result int32 = 0
	length := len(num)
	for i:=0;i<length;i++ {
		numL := int32(num[i])
		result = result | ((numL) << ((length-i-1) * 8))
	}
	return result
}

func bytesToFloat(bytes []byte) float32 {
	bit := bytesToInt(bytes)
	float := math.Float32frombits(uint32(bit))
	return float
}

func (dsr *DataStream) ReadFloat() float32 {
	resByte := dsr.getHeadBytes(4)
	return bytesToFloat(resByte)

}

func (dsr *DataStream) ReadInt() int {
	resBytes := dsr.getHeadBytes(4)
	return int(bytesToInt(resBytes))
}

func (dsr *DataStream) ReadByte() byte {
	resByte := dsr.buffer[dsr.head]
	dsr.head++
	return resByte
}

func (dsr *DataStream) ReadString() string {
	bytesCount := int(dsr.ReadByte())
	resArr := dsr.getHeadBytes(bytesCount)
	return string(resArr)
}