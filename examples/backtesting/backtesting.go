package main

import (
	"context"

	"github.com/rodrigo-brito/ninjabot"
	"github.com/rodrigo-brito/ninjabot/examples/strategies"
	"github.com/rodrigo-brito/ninjabot/exchange"
	"github.com/rodrigo-brito/ninjabot/plot"
	"github.com/rodrigo-brito/ninjabot/plot/indicator"
	"github.com/rodrigo-brito/ninjabot/storage"
	"github.com/rodrigo-brito/ninjabot/tools/log"
)

// This example shows how to use backtesting with NinjaBot
// Backtesting is a simulation of the strategy in historical data (from CSV)
func main() {
	ctx := context.Background()

	// bot settings (eg: pairs, telegram, etc)
	settings := ninjabot.Settings{
		Pairs: []string{
			"GALUSDT",
			"ACHUSDT",
			"FRONTUSDT",
			"ETHUSDT",
			"BTCUSDT",
			"BICOUSDT",
			"TFUELUSDT",
			"RNDRUSDT",
			"CLVUSDT",
			"SPELLUSDT",
			"RPLUSDT",
			"DYDXUSDT",
			"CYBERUSDT",
			"SNXUSDT",
			"SOLUSDT",
			"RLCUSDT",
			"WOOUSDT",
			"BLZUSDT",
			"CELRUSDT",
			"KEYUSDT",
			"BETAUSDT",
			"AUDIOUSDT",
			"VIBUSDT",
			"TVKUSDT",
			"SSVUSDT",
			"CKBUSDT",
			"LPTUSDT",
			"SHIBUSDT",
			"DFUSDT",
			"LAZIOUSDT",
			"COMPUSDT",
			"IOSTUSDT",
			"XVSUSDT",
			"REQUSDT",
			"LTCUSDT",
			"FLUXUSDT",
			"LEVERUSDT",
			"OCEANUSDT",
			"STGUSDT",
			"QTUMUSDT",
			"QNTUSDT",
			"NEARUSDT",
			"MDTUSDT",
			"NULSUSDT",
			"BONDUSDT",
			"AXSUSDT",
			"GMXUSDT",
			"BCHUSDT",
			"DENTUSDT",
			"ROSEUSDT",
			"UNIUSDT",
			"TRUUSDT",
			"VTHOUSDT",
			"TRXUSDT",
			"RSRUSDT",
			"ZILUSDT",
			"ANTUSDT",
			"MKRUSDT",
			"RUNEUSDT",
			"ATAUSDT",
			"ICPUSDT",
			"MDXUSDT",
			"OGNUSDT",
			"WINGUSDT",
			"PAXGUSDT",
			"FDUSDUSDT",
			"USDCUSDT",
			"TUSDUSDT",
			"MATICUSDT",
			"PUNDIXUSDT",
			"HBARUSDT",
			"ETCUSDT",
			"ENSUSDT",
			"HOTUSDT",
			"MCUSDT",
			"BANDUSDT",
			"KNCUSDT",
			"TWTUSDT",
			"XTZUSDT",
			"DUSKUSDT",
			"EOSUSDT",
			"LUNAUSDT",
			"FXSUSDT",
			"ALGOUSDT",
			"RDNTUSDT",
			"TRBUSDT",
			"COMBOUSDT",
			"APTUSDT",
			"KDAUSDT",
			"IOTAUSDT",
			"JASMYUSDT",
			"CAKEUSDT",
			"DOGEUSDT",
			"AAVEUSDT",
			"WAVESUSDT",
			"UMAUSDT",
			"WLDUSDT",
			"LINKUSDT",
			"BATUSDT",
			"PONDUSDT",
			"APEUSDT",
			"MAGICUSDT",
			"CTKUSDT",
			"DOTUSDT",
			"VETUSDT",
			"ADAUSDT",
			"EDUUSDT",
			"STORJUSDT",
			"GRTUSDT",
			"IOTXUSDT",
			"STPTUSDT",
			"KP3RUSDT",
			"KSMUSDT",
			"AVAXUSDT",
			"AVAUSDT",
			"SANTOSUSDT",
			"PYRUSDT",
			"ARBUSDT",
			"PENDLEUSDT",
			"ZECUSDT",
			"LUNCUSDT",
			"MANAUSDT",
			"CELOUSDT",
			"SFPUSDT",
			"CTXCUSDT",
			"HIVEUSDT",
			"CHZUSDT",
			"ALCXUSDT",
			"CTSIUSDT",
			"OSMOUSDT",
			"BNBUSDT",
			"TUSDT",
			"SLPUSDT",
			"ENJUSDT",
			"GALAUSDT",
			"API3USDT",
			"SUIUSDT",
			"JOEUSDT",
			"WAXPUSDT",
			"LRCUSDT",
			"NEOUSDT",
			"PEOPLEUSDT",
			"TLMUSDT",
			"SKLUSDT",
			"ATOMUSDT",
			"DASHUSDT",
			"KAVAUSDT",
			"MINAUSDT",
			"SXPUSDT",
			"INJUSDT",
			"1INCHUSDT",
			"GLMRUSDT",
			"REEFUSDT",
			"SUSHIUSDT",
			"CRVUSDT",
			"YFIUSDT",
			"CHRUSDT",
			"NKNUSDT",
			"FILUSDT",
			"ONEUSDT",
			"EGLDUSDT",
			"THETAUSDT",
			"FETUSDT",
			"UNFIUSDT",
			"COTIUSDT",
			"QUICKUSDT",
			"GLMUSDT",
			"ANKRUSDT",
			"ONTUSDT",
			"GNSUSDT",
			"HARDUSDT",
			"GMTUSDT",
			"ARKMUSDT",
			"SEIUSDT",
			"QIUSDT",
			"XECUSDT",
			"BELUSDT",
			"HFTUSDT",
			"DODOUSDT",
			"LINAUSDT",
			"SUPERUSDT",
			"FTMUSDT",
			"STXUSDT",
			"C98USDT",
			"FLMUSDT",
			"XLMUSDT",
			"CFXUSDT",
			"XEMUSDT",
			"ALICEUSDT",
			"KLAYUSDT",
			"ZRXUSDT",
			"PHBUSDT",
			"ERNUSDT",
			"IDEXUSDT",
			"EPXUSDT",
			"USTCUSDT",
			"ARPAUSDT",
			"BURGERUSDT",
			"MASKUSDT",
			"XVGUSDT",
			"XRPUSDT",
			"PEPEUSDT",
			"RIFUSDT",
			"HOOKUSDT",
			"BAKEUSDT",
			"YGGUSDT",
			"TOMOUSDT",
			"AMBUSDT",
			"PERPUSDT",
			"RVNUSDT",
			"BNTUSDT",
			"LDOUSDT",
			"BSWUSDT",
			"MAVUSDT",
			"HIGHUSDT",
			"MBLUSDT",
			"OAXUSDT",
			"DARUSDT",
			"MTLUSDT",
			"LITUSDT",
			"SANDUSDT",
			"GTCUSDT",
			"FIDAUSDT",
			"AGIXUSDT",
			"AGLDUSDT",
			"ICXUSDT",
			"STMXUSDT",
			"PHAUSDT",
			"FLOKIUSDT",
			"AUCTIONUSDT",
			"ARUSDT",
			"TROYUSDT",
			"LQTYUSDT",
			"SYNUSDT",
			"ASTRUSDT",
			"NMRUSDT",
			"RADUSDT",
			"POLYXUSDT",
			"FLOWUSDT",
			"IDUSDT",
			"OPUSDT",
			"BNXUSDT",
			"IMXUSDT",
			"WLDUSDT",
			"LINKUSDT",
			"BATUSDT",
			"PONDUSDT",
			"APEUSDT",
			"MAGICUSDT",
			"CTKUSDT",
			"DOTUSDT",
			"VETUSDT",
			"ADAUSDT",
			"EDUUSDT",
			"STORJUSDT",
			"GRTUSDT",
			"IOTXUSDT",
			"STPTUSDT",
			"KP3RUSDT",
			"KSMUSDT",
			"AVAXUSDT",
			"AVAUSDT",
			"SANTOSUSDT",
			"PYRUSDT",
			"ARBUSDT",
			"PENDLEUSDT",
			"ZECUSDT",
			"LUNCUSDT",
			"MANAUSDT",
			"CELOUSDT",
			"SFPUSDT",
			"CTXCUSDT",
			"HIVEUSDT",
			"CHZUSDT",
			"ALCXUSDT",
			"CTSIUSDT",
			"OSMOUSDT",
			"BNBUSDT",
			"TUSDT",
			"SLPUSDT",
			"ENJUSDT",
			"GALAUSDT",
			"API3USDT",
			"SUIUSDT",
			"JOEUSDT",
			"WAXPUSDT",
			"LRCUSDT",
			"NEOUSDT",
			"PEOPLEUSDT",
			"TLMUSDT",
			"SKLUSDT",
			"ATOMUSDT",
			"DASHUSDT",
			"KAVAUSDT",
			"MINAUSDT",
			"SXPUSDT",
			"INJUSDT",
			"1INCHUSDT",
			"GLMRUSDT",
			"REEFUSDT",
			"SUSHIUSDT",
			"CRVUSDT",
			"YFIUSDT",
			"CHRUSDT",
			"NKNUSDT",
			"FILUSDT",
			"ONEUSDT",
			"EGLDUSDT",
			"THETAUSDT",
			"FETUSDT",
			"UNFIUSDT",
		},
	}

	// initialize your strategy
	strategy := new(strategies.CrossEMA)

	// load historical data from CSV files
	csvFeed, err := exchange.NewCSVFeed(
		strategy.Timeframe(),
		exchange.PairFeed{
			Pair:      "GALUSDT",
			File:      "/ninja/GAL-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ACHUSDT",
			File:      "/ninja/ACH-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "FRONTUSDT",
			File:      "/ninja/FRONT-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ETHUSDT",
			File:      "/ninja/ETH-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "BTCUSDT",
			File:      "/ninja/BTC-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "BICOUSDT",
			File:      "/ninja/BICO-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "TFUELUSDT",
			File:      "/ninja/TFUEL-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "RNDRUSDT",
			File:      "/ninja/RNDR-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "CLVUSDT",
			File:      "/ninja/CLV-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "SPELLUSDT",
			File:      "/ninja/SPELL-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "RPLUSDT",
			File:      "/ninja/RPL-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "DYDXUSDT",
			File:      "/ninja/DYDX-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "CYBERUSDT",
			File:      "/ninja/CYBER-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "SNXUSDT",
			File:      "/ninja/SNX-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "SOLUSDT",
			File:      "/ninja/SOL-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "RLCUSDT",
			File:      "/ninja/RLC-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "WOOUSDT",
			File:      "/ninja/WOO-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "BLZUSDT",
			File:      "/ninja/BLZ-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "CELRUSDT",
			File:      "/ninja/CELR-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "KEYUSDT",
			File:      "/ninja/KEY-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "BETAUSDT",
			File:      "/ninja/BETA-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "AUDIOUSDT",
			File:      "/ninja/AUDIO-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "VIBUSDT",
			File:      "/ninja/VIB-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "TVKUSDT",
			File:      "/ninja/TVK-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "SSVUSDT",
			File:      "/ninja/SSV-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "CKBUSDT",
			File:      "/ninja/CKB-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "LPTUSDT",
			File:      "/ninja/LPT-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "SHIBUSDT",
			File:      "/ninja/SHIB-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "DFUSDT",
			File:      "/ninja/DF-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "LAZIOUSDT",
			File:      "/ninja/LAZIO-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "COMPUSDT",
			File:      "/ninja/COMP-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "IOSTUSDT",
			File:      "/ninja/IOST-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "XVSUSDT",
			File:      "/ninja/XVS-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "REQUSDT",
			File:      "/ninja/REQ-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "LTCUSDT",
			File:      "/ninja/LTC-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "FLUXUSDT",
			File:      "/ninja/FLUX-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "LEVERUSDT",
			File:      "/ninja/LEVER-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "OCEANUSDT",
			File:      "/ninja/OCEAN-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "STGUSDT",
			File:      "/ninja/STG-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "QTUMUSDT",
			File:      "/ninja/QTUM-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "QNTUSDT",
			File:      "/ninja/QNT-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "NEARUSDT",
			File:      "/ninja/NEAR-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "MDTUSDT",
			File:      "/ninja/MDT-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "NULSUSDT",
			File:      "/ninja/NULS-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "BONDUSDT",
			File:      "/ninja/BOND-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "AXSUSDT",
			File:      "/ninja/AXS-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "GMXUSDT",
			File:      "/ninja/GMX-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "BCHUSDT",
			File:      "/ninja/BCH-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "DENTUSDT",
			File:      "/ninja/DENT-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ROSEUSDT",
			File:      "/ninja/ROSE-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "UNIUSDT",
			File:      "/ninja/UNI-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "TRUUSDT",
			File:      "/ninja/TRU-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "VTHOUSDT",
			File:      "/ninja/VTHO-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "TRXUSDT",
			File:      "/ninja/TRX-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "RSRUSDT",
			File:      "/ninja/RSR-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ZILUSDT",
			File:      "/ninja/ZIL-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ANTUSDT",
			File:      "/ninja/ANT-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "MKRUSDT",
			File:      "/ninja/MKR-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "RUNEUSDT",
			File:      "/ninja/RUNE-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ATAUSDT",
			File:      "/ninja/ATA-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ICPUSDT",
			File:      "/ninja/ICP-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "MDXUSDT",
			File:      "/ninja/MDX-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "OGNUSDT",
			File:      "/ninja/OGN-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "WINGUSDT",
			File:      "/ninja/WING-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "PAXGUSDT",
			File:      "/ninja/PAXG-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "FDUSDUSDT",
			File:      "/ninja/FDUSD-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "USDCUSDT",
			File:      "/ninja/USDC-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "TUSDUSDT",
			File:      "/ninja/TUSD-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "MATICUSDT",
			File:      "/ninja/MATIC-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "PUNDIXUSDT",
			File:      "/ninja/PUNDIX-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "HBARUSDT",
			File:      "/ninja/HBAR-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ETCUSDT",
			File:      "/ninja/ETC-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ENSUSDT",
			File:      "/ninja/ENS-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "HOTUSDT",
			File:      "/ninja/HOT-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "MCUSDT",
			File:      "/ninja/MC-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "BANDUSDT",
			File:      "/ninja/BAND-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "KNCUSDT",
			File:      "/ninja/KNC-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "TWTUSDT",
			File:      "/ninja/TWT-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "XTZUSDT",
			File:      "/ninja/XTZ-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "DUSKUSDT",
			File:      "/ninja/DUSK-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "EOSUSDT",
			File:      "/ninja/EOS-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "LUNAUSDT",
			File:      "/ninja/LUNA-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "FXSUSDT",
			File:      "/ninja/FXS-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ALGOUSDT",
			File:      "/ninja/ALGO-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "RDNTUSDT",
			File:      "/ninja/RDNT-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "TRBUSDT",
			File:      "/ninja/TRB-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "COMBOUSDT",
			File:      "/ninja/COMBO-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "APTUSDT",
			File:      "/ninja/APT-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "KDAUSDT",
			File:      "/ninja/KDA-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "IOTAUSDT",
			File:      "/ninja/IOTA-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "JASMYUSDT",
			File:      "/ninja/JASMY-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "CAKEUSDT",
			File:      "/ninja/CAKE-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "DOGEUSDT",
			File:      "/ninja/DOGE-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "AAVEUSDT",
			File:      "/ninja/AAVE-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "WAVESUSDT",
			File:      "/ninja/WAVES-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "UMAUSDT",
			File:      "/ninja/UMA-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "WLDUSDT",
			File:      "/ninja/WLD-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "LINKUSDT",
			File:      "/ninja/LINK-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "BATUSDT",
			File:      "/ninja/BAT-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "PONDUSDT",
			File:      "/ninja/POND-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "APEUSDT",
			File:      "/ninja/APE-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "MAGICUSDT",
			File:      "/ninja/MAGIC-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "CTKUSDT",
			File:      "/ninja/CTK-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "DOTUSDT",
			File:      "/ninja/DOT-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "VETUSDT",
			File:      "/ninja/VET-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ADAUSDT",
			File:      "/ninja/ADA-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "EDUUSDT",
			File:      "/ninja/EDU-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "STORJUSDT",
			File:      "/ninja/STORJ-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "GRTUSDT",
			File:      "/ninja/GRT-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "IOTXUSDT",
			File:      "/ninja/IOTX-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "STPTUSDT",
			File:      "/ninja/STPT-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "KP3RUSDT",
			File:      "/ninja/KP3R-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "KSMUSDT",
			File:      "/ninja/KSM-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "AVAXUSDT",
			File:      "/ninja/AVAX-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "AVAUSDT",
			File:      "/ninja/AVA-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "SANTOSUSDT",
			File:      "/ninja/SANTOS-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "PYRUSDT",
			File:      "/ninja/PYR-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ARBUSDT",
			File:      "/ninja/ARB-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "PENDLEUSDT",
			File:      "/ninja/PENDLE-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ZECUSDT",
			File:      "/ninja/ZEC-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "LUNCUSDT",
			File:      "/ninja/LUNC-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "MANAUSDT",
			File:      "/ninja/MANA-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "CELOUSDT",
			File:      "/ninja/CELO-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "SFPUSDT",
			File:      "/ninja/SFP-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "CTXCUSDT",
			File:      "/ninja/CTXC-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "HIVEUSDT",
			File:      "/ninja/HIVE-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "CHZUSDT",
			File:      "/ninja/CHZ-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ALCXUSDT",
			File:      "/ninja/ALCX-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "CTSIUSDT",
			File:      "/ninja/CTSI-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "OSMOUSDT",
			File:      "/ninja/OSMO-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "BNBUSDT",
			File:      "/ninja/BNB-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "TUSDT",
			File:      "/ninja/T-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "SLPUSDT",
			File:      "/ninja/SLP-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ENJUSDT",
			File:      "/ninja/ENJ-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "GALAUSDT",
			File:      "/ninja/GALA-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "API3USDT",
			File:      "/ninja/API3-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "SUIUSDT",
			File:      "/ninja/SUI-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "JOEUSDT",
			File:      "/ninja/JOE-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "WAXPUSDT",
			File:      "/ninja/WAXP-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "LRCUSDT",
			File:      "/ninja/LRC-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "NEOUSDT",
			File:      "/ninja/NEO-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "PEOPLEUSDT",
			File:      "/ninja/PEOPLE-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "TLMUSDT",
			File:      "/ninja/TLM-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "SKLUSDT",
			File:      "/ninja/SKL-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ATOMUSDT",
			File:      "/ninja/ATOM-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "DASHUSDT",
			File:      "/ninja/DASH-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "KAVAUSDT",
			File:      "/ninja/KAVA-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "MINAUSDT",
			File:      "/ninja/MINA-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "SXPUSDT",
			File:      "/ninja/SXP-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "INJUSDT",
			File:      "/ninja/INJ-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "1INCHUSDT",
			File:      "/ninja/1INCH-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "GLMRUSDT",
			File:      "/ninja/GLMR-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "REEFUSDT",
			File:      "/ninja/REEF-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "SUSHIUSDT",
			File:      "/ninja/SUSHI-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "CRVUSDT",
			File:      "/ninja/CRV-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "YFIUSDT",
			File:      "/ninja/YFI-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "CHRUSDT",
			File:      "/ninja/CHR-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "NKNUSDT",
			File:      "/ninja/NKN-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "FILUSDT",
			File:      "/ninja/FIL-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ONEUSDT",
			File:      "/ninja/ONE-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "EGLDUSDT",
			File:      "/ninja/EGLD-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "THETAUSDT",
			File:      "/ninja/THETA-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "FETUSDT",
			File:      "/ninja/FET-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "UNFIUSDT",
			File:      "/ninja/UNFI-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "COTIUSDT",
			File:      "/ninja/COTI-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "QUICKUSDT",
			File:      "/ninja/QUICK-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "GLMUSDT",
			File:      "/ninja/GLM-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ANKRUSDT",
			File:      "/ninja/ANKR-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ONTUSDT",
			File:      "/ninja/ONT-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "GNSUSDT",
			File:      "/ninja/GNS-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "HARDUSDT",
			File:      "/ninja/HARD-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "GMTUSDT",
			File:      "/ninja/GMT-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ARKMUSDT",
			File:      "/ninja/ARKM-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "SEIUSDT",
			File:      "/ninja/SEI-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "QIUSDT",
			File:      "/ninja/QI-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "XECUSDT",
			File:      "/ninja/XEC-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "BELUSDT",
			File:      "/ninja/BEL-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "HFTUSDT",
			File:      "/ninja/HFT-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "DODOUSDT",
			File:      "/ninja/DODO-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "LINAUSDT",
			File:      "/ninja/LINA-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "SUPERUSDT",
			File:      "/ninja/SUPER-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "FTMUSDT",
			File:      "/ninja/FTM-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "STXUSDT",
			File:      "/ninja/STX-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "C98USDT",
			File:      "/ninja/C98-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "FLMUSDT",
			File:      "/ninja/FLM-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "XLMUSDT",
			File:      "/ninja/XLM-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "CFXUSDT",
			File:      "/ninja/CFX-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "XEMUSDT",
			File:      "/ninja/XEM-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ALICEUSDT",
			File:      "/ninja/ALICE-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "KLAYUSDT",
			File:      "/ninja/KLAY-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ZRXUSDT",
			File:      "/ninja/ZRX-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "PHBUSDT",
			File:      "/ninja/PHB-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ERNUSDT",
			File:      "/ninja/ERN-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "IDEXUSDT",
			File:      "/ninja/IDEX-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "EPXUSDT",
			File:      "/ninja/EPX-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "USTCUSDT",
			File:      "/ninja/USTC-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ARPAUSDT",
			File:      "/ninja/ARPA-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "BURGERUSDT",
			File:      "/ninja/BURGER-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "MASKUSDT",
			File:      "/ninja/MASK-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "XVGUSDT",
			File:      "/ninja/XVG-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "XRPUSDT",
			File:      "/ninja/XRP-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "PEPEUSDT",
			File:      "/ninja/PEPE-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "RIFUSDT",
			File:      "/ninja/RIF-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "HOOKUSDT",
			File:      "/ninja/HOOK-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "BAKEUSDT",
			File:      "/ninja/BAKE-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "YGGUSDT",
			File:      "/ninja/YGG-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "TOMOUSDT",
			File:      "/ninja/TOMO-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "AMBUSDT",
			File:      "/ninja/AMB-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "PERPUSDT",
			File:      "/ninja/PERP-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "RVNUSDT",
			File:      "/ninja/RVN-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "BNTUSDT",
			File:      "/ninja/BNT-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "LDOUSDT",
			File:      "/ninja/LDO-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "BSWUSDT",
			File:      "/ninja/BSW-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "MAVUSDT",
			File:      "/ninja/MAV-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "HIGHUSDT",
			File:      "/ninja/HIGH-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "MBLUSDT",
			File:      "/ninja/MBL-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "OAXUSDT",
			File:      "/ninja/OAX-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "DARUSDT",
			File:      "/ninja/DAR-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "MTLUSDT",
			File:      "/ninja/MTL-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "LITUSDT",
			File:      "/ninja/LIT-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "SANDUSDT",
			File:      "/ninja/SAND-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "GTCUSDT",
			File:      "/ninja/GTC-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "FIDAUSDT",
			File:      "/ninja/FIDA-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "AGIXUSDT",
			File:      "/ninja/AGIX-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "AGLDUSDT",
			File:      "/ninja/AGLD-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ICXUSDT",
			File:      "/ninja/ICX-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "STMXUSDT",
			File:      "/ninja/STMX-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "PHAUSDT",
			File:      "/ninja/PHA-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "FLOKIUSDT",
			File:      "/ninja/FLOKI-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "AUCTIONUSDT",
			File:      "/ninja/AUCTION-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ARUSDT",
			File:      "/ninja/AR-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "TROYUSDT",
			File:      "/ninja/TROY-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "LQTYUSDT",
			File:      "/ninja/LQTY-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "SYNUSDT",
			File:      "/ninja/SYN-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ASTRUSDT",
			File:      "/ninja/ASTR-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "NMRUSDT",
			File:      "/ninja/NMR-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "RADUSDT",
			File:      "/ninja/RAD-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "POLYXUSDT",
			File:      "/ninja/POLYX-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "FLOWUSDT",
			File:      "/ninja/FLOW-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "IDUSDT",
			File:      "/ninja/ID-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "OPUSDT",
			File:      "/ninja/OP-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "BNXUSDT",
			File:      "/ninja/BNX-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "IMXUSDT",
			File:      "/ninja/IMX-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "WLDUSDT",
			File:      "/ninja/WLD-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "LINKUSDT",
			File:      "/ninja/LINK-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "BATUSDT",
			File:      "/ninja/BAT-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "PONDUSDT",
			File:      "/ninja/POND-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "APEUSDT",
			File:      "/ninja/APE-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "MAGICUSDT",
			File:      "/ninja/MAGIC-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "CTKUSDT",
			File:      "/ninja/CTK-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "DOTUSDT",
			File:      "/ninja/DOT-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "VETUSDT",
			File:      "/ninja/VET-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ADAUSDT",
			File:      "/ninja/ADA-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "EDUUSDT",
			File:      "/ninja/EDU-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "STORJUSDT",
			File:      "/ninja/STORJ-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "GRTUSDT",
			File:      "/ninja/GRT-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "IOTXUSDT",
			File:      "/ninja/IOTX-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "STPTUSDT",
			File:      "/ninja/STPT-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "KP3RUSDT",
			File:      "/ninja/KP3R-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "KSMUSDT",
			File:      "/ninja/KSM-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "AVAXUSDT",
			File:      "/ninja/AVAX-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "AVAUSDT",
			File:      "/ninja/AVA-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "SANTOSUSDT",
			File:      "/ninja/SANTOS-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "PYRUSDT",
			File:      "/ninja/PYR-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ARBUSDT",
			File:      "/ninja/ARB-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "PENDLEUSDT",
			File:      "/ninja/PENDLE-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ZECUSDT",
			File:      "/ninja/ZEC-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "LUNCUSDT",
			File:      "/ninja/LUNC-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "MANAUSDT",
			File:      "/ninja/MANA-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "CELOUSDT",
			File:      "/ninja/CELO-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "SFPUSDT",
			File:      "/ninja/SFP-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "CTXCUSDT",
			File:      "/ninja/CTXC-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "HIVEUSDT",
			File:      "/ninja/HIVE-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "CHZUSDT",
			File:      "/ninja/CHZ-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ALCXUSDT",
			File:      "/ninja/ALCX-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "CTSIUSDT",
			File:      "/ninja/CTSI-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "OSMOUSDT",
			File:      "/ninja/OSMO-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "BNBUSDT",
			File:      "/ninja/BNB-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "TUSDT",
			File:      "/ninja/T-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "SLPUSDT",
			File:      "/ninja/SLP-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ENJUSDT",
			File:      "/ninja/ENJ-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "GALAUSDT",
			File:      "/ninja/GALA-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "API3USDT",
			File:      "/ninja/API3-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "SUIUSDT",
			File:      "/ninja/SUI-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "JOEUSDT",
			File:      "/ninja/JOE-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "WAXPUSDT",
			File:      "/ninja/WAXP-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "LRCUSDT",
			File:      "/ninja/LRC-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "NEOUSDT",
			File:      "/ninja/NEO-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "PEOPLEUSDT",
			File:      "/ninja/PEOPLE-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "TLMUSDT",
			File:      "/ninja/TLM-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "SKLUSDT",
			File:      "/ninja/SKL-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ATOMUSDT",
			File:      "/ninja/ATOM-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "DASHUSDT",
			File:      "/ninja/DASH-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "KAVAUSDT",
			File:      "/ninja/KAVA-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "MINAUSDT",
			File:      "/ninja/MINA-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "SXPUSDT",
			File:      "/ninja/SXP-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "INJUSDT",
			File:      "/ninja/INJ-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "1INCHUSDT",
			File:      "/ninja/1INCH-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "GLMRUSDT",
			File:      "/ninja/GLMR-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "REEFUSDT",
			File:      "/ninja/REEF-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "SUSHIUSDT",
			File:      "/ninja/SUSHI-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "CRVUSDT",
			File:      "/ninja/CRV-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "YFIUSDT",
			File:      "/ninja/YFI-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "CHRUSDT",
			File:      "/ninja/CHR-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "NKNUSDT",
			File:      "/ninja/NKN-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "FILUSDT",
			File:      "/ninja/FIL-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "ONEUSDT",
			File:      "/ninja/ONE-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "EGLDUSDT",
			File:      "/ninja/EGLD-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "THETAUSDT",
			File:      "/ninja/THETA-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "FETUSDT",
			File:      "/ninja/FET-1h.csv",
			Timeframe: "1h",
		},

		exchange.PairFeed{
			Pair:      "UNFIUSDT",
			File:      "/ninja/UNFI-1h.csv",
			Timeframe: "1h",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// initialize a database in memory
	storage, err := storage.FromMemory()
	if err != nil {
		log.Fatal(err)
	}

	// create a paper wallet for simulation, initializing with 1.000 USDT
	wallet := exchange.NewPaperWallet(
		ctx,
		"USDT",
		exchange.WithPaperAsset("USDT", 1000),
		exchange.WithDataFeed(csvFeed),
	)

	// create a chart  with indicators from the strategy and a custom additional RSI indicator
	chart, err := plot.NewChart(
		plot.WithStrategyIndicators(strategy),
		plot.WithCustomIndicators(
			indicator.RSI(14, "purple"),
		),
		plot.WithPaperWallet(wallet),
	)
	if err != nil {
		log.Fatal(err)
	}

	// initializer Ninjabot with the objects created before
	bot, err := ninjabot.NewBot(
		ctx,
		settings,
		wallet,
		strategy,
		ninjabot.WithBacktest(wallet), // Required for Backtest mode
		ninjabot.WithStorage(storage),

		// connect bot feed (candle and orders) to the chart
		ninjabot.WithCandleSubscription(chart),
		ninjabot.WithOrderSubscription(chart),
		ninjabot.WithLogLevel(log.WarnLevel),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Initializer simulation
	err = bot.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Print bot results
	bot.Summary()

	// Display candlesticks chart in local browser
	err = chart.Start()
	if err != nil {
		log.Fatal(err)
	}
}
