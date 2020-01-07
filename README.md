开始
--------------

`chaos/main.go`是程序的入口。

你需要先初始化一个`App`作为`Moudle`的容器：

`app, err := micserver.SetupApp(GetInitManger().GetConfigPath())`

初始化管理器`GetInitManger()`需要做的事情你可以直接参照`chaos/init.go`的实现。

然后，只需要将实现了`module.IModule`接口的模块传入App运行入口即可：

`app.RunAndBlock(GetInitManger().GetProgramModuleList())`

App会先初始化各个模块的底层连接，然后调用模块的`AfterInitModule()`方法，你可以在这里面执行你业务的初始化操作。

配置文件
--------------

配置文件是必须的，最简单的配置文件如`config/config.json`所示。由于JSON没有注释语法，以`//`为名的键视为注释。

网关模块(Gateway)
--------------

作为服务器，需要向客户端提供一个入口，本例中的网关模块位于`gatemodule`下。你可以在这里找到一个最简单的网关模块实现，以及如何劫持底层协议实现使用`websocket`与客户端通信。

会话(Session)
--------------

客户端在连接到服务器网络后，除了`Gate`能取到客户端的实际连接`Client`外，其他模块只能通过客户端的`Session`操作客户端，你可以在`loginmodule/msghandler_client.go`中找到`Session`最基础的应用，可以在`Session`中保存一些重要且简单的数据，或者通过`Session.SendMsg()`向客户端连接发送数据。如`playermodule/boxes/player.go`中的实现：

    func (this *Player) SendMsg(msg interface{}) {
        btop := ccmd.GetSCTopLayer(msg)
        // 使用当前Module作为发送者，向该Session绑定的gate类型模块发送数据，由管理该连接的Gate向真正的客户端连接发送数据。
        this.Session.SendMsg(this.mod, "gate", 0, btop)
    }

ROC
--------------

Remote Object Call.

ROC ，是 micserver 重要的分布式目标调用思想。如果一个对象，例如房间/商品/玩家/工会，需要提供一个可供远程执行的操作，这在以前称之为 RPC 调用，在 micserver 中，任意一个构造了这种对象的模块，均可以通过 BaseModule.GetROC(objtype).GetOrRegObj(obj) 来
注册一个 ROC 对象，在其他模块中，只需要通过 BaseModule.ROCCallNR 等方法，提供
目标对象的类型及ID，即可发起针对该对象的远程操作。因此，在任意模块中，发起的任意针对其他模块的调用，都不应该使用模块ID来操作，因为使用统一的 ROC 至少包含以下好处：
* 无需知道目标对象在哪个模块上；
* 只需要关心目标对象的ID（目标的类型你当然是知道的）；
* 在模块更新时，可以统一将该模块的 ROC 对象迁移到新版本模块中实现热更；
* 可以将对象存储到数据库并且在其他模块中加载（基于第一点好处）；
* 对象的位置/调用路由等由底层系统维护，可提供一个高可用/强一致/易维护的分布式网络。

你需要先在`Module`中注册一种类型的ROC，例如`playermodule/module.go`中：

    func (this *PlayerModule) AfterInitModule() {
        this.BaseModule.AfterInitModule()
        // 初始化Player ROC
        this.BaseModule.NewROC(ccmd.ROCTypePlayer)
    }

然后就是在你需要增加一个`ccmd.ROCTypePlayer`类型的ROC对象的时候，注册这个类型的对象。你可以在`playermodule/manager/playerdoc_manager.go`中找到最简单的ROC对象注册方式：

`vi, isLoad := this.playerRoc.GetOrRegObj(uuid, playerobj)`

其中，`playerobj`必须实现`roc.IObj`接口，你可以在`playermodule/boxes/player.go`中找到这些实现：

    // ROC对象接口实现
    func (this *Player) GetROCObjID() string {
        return this.Account.UUID
    }
    // ROC对象接口实现
    func (this *Player) GetROCObjType() roc.ROCObjType {
        return ccmd.ROCTypePlayer
    }
    // ROC对象接口实现
    func (this *Player) OnROCCall(path *roc.ROCPath, arg []byte) ([]byte, error) {
        this.Info("ROC调用执行:[%s],%+v", path.String(), arg)
        switch path.Move() {
        case "GateClose":
            this.OnGateClose()
        }
        return nil, nil
    }

只要ROC对象完成了注册`GetOrRegObj()`，在其他任何连接到该模块的模块中，都可以通过以下方式发起调用到`playerobj.OnROCCall()`：

`module.ROCCallNR(roc.O(ccmd.ROCTypePlayer,uuid).F("Func"),data)`

你可以在`gatemodule/module.go`的`OnCloseClient()`中找到最简单的示例。
