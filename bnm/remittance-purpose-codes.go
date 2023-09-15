/*
 * Project: Bank Negara Malaysia Master Data
 * Filename: /remittance-purpose-codes.go
 * Created Date: Thursday September 14th 2023 12:06:54 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Thursday September 14th 2023 12:42:30 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package bnm

var BNMRemitPurposeDescription = map[string]string{
	"06000": "MANUFACTURED GOODS",
	"07000": "MACHINERY, NON-CUSTOMISED PACKAGED SOFTWARE AND TRANSPORT EQUIPMENT",
	"08000": "MISCELLANEOUS MANUFACTURED ARTICLES",
	"10010": "GOODS FOR PROCESSING/MANUFACTURING SERVICES",
	"11200": "PASSENGER FARE",
	"13110": "GOODS AND SERVICES PURCHASE BY TRAVELLERS",
	"13210": "GOODS AND SERVICES PURCHASE THROUGH BUSINESS AND OFFICIAL TRAVEL",
	"13400": "TRAVEL FOR MEDICAL TREATMENT",
	"13500": "EDUCATION-RELATED",
	"14100": "INVESTMENT INCOME",
	"14310": "WAGES AND SALARIES IN CASH",
	"16210": "CONSTRUCTION AND INSTALLATION SERVICES IN MALAYSIA",
	"16791": "TRADE-RELATED SERVICES",
	"16820": "HEALTH SERVICES",
	"16830": "EDUCATION SERVICES",
	"16850": "OTHER PERSONAL SERVICES",
	"16999": "OTHER BUSINESS SERVICES NCE",
	"21210": "GRANTS AND GIFTS",
	"21220": "WORKERS' REMITTANCES",
	"31110": "TERM LOAN - LONG TERM",
	"31120": "TERM LOAN - SHORT TERM",
	"31410": "FINANCIAL LEASE - LONG TERM",
	"35000": "EQUITY CAPITAL",
	"39130": "PURCHASE/SALE OF REAL ESTATE IN HOST COUNTRY",
	"39210": "PLACEMENT/WITHDRAWAL OF DEPOSITS OF RESIDENTS WITH/FROM FINANCIAL INSTITUTIONS ABROAD",
	"39220": "PLACEMENT/WITHDRAWAL OF DEPOSITS OF RESIDENTS WITH/FROM OFFSHORE FINANCIAL INSTITUTIONS IN LABUAN",
}

var BNMRemitPurposeCode = map[string]string{
	"MANUFACTURED GOODS": "06000",
	"MACHINERY, NON-CUSTOMISED PACKAGED SOFTWARE AND TRANSPORT EQUIPMENT": "07000",
	"MISCELLANEOUS MANUFACTURED ARTICLES":                                 "08000",
	"GOODS FOR PROCESSING/MANUFACTURING SERVICES":                         "10010",
	"PASSENGER FARE": "11200",
	"GOODS AND SERVICES PURCHASE BY TRAVELLERS":                        "13110",
	"GOODS AND SERVICES PURCHASE THROUGH BUSINESS AND OFFICIAL TRAVEL": "13210",
	"TRAVEL FOR MEDICAL TREATMENT":                                     "13400",
	"EDUCATION-RELATED":                                                "13500",
	"INVESTMENT INCOME":                                                "14100",
	"WAGES AND SALARIES IN CASH":                                       "14310",
	"CONSTRUCTION AND INSTALLATION SERVICES IN MALAYSIA":               "16210",
	"TRADE-RELATED SERVICES":                                           "16791",
	"HEALTH SERVICES":                                                  "16820",
	"EDUCATION SERVICES":                                               "16830",
	"OTHER PERSONAL SERVICES":                                          "16850",
	"OTHER BUSINESS SERVICES NCE":                                      "16999",
	"GRANTS AND GIFTS":                                                 "21210",
	"WORKERS' REMITTANCES":                                             "21220",
	"TERM LOAN - LONG TERM":                                            "31110",
	"TERM LOAN - SHORT TERM":                                           "31120",
	"FINANCIAL LEASE - LONG TERM":                                      "31410",
	"EQUITY CAPITAL":                                                   "35000",
	"PURCHASE/SALE OF REAL ESTATE IN HOST COUNTRY":                     "39130",
	"PLACEMENT/WITHDRAWAL OF DEPOSITS OF RESIDENTS WITH/FROM FINANCIAL INSTITUTIONS ABROAD":             "39210",
	"PLACEMENT/WITHDRAWAL OF DEPOSITS OF RESIDENTS WITH/FROM OFFSHORE FINANCIAL INSTITUTIONS IN LABUAN": "39220",
}
