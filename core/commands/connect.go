package cmd

import (
	proxy "bushuray-core/lib/proxy/mainproxy"
	"bushuray-core/structs"
	"log"
)

func (cmd *Cmd) Disconnect(data structs.DisconnectData, proxy_manager *proxy.ProxyManager) {
	ConnectionMutex.Lock()
	defer ConnectionMutex.Unlock()

	proxy_manager.Stop()
}

func (cmd *Cmd) Connect(data structs.ConnectData, proxy_manager *proxy.ProxyManager) {
	ConnectionMutex.Lock()
	defer ConnectionMutex.Unlock()

	profile, err := cmd.DB.GetProfile(data.Profile.GroupId, data.Profile.Id)
	if err != nil {
		log.Println(err.Error())
		cmd.warn("connect-failed", "Failed to connect")
		return
	}

	if err := proxy_manager.Connect(profile); err != nil {
		log.Println(err.Error())
		cmd.warn("connect-failed", "Failed to connect")
		return
	}
}
