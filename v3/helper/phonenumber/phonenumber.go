package phonenumber

import (
	"fmt"
	"strings"
)

type Country struct {
	ISOCode     string
	CountryCode string
	Name        string
}

var countries = []Country{
	{"ABK", "7", "Abkhazia"},
	{"AFG", "93", "Afghanistan"},
	{"ALB", "355", "Albania"},
	{"DZA", "213", "Algeria"},
	{"ASM", "1", "American Samoa"},
	{"AND", "376", "Andorra"},
	{"AGO", "244", "Angola"},
	{"AIA", "1264", "Anguilla"},
	{"ATA", "672", "Antarctica"},
	{"ATG", "1268", "Antigua and Barbuda"},
	{"ARG", "54", "Argentina"},
	{"ARM", "374", "Armenia"},
	{"ABW", "297", "Aruba"},
	{"AUS", "61", "Australia"},
	{"AUT", "43", "Austria"},
	{"AZE", "994", "Azerbaijan"},
	{"BHS", "1242", "Bahamas"},
	{"BHR", "973", "Bahrain"},
	{"BGD", "880", "Bangladesh"},
	{"BRB", "1246", "Barbados"},
	{"BLR", "375", "Belarus"},
	{"BEL", "32", "Belgium"},
	{"BLZ", "501", "Belize"},
	{"BEN", "229", "Benin"},
	{"BMU", "1441", "Bermuda"},
	{"BTN", "975", "Bhutan"},
	{"BOL", "591", "Bolivia"},
	{"BIH", "387", "Bosnia and Herzegovina"},
	{"BWA", "267", "Botswana"},
	{"BRA", "55", "Brazil"},
	{"BRN", "673", "Brunei Darussalam"},
	{"BGR", "359", "Bulgaria"},
	{"BFA", "226", "Burkina Faso"},
	{"BDI", "257", "Burundi"},
	{"KHM", "855", "Cambodia"},
	{"CMR", "237", "Cameroon"},
	{"CAN", "1", "Canada"},
	{"CPV", "238", "Cape Verde"},
	{"CYM", "1345", "Cayman Islands"},
	{"CAF", "236", "Central African Republic"},
	{"TCD", "235", "Chad"},
	{"CHL", "56", "Chile"},
	{"CHN", "86", "China"},
	{"COL", "57", "Colombia"},
	{"COM", "269", "Comoros"},
	{"COG", "242", "Congo"},
	{"COK", "682", "Cook Islands"},
	{"CRI", "506", "Costa Rica"},
	{"HRV", "385", "Croatia"},
	{"CUB", "53", "Cuba"},
	{"CUW", "599", "Curacao"},
	{"CYP", "357", "Cyprus"},
	{"CYN", "90", "Cyprus, Northern"},
	{"CZE", "420", "Czech Republic"},
	{"DNK", "45", "Denmark"},
	{"DJI", "253", "Djibouti"},
	{"DMA", "1767", "Dominica"},
	{"DOM", "1829", "Dominican Republic"},
	{"ECU", "593", "Ecuador"},
	{"EGY", "20", "Egypt"},
	{"SLV", "503", "El Salvador"},
	{"GNQ", "240", "Equatorial Guinea"},
	{"ERI", "291", "Eritrea"},
	{"EST", "372", "Estonia"},
	{"ETH", "251", "Ethiopia"},
	{"FLK", "500", "Falkland Islands"},
	{"FRO", "298", "Faroe Islands"},
	{"FJI", "679", "Fiji"},
	{"FIN", "358", "Finland"},
	{"FRA", "33", "France"},
	{"GUF", "594", "French Guiana"},
	{"PYF", "689", "French Polynesia"},
	{"GAB", "241", "Gabon"},
	{"GMB", "220", "Gambia"},
	{"GEO", "995", "Georgia"},
	{"DEU", "49", "Germany"},
	{"GHA", "233", "Ghana"},
	{"GIB", "350", "Gibraltar"},
	{"GRC", "30", "Greece"},
	{"GRL", "299", "Greenland"},
	{"GRD", "1473", "Grenada"},
	{"GLP", "590", "Guadeloupe"},
	{"GUM", "1671", "Guam"},
	{"GTM", "502", "Guatemala"},
	{"GGY", "44", "Guernsey"},
	{"GIN", "224", "Guinea"},
	{"GNB", "245", "Guinea-Bissau"},
	{"GUY", "592", "Guyana"},
	{"HTI", "509", "Haiti"},
	{"HND", "504", "Honduras"},
	{"HKG", "852", "Hong Kong"},
	{"HUN", "36", "Hungary"},
	{"ISL", "354", "Iceland"},
	{"IND", "91", "India"},
	{"INT", "883", "International National Rate Service"},
	{"IDN", "62", "Indonesia"},
	{"IRN", "98", "Iran"},
	{"IRQ", "964", "Iraq"},
	{"IRL", "353", "Ireland"},
	{"IMN", "44", "Isle of Man"},
	{"ISR", "972", "Israel"},
	{"ITA", "39", "Italy"},
	{"CIV", "225", "Ivory Coast"},
	{"JAM", "1876", "Jamaica"},
	{"JPN", "81", "Japan"},
	{"JEY", "44", "Jersey"},
	{"JOR", "962", "Jordan"},
	{"KAZ", "7", "Kazakhstan"},
	{"KEN", "254", "Kenya"},
	{"KIR", "686", "Kiribati"},
	{"KOS", "383", "Kosovo"},
	{"KWT", "965", "Kuwait"},
	{"KGZ", "996", "Kyrgyzstan"},
	{"LAO", "856", "Laos"},
	{"LVA", "371", "Latvia"},
	{"LBN", "961", "Lebanon"},
	{"LSO", "266", "Lesotho"},
	{"LBR", "231", "Liberia"},
	{"LBY", "218", "Libya"},
	{"LIE", "423", "Liechtenstein"},
	{"LTU", "370", "Lithuania"},
	{"LUX", "352", "Luxembourg"},
	{"MAC", "853", "Macau"},
	{"MKD", "389", "Macedonia"},
	{"MDG", "261", "Madagascar"},
	{"MWI", "265", "Malawi"},
	{"MYS", "60", "Malaysia"},
	{"MDV", "960", "Maldives"},
	{"MLI", "223", "Mali"},
	{"MLT", "356", "Malta"},
	{"MHL", "692", "Marshall Islands"},
	{"MRT", "222", "Mauritania"},
	{"MUS", "230", "Mauritius"},
	{"MEX", "52", "Mexico"},
	{"FSM", "691", "Micronesia"},
	{"MDA", "373", "Moldova"},
	{"MCO", "377", "Monaco"},
	{"MNG", "976", "Mongolia"},
	{"MNE", "382", "Montenegro"},
	{"MSR", "1664", "Montserrat"},
	{"MAR", "212", "Morocco"},
	{"MOZ", "258", "Mozambique"},
	{"MMR", "95", "Myanmar"},
	{"NAM", "264", "Namibia"},
	{"NRU", "674", "Nauru"},
	{"NPL", "977", "Nepal"},
	{"NLD", "31", "Netherlands"},
	{"NCL", "687", "New Caledonia"},
	{"NZL", "64", "New Zealand"},
	{"NIC", "505", "Nicaragua"},
	{"NER", "227", "Niger"},
	{"NGA", "234", "Nigeria"},
	{"NIU", "683", "Niue"},
	{"NFK", "672", "Norfolk Island"},
	{"MNP", "1670", "Northern Mariana Islands"},
	{"NOR", "47", "Norway"},
	{"OMN", "968", "Oman"},
	{"PAK", "92", "Pakistan"},
	{"PLW", "680", "Palau"},
	{"PSE", "972", "Palestine"},
	{"PAN", "507", "Panama"},
	{"PNG", "675", "Papua New Guinea"},
	{"PRY", "595", "Paraguay"},
	{"PER", "51", "Peru"},
	{"PHL", "63", "Philippines"},
	{"POL", "48", "Poland"},
	{"PRT", "351", "Portugal"},
	{"PRI", "1787", "Puerto Rico"},
	{"QAT", "974", "Qatar"},
	{"COD", "243", "Republic Democratic Of Congo"},
	{"REU", "262", "Reunion"},
	{"ROU", "40", "Romania"},
	{"RUS", "7", "Russia"},
	{"RWA", "250", "Rwanda"},
	{"KNA", "1869", "Saint Kitts and Nevis"},
	{"LCA", "1758", "Saint Lucia"},
	{"SPM", "508", "Saint Pierre and Miquelon"},
	{"VCT", "1784", "Saint Vincent and the Grenadines"},
	{"WSM", "685", "Samoa"},
	{"SMR", "378", "San Marino"},
	{"STP", "239", "Sao Tome and Principe"},
	{"SAU", "966", "Saudi Arabia"},
	{"SEN", "221", "Senegal"},
	{"SRB", "381", "Serbia"},
	{"SYC", "248", "Seychelles"},
	{"SLE", "232", "Sierra Leone"},
	{"SGP", "65", "Singapore"},
	{"SVK", "421", "Slovakia"},
	{"SVN", "386", "Slovenia"},
	{"SLB", "677", "Solomon Islands"},
	{"SOM", "252", "Somalia"},
	{"ZAF", "27", "South Africa"},
	{"KOR", "82", "South Korea"},
	{"SSD", "211", "South Sudan"},
	{"ESP", "34", "Spain"},
	{"LKA", "94", "Sri Lanka"},
	{"SDN", "249", "Sudan"},
	{"SUR", "597", "Suriname"},
	{"SWZ", "268", "Swaziland"},
	{"SWE", "46", "Sweden"},
	{"CHE", "41", "Switzerland"},
	{"SYR", "963", "Syria"},
	{"TWN", "886", "Taiwan"},
	{"TJK", "992", "Tajikistan"},
	{"TZA", "255", "Tanzania"},
	{"THA", "66", "Thailand"},
	{"TLS", "670", "Timor-Leste"},
	{"TGO", "228", "Togo"},
	{"TON", "676", "Tonga"},
	{"TTO", "1868", "Trinidad and Tobago"},
	{"TUN", "216", "Tunisia"},
	{"TUR", "90", "Turkey"},
	{"TKM", "993", "Turkmenistan"},
	{"TCA", "1649", "Turks and Caicos Islands"},
	{"TUV", "688", "Tuvalu"},
	{"UGA", "256", "Uganda"},
	{"UKR", "380", "Ukraine"},
	{"ARE", "971", "United Arab Emirates"},
	{"GBR", "44", "United Kingdom"},
	{"USA", "1", "United States"},
	{"URY", "598", "Uruguay"},
	{"UZB", "998", "Uzbekistan"},
	{"VUT", "678", "Vanuatu"},
	{"VEN", "58", "Venezuela"},
	{"VNM", "84", "Vietnam"},
	{"VGB", "1284", "Virgin Islands, British"},
	{"WLF", "681", "Wallis and Futuna Islands"},
	{"YEM", "967", "Yemen"},
	{"ZMB", "260", "Zambia"},
	{"ZWE", "263", "Zimbabwe"},
}

func GetCountryByPhoneNumber(phoneNumber string) (*Country, error) {
	phoneNumber = strings.Replace(phoneNumber, "+", "", 1)

	for _, country := range countries {
		countryCodeLength := len(country.CountryCode)
		phoneNumberCountryCode := phoneNumber[0:countryCodeLength]

		if country.CountryCode == phoneNumberCountryCode {
			return &country, nil
		}
	}

	return nil, fmt.Errorf("phone number not recognized")
}

func GetAllCountries() []Country {
	return countries
}