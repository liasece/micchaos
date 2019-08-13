package command
import (
	"encoding/binary"
	"math"
	"encoding/json"
)
const (
	AccountInfoID = 36
	CS_AccountLoginID = 37
	SC_ResAccountLoginID = 38
	CS_AccountRegisterID = 39
	SC_ResAccountRigsterID = 40
	CS_EnterGameID = 41
	SC_RedEnterGameID = 42
)
const (
	AccountInfoName = "command.AccountInfo"
	CS_AccountLoginName = "command.CS_AccountLogin"
	SC_ResAccountLoginName = "command.SC_ResAccountLogin"
	CS_AccountRegisterName = "command.CS_AccountRegister"
	SC_ResAccountRigsterName = "command.SC_ResAccountRigster"
	CS_EnterGameName = "command.CS_EnterGame"
	SC_RedEnterGameName = "command.SC_RedEnterGame"
)
func (this *AccountInfo) WriteBinary(data []byte) int {
	return WriteMsgAccountInfoByObj(data,this)
}
func (this *CS_AccountLogin) WriteBinary(data []byte) int {
	return WriteMsgCS_AccountLoginByObj(data,this)
}
func (this *SC_ResAccountLogin) WriteBinary(data []byte) int {
	return WriteMsgSC_ResAccountLoginByObj(data,this)
}
func (this *CS_AccountRegister) WriteBinary(data []byte) int {
	return WriteMsgCS_AccountRegisterByObj(data,this)
}
func (this *SC_ResAccountRigster) WriteBinary(data []byte) int {
	return WriteMsgSC_ResAccountRigsterByObj(data,this)
}
func (this *CS_EnterGame) WriteBinary(data []byte) int {
	return WriteMsgCS_EnterGameByObj(data,this)
}
func (this *SC_RedEnterGame) WriteBinary(data []byte) int {
	return WriteMsgSC_RedEnterGameByObj(data,this)
}
func (this *AccountInfo) ReadBinary(data []byte) int {
	size,_ := ReadMsgAccountInfoByBytes(data, this)
	return size
}
func (this *CS_AccountLogin) ReadBinary(data []byte) int {
	size,_ := ReadMsgCS_AccountLoginByBytes(data, this)
	return size
}
func (this *SC_ResAccountLogin) ReadBinary(data []byte) int {
	size,_ := ReadMsgSC_ResAccountLoginByBytes(data, this)
	return size
}
func (this *CS_AccountRegister) ReadBinary(data []byte) int {
	size,_ := ReadMsgCS_AccountRegisterByBytes(data, this)
	return size
}
func (this *SC_ResAccountRigster) ReadBinary(data []byte) int {
	size,_ := ReadMsgSC_ResAccountRigsterByBytes(data, this)
	return size
}
func (this *CS_EnterGame) ReadBinary(data []byte) int {
	size,_ := ReadMsgCS_EnterGameByBytes(data, this)
	return size
}
func (this *SC_RedEnterGame) ReadBinary(data []byte) int {
	size,_ := ReadMsgSC_RedEnterGameByBytes(data, this)
	return size
}
func MsgIdToString(id uint16) string {
	switch(id ) {
		case AccountInfoID: 
		return AccountInfoName
		case CS_AccountLoginID: 
		return CS_AccountLoginName
		case SC_ResAccountLoginID: 
		return SC_ResAccountLoginName
		case CS_AccountRegisterID: 
		return CS_AccountRegisterName
		case SC_ResAccountRigsterID: 
		return SC_ResAccountRigsterName
		case CS_EnterGameID: 
		return CS_EnterGameName
		case SC_RedEnterGameID: 
		return SC_RedEnterGameName
		default:
		return ""
	}
}
func StringToMsgId(msgname string) uint16 {
	switch(msgname ) {
		case AccountInfoName: 
		return AccountInfoID
		case CS_AccountLoginName: 
		return CS_AccountLoginID
		case SC_ResAccountLoginName: 
		return SC_ResAccountLoginID
		case CS_AccountRegisterName: 
		return CS_AccountRegisterID
		case SC_ResAccountRigsterName: 
		return SC_ResAccountRigsterID
		case CS_EnterGameName: 
		return CS_EnterGameID
		case SC_RedEnterGameName: 
		return SC_RedEnterGameID
		default:
		return 0
	}
}
func MsgIdToType(id uint16) rune {
	switch(id ) {
		case AccountInfoID: 
		return rune('A')
		case CS_AccountLoginID: 
		return rune('C')
		case SC_ResAccountLoginID: 
		return rune('S')
		case CS_AccountRegisterID: 
		return rune('C')
		case SC_ResAccountRigsterID: 
		return rune('S')
		case CS_EnterGameID: 
		return rune('C')
		case SC_RedEnterGameID: 
		return rune('S')
		default:
		return rune(0)
	}
}
func (this *AccountInfo) GetMsgId() uint16 {
	return AccountInfoID
}
func (this *CS_AccountLogin) GetMsgId() uint16 {
	return CS_AccountLoginID
}
func (this *SC_ResAccountLogin) GetMsgId() uint16 {
	return SC_ResAccountLoginID
}
func (this *CS_AccountRegister) GetMsgId() uint16 {
	return CS_AccountRegisterID
}
func (this *SC_ResAccountRigster) GetMsgId() uint16 {
	return SC_ResAccountRigsterID
}
func (this *CS_EnterGame) GetMsgId() uint16 {
	return CS_EnterGameID
}
func (this *SC_RedEnterGame) GetMsgId() uint16 {
	return SC_RedEnterGameID
}
func (this *AccountInfo) GetMsgName() string {
	return AccountInfoName
}
func (this *CS_AccountLogin) GetMsgName() string {
	return CS_AccountLoginName
}
func (this *SC_ResAccountLogin) GetMsgName() string {
	return SC_ResAccountLoginName
}
func (this *CS_AccountRegister) GetMsgName() string {
	return CS_AccountRegisterName
}
func (this *SC_ResAccountRigster) GetMsgName() string {
	return SC_ResAccountRigsterName
}
func (this *CS_EnterGame) GetMsgName() string {
	return CS_EnterGameName
}
func (this *SC_RedEnterGame) GetMsgName() string {
	return SC_RedEnterGameName
}
func (this *AccountInfo) GetSize() int {
	return GetSizeAccountInfo(this)
}
func (this *CS_AccountLogin) GetSize() int {
	return GetSizeCS_AccountLogin(this)
}
func (this *SC_ResAccountLogin) GetSize() int {
	return GetSizeSC_ResAccountLogin(this)
}
func (this *CS_AccountRegister) GetSize() int {
	return GetSizeCS_AccountRegister(this)
}
func (this *SC_ResAccountRigster) GetSize() int {
	return GetSizeSC_ResAccountRigster(this)
}
func (this *CS_EnterGame) GetSize() int {
	return GetSizeCS_EnterGame(this)
}
func (this *SC_RedEnterGame) GetSize() int {
	return GetSizeSC_RedEnterGame(this)
}
func (this *AccountInfo) GetJson() string {
	json,_ := json.Marshal(this)
	return string(json)
}
func (this *CS_AccountLogin) GetJson() string {
	json,_ := json.Marshal(this)
	return string(json)
}
func (this *SC_ResAccountLogin) GetJson() string {
	json,_ := json.Marshal(this)
	return string(json)
}
func (this *CS_AccountRegister) GetJson() string {
	json,_ := json.Marshal(this)
	return string(json)
}
func (this *SC_ResAccountRigster) GetJson() string {
	json,_ := json.Marshal(this)
	return string(json)
}
func (this *CS_EnterGame) GetJson() string {
	json,_ := json.Marshal(this)
	return string(json)
}
func (this *SC_RedEnterGame) GetJson() string {
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
func ReadMsgCS_AccountLoginByBytes(indata []byte, obj *CS_AccountLogin) (int,*CS_AccountLogin ) {
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
		obj=&CS_AccountLogin{}
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
func WriteMsgCS_AccountLoginByObj(data []byte, obj *CS_AccountLogin) int {
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
func GetSizeCS_AccountLogin(obj *CS_AccountLogin) int {
	if obj == nil {
		return 4
	}
	return 4 + 4 + len(obj.LoginName) + 4 + len(obj.PassWordMD5)
}
func ReadMsgSC_ResAccountLoginByBytes(indata []byte, obj *SC_ResAccountLogin) (int,*SC_ResAccountLogin ) {
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
		obj=&SC_ResAccountLogin{}
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
func WriteMsgSC_ResAccountLoginByObj(data []byte, obj *SC_ResAccountLogin) int {
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
func GetSizeSC_ResAccountLogin(obj *SC_ResAccountLogin) int {
	if obj == nil {
		return 4
	}
	return 4 + 4 + 4 + len(obj.Message) + 4 + len(obj.ConnectID) + obj.Account.GetSize()
}
func ReadMsgCS_AccountRegisterByBytes(indata []byte, obj *CS_AccountRegister) (int,*CS_AccountRegister ) {
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
		obj=&CS_AccountRegister{}
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
func WriteMsgCS_AccountRegisterByObj(data []byte, obj *CS_AccountRegister) int {
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
func GetSizeCS_AccountRegister(obj *CS_AccountRegister) int {
	if obj == nil {
		return 4
	}
	return 4 + 4 + len(obj.LoginName) + 4 + len(obj.PassWordMD5)
}
func ReadMsgSC_ResAccountRigsterByBytes(indata []byte, obj *SC_ResAccountRigster) (int,*SC_ResAccountRigster ) {
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
		obj=&SC_ResAccountRigster{}
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
func WriteMsgSC_ResAccountRigsterByObj(data []byte, obj *SC_ResAccountRigster) int {
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
func GetSizeSC_ResAccountRigster(obj *SC_ResAccountRigster) int {
	if obj == nil {
		return 4
	}
	return 4 + 4 + 4 + len(obj.Message) + 4 + len(obj.ConnectID) + obj.Account.GetSize()
}
func ReadMsgCS_EnterGameByBytes(indata []byte, obj *CS_EnterGame) (int,*CS_EnterGame ) {
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
		obj=&CS_EnterGame{}
	}
	if offset + objsize > len(indata ) {
		return offset,obj
	}
	endpos := offset+objsize
	return endpos,obj
}
func WriteMsgCS_EnterGameByObj(data []byte, obj *CS_EnterGame) int {
	if obj == nil {
		binary.BigEndian.PutUint32(data[0:4],0)
		return 4
	}
	objsize := obj.GetSize() - 4
	offset := 0
	binary.BigEndian.PutUint32(data[offset:offset+4],uint32(objsize))
	offset += 4
	return offset
}
func GetSizeCS_EnterGame(obj *CS_EnterGame) int {
	if obj == nil {
		return 4
	}
	return 4 + 0
}
func ReadMsgSC_RedEnterGameByBytes(indata []byte, obj *SC_RedEnterGame) (int,*SC_RedEnterGame ) {
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
		obj=&SC_RedEnterGame{}
	}
	if offset + objsize > len(indata ) {
		return offset,obj
	}
	endpos := offset+objsize
	return endpos,obj
}
func WriteMsgSC_RedEnterGameByObj(data []byte, obj *SC_RedEnterGame) int {
	if obj == nil {
		binary.BigEndian.PutUint32(data[0:4],0)
		return 4
	}
	objsize := obj.GetSize() - 4
	offset := 0
	binary.BigEndian.PutUint32(data[offset:offset+4],uint32(objsize))
	offset += 4
	return offset
}
func GetSizeSC_RedEnterGame(obj *SC_RedEnterGame) int {
	if obj == nil {
		return 4
	}
	return 4 + 0
}
