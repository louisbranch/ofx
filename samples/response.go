package samples

import "github.com/luizbranco/ofx"

var Response = &ofx.OFX{
	SignOn: ofx.SignOn{
		SignOnResponse: ofx.SignOnResponse{
			Status: ofx.Status{
				Code:     0,
				Severity: ofx.Info,
			},
			DateTimeServer: ofx.DateTime("20071015021529.000[-8:PST]"),
			Language:       ofx.Language("ENG"),
			DateAccountUp:  ofx.DateTime("19900101000000"),
			FinancialInstitution: ofx.FinancialInstitution{
				Organization: "MYBANK",
				ID:           "01234",
			},
		},
	},
	Banking: ofx.Banking{
		BankingResponse: []ofx.BankingResponse{
			{
				TransactionUID: ofx.TransactionUID("23382938"),
				Status: ofx.Status{
					Code:     0,
					Severity: ofx.Info,
				},
				BankStatementResponse: ofx.BankStatementResponse{
					CurrencyDefault: ofx.Currency("USD"),
					BankingAccount: ofx.BankingAccount{
						BankID:      "987654321",
						ID:          "098-121",
						AccountType: ofx.Savings,
					},
					BankTransactionsList: ofx.BankTransactionsList{
						DateStart: ofx.Date("20070101"),
						DateEnd:   ofx.Date("20071015"),
						Transactions: []ofx.Transaction{
							{
								TransactionType: ofx.Credit,
								DatePosted:      ofx.DateTime("20070315"),
								DateUser:        ofx.DateTime("20070315"),
								Amount:          200,
								FITID:           "980315001",
								Name:            "DEPOSIT",
								Memo:            "automatic deposit",
							},
							{
								TransactionType: ofx.Credit,
								DatePosted:      ofx.DateTime("20070329"),
								DateUser:        ofx.DateTime("20070329"),
								Amount:          150,
								FITID:           "980310001",
								Name:            "TRANSFER",
								Memo:            "Transfer from checking",
							},
							{
								TransactionType: ofx.EletronicPayment,
								DatePosted:      ofx.DateTime("20070709"),
								DateUser:        ofx.DateTime("20070709"),
								Amount:          -100,
								FITID:           "980309001",
								CheckNumber:     "1025",
								Name:            "John Hancock",
							},
						},
					},
					LedgerBalance: ofx.Balance{
						Amount:        5250,
						EffectiveDate: ofx.DateTime("20071015021529.000[-8:PST]"),
					},
					AvailableBalance: ofx.Balance{
						Amount:        5250,
						EffectiveDate: ofx.DateTime("20071015021529.000[-8:PST]"),
					},
				},
			},
		},
	},
}
