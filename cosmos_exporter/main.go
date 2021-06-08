package main

import (
 "encoding/json"
 "io/ioutil"
 "log"
 "net/http"
 "strconv"
 "time"
 "github.com/prometheus/client_golang/prometheus"
 "github.com/prometheus/client_golang/prometheus/promhttp"
)

// Get number of blockchain peers from REST API of Gaia
func getNumberPeers() float64 {

	// Inner struct for JSON /net_info 'n_peers'
	type Peers struct {
		Peers string `json:"n_peers"`
    }

    // Outer struct for JSON /net_info 'result'
	type Result struct {
		Result Peers `json:"result"`
    }

	url := "http://localhost:26657/net_info"

	gaiaClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := gaiaClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var peers_number Result
	jsonErr := json.Unmarshal(body, &peers_number)
        if jsonErr != nil {
		log.Fatalf("Unable to parse value: %q, error: %s", string(body), jsonErr.Error())
	}

	numberPeersValue, err := strconv.ParseFloat(peers_number.Result.Peers, 64)
	return float64(numberPeersValue)
}

// Get current block number from REST API of Gaia
func getCurrentBlock() float64 {

	// Inner struct for JSON /status 'block'
	type Block struct {
		Block string `json:"block"`
    }

	// Inner struct for JSON /status 'protocol_version'
	type ProtocolVersion struct {
		ProtocolVersion Block `json:"protocol_version"`
    }

	// Inner struct for JSON /status 'node_info'
	type NodeInfo struct {
		NodeInfo ProtocolVersion `json:"node_info"`
    }

    // Outer struct for JSON /status 'result'
	type Result struct {
		Result NodeInfo `json:"result"`
    }

	url := "http://localhost:26657/status"

	gaiaClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := gaiaClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var block_number Result
	jsonErr := json.Unmarshal(body, &block_number)
        if jsonErr != nil {
		log.Fatalf("Unable to parse value: %q, error: %s", string(body), jsonErr.Error())
	}

	blockNumberValue, err := strconv.ParseFloat(block_number.Result.NodeInfo.ProtocolVersion.Block, 64)
	return float64(blockNumberValue)
}

// Get time for current block from REST API of Gaia
func getCurrentBlockTime() string {

	// Inner struct for JSON /status 'protocol_version'
	type BlockTime struct {
		BlockTime string `json:"latest_block_time"`
    }

	// Inner struct for JSON /status 'sync_info'
	type SyncInfo struct {
		SyncInfo BlockTime `json:"sync_info"`
    }

    // Outer struct for JSON /status 'result'
	type Result struct {
		Result SyncInfo `json:"result"`
    }

	url := "http://localhost:26657/status"

	gaiaClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := gaiaClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var block_time Result
	jsonErr := json.Unmarshal(body, &block_time)
        if jsonErr != nil {
		log.Fatalf("Unable to parse value: %q, error: %s", string(body), jsonErr.Error())
	}

	return string(block_time.Result.SyncInfo.BlockTime)
}

// Calculate time difference in seconds
func getTimeDifference(timestamp string) float64 {

	// Get time now
	dt1 := time.Now()
	// Get current block time and convert into Time struct
	layout := "2006-01-02T15:04:05Z"
	dt2, err := time.Parse(layout, timestamp)
	if err != nil {
		log.Fatal(err)
	}

	return (dt1.Sub(dt2).Seconds())
}

//Define a struct for the collector that contains pointers
//to prometheus descriptors for each metric you wish to expose.
//Note you can also include fields of other types if they provide utility
//but we just won't be exposing them as metrics.
type cosmosCollector struct {
	numberPeersMetric *prometheus.Desc
	blockNumberMetric *prometheus.Desc
	blockTimeSyncMetric *prometheus.Desc
}

//You must create a constructor for you collector that
//initializes every descriptor and returns a pointer to the collector
func newCosmosCollector() *cosmosCollector {
	return &cosmosCollector{
		numberPeersMetric: prometheus.NewDesc("cosmos_peers_number",
			"Number of peers in blockchain network",
			nil, nil,
		),
		blockNumberMetric: prometheus.NewDesc("cosmos_block_number",
			"Current block number at the node",
			nil, nil,
		),
		blockTimeSyncMetric: prometheus.NewDesc("cosmos_block_time_out_sync_seconds",
			"Time out of sync between current block time and time now",
			nil, nil,
		),
	}
}

//Each and every collector must implement the Describe function.
//It essentially writes all descriptors to the prometheus desc channel.
func (collector *cosmosCollector) Describe(ch chan<- *prometheus.Desc) {

	//Update this section with the each metric you create for a given collector
	ch <- collector.numberPeersMetric
	ch <- collector.blockNumberMetric
	ch <- collector.blockTimeSyncMetric
}

//Collect implements required collect function for all promehteus collectors
func (collector *cosmosCollector) Collect(ch chan<- prometheus.Metric) {

	//Implement logic here to determine proper metric value to return to prometheus
	//for each descriptor or call other functions that do so.
	numberPeersMetricValue := getNumberPeers()
	blockNumberMetricValue := getCurrentBlock()
	blockTimeSyncMetricValue := getTimeDifference(getCurrentBlockTime())

	//Write latest value for each metric in the prometheus metric channel.
	//Note that you can pass CounterValue, GaugeValue, or UntypedValue types here.
	ch <- prometheus.MustNewConstMetric(collector.numberPeersMetric, prometheus.CounterValue, numberPeersMetricValue)
	ch <- prometheus.MustNewConstMetric(collector.blockNumberMetric, prometheus.CounterValue, blockNumberMetricValue)
	ch <- prometheus.MustNewConstMetric(collector.blockTimeSyncMetric, prometheus.CounterValue, blockTimeSyncMetricValue)
	log.Println("Endpoint scraped")

}

func main() {
   cosmos := newCosmosCollector()
   prometheus.MustRegister(cosmos)

   http.Handle("/metrics", promhttp.Handler())
   log.Println("Start listening on port :9201")
   log.Fatal(http.ListenAndServe(":9201", nil))
}
