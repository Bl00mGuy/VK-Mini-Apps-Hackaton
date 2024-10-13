package factory

import "github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/controller"

func CreatePingHandler() *controller.PingHandler {
	return controller.NewPingHandler()
}
