package examples

import "github.com/luizbranco/ofx"

var Ing = &ofx.OFX{
	SignOn: ofx.SignOn{
		SignOnResponse: ofx.SignOnResponse{
			Status: ofx.Status{
				Code:     0,
				Severity: ofx.Info,
				Message:  "Authentication Successful.",
			},
			DateTimeServer: ofx.DateTime("20080223132939.718[-5:EST]"),
			Language:       ofx.Language("ENG"),
			FinancialInstitution: ofx.FinancialInstitution{
				Organization: "INGDIRECTCanada",
				ID:           "10951",
			},
		},
	},
	Banking: ofx.Banking{
		BankingResponse: []ofx.BankingResponse{
			{
				TransactionUID: ofx.TransactionUID("0"),
				Status: ofx.Status{
					Code:     0,
					Severity: ofx.Info,
				},
				BankStatementResponse: ofx.BankStatementResponse{
					CurrencyDefault: ofx.CurrencySymbol("CAD"),
					BankingAccount: ofx.BankingAccount{
						BankID:      "0614",
						ID:          "123456",
						AccountType: ofx.Savings,
					},
					BankTransactionsList: ofx.BankTransactionsList{
						DateStart: ofx.Date("20031231190000.000[-5:EST]"),
						DateEnd:   ofx.Date("20080222190000.000[-5:EST]"),
						Transactions: []ofx.Transaction{
							{
								TransactionType: ofx.Credit,
								DatePosted:      ofx.DateTime("20050923120000.000"),
								Amount:          100,
								FITID:           "4",
								Name:            "Dépôt",
								Memo:            "81530219 3893476",
							},
						},
					},
					LedgerBalance: ofx.Balance{
						Amount:        4855.99,
						EffectiveDate: ofx.DateTime("20080223132935.218[-5:EST]"),
					},
					AvailableBalance: ofx.Balance{
						Amount:        4855.99,
						EffectiveDate: ofx.DateTime("20080223132935.218[-5:EST]"),
					},
				},
			},
		},
	},
}
