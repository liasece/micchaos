package command
import (
	"encoding/binary"
	"math"
	"encoding/json"
)
const (
	CS_LoginID = 36
	SC_ResLoginID = 37
	UserInfoID = 38
)
const (
	CS_LoginName = "command.CS_Login"
	SC_ResLoginName = "command.SC_ResLogin"
	UserInfoName = "command.UserInfo"
)
func (this *CS_Login) WriteBinary(data []byte) int {
	return WriteMsgCS_LoginByObj(data,this)
}
func (this *SC_ResLogin) WriteBinary(data []byte) int {
	return WriteMsgSC_ResLoginByObj(data,this)
}
func (this *UserInfo) WriteBinary(data []byte) int {
	return WriteMsgUserInfoByObj(data,this)
}
func (this *CS_Login) ReadBinary(data []byte) int {
	size,_ := ReadMsgCS_LoginByBytes(data, this)
	return size
}
func (this *SC_ResLogin) ReadBinary(data []byte) int {
	size,_ := ReadMsgSC_ResLoginByBytes(data, this)
	return size
}
func (this *UserInfo) ReadBinary(data []byte) int {
	size,_ := ReadMsgUserInfoByBytes(data, this)
	return size
}
func MsgIdToString(id uint16) string {
	switch(id ) {
		case CS_LoginID: 
		return CS_LoginName
		case SC_ResLoginID: 
		return SC_ResLoginName
		case UserInfoID: 
		return UserInfoName
		default:
		return ""
	}
}
func StringToMsgId(msgname string) uint16 {
	switch(msgname ) {
		case CS_LoginName: 
		return CS_LoginID
		case SC_ResLoginName: 
		return SC_ResLoginID
		case UserInfoName: 
		return UserInfoID
		default:
		return 0
	}
}
func MsgIdToType(id uint16) rune {
	switch(id ) {
		case CS_LoginID: 
		return rune('C')
		case SC_ResLoginID: 
		return rune('S')
		case UserInfoID: 
		return rune('U')
		default:
		return rune(0)
	}
}
func (this *CS_Login) GetMsgId() uint16 {
	return CS_LoginID
}
func (this *SC_ResLogin) GetMsgId() uint16 {
	return SC_ResLoginID
}
func (this *UserInfo) GetMsgId() uint16 {
	return UserInfoID
}
func (this *CS_Login) GetMsgName() string {
	return CS_LoginName
}
func (this *SC_ResLogin) GetMsgName() string {
	return SC_ResLoginName
}
func (this *UserInfo) GetMsgName() string {
	return UserInfoName
}
func (this *CS_Login) GetSize() int {
	return GetSizeCS_Login(this)
}
func (this *SC_ResLogin) GetSize() int {
	return GetSizeSC_ResLogin(this)
}
func (this *UserInfo) GetSize() int {
	return GetSizeUserInfo(this)
}
func (this *CS_Login) GetJson() string {
	json,_ := json.Marshal(this)
	return string(json)
}
func (this *SC_ResLogin) GetJson() string {
	json,_ := json.Marshal(this)
	return string(json)
}
func (this *UserInfo) GetJson() string {
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
func ReadMsgCS_LoginByBytes(indata []byte, obj *CS_Login) (int,*CS_Login ) {
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
		obj=&CS_Login{}
	}
	if offset + objsize > len(indata ) {
		return offset,obj
	}
	endpos := offset+objsize
	data := indata[offset:offset+objsize]
	offset = 0
	data__len := len(data)
	if offset + 4 + len(obj.Account) > data__len{
		return endpos,obj
	}
	obj.Account = readBinaryString(data[offset:])
	offset += 4 + len(obj.Account)
	return endpos,obj
}
func WriteMsgCS_LoginByObj(data []byte, obj *CS_Login) int {
	if obj == nil {
		binary.BigEndian.PutUint32(data[0:4],0)
		return 4
	}
	objsize := obj.GetSize() - 4
	offset := 0
	binary.BigEndian.PutUint32(data[offset:offset+4],uint32(objsize))
	offset += 4
	writeBinaryString(data[offset:],obj.Account)
	offset += 4 + len(obj.Account)
	return offset
}
func GetSizeCS_Login(obj *CS_Login) int {
	if obj == nil {
		return 4
	}
	return 4 + 4 + len(obj.Account)
}
func ReadMsgSC_ResLoginByBytes(indata []byte, obj *SC_ResLogin) (int,*SC_ResLogin ) {
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
		obj=&SC_ResLogin{}
	}
	if offset + objsize > len(indata ) {
		return offset,obj
	}
	endpos := offset+objsize
	data := indata[offset:offset+objsize]
	offset = 0
	data__len := len(data)
	if offset + 4 > data__len{
		return endpos,obj
	}
	obj.Code = readBinaryInt32(data[offset:offset+4])
	offset+=4
	if offset + obj.UserInfo.GetSize() > data__len{
		return endpos,obj
	}
	rsize_UserInfo := 0
	rsize_UserInfo,obj.UserInfo = ReadMsgUserInfoByBytes(data[offset:], nil)
	offset += rsize_UserInfo
	return endpos,obj
}
func WriteMsgSC_ResLoginByObj(data []byte, obj *SC_ResLogin) int {
	if obj == nil {
		binary.BigEndian.PutUint32(data[0:4],0)
		return 4
	}
	objsize := obj.GetSize() - 4
	offset := 0
	binary.BigEndian.PutUint32(data[offset:offset+4],uint32(objsize))
	offset += 4
	writeBinaryInt32(data[offset:offset+4], obj.Code)
	offset+=4
	offset += WriteMsgUserInfoByObj(data[offset:], obj.UserInfo)
	return offset
}
func GetSizeSC_ResLogin(obj *SC_ResLogin) int {
	if obj == nil {
		return 4
	}
	return 4 + 4 + obj.UserInfo.GetSize()
}
func ReadMsgUserInfoByBytes(indata []byte, obj *UserInfo) (int,*UserInfo ) {
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
		obj=&UserInfo{}
	}
	if offset + objsize > len(indata ) {
		return offset,obj
	}
	endpos := offset+objsize
	data := indata[offset:offset+objsize]
	offset = 0
	data__len := len(data)
	if offset + 4 + len(obj.UUID) > data__len{
		return endpos,obj
	}
	obj.UUID = readBinaryString(data[offset:])
	offset += 4 + len(obj.UUID)
	if offset + 4 + len(obj.Name) > data__len{
		return endpos,obj
	}
	obj.Name = readBinaryString(data[offset:])
	offset += 4 + len(obj.Name)
	return endpos,obj
}
func WriteMsgUserInfoByObj(data []byte, obj *UserInfo) int {
	if obj == nil {
		binary.BigEndian.PutUint32(data[0:4],0)
		return 4
	}
	objsize := obj.GetSize() - 4
	offset := 0
	binary.BigEndian.PutUint32(data[offset:offset+4],uint32(objsize))
	offset += 4
	writeBinaryString(data[offset:],obj.UUID)
	offset += 4 + len(obj.UUID)
	writeBinaryString(data[offset:],obj.Name)
	offset += 4 + len(obj.Name)
	return offset
}
func GetSizeUserInfo(obj *UserInfo) int {
	if obj == nil {
		return 4
	}
	return 4 + 4 + len(obj.UUID) + 4 + len(obj.Name)
}
