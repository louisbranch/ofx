package samples

import "github.com/luizbranco/ofx"

var Fortis = &ofx.OFX{
	SignOn: ofx.SignOn{
		SignOnResponse: ofx.SignOnResponse{
			Status: ofx.Status{
				Code:     0,
				Severity: ofx.Info,
			},
			DateTimeServer: ofx.DateTime("20080907"),
			Language:       ofx.Language("ENG"),
			FinancialInstitution: ofx.FinancialInstitution{
				Organization: "FORTIS",
				ID:           "FORTIS BANK",
			},
		},
	},
	Banking: ofx.Banking{
		BankingResponse: []ofx.BankingResponse{
			{
				TransactionUID: ofx.TransactionUID("1"),
				Status: ofx.Status{
					Code:     0,
					Severity: ofx.Info,
				},
				BankStatementResponse: ofx.BankStatementResponse{
					CurrencyDefault: ofx.Currency("EUR"),
					BankingAccount: ofx.BankingAccount{
						BankID:      "FORTIS",
						ID:          "001-5587496-84",
						AccountType: ofx.Checking,
					},
					BankTransactionsList: ofx.BankTransactionsList{
						DateStart: ofx.Date("20080627"),
						DateEnd:   ofx.Date("20080905"),
						Transactions: []ofx.Transaction{
							{
								TransactionType: ofx.Debit,
								DatePosted:      ofx.DateTime("20080905"),
								Amount:          -2.6,
								FITID:           "20080026",
								CheckNumber:     "20080026",
								Name:            "REDEVANCE MENSUELLE",
								Memo:            "REDEVANCE MENSUELLE                      01-09            2,60-EASY PACK (SINGLE) EXECUTE LE 04-09 DATE VALEUR : 01/09/2008",
							},
						},
					},
					LedgerBalance: ofx.Balance{
						Amount:        7490.73,
						EffectiveDate: ofx.DateTime("20080905"),
					},
				},
			},
		},
	},
}
