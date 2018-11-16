实用工具包
1.Config 类
	一般程序启动时加载配置文件：Cfg:=Tools.LoadConfig(path),
	使用Tools.Cfgs定义Cfg变量,一般情况把Cfg设为全局变量，或者传入需要使用配置数据的函数
	运行过程中使用:Cfg.SetInt(section,key),Cfg.GetFloat(section,key),
	Cfg.GetBool(section,key),Cfg.GetValue(section,key)可以获取需要的数据
	运行过程中使用Set类方法可以改变配置数据
	运行结束后可以使用Cfg.SaveConfig()保存配置文件,
	SaveConfigAs(path string),配置文件另存到指定路径