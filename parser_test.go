package ofx

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func TestParse(t *testing.T) {

	egs := []struct {
		file string
		res  *OFX
	}{
		{"ccstmtrs.ofx", CCStmtRs},
		{"desjardins.ofx", Desjardins},
		{"desjardins2.ofx", Desjardins2},
		{"desjardins3.ofx", Desjardins3},
		{"response.ofx", Response},
		{"ing.qfx", Ing},
		{"fortis.ofx", Fortis},
	}

	for _, eg := range egs {
		in, err := os.Open(path.Join("samples", eg.file))
		if err != nil {
			t.Fatal(err)
		}

		res, err := Parse(in)

		if err != nil {
			t.Errorf("Parse unexpected error file %s, %s", eg.file, err)
		}
		if !reflect.DeepEqual(eg.res, res) {
			t.Errorf("Parse expected file %s\n", eg.file)
			spew.Dump(eg.res, res)
		}
	}
}
func TestEndTags(t *testing.T) {
	egs := []bool{true, true, true}

	for i, ok := range egs {
		var buf bytes.Buffer
		in, _ := os.Open(path.Join("samples", fmt.Sprintf("in_%d.ofx", i)))
		f, _ := os.Open(path.Join("samples", fmt.Sprintf("out_%d.ofx", i)))

		tr := charmap.Windows1252
		r := transform.NewReader(f, tr.NewDecoder())
		out, _ := ioutil.ReadAll(r)

		for _, l := range bytes.Split(out, []byte("\n")) {
			buf.Write(bytes.TrimSpace(l))
		}
		out = buf.Bytes()

		res, err := endTags(in)

		if ok {
			if err != nil {
				t.Errorf("EndTags unexpected error file %d, %s", i, err)
			}
			if !bytes.Equal(out, res) {
				t.Errorf("EndTags expected file %d\n%s\ngot\n%s", i, out, res)
			}
		} else {
			if err == nil {
				t.Errorf("EndTags expected error file %d, got none", i)
			}
		}
	}
}

var CCStmtRs = &OFX{
	SignOn: SignOn{
		SignOnResponse: SignOnResponse{
			Status: Status{
				Code:     0,
				Severity: Info,
			},
			DateTimeServer: DateTime("20130820000000"),
			Language:       Language("FRA"),
		},
	},
	CreditCard: CreditCard{
		CreditCardResponse: []CreditCardResponse{
			{
				TransactionUID: TransactionUID("20130820000000"),
				Status: Status{
					Code:     0,
					Severity: Info,
				},
				CreditCardStatementResponse: CreditCardStatementResponse{
					CurrencyDefault: Currency("EUR"),
					CreditCardAccount: CreditCardAccount{
						ID: "4XXXXXXXXXXXXXX5",
					},
					BankTransactionsList: BankTransactionsList{
						DateStart: Date("20130620000000"),
						DateEnd:   Date("20130812000000"),
						Transactions: []Transaction{
							{
								TransactionType: PointOfSale,
								DatePosted:      DateTime("20130809"),
								DateUser:        DateTime("20130809"),
								Amount:          -27.24,
								FITID:           "LF5NP7L3LF",
								ReferenceNumber: "INTERMARCHE",
								Name:            "INTERMARCHE",
							},
							{
								TransactionType: PointOfSale,
								DatePosted:      DateTime("20130622"),
								DateUser:        DateTime("20130622"),
								Amount:          -45.60,
								FITID:           "LF5DOD0GLF",
								ReferenceNumber: "PIZZA D EL ARTE",
								Name:            "PIZZA D EL ARTE",
							},
						},
					},
					LedgerBalance: Balance{
						Amount:        -227.20,
						EffectiveDate: DateTime("20130812000000"),
					},
					AvailableBalance: Balance{
						Amount:        0,
						EffectiveDate: DateTime("20130812000000"),
					},
				},
			},
			{
				TransactionUID: TransactionUID("20130820000000"),
				Status: Status{
					Code:     0,
					Severity: Info,
				},
				CreditCardStatementResponse: CreditCardStatementResponse{
					CurrencyDefault: Currency("EUR"),
					CreditCardAccount: CreditCardAccount{
						ID: "4XXXXXXXXXXXXXX9",
					},
					BankTransactionsList: BankTransactionsList{
						DateStart: Date("20130622000000"),
						DateEnd:   Date("20130819000000"),
						Transactions: []Transaction{
							{
								TransactionType: PointOfSale,
								DatePosted:      DateTime("20130722"),
								DateUser:        DateTime("20130722"),
								Amount:          -4.99,
								FITID:           "LF5N_%FELO",
								ReferenceNumber: "ELSWORD ORG GAME",
								Name:            "ELSWORD ORG GAME",
							},
							{
								TransactionType: PointOfSale,
								DatePosted:      DateTime("20130727"),
								DateUser:        DateTime("20130727"),
								Amount:          -84.30,
								FITID:           "LF5N_M5ALO",
								ReferenceNumber: "AG DISTRIBUTION",
								Name:            "AG DISTRIBUTION",
							},
							{
								TransactionType: PointOfSale,
								DatePosted:      DateTime("20130727"),
								DateUser:        DateTime("20130727"),
								Amount:          -151.10,
								FITID:           "LF5N_M5ALF",
								ReferenceNumber: "E LECLERC",
								Name:            "E LECLERC",
							},
						},
					},
					LedgerBalance: Balance{
						Amount:        -334.43,
						EffectiveDate: DateTime("20130819000000"),
					},
					AvailableBalance: Balance{
						Amount:        0.0,
						EffectiveDate: DateTime("20130819000000"),
					},
				},
			},
		},
	},
}

