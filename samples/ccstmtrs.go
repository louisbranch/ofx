package samples

import "github.com/luizbranco/ofx"

var CCStmtRs = &ofx.OFX{
	SignOn: ofx.SignOn{
		SignOnResponse: ofx.SignOnResponse{
			Status: ofx.Status{
				Code:     0,
				Severity: ofx.Info,
			},
			DateTimeServer: ofx.DateTime("20130820000000"),
			Language:       ofx.Language("FRA"),
		},
	},
	CreditCard: ofx.CreditCard{
		CreditCardResponse: []ofx.CreditCardResponse{
			{
				TransactionUID: ofx.TransactionUID("20130820000000"),
				Status: ofx.Status{
					Code:     0,
					Severity: ofx.Info,
				},
				CreditCardStatementResponse: ofx.CreditCardStatementResponse{
					CurrencyDefault: ofx.CurrencySymbol("EUR"),
					CreditCardAccount: ofx.CreditCardAccount{
						ID: "4XXXXXXXXXXXXXX5",
					},
					BankTransactionsList: ofx.BankTransactionsList{
						DateStart: ofx.Date("20130620000000"),
						DateEnd:   ofx.Date("20130812000000"),
						Transactions: []ofx.Transaction{
							{
								TransactionType: ofx.PointOfSale,
								DatePosted:      ofx.DateTime("20130809"),
								DateUser:        ofx.DateTime("20130809"),
								Amount:          -27.24,
								FITID:           "LF5NP7L3LF",
								ReferenceNumber: "INTERMARCHE",
								Name:            "INTERMARCHE",
							},
							{
								TransactionType: ofx.PointOfSale,
								DatePosted:      ofx.DateTime("20130622"),
								DateUser:        ofx.DateTime("20130622"),
								Amount:          -45.60,
								FITID:           "LF5DOD0GLF",
								ReferenceNumber: "PIZZA D EL ARTE",
								Name:            "PIZZA D EL ARTE",
							},
						},
					},
					LedgerBalance: ofx.Balance{
						Amount:        -227.20,
						EffectiveDate: ofx.DateTime("20130812000000"),
					},
					AvailableBalance: ofx.Balance{
						Amount:        0,
						EffectiveDate: ofx.DateTime("20130812000000"),
					},
				},
			},
			{
				TransactionUID: ofx.TransactionUID("20130820000000"),
				Status: ofx.Status{
					Code:     0,
					Severity: ofx.Info,
				},
				CreditCardStatementResponse: ofx.CreditCardStatementResponse{
					CurrencyDefault: ofx.CurrencySymbol("EUR"),
					CreditCardAccount: ofx.CreditCardAccount{
						ID: "4XXXXXXXXXXXXXX9",
					},
					BankTransactionsList: ofx.BankTransactionsList{
						DateStart: ofx.Date("20130622000000"),
						DateEnd:   ofx.Date("20130819000000"),
						Transactions: []ofx.Transaction{
							{
								TransactionType: ofx.PointOfSale,
								DatePosted:      ofx.DateTime("20130722"),
								DateUser:        ofx.DateTime("20130722"),
								Amount:          -4.99,
								FITID:           "LF5N_%FELO",
								ReferenceNumber: "ELSWORD ORG GAME",
								Name:            "ELSWORD ORG GAME",
							},
							{
								TransactionType: ofx.PointOfSale,
								DatePosted:      ofx.DateTime("20130727"),
								DateUser:        ofx.DateTime("20130727"),
								Amount:          -84.30,
								FITID:           "LF5N_M5ALO",
								ReferenceNumber: "AG DISTRIBUTION",
								Name:            "AG DISTRIBUTION",
							},
							{
								TransactionType: ofx.PointOfSale,
								DatePosted:      ofx.DateTime("20130727"),
								DateUser:        ofx.DateTime("20130727"),
								Amount:          -151.10,
								FITID:           "LF5N_M5ALF",
								ReferenceNumber: "E LECLERC",
								Name:            "E LECLERC",
							},
						},
					},
					LedgerBalance: ofx.Balance{
						Amount:        -334.43,
						EffectiveDate: ofx.DateTime("20130819000000"),
					},
					AvailableBalance: ofx.Balance{
						Amount:        0.0,
						EffectiveDate: ofx.DateTime("20130819000000"),
					},
				},
			},
		},
	},
}
