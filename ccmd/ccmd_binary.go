package ccmd

const (
	SC_TopLayerName = "ccmd.SC_TopLayer"
	CS_TopLayerName = "ccmd.CS_TopLayer"
	AccountInfoName = "ccmd.AccountInfo"
	CS_AccountLoginName = "ccmd.CS_AccountLogin"
	SC_ResAccountLoginName = "ccmd.SC_ResAccountLogin"
	CS_AccountRegisterName = "ccmd.CS_AccountRegister"
	SC_ResAccountRigsterName = "ccmd.SC_ResAccountRigster"
	CS_EnterGameName = "ccmd.CS_EnterGame"
	SC_ResEnterGameName = "ccmd.SC_ResEnterGame"
	SC_TipsName = "ccmd.SC_Tips"
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

func (this *SC_Tips) GetMsgName() string {
	return SC_TipsName
	            }

