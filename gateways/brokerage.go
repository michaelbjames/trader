package gateways

import "local/trader/models"

func New(gateway models.BrokerageName, conf interface{}) models.Brokerage {
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
