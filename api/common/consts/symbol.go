package consts

import (
	"slices"
	"strings"
)

// Symbol identifies a tradable product on GMO Coin. Single-currency values
// (e.g. SymbolBTC) are spot trading symbols, while pair values suffixed with
// "JPY" (e.g. SymbolBTCJPY) are leveraged trading symbols.
type Symbol string

const (
	// SymbolBTC is the spot trading symbol for Bitcoin.
	SymbolBTC = Symbol("BTC")
	// SymbolETH is the spot trading symbol for Ethereum.
	SymbolETH = Symbol("ETH")
	// SymbolBCH is the spot trading symbol for Bitcoin Cash.
	SymbolBCH = Symbol("BCH")
	// SymbolLTC is the spot trading symbol for Litecoin.
	SymbolLTC = Symbol("LTC")
	// SymbolXRP is the spot trading symbol for Ripple.
	SymbolXRP = Symbol("XRP")
	// SymbolXLM is the spot trading symbol for Stellar Lumens.
	SymbolXLM = Symbol("XLM")
	// SymbolOMG is the spot trading symbol for OMG Network.
	SymbolOMG = Symbol("OMG")
	// SymbolXTZ is the spot trading symbol for Tezos.
	SymbolXTZ = Symbol("XTZ")
	// SymbolDOT is the spot trading symbol for Polkadot.
	SymbolDOT = Symbol("DOT")
	// SymbolATOM is the spot trading symbol for Cosmos.
	SymbolATOM = Symbol("ATOM")
	// SymbolDAI is the spot trading symbol for Dai.
	SymbolDAI = Symbol("DAI")
	// SymbolFCR is the spot trading symbol for FC Ryukyu Coin.
	SymbolFCR = Symbol("FCR")
	// SymbolADA is the spot trading symbol for Cardano.
	SymbolADA = Symbol("ADA")
	// SymbolLINK is the spot trading symbol for Chainlink.
	SymbolLINK = Symbol("LINK")
	// SymbolDOGE is the spot trading symbol for Dogecoin.
	SymbolDOGE = Symbol("DOGE")
	// SymbolSOL is the spot trading symbol for Solana.
	SymbolSOL = Symbol("SOL")
	// SymbolASTR is the spot trading symbol for Astar.
	SymbolASTR = Symbol("ASTR")
	// SymbolAVAX is the spot trading symbol for Avalanche.
	SymbolAVAX = Symbol("AVAX")
	// SymbolNAC is the spot trading symbol for Nippon Active Coin.
	SymbolNAC = Symbol("NAC")
	// SymbolBTCJPY is the leveraged trading symbol for the BTC/JPY pair.
	SymbolBTCJPY = Symbol("BTC_JPY")
	// SymbolETHJPY is the leveraged trading symbol for the ETH/JPY pair.
	SymbolETHJPY = Symbol("ETH_JPY")
	// SymbolBCHJPY is the leveraged trading symbol for the BCH/JPY pair.
	SymbolBCHJPY = Symbol("BCH_JPY")
	// SymbolLTCJPY is the leveraged trading symbol for the LTC/JPY pair.
	SymbolLTCJPY = Symbol("LTC_JPY")
	// SymbolXRPJPY is the leveraged trading symbol for the XRP/JPY pair.
	SymbolXRPJPY = Symbol("XRP_JPY")
	// SymbolDOTJPY is the leveraged trading symbol for the DOT/JPY pair.
	SymbolDOTJPY = Symbol("DOT_JPY")
	// SymbolATOMJPY is the leveraged trading symbol for the ATOM/JPY pair.
	SymbolATOMJPY = Symbol("ATOM_JPY")
	// SymbolADAJPY is the leveraged trading symbol for the ADA/JPY pair.
	SymbolADAJPY = Symbol("ADA_JPY")
	// SymbolLINKJPY is the leveraged trading symbol for the LINK/JPY pair.
	SymbolLINKJPY = Symbol("LINK_JPY")
	// SymbolDOGEJPY is the leveraged trading symbol for the DOGE/JPY pair.
	SymbolDOGEJPY = Symbol("DOGE_JPY")
	// SymbolSOLJPY is the leveraged trading symbol for the SOL/JPY pair.
	SymbolSOLJPY = Symbol("SOL_JPY")
	// SymbolJPY is the symbol for Japanese Yen.
	SymbolJPY = Symbol("JPY")
	// SymbolFLR is the spot trading symbol for Flare.
	SymbolFLR = Symbol("FLR")
	// SymbolFIL is the spot trading symbol for Filecoin.
	SymbolFIL = Symbol("FIL")
	// SymbolSAND is the spot trading symbol for The Sandbox.
	SymbolSAND = Symbol("SAND")
	// SymbolCHZ is the spot trading symbol for Chiliz.
	SymbolCHZ = Symbol("CHZ")
	// SymbolSUI is the spot trading symbol for Sui.
	SymbolSUI = Symbol("SUI")
	// SymbolSUIJPY is the leveraged trading symbol for the SUI/JPY pair.
	SymbolSUIJPY = Symbol("SUI_JPY")
	// SymbolWILD is the spot trading symbol for Wilder World.
	SymbolWILD = Symbol("WILD")
	// SymbolZPG is the spot trading symbol for Zipangcoin.
	SymbolZPG = Symbol("ZPG")
	// SymbolZPGAG is the spot trading symbol for Zipangcoin Silver.
	SymbolZPGAG = Symbol("ZPGAG")
	// SymbolZPGPT is the spot trading symbol for Zipangcoin Platinum.
	SymbolZPGPT = Symbol("ZPGPT")
	// SymbolNONE is the zero value representing an unspecified symbol.
	SymbolNONE = Symbol("")
)

var allSymbols = []Symbol{
	SymbolBTC,
	SymbolETH,
	SymbolBCH,
	SymbolLTC,
	SymbolXRP,
	SymbolXLM,
	SymbolOMG,
	SymbolXTZ,
	SymbolDOT,
	SymbolATOM,
	SymbolDAI,
	SymbolFCR,
	SymbolADA,
	SymbolLINK,
	SymbolDOGE,
	SymbolSOL,
	SymbolASTR,
	SymbolAVAX,
	SymbolNAC,
	SymbolBTCJPY,
	SymbolETHJPY,
	SymbolBCHJPY,
	SymbolLTCJPY,
	SymbolXRPJPY,
	SymbolDOTJPY,
	SymbolATOMJPY,
	SymbolADAJPY,
	SymbolLINKJPY,
	SymbolDOGEJPY,
	SymbolSOLJPY,
	SymbolJPY,
	SymbolFLR,
	SymbolFIL,
	SymbolSAND,
	SymbolCHZ,
	SymbolSUI,
	SymbolSUIJPY,
	SymbolWILD,
	SymbolZPG,
	SymbolZPGAG,
	SymbolZPGPT,
}

func (c *Symbol) UnmarshalJSON(d []byte) error {
	symbol := Symbol(strings.ReplaceAll(string(d), `"`, ``))
	if !slices.Contains(allSymbols, symbol) {
		return ErrUnsupportedSymbol
	}
	*c = symbol
	return nil
}
