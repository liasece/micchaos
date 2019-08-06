package command
import (
	"encoding/binary"
	"math"
	"encoding/json"
)
const (
	TestID = 36
)
const (
	TestName = "command.Test"
)
func (this *Test) WriteBinary(data []byte) int {
	return WriteMsgTestByObj(data,this)
}
func (this *Test) ReadBinary(data []byte) int {
	size,_ := ReadMsgTestByBytes(data, this)
	return size
}
func MsgIdToString(id uint16) string {
	switch(id ) {
		case TestID: 
		return TestName
		default:
		return ""
	}
}
func StringToMsgId(msgname string) uint16 {
	switch(msgname ) {
		case TestName: 
		return TestID
		default:
		return 0
	}
}
func MsgIdToType(id uint16) rune {
	switch(id ) {
		case TestID: 
		return rune('T')
		default:
		return rune(0)
	}
}
func (this *Test) GetMsgId() uint16 {
	return TestID
}
func (this *Test) GetMsgName() string {
	return TestName
}
func (this *Test) GetSize() int {
	return GetSizeTest(this)
}
func (this *Test) GetJson() string {
	json,_ := json.Marshal(this)
	return string(json)
}
func readBinaryString(data []byte) string {
	strfunclen := binary.BigEndian.Uint32(data[:4])
	if int(strfunclen) + 4 > len(data ) {
		return ""
	}
	return string(data[4:4+strfunclen])
}
func writeBinaryString(data []byte,obj string) int {
	objlen := len(obj)
	binary.BigEndian.PutUint32(data[:4],uint32(objlen))
	copy(data[4:4+objlen], obj)
	return 4+objlen
}
func bool2int(value bool) int {
	if value {
		return 1
	}
	return 0
}
func readBinaryInt64(data []byte) int64 {
	// 大端模式
	num := int64(0)
	num |= int64(data[7]) << 0
	num |= int64(data[6]) << 8
	num |= int64(data[5]) << 16
	num |= int64(data[4]) << 24
	num |= int64(data[3]) << 32
	num |= int64(data[2]) << 40
	num |= int64(data[1]) << 48
	num |= int64(data[0]) << 56
	return num
}
func writeBinaryInt64(data []byte, num int64 ) {
	// 大端模式
	data[7] = byte((num >> 0) & 0xff)
	data[6] = byte((num >> 8) & 0xff)
	data[5] = byte((num >> 16) & 0xff)
	data[4] = byte((num >> 24) & 0xff)
	data[3] = byte((num >> 32) & 0xff)
	data[2] = byte((num >> 40) & 0xff)
	data[1] = byte((num >> 48) & 0xff)
	data[0] = byte((num >> 56) & 0xff)
}
func readBinaryInt32(data []byte) int32 {
	// 大端模式
	num := int32(0)
	num |= int32(data[3]) << 0
	num |= int32(data[2]) << 8
	num |= int32(data[1]) << 16
	num |= int32(data[0]) << 24
	return num
}
func writeBinaryInt32(data []byte, num int32 ) {
	// 大端模式
	data[3] = byte((num >> 0) & 0xff)
	data[2] = byte((num >> 8) & 0xff)
	data[1] = byte((num >> 16) & 0xff)
	data[0] = byte((num >> 24) & 0xff)
}
func readBinaryInt(data []byte) int {
	return int(readBinaryInt32(data))
}
func writeBinaryInt(data []byte, num int ) {
	writeBinaryInt32(data,int32(num))
}
func readBinaryInt16(data []byte) int16 {
	// 大端模式
	num := int16(0)
	num |= int16(data[1]) << 0
	num |= int16(data[0]) << 8
	return num
}
func writeBinaryInt16(data []byte, num int16 ) {
	// 大端模式
	data[1] = byte((num >> 0) & 0xff)
	data[0] = byte((num >> 8) & 0xff)
}
func readBinaryInt8(data []byte) int8 {
	// 大端模式
	num := int8(0)
	num |= int8(data[0]) << 0
	return num
}
func writeBinaryInt8(data []byte, num int8 ) {
	// 大端模式
	data[0] = byte(num)
}
func readBinaryBool(data []byte) bool {
	// 大端模式
	num := int8(0)
	num |= int8(data[0]) << 0
	return num>0
}
func writeBinaryBool(data []byte, num bool ) {
	// 大端模式
	if num == true {
		data[0] = byte(1)
	} else {
		data[0] = byte(0)
	}
}
func readBinaryUint8(data []byte) uint8 {
	return uint8(data[0])
}
func writeBinaryUint8(data []byte, num uint8 ) {
	data[0] = byte(num)
}
func readBinaryUint(data []byte) uint {
	return uint(binary.BigEndian.Uint32(data))
}
func writeBinaryUint(data []byte, num uint ) {
	binary.BigEndian.PutUint32(data,uint32(num))
}
func writeBinaryFloat32(data []byte, num float32 ) {
	bits := math.Float32bits(num)
	binary.BigEndian.PutUint32(data,bits)
}
func readBinaryFloat32(data []byte) float32 {
	bits := binary.BigEndian.Uint32(data)
	return math.Float32frombits(bits)
}
func writeBinaryFloat64(data []byte, num float64 ) {
	bits := math.Float64bits(num)
	binary.BigEndian.PutUint64(data,bits)
}
func readBinaryFloat64(data []byte) float64 {
	bits := binary.BigEndian.Uint64(data)
	return math.Float64frombits(bits)
}
func ReadMsgTestByBytes(indata []byte, obj *Test) (int,*Test ) {
	offset := 0
	if len(indata) < 4 {
		return 0,nil
	}
	objsize := int(binary.BigEndian.Uint32(indata))
	offset += 4
	if objsize == 0 {
		return 4,nil
	}
	if obj == nil{
		obj=&Test{}
	}
	if offset + objsize > len(indata ) {
		return offset,obj
	}
	endpos := offset+objsize
	data := indata[offset:offset+objsize]
	offset = 0
	data__len := len(data)
	if offset + 8 > data__len{
		return endpos,obj
	}
	obj.Seq = readBinaryInt64(data[offset:offset+8])
	offset+=8
	if offset + 4 > data__len{
		return endpos,obj
	}
	Data_slen := int(binary.BigEndian.Uint32(data[offset:offset+4]))
	offset += 4
	if Data_slen != 0xffffffff {
		if offset + Data_slen > data__len {
			return endpos,obj
		}
		obj.Data = make([]byte,Data_slen)
		copy(obj.Data, data[offset:offset+Data_slen])
		offset += Data_slen
	}
	return endpos,obj
}
func WriteMsgTestByObj(data []byte, obj *Test) int {
	if obj == nil {
		binary.BigEndian.PutUint32(data[0:4],0)
		return 4
	}
	objsize := obj.GetSize() - 4
	offset := 0
	binary.BigEndian.PutUint32(data[offset:offset+4],uint32(objsize))
	offset += 4
	writeBinaryInt64(data[offset:offset+8], obj.Seq)
	offset+=8
	if obj.Data == nil {
		binary.BigEndian.PutUint32(data[offset:offset+4],0xffffffff)
	} else {
		binary.BigEndian.PutUint32(data[offset:offset+4],uint32(len(obj.Data)))
	}
	offset += 4
	Data_slen := len(obj.Data)
	copy(data[offset:offset+Data_slen], obj.Data)
	offset += Data_slen
	return offset
}
func GetSizeTest(obj *Test) int {
	if obj == nil {
		return 4
	}
	return 4 + 8 + 4 + len(obj.Data) * 1
}
