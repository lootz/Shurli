package sagoutil

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/satindergrewal/kmdgo"
)

// ConfigCoins type holds the values of coin information
type ConfigCoins struct {
	Coin     string `json:"coin"`
	Ticker   string `json:"ticker"`
	Name     string `json:"name"`
	Shielded bool   `json:"shielded"`
}

// SubAtomicConfig holds the app's confugration settings
type SubAtomicConfig struct {
	Chains       []ConfigCoins     `json:"chains"`
	SubatomicExe string            `json:"subatomic_exe"`
	SubatomicDir string            `json:"subatomic_dir"`
	Explorers    map[string]string `json:"explorers"`
}

//SubAtomicConfInfo returns application's config params
func SubAtomicConfInfo() SubAtomicConfig {
	var conf SubAtomicConfig
	content, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(content, &conf)
	// fmt.Println(conf.Explorers["zVRSC"])
	return conf
}

//StrToAppType converts and returns slice of string as slice of kmdgo.AppType
func StrToAppType(chain []ConfigCoins) []kmdgo.AppType {
	var chainskmd []kmdgo.AppType
	for _, v := range chain {
		// fmt.Println(v.Coin)
		chainskmd = append(chainskmd, kmdgo.AppType(v.Coin))
	}
	return chainskmd
}

// GetCoinConfInfo returns single coin info from config.json cofiguration list
func GetCoinConfInfo(coin string) ConfigCoins {
	var conf SubAtomicConfig = SubAtomicConfInfo()
	var confChains []ConfigCoins = conf.Chains
	// fmt.Println(confChains)

	if coin == "KMD" || coin == "Komodo" {
		coin = "komodo"
	}

	// fmt.Println("getCoinConfInfo:", coin)

	var coininfo ConfigCoins
	for _, v := range confChains {
		// fmt.Println(v)
		if coin == v.Coin {
			coininfo = v
			return coininfo
		}
	}

	return coininfo
}