package command
import (
	"encoding/binary"
	"math"
	"encoding/json"
)
const (
	CS_LoginID = 36
	SC_ResLoginID = 37
	AccountInfoID = 38
	CS_RegisterID = 39
	SC_ResRigsterID = 40
)
const (
	CS_LoginName = "command.CS_Login"
	SC_ResLoginName = "command.SC_ResLogin"
	AccountInfoName = "command.AccountInfo"
	CS_RegisterName = "command.CS_Register"
	SC_ResRigsterName = "command.SC_ResRigster"
)
func (this *CS_Login) WriteBinary(data []byte) int {
	return WriteMsgCS_LoginByObj(data,this)
}
func (this *SC_ResLogin) WriteBinary(data []byte) int {
	return WriteMsgSC_ResLoginByObj(data,this)
}
func (this *AccountInfo) WriteBinary(data []byte) int {
	return WriteMsgAccountInfoByObj(data,this)
}
func (this *CS_Register) WriteBinary(data []byte) int {
	return WriteMsgCS_RegisterByObj(data,this)
}
func (this *SC_ResRigster) WriteBinary(data []byte) int {
	return WriteMsgSC_ResRigsterByObj(data,this)
}
func (this *CS_Login) ReadBinary(data []byte) int {
	size,_ := ReadMsgCS_LoginByBytes(data, this)
	return size
}
func (this *SC_ResLogin) ReadBinary(data []byte) int {
	size,_ := ReadMsgSC_ResLoginByBytes(data, this)
	return size
}
func (this *AccountInfo) ReadBinary(data []byte) int {
	size,_ := ReadMsgAccountInfoByBytes(data, this)
	return size
}
func (this *CS_Register) ReadBinary(data []byte) int {
	size,_ := ReadMsgCS_RegisterByBytes(data, this)
	return size
}
func (this *SC_ResRigster) ReadBinary(data []byte) int {
	size,_ := ReadMsgSC_ResRigsterByBytes(data, this)
	return size
}
func MsgIdToString(id uint16) string {
	switch(id ) {
		case CS_LoginID: 
		return CS_LoginName
		case SC_ResLoginID: 
		return SC_ResLoginName
		case AccountInfoID: 
		return AccountInfoName
		case CS_RegisterID: 
		return CS_RegisterName
		case SC_ResRigsterID: 
		return SC_ResRigsterName
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
		case AccountInfoName: 
		return AccountInfoID
		case CS_RegisterName: 
		return CS_RegisterID
		case SC_ResRigsterName: 
		return SC_ResRigsterID
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
		case AccountInfoID: 
		return rune('A')
		case CS_RegisterID: 
		return rune('C')
		case SC_ResRigsterID: 
		return rune('S')
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
func (this *AccountInfo) GetMsgId() uint16 {
	return AccountInfoID
}
func (this *CS_Register) GetMsgId() uint16 {
	return CS_RegisterID
}
func (this *SC_ResRigster) GetMsgId() uint16 {
	return SC_ResRigsterID
}
func (this *CS_Login) GetMsgName() string {
	return CS_LoginName
}
func (this *SC_ResLogin) GetMsgName() string {
	return SC_ResLoginName
}
func (this *AccountInfo) GetMsgName() string {
	return AccountInfoName
}
func (this *CS_Register) GetMsgName() string {
	return CS_RegisterName
}
func (this *SC_ResRigster) GetMsgName() string {
	return SC_ResRigsterName
}
func (this *CS_Login) GetSize() int {
	return GetSizeCS_Login(this)
}
func (this *SC_ResLogin) GetSize() int {
	return GetSizeSC_ResLogin(this)
}
func (this *AccountInfo) GetSize() int {
	return GetSizeAccountInfo(this)
}
func (this *CS_Register) GetSize() int {
	return GetSizeCS_Register(this)
}
func (this *SC_ResRigster) GetSize() int {
	return GetSizeSC_ResRigster(this)
}
func (this *CS_Login) GetJson() string {
	json,_ := json.Marshal(this)
	return string(json)
}
func (this *SC_ResLogin) GetJson() string {
	json,_ := json.Marshal(this)
	return string(json)
}
func (this *AccountInfo) GetJson() string {
	json,_ := json.Marshal(this)
	return string(json)
}
func (this *CS_Register) GetJson() string {
	json,_ := json.Marshal(this)
	return string(json)
}
func (this *SC_ResRigster) GetJson() string {
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
	if offset + 4 + len(obj.LoginName) > data__len{
		return endpos,obj
	}
	obj.LoginName = readBinaryString(data[offset:])
	offset += 4 + len(obj.LoginName)
	if offset + 4 + len(obj.PassWordMD5) > data__len{
		return endpos,obj
	}
	obj.PassWordMD5 = readBinaryString(data[offset:])
	offset += 4 + len(obj.PassWordMD5)
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
	writeBinaryString(data[offset:],obj.LoginName)
	offset += 4 + len(obj.LoginName)
	writeBinaryString(data[offset:],obj.PassWordMD5)
	offset += 4 + len(obj.PassWordMD5)
	return offset
}
func GetSizeCS_Login(obj *CS_Login) int {
	if obj == nil {
		return 4
	}
	return 4 + 4 + len(obj.LoginName) + 4 + len(obj.PassWordMD5)
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
	if offset + 4 + len(obj.Message) > data__len{
		return endpos,obj
	}
	obj.Message = readBinaryString(data[offset:])
	offset += 4 + len(obj.Message)
	if offset + 4 + len(obj.ConnectID) > data__len{
		return endpos,obj
	}
	obj.ConnectID = readBinaryString(data[offset:])
	offset += 4 + len(obj.ConnectID)
	if offset + obj.Account.GetSize() > data__len{
		return endpos,obj
	}
	rsize_Account := 0
	rsize_Account,obj.Account = ReadMsgAccountInfoByBytes(data[offset:], nil)
	offset += rsize_Account
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
	writeBinaryString(data[offset:],obj.Message)
	offset += 4 + len(obj.Message)
	writeBinaryString(data[offset:],obj.ConnectID)
	offset += 4 + len(obj.ConnectID)
	offset += WriteMsgAccountInfoByObj(data[offset:], obj.Account)
	return offset
}
func GetSizeSC_ResLogin(obj *SC_ResLogin) int {
	if obj == nil {
		return 4
	}
	return 4 + 4 + 4 + len(obj.Message) + 4 + len(obj.ConnectID) + obj.Account.GetSize()
}
func ReadMsgAccountInfoByBytes(indata []byte, obj *AccountInfo) (int,*AccountInfo ) {
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
		obj=&AccountInfo{}
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
	if offset + 4 + len(obj.LoginName) > data__len{
		return endpos,obj
	}
	obj.LoginName = readBinaryString(data[offset:])
	offset += 4 + len(obj.LoginName)
	return endpos,obj
}
func WriteMsgAccountInfoByObj(data []byte, obj *AccountInfo) int {
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
	writeBinaryString(data[offset:],obj.LoginName)
	offset += 4 + len(obj.LoginName)
	return offset
}
func GetSizeAccountInfo(obj *AccountInfo) int {
	if obj == nil {
		return 4
	}
	return 4 + 4 + len(obj.UUID) + 4 + len(obj.LoginName)
}
func ReadMsgCS_RegisterByBytes(indata []byte, obj *CS_Register) (int,*CS_Register ) {
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
		obj=&CS_Register{}
	}
	if offset + objsize > len(indata ) {
		return offset,obj
	}
	endpos := offset+objsize
	data := indata[offset:offset+objsize]
	offset = 0
	data__len := len(data)
	if offset + 4 + len(obj.LoginName) > data__len{
		return endpos,obj
	}
	obj.LoginName = readBinaryString(data[offset:])
	offset += 4 + len(obj.LoginName)
	if offset + 4 + len(obj.PassWordMD5) > data__len{
		return endpos,obj
	}
	obj.PassWordMD5 = readBinaryString(data[offset:])
	offset += 4 + len(obj.PassWordMD5)
	return endpos,obj
}
func WriteMsgCS_RegisterByObj(data []byte, obj *CS_Register) int {
	if obj == nil {
		binary.BigEndian.PutUint32(data[0:4],0)
		return 4
	}
	objsize := obj.GetSize() - 4
	offset := 0
	binary.BigEndian.PutUint32(data[offset:offset+4],uint32(objsize))
	offset += 4
	writeBinaryString(data[offset:],obj.LoginName)
	offset += 4 + len(obj.LoginName)
	writeBinaryString(data[offset:],obj.PassWordMD5)
	offset += 4 + len(obj.PassWordMD5)
	return offset
}
func GetSizeCS_Register(obj *CS_Register) int {
	if obj == nil {
		return 4
	}
	return 4 + 4 + len(obj.LoginName) + 4 + len(obj.PassWordMD5)
}
func ReadMsgSC_ResRigsterByBytes(indata []byte, obj *SC_ResRigster) (int,*SC_ResRigster ) {
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
		obj=&SC_ResRigster{}
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
	if offset + 4 + len(obj.Message) > data__len{
		return endpos,obj
	}
	obj.Message = readBinaryString(data[offset:])
	offset += 4 + len(obj.Message)
	if offset + 4 + len(obj.ConnectID) > data__len{
		return endpos,obj
	}
	obj.ConnectID = readBinaryString(data[offset:])
	offset += 4 + len(obj.ConnectID)
	if offset + obj.Account.GetSize() > data__len{
		return endpos,obj
	}
	rsize_Account := 0
	rsize_Account,obj.Account = ReadMsgAccountInfoByBytes(data[offset:], nil)
	offset += rsize_Account
	return endpos,obj
}
func WriteMsgSC_ResRigsterByObj(data []byte, obj *SC_ResRigster) int {
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
	writeBinaryString(data[offset:],obj.Message)
	offset += 4 + len(obj.Message)
	writeBinaryString(data[offset:],obj.ConnectID)
	offset += 4 + len(obj.ConnectID)
	offset += WriteMsgAccountInfoByObj(data[offset:], obj.Account)
	return offset
}
func GetSizeSC_ResRigster(obj *SC_ResRigster) int {
	if obj == nil {
		return 4
	}
	return 4 + 4 + 4 + len(obj.Message) + 4 + len(obj.ConnectID) + obj.Account.GetSize()
}
