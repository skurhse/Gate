package core

import (
	"errors"
	"github.com/fatih/color"
	"github.com/mrNobody95/Gate/brokerages"
	"github.com/mrNobody95/Gate/indicators"
	"github.com/mrNobody95/Gate/models"
	"github.com/mrNobody95/Gate/storage"
	"github.com/mrNobody95/Gate/strategies"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strconv"
	"time"
)

type Node struct {
	Markets         []models.Market
	Strategy        *strategies.Strategy
	Requests        brokerages.BrokerageRequests
	Brokerage       *models.Brokerage
	Resolutions     []models.Resolution
	PivotResolution *models.Resolution
	FakeTrading     bool
	EnableTrading   bool
	IndicatorConfig *indicators.Configuration
	dataChannel     chan models.Candle
	//NetworkManager   interface{}
	//IndicatorConfig  indicators.Configuration
	//PivotResolution  map[models.Market]models.Resolution
	//HelperResolution map[models.Market]models.Resolution
	//BufferedCandles  map[models.Market]map[string][]models.Candle
}

func (node *Node) Validate() error {
	color.HiGreen("Validating node")
	if node.Brokerage == nil {
		return errors.New("you must declared working brokerage")
	}
	if node.Requests == nil {
		return errors.New("brokerage request API not declared")
	}
	if node.Strategy == nil {
		return errors.New("trading strategy must be declared")
	}
	return nil
}

func (node *Node) Start() error {
	color.HiGreen("Starting node for %s", node.Brokerage.Name)
	node.dataChannel = make(chan models.Candle, 100)
	for _, market := range node.Markets {
		go func(market *models.Market) {
			pool, err := storage.NewPool(node.Strategy.BufferedCandleCount, market.Id, node.PivotResolution.Id)
			if err != nil {
				color.Red("create candle pool failed for market %s in timeframe %s: %v",
					market.Name, node.PivotResolution.Label, err.Error())
				return
			}
			thread := MarketThread{
				Node:            node,
				Market:          market,
				StartFrom:       market.StartTime,
				Resolution:      node.PivotResolution,
				IndicatorConfig: node.IndicatorConfig,
				CandlePool:      pool,
			}
			if err = thread.CollectPrimaryData(); err != nil {
				log.WithError(err).Errorf("faile to collect primary data for %s in timeframe %s",
					thread.Market.Name, thread.PivotResolution.Label)
				return
			}
			thread.PeriodicOHLC()
		}(&market)
	}
	return nil
}

func (node *Node) LoadConfig(path string) error {
	color.HiGreen("Loading YAML config")
	if path == "" {
		path = node.Brokerage.DefaultConfigPath
	}
	confBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var config NodeConfig
	if err = yaml.Unmarshal(confBytes, &config); err != nil {
		return err
	}
	if config.LoadMarketsOnline {
		resp := node.Requests.MarketList()
		if resp.Error != nil {
			return resp.Error
		}
		config.Markets = resp.Markets
	}
	for _, market := range config.Markets {
		if market.StartTimeString != "" {
			t, err := strconv.ParseInt(market.StartTimeString, 10, 64)
			if err != nil {
				return err
			}
			market.StartTime = time.Unix(t, 0)
		} else {
			market.StartTime = time.Unix(1594512000, 0)
		}
		market.BrokerageRefer = node.Brokerage.Id
		err = market.CreateOrLoad()
		if err != nil {
			return err
		}
		node.Markets = append(node.Markets, market)
	}
	for _, resolution := range config.Resolutions {
		resolution.BrokerageRefer = node.Brokerage.Id
		err = resolution.CreateOrLoad()
		if err != nil {
			return err
		}
		node.Resolutions = append(node.Resolutions, resolution)
	}

	node.PivotResolution = &node.Resolutions[0]
	node.Strategy = &config.Strategy
	node.FakeTrading = config.FakeTrading
	node.EnableTrading = config.EnableTrading
	node.IndicatorConfig = &config.IndicatorConfigs
	return nil
}

//func (node *Node) cliPrinter() {
//	dataMap := make(map[uint16]map[uint]models.Candle)
//	writer := uilive.New()
//	writer.Start()
//	for candle := range node.dataChannel {
//		_, ok := dataMap[candle.Market.Id][candle.Resolution.Id]
//		if !ok {
//			dataMap[candle.Market.Id] = make(map[uint]models.Candle)
//		}
//		dataMap[candle.Market.Id][candle.Resolution.Id] = candle
//		output := "+------------+------------+--------------------+--------------------+--------------------+--------------------+--------------------+\n"
//		output += "|   Market   | Resolution |         RSI        |     Stochastic     | Bollinger band(Up) |Bollinger band(Down)|Bollinger band(Midl)|\n"
//		output += "+------------+------------+--------------------+--------------------+--------------------+--------------------+--------------------+\n"
//
//		for _, resolution := range dataMap {
//			for _, data := range resolution {
//				output += fmt.Sprintf("| %-10s | %-10s | %-18f | %-18f | %-18f | %-18f | %-18f |\n+------------+------------+--------------------+--------------------+--------------------+--------------------+--------------------+\n",
//					data.Market.Name, data.Resolution.Label,
//					data.MACD.MACD, data.MACD.Signal, data.ADX.DIPositive, data.ADX.DX, data.MA)
//			}
//		}
//
//		fmt.Fprint(writer, output)
//	}
//	color.Blue("closing printer channel")
//	writer.Stop()
//}
