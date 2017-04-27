package gateways

import "local/trader/models"

func Trade(action *models.Action) {
	if action == nil {
		return
	}
	return
}

func New(gateway models.Broker, conf interface{}) models.Brokerage {
	switch gateway {
	case models.BrokerAnalysis:
	default:
		analysisConf, ok := conf.(models.AnalysisInit)
		if !ok {
			return nil
		}
		return NewAnalysisBroker(analysisConf)
	}
	return nil
}
