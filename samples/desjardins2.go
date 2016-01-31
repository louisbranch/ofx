package samples

import "github.com/luizbranco/ofx"

var Desjardins2 = &ofx.OFX{
	SignOn: ofx.SignOn{
		SignOnResponse: ofx.SignOnResponse{
			Status: ofx.Status{
				Code:     0,
				Severity: ofx.Info,
				Message:  "OK",
			},
			DateTimeServer: ofx.DateTime("20080212"),
			UserKey:        "61775520C2D1F127",
			Language:       ofx.Language("ENG"),
		},
	},
	Banking: ofx.Banking{
		BankingResponse: []ofx.BankingResponse{
			{
				TransactionUID: ofx.TransactionUID("DESJ-123454563456"),
				Status: ofx.Status{
					Code:     0,
					Severity: ofx.Info,
					Message:  "OK",
				},
				BankStatementResponse: ofx.BankStatementResponse{
					CurrencyDefault: ofx.Currency("CAD"),
					BankingAccount: ofx.BankingAccount{
						BankID:      "700000100",
						BranchID:    "0389347",
						ID:          "Another account",
						AccountType: ofx.Savings,
					},
					BankTransactionsList: ofx.BankTransactionsList{
						DateStart: ofx.Date("20080113000000"),
						DateEnd:   ofx.Date("20080212000000"),
					},
					LedgerBalance: ofx.Balance{
						Amount:        0,
						EffectiveDate: ofx.DateTime("20080212160227"),
					},
					AvailableBalance: ofx.Balance{
						Amount:        0,
						EffectiveDate: ofx.DateTime("20080212160227"),
					},
				},
			},
			{
				TransactionUID: ofx.TransactionUID("DESJ-2008021216022711068"),
				Status: ofx.Status{
					Code:     0,
					Severity: ofx.Info,
					Message:  "OK",
				},
				BankStatementResponse: ofx.BankStatementResponse{
					CurrencyDefault: ofx.Currency("USD"),
					BankingAccount: ofx.BankingAccount{
						BankID:      "1234",
						BranchID:    "5678",
						ID:          "815-30219-11111-EOP",
						AccountType: ofx.Checking,
					},
					BankTransactionsList: ofx.BankTransactionsList{
						DateStart: ofx.Date("20080113000000"),
						DateEnd:   ofx.Date("20080212000000"),
						Transactions: []ofx.Transaction{
							{
								TransactionType: ofx.Debit,
								DatePosted:      ofx.DateTime("20080201"),
								Amount:          -2600,
								FITID:           "uWXMTAMqA",
								Name:            "Retrait au comptoir/",
								Memo:            "RSL",
							},
							{
								TransactionType: ofx.Credit,
								DatePosted:      ofx.DateTime("20080229"),
								Amount:          300,
								FITID:           "abcdef",
								Name:            "Added entry",
								Memo:            "DSL",
							},
						},
					},
				},
			},
		},
	},
}