var Desjardins = &OFX{
	SignOn: SignOn{
		SignOnResponse: SignOnResponse{
			Status: Status{
				Code:     0,
				Severity: Info,
				Message:  "OK",
			},
			DateTimeServer: DateTime("20080212"),
			UserKey:        "61775520C2D1F127",
			Language:       Language("ENG"),
		},
	},
	Banking: Banking{
		BankingResponse: []BankingResponse{
			{
				TransactionUID: TransactionUID("DESJ-2008021216022711068"),
				Status: Status{
					Code:     0,
					Severity: Info,
					Message:  "OK",
				},
				BankStatementResponse: BankStatementResponse{
					CurrencyDefault: Currency("CAD"),
					BankingAccount: BankingAccount{
						BankID:      "700000100",
						BranchID:    "0389347",
						ID:          "815-30219-12345-EOP",
						AccountType: Checking,
					},
					BankTransactionsList: BankTransactionsList{
						DateStart: Date("20080113000000"),
						DateEnd:   Date("20080212000000"),
						Transactions: []Transaction{
							{
								TransactionType: Check,
								DatePosted:      DateTime("20080114"),
								Amount:          -100,
								FITID:           "Th3DJACES",
								CheckNumber:     "7",
								Name:            "Chèque/",
								Memo:            "DCN",
							},
							{
								TransactionType: Debit,
								DatePosted:      DateTime("20080118"),
								Amount:          -200,
								FITID:           "XwT2KAOtO",
								Name:            "Retrait au GA/CP CREMAZIE DE MTL",
								Memo:            "RGA",
							},
							{
								TransactionType: Debit,
								DatePosted:      DateTime("20080122"),
								Amount:          -100.14,
								FITID:           "XN5ONAQET",
								Name:            "Paiement facture - AccèsD Intern",
								Memo:            "PWW",
							},
							{
								TransactionType: Debit,
								DatePosted:      DateTime("20080122"),
								Amount:          -1090.18,
								FITID:           "6Z5ONAQSd",
								Name:            "Paiement facture - AccèsD Intern",
								Memo:            "PWW",
							},
							{
								TransactionType: Check,
								DatePosted:      DateTime("20080122"),
								Amount:          -93.24,
								FITID:           "TsC1NACD6",
								CheckNumber:     "9",
								Name:            "Chèque/",
								Memo:            "DCN",
							},
							{
								TransactionType: Check,
								DatePosted:      DateTime("20080122"),
								Amount:          -600,
								FITID:           "UsC1NACD6",
								CheckNumber:     "8",
								Name:            "Chèque/",
								Memo:            "DCN",
							},
							{
								TransactionType: Credit,
								DatePosted:      DateTime("20080123"),
								Amount:          1000,
								FITID:           "CtC1NAPjr",
								Name:            "Placement/ING",
								Memo:            "DI",
							},
							{
								TransactionType: Credit,
								DatePosted:      DateTime("20080128"),
								Amount:          700,
								FITID:           "gnxzQAPk2",
								Name:            "Placement/ING",
								Memo:            "DI",
							},
							{
								TransactionType: Check,
								DatePosted:      DateTime("20080128"),
								Amount:          -170,
								FITID:           "UH7ZRAEYE",
								CheckNumber:     "10",
								Name:            "Chèque/",
								Memo:            "DCN",
							},
							{
								TransactionType: Credit,
								DatePosted:      DateTime("20080131"),
								Amount:          0.13,
								FITID:           "PBOmSAWsd",
								Name:            "Intérêt sur EOP/",
								Memo:            "INT",
							},
							{
								TransactionType: Credit,
								DatePosted:      DateTime("20080201"),
								Amount:          2613.03,
								FITID:           ":WXMTAMse",
								Name:            "Dépôt au comptoir/",
								Memo:            "DSL",
							},
							{
								TransactionType: Debit,
								DatePosted:      DateTime("20080201"),
								Amount:          -200,
								FITID:           "VRXMTAMun",
								Name:            "Retrait au GA/C.D.S. STE-CECILE",
								Memo:            "RGA",
							},
							{
								TransactionType: Debit,
								DatePosted:      DateTime("20080201"),
								Amount:          -173.65,
								FITID:           "YrgyTATGw",
								Name:            "Virement-remboursement/à PR 01",
								Memo:            "VAP",
							},
							{
								TransactionType: Check,
								DatePosted:      DateTime("20080204"),
								Amount:          -100,
								FITID:           "M38kVACFZ",
								CheckNumber:     "11",
								Name:            "Chèque/",
								Memo:            "DCN",
							},
							{
								TransactionType: Check,
								DatePosted:      DateTime("20080205"),
								Amount:          -1450,
								FITID:           "pYGLWAEO3",
								CheckNumber:     "2",
								Name:            "Chèque/",
								Memo:            "DCN",
							},
							{
								TransactionType: Debit,
								DatePosted:      DateTime("20080208"),
								Amount:          -47.41,
								FITID:           "vKZXXAQEM",
								Name:            "Achat/STATION CENTRALE D'AUT",
								Memo:            "ACH",
							},
							{
								TransactionType: Debit,
								DatePosted:      DateTime("20080211"),
								Amount:          -18.97,
								FITID:           "tV1JZASI1",
								Name:            "Assurance/MELOCHE MONNEX/SECURIT",
								Memo:            "RA",
							},
						},
					},
					LedgerBalance: Balance{
						Amount:        3925.84,
						EffectiveDate: DateTime("20080212160227"),
					},
					AvailableBalance: Balance{
						Amount:        3925.84,
						EffectiveDate: DateTime("20080212160227"),
					},
				},
			},
			{
				TransactionUID: TransactionUID("DESJ-2008021216022711068"),
				Status: Status{
					Code:     0,
					Severity: Info,
					Message:  "OK",
				},
				BankStatementResponse: BankStatementResponse{
					CurrencyDefault: Currency("CAD"),
					BankingAccount: BankingAccount{
						BankID:      "700000100",
						BranchID:    "0389347",
						ID:          "815-30219-54321-ES1",
						AccountType: Savings,
					},
					BankTransactionsList: BankTransactionsList{
						DateStart: Date("20080113000000"),
						DateEnd:   Date("20080212000000"),
					},
					LedgerBalance: Balance{
						Amount:        0,
						EffectiveDate: DateTime("20080212160227"),
					},
					AvailableBalance: Balance{
						Amount:        0,
						EffectiveDate: DateTime("20080212160227"),
					},
				},
			},
			{
				TransactionUID: TransactionUID("DESJ-2008021216022711068"),
				Status: Status{
					Code:     0,
					Severity: Info,
					Message:  "OK",
				},
				BankStatementResponse: BankStatementResponse{
					CurrencyDefault: Currency("USD"),
					BankingAccount: BankingAccount{
						BankID:      "1234",
						BranchID:    "5678",
						ID:          "815-30219-11111-EOP",
						AccountType: Checking,
					},
					BankTransactionsList: BankTransactionsList{
						DateStart: Date("20080113000000"),
						DateEnd:   Date("20080212000000"),
						Transactions: []Transaction{
							{
								TransactionType: Credit,
								DatePosted:      DateTime("20080131"),
								Amount:          0.02,
								FITID:           "e7NmSAWtn",
								Name:            "Intérêt sur EOP/",
								Memo:            "INT",
							},
							{
								TransactionType: Credit,
								DatePosted:      DateTime("20080201"),
								Amount:          5029.50,
								FITID:           "qWXMTAMiq",
								Name:            "Dépôt au comptoir/",
								Memo:            "DSL",
							},
							{
								TransactionType: Debit,
								DatePosted:      DateTime("20080201"),
								Amount:          -2665,
								FITID:           "uWXMTAMqA",
								Name:            "Retrait au comptoir/",
								Memo:            "RSL",
							},
						},
					},
					LedgerBalance: Balance{
						Amount:        3046.90,
						EffectiveDate: DateTime("20080212160227"),
					},
					AvailableBalance: Balance{
						Amount:        3046.90,
						EffectiveDate: DateTime("20080212160227"),
					},
				},
			},
		},
	},
}

