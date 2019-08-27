package command
const (
	SC_TopLayerName = "command.SC_TopLayer"
	CS_TopLayerName = "command.CS_TopLayer"
	AccountInfoName = "command.AccountInfo"
	CS_AccountLoginName = "command.CS_AccountLogin"
	SC_ResAccountLoginName = "command.SC_ResAccountLogin"
	CS_AccountRegisterName = "command.CS_AccountRegister"
	SC_ResAccountRigsterName = "command.SC_ResAccountRigster"
	CS_EnterGameName = "command.CS_EnterGame"
	SC_ResEnterGameName = "command.SC_ResEnterGame"
)
func (this *SC_TopLayer) GetMsgName() string {
	return SC_TopLayerName
}
func (this *CS_TopLayer) GetMsgName() string {
	return CS_TopLayerName
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
func (this *SC_ResEnterGame) GetMsgName() string {
	return SC_ResEnterGameName
}
