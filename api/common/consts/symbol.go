package consts

import (
	"slices"
	"strings"
)

// Symbol ...
type Symbol string

const (
	// SymbolBTC ...
	SymbolBTC = Symbol("BTC")
	// SymbolETH ...
	SymbolETH = Symbol("ETH")
	// SymbolBCH ...
	SymbolBCH = Symbol("BCH")
	// SymbolLTC ...
	SymbolLTC = Symbol("LTC")
	// SymbolXRP ...
	SymbolXRP = Symbol("XRP")
	// SymbolXLM ...
	SymbolXLM = Symbol("XLM")
	// SymbolOMG ...
	SymbolOMG = Symbol("OMG")
	// SymbolXTZ ...
	SymbolXTZ = Symbol("XTZ")
	// SymbolDOT ...
	SymbolDOT = Symbol("DOT")
	// SymbolATOM ...
	SymbolATOM = Symbol("ATOM")
	// SymbolDAI ...
	SymbolDAI = Symbol("DAI")
	// SymbolFCR ...
	SymbolFCR = Symbol("FCR")
	// SymbolADA ...
	SymbolADA = Symbol("ADA")
	// SymbolLINK ...
	SymbolLINK = Symbol("LINK")
	// SymbolDOGE ...
	SymbolDOGE = Symbol("DOGE")
	// SymbolSOL ...
	SymbolSOL = Symbol("SOL")
	// SymbolASTR ...
	SymbolASTR = Symbol("ASTR")
	// SymbolAVAX ...
	SymbolAVAX = Symbol("AVAX")
	// SymbolNAC ...
	SymbolNAC = Symbol("NAC")
	// SymbolBTCJPY ...
	SymbolBTCJPY = Symbol("BTC_JPY")
	// SymbolETHJPY ...
	SymbolETHJPY = Symbol("ETH_JPY")
	// SymbolBCHJPY ...
	SymbolBCHJPY = Symbol("BCH_JPY")
	// SymbolLTCJPY ...
	SymbolLTCJPY = Symbol("LTC_JPY")
	// SymbolXRPJPY ...
	SymbolXRPJPY = Symbol("XRP_JPY")
	// SymbolDOTJPY ...
	SymbolDOTJPY = Symbol("DOT_JPY")
	// SymbolATOMJPY ...
	SymbolATOMJPY = Symbol("ATOM_JPY")
	// SymbolADAJPY ...
	SymbolADAJPY = Symbol("ADA_JPY")
	// SymbolLINKJPY ...
	SymbolLINKJPY = Symbol("LINK_JPY")
	// SymbolDOGEJPY ...
	SymbolDOGEJPY = Symbol("DOGE_JPY")
	// SymbolSOLJPY ...
	SymbolSOLJPY = Symbol("SOL_JPY")
	// SymbolJPY ...
	SymbolJPY = Symbol("JPY")
	// SymbolFLR ...
	SymbolFLR = Symbol("FLR")
	// SymbolFIL ...
	SymbolFIL = Symbol("FIL")
	// SymbolSAND ...
	SymbolSAND = Symbol("SAND")
	// SymbolCHZ ...
	SymbolCHZ = Symbol("CHZ")
	// SymbolNONE ...
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
}

func (c *Symbol) UnmarshalJSON(d []byte) error {
	symbol := Symbol(strings.ReplaceAll(string(d), `"`, ``))
	if !slices.Contains(allSymbols, symbol) {
		return ErrUnsupportedSymbol
	}
	*c = symbol
	return nil
}
