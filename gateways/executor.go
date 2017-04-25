package gateways

import "local/trader/models"

func ExecuteTrade(action *models.Action) {
	if action == nil {
		return
	}
	return
}

func New(gateway models.Broker) models.Executor {

}
