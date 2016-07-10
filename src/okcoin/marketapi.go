/*
  btcbot is a Bitcoin trading bot for HUOBI.com written
  in golang, it features multiple trading methods using
  technical analysis.

  Disclaimer:

  USE AT YOUR OWN RISK!

  The author of this project is NOT responsible for any damage or loss caused
  by this software. There can be bugs and the bot may not perform as expected
  or specified. Please consider testing it first with paper trading /
  backtesting on historical data. Also look at the code to see what how
  it's working.

  Weibo:http://weibo.com/bocaicfa
*/

package okcoin

import (
	. "common"
	. "config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"logger"
	"net/http"
	"time"
	"util"
)

func (w *Okcoin) AnalyzeKLinePeroid(symbol string, period string) (ret bool, records []Record) {
	ret = false
	now := time.Now().Unix() * 1000

	req, err := http.NewRequest("GET", fmt.Sprintf(Config["ok_kline_url"], symbol, period, now), nil)
	if err != nil {
		logger.Fatal(err)
		return
	}

	req.Header.Set("Referer", Config["ok_base_url"])
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36")

	logger.Traceln(req)

	c := util.NewTimeoutClient()
	logger.Tracef("okHTTP req begin")
	resp, err := c.Do(req)
	logger.Tracef("okHTTP req end")
	if err != nil {
		logger.Traceln(err)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		logger.Tracef("HTTP returned status %v", resp)
		return
	}

	var body string
	contentEncoding := resp.Header.Get("Content-Encoding")

	logger.Tracef("HTTP returned Content-Encoding %s", contentEncoding)
	logger.Traceln(resp.Header.Get("Content-Type"))

	switch contentEncoding {
	case "gzip":
		body = util.DumpGZIP(resp.Body)

	default:
		bodyByte, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Errorln("read the http stream failed")
			return
		} else {
			body = string(bodyByte)

			ioutil.WriteFile(fmt.Sprintf("cache/okTradeKLine_%s.data", period), bodyByte, 0644)
		}
	}

	return analyzePeroidLine(body)
}

func convert2Records(_records []interface{}) (records []Record) {
	records = make([]Record, len(_records))
	for k, v := range _records {
		if vt, ok := v.([]interface{}); ok {
			for ik, iv := range vt {
				switch ik {
				case 0:
					const layout = "2006-01-02 15:04:05"

					Time := int64(util.InterfaceToFloat64(iv))

					t := time.Unix(Time/1000, 0)

					records[k].TimeStr = t.Format(layout)
					records[k].Time = Time
				case 1:
					records[k].Open = util.InterfaceToFloat64(iv)
				case 2:
					records[k].High = util.InterfaceToFloat64(iv)
				case 3:
					records[k].Low = util.InterfaceToFloat64(iv)
				case 4:
					records[k].Close = util.InterfaceToFloat64(iv)
				case 5:
					records[k].Volumn = util.InterfaceToFloat64(iv)
				}
			}
		}
	}
	return
}
func analyzePeroidLine(content string) (ret bool, records []Record) {
	logger.Traceln("Okcoin analyzePeroidLine begin....")
	ret = false
	var _recordBook []interface{}
	if err := json.Unmarshal([]byte(content), &_recordBook); err != nil {
		logger.Infoln(err)
		return
	}
	records = convert2Records(_recordBook)
	logger.Traceln("Okcoin parsePeroidArray end....")
	ret = true
	return
}