var Desjardins2 = &OFX{
	SignOn: SignOn{
		SignOnResponse: SignOnResponse{
			Status: Status{
				Code:     0,
				Severity: Info,
				Message:  "OK",
			},
			DateTimeServer: DateTime("20080212"),
			UserKey:        "61775520C2D1F127",
			Language:       Language("ENG"),
		},
	},
	Banking: Banking{
		BankingResponse: []BankingResponse{
			{
				TransactionUID: TransactionUID("DESJ-123454563456"),
				Status: Status{
					Code:     0,
					Severity: Info,
					Message:  "OK",
				},
				BankStatementResponse: BankStatementResponse{
					CurrencyDefault: Currency("CAD"),
					BankingAccount: BankingAccount{
						BankID:      "700000100",
						BranchID:    "0389347",
						ID:          "Another account",
						AccountType: Savings,
					},
					BankTransactionsList: BankTransactionsList{
						DateStart: Date("20080113000000"),
						DateEnd:   Date("20080212000000"),
					},
					LedgerBalance: Balance{
						Amount:        0,
						EffectiveDate: DateTime("20080212160227"),
					},
					AvailableBalance: Balance{
						Amount:        0,
						EffectiveDate: DateTime("20080212160227"),
					},
				},
			},
			{
				TransactionUID: TransactionUID("DESJ-2008021216022711068"),
				Status: Status{
					Code:     0,
					Severity: Info,
					Message:  "OK",
				},
				BankStatementResponse: BankStatementResponse{
					CurrencyDefault: Currency("USD"),
					BankingAccount: BankingAccount{
						BankID:      "1234",
						BranchID:    "5678",
						ID:          "815-30219-11111-EOP",
						AccountType: Checking,
					},
					BankTransactionsList: BankTransactionsList{
						DateStart: Date("20080113000000"),
						DateEnd:   Date("20080212000000"),
						Transactions: []Transaction{
							{
								TransactionType: Debit,
								DatePosted:      DateTime("20080201"),
								Amount:          -2600,
								FITID:           "uWXMTAMqA",
								Name:            "Retrait au comptoir/",
								Memo:            "RSL",
							},
							{
								TransactionType: Credit,
								DatePosted:      DateTime("20080229"),
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

var Desjardins3 = &OFX{
	SignOn: SignOn{
		SignOnResponse: SignOnResponse{
			Status: Status{
				Code:     0,
				Severity: Info,
				Message:  "OK",
			},
			DateTimeServer: DateTime("20080212"),
			UserKey:        "61775520C2D1F127",
			Language:       Language("ENG"),
		},
	},
	Banking: Banking{
		BankingResponse: []BankingResponse{
			{
				TransactionUID: TransactionUID("DESJ-2008021216022711068"),
				Status: Status{
					Code:     0,
					Severity: Info,
					Message:  "OK",
				},
				BankStatementResponse: BankStatementResponse{
					CurrencyDefault: Currency("USD"),
					BankingAccount: BankingAccount{
						BankID:      "1234",
						BranchID:    "5678",
						ID:          "NEW_ACCOUNT",
						AccountType: Checking,
					},
					BankTransactionsList: BankTransactionsList{
						DateStart: Date("20080113000000"),
						DateEnd:   Date("20080212000000"),
						Transactions: []Transaction{
							{
								TransactionType: Credit,
								DatePosted:      DateTime("20080229"),
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

var Fortis = &OFX{
	SignOn: SignOn{
		SignOnResponse: SignOnResponse{
			Status: Status{
				Code:     0,
				Severity: Info,
			},
			DateTimeServer: DateTime("20080907"),
			Language:       Language("ENG"),
			FinancialInstitution: FinancialInstitution{
				Organization: "FORTIS",
				ID:           "FORTIS BANK",
			},
		},
	},
	Banking: Banking{
		BankingResponse: []BankingResponse{
			{
				TransactionUID: TransactionUID("1"),
				Status: Status{
					Code:     0,
					Severity: Info,
				},
				BankStatementResponse: BankStatementResponse{
					CurrencyDefault: Currency("EUR"),
					BankingAccount: BankingAccount{
						BankID:      "FORTIS",
						ID:          "001-5587496-84",
						AccountType: Checking,
					},
					BankTransactionsList: BankTransactionsList{
						DateStart: Date("20080627"),
						DateEnd:   Date("20080905"),
						Transactions: []Transaction{
							{
								TransactionType: Debit,
								DatePosted:      DateTime("20080905"),
								Amount:          -2.6,
								FITID:           "20080026",
								CheckNumber:     "20080026",
								Name:            "REDEVANCE MENSUELLE",
								Memo:            "REDEVANCE MENSUELLE                      01-09            2,60-EASY PACK (SINGLE) EXECUTE LE 04-09 DATE VALEUR : 01/09/2008",
							},
						},
					},
					LedgerBalance: Balance{
						Amount:        7490.73,
						EffectiveDate: DateTime("20080905"),
					},
				},
			},
		},
	},
}

var Ing = &OFX{
	SignOn: SignOn{
		SignOnResponse: SignOnResponse{
			Status: Status{
				Code:     0,
				Severity: Info,
				Message:  "Authentication Successful.",
			},
			DateTimeServer: DateTime("20080223132939.718[-5:EST]"),
			Language:       Language("ENG"),
			FinancialInstitution: FinancialInstitution{
				Organization: "INGDIRECTCanada",
				ID:           "10951",
			},
		},
	},
	Banking: Banking{
		BankingResponse: []BankingResponse{
			{
				TransactionUID: TransactionUID("0"),
				Status: Status{
					Code:     0,
					Severity: Info,
				},
				BankStatementResponse: BankStatementResponse{
					CurrencyDefault: Currency("CAD"),
					BankingAccount: BankingAccount{
						BankID:      "0614",
						ID:          "123456",
						AccountType: Savings,
					},
					BankTransactionsList: BankTransactionsList{
						DateStart: Date("20031231190000.000[-5:EST]"),
						DateEnd:   Date("20080222190000.000[-5:EST]"),
						Transactions: []Transaction{
							{
								TransactionType: Credit,
								DatePosted:      DateTime("20050923120000.000"),
								Amount:          100,
								FITID:           "4",
								Name:            "Dépôt",
								Memo:            "81530219 3893476",
							},
						},
					},
					LedgerBalance: Balance{
						Amount:        4855.99,
						EffectiveDate: DateTime("20080223132935.218[-5:EST]"),
					},
					AvailableBalance: Balance{
						Amount:        4855.99,
						EffectiveDate: DateTime("20080223132935.218[-5:EST]"),
					},
				},
			},
		},
	},
}

var Response = &OFX{
	SignOn: SignOn{
		SignOnResponse: SignOnResponse{
			Status: Status{
				Code:     0,
				Severity: Info,
			},
			DateTimeServer: DateTime("20071015021529.000[-8:PST]"),
			Language:       Language("ENG"),
			DateAccountUp:  DateTime("19900101000000"),
			FinancialInstitution: FinancialInstitution{
				Organization: "MYBANK",
				ID:           "01234",
			},
		},
	},
	Banking: Banking{
		BankingResponse: []BankingResponse{
			{
				TransactionUID: TransactionUID("23382938"),
				Status: Status{
					Code:     0,
					Severity: Info,
				},
				BankStatementResponse: BankStatementResponse{
					CurrencyDefault: Currency("USD"),
					BankingAccount: BankingAccount{
						BankID:      "987654321",
						ID:          "098-121",
						AccountType: Savings,
					},
					BankTransactionsList: BankTransactionsList{
						DateStart: Date("20070101"),
						DateEnd:   Date("20071015"),
						Transactions: []Transaction{
							{
								TransactionType: Credit,
								DatePosted:      DateTime("20070315"),
								DateUser:        DateTime("20070315"),
								Amount:          200,
								FITID:           "980315001",
								Name:            "DEPOSIT",
								Memo:            "automatic deposit",
							},
							{
								TransactionType: Credit,
								DatePosted:      DateTime("20070329"),
								DateUser:        DateTime("20070329"),
								Amount:          150,
								FITID:           "980310001",
								Name:            "TRANSFER",
								Memo:            "Transfer from checking",
							},
							{
								TransactionType: EletronicPayment,
								DatePosted:      DateTime("20070709"),
								DateUser:        DateTime("20070709"),
								Amount:          -100,
								FITID:           "980309001",
								CheckNumber:     "1025",
								Name:            "John Hancock",
							},
						},
					},
					LedgerBalance: Balance{
						Amount:        5250,
						EffectiveDate: DateTime("20071015021529.000[-8:PST]"),
					},
					AvailableBalance: Balance{
						Amount:        5250,
						EffectiveDate: DateTime("20071015021529.000[-8:PST]"),
					},
				},
			},
		},
	},
}
