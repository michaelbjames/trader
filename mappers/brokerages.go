package mappers

import (
	"errors"
	"local/trader/models"
)

func ParseBroker(s string) (models.BrokerageName, error) {
	switch models.BrokerageName(s) {
	case models.BrokerAnalysis:
	case models.BrokerRobinhood:
	case models.BrokerCoindesk:
		return models.BrokerageName(s), nil
	}
	return models.BrokerageName(""), errors.New("broker not available")
}
