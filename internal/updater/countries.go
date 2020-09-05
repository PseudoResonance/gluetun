package updater

func getCountryCodes() map[string]string { //nolint:dupl
	return map[string]string{
		"af": "Afghanistan",
		"ax": "Aland Islands",
		"al": "Albania",
		"dz": "Algeria",
		"as": "American Samoa",
		"ad": "Andorra",
		"ao": "Angola",
		"ai": "Anguilla",
		"aq": "Antarctica",
		"ag": "Antigua and Barbuda",
		"ar": "Argentina",
		"am": "Armenia",
		"aw": "Aruba",
		"au": "Australia",
		"at": "Austria",
		"az": "Azerbaijan",
		"bs": "Bahamas",
		"bh": "Bahrain",
		"bd": "Bangladesh",
		"bb": "Barbados",
		"by": "Belarus",
		"be": "Belgium",
		"bz": "Belize",
		"bj": "Benin",
		"bm": "Bermuda",
		"bt": "Bhutan",
		"bo": "Bolivia",
		"bq": "Bonaire",
		"ba": "Bosnia and Herzegovina",
		"bw": "Botswana",
		"bv": "Bouvet Island",
		"br": "Brazil",
		"io": "British Indian Ocean Territory",
		"vg": "British Virgin Islands",
		"bn": "Brunei Darussalam",
		"bg": "Bulgaria",
		"bf": "Burkina Faso",
		"bi": "Burundi",
		"kh": "Cambodia",
		"cm": "Cameroon",
		"ca": "Canada",
		"cv": "Cape Verde",
		"ky": "Cayman Islands",
		"cf": "Central African Republic",
		"td": "Chad",
		"cl": "Chile",
		"cn": "China",
		"cx": "Christmas Island",
		"cc": "Cocos Islands",
		"co": "Colombia",
		"km": "Comoros",
		"cg": "Congo",
		"ck": "Cook Islands",
		"cr": "Costa Rica",
		"ci": "Cote d'Ivoire",
		"hr": "Croatia",
		"cu": "Cuba",
		"cw": "Curacao",
		"cy": "Cyprus",
		"cz": "Czech Republic",
		"cd": "Democratic Republic of the Congo",
		"dk": "Denmark",
		"dj": "Djibouti",
		"dm": "Dominica",
		"do": "Dominican Republic",
		"ec": "Ecuador",
		"eg": "Egypt",
		"sv": "El Salvador",
		"gq": "Equatorial Guinea",
		"er": "Eritrea",
		"ee": "Estonia",
		"et": "Ethiopia",
		"fk": "Falkland Islands",
		"fo": "Faroe Islands",
		"fj": "Fiji",
		"fi": "Finland",
		"fr": "France",
		"gf": "French Guiana",
		"pf": "French Polynesia",
		"tf": "French Southern Territories",
		"ga": "Gabon",
		"gm": "Gambia",
		"ge": "Georgia",
		"de": "Germany",
		"gh": "Ghana",
		"gi": "Gibraltar",
		"gr": "Greece",
		"gl": "Greenland",
		"gd": "Grenada",
		"gp": "Guadeloupe",
		"gu": "Guam",
		"gt": "Guatemala",
		"gg": "Guernsey",
		"gw": "Guinea-Bissau",
		"gn": "Guinea",
		"gy": "Guyana",
		"ht": "Haiti",
		"hm": "Heard Island and McDonald Islands",
		"hn": "Honduras",
		"hk": "Hong Kong",
		"hu": "Hungary",
		"is": "Iceland",
		"in": "India",
		"id": "Indonesia",
		"ir": "Iran",
		"iq": "Iraq",
		"ie": "Ireland",
		"im": "Isle of Man",
		"il": "Israel",
		"it": "Italy",
		"jm": "Jamaica",
		"jp": "Japan",
		"je": "Jersey",
		"jo": "Jordan",
		"kz": "Kazakhstan",
		"ke": "Kenya",
		"ki": "Kiribati",
		"kr": "Korea",
		"kw": "Kuwait",
		"kg": "Kyrgyzstan",
		"la": "Lao People's Democratic Republic",
		"lv": "Latvia",
		"lb": "Lebanon",
		"ls": "Lesotho",
		"lr": "Liberia",
		"ly": "Libya",
		"li": "Liechtenstein",
		"lt": "Lithuania",
		"lu": "Luxembourg",
		"mo": "Macao",
		"mk": "Macedonia",
		"mg": "Madagascar",
		"mw": "Malawi",
		"my": "Malaysia",
		"mv": "Maldives",
		"ml": "Mali",
		"mt": "Malta",
		"mh": "Marshall Islands",
		"mq": "Martinique",
		"mr": "Mauritania",
		"mu": "Mauritius",
		"yt": "Mayotte",
		"mx": "Mexico",
		"fm": "Micronesia",
		"md": "Moldova",
		"mc": "Monaco",
		"mn": "Mongolia",
		"me": "Montenegro",
		"ms": "Montserrat",
		"ma": "Morocco",
		"mz": "Mozambique",
		"mm": "Myanmar",
		"na": "Namibia",
		"nr": "Nauru",
		"np": "Nepal",
		"nl": "Netherlands",
		"nc": "New Caledonia",
		"nz": "New Zealand",
		"ni": "Nicaragua",
		"ne": "Niger",
		"ng": "Nigeria",
		"nu": "Niue",
		"nf": "Norfolk Island",
		"mp": "Northern Mariana Islands",
		"no": "Norway",
		"om": "Oman",
		"pk": "Pakistan",
		"pw": "Palau",
		"ps": "Palestine, State of",
		"pa": "Panama",
		"pg": "Papua New Guinea",
		"py": "Paraguay",
		"pe": "Peru",
		"ph": "Philippines",
		"pn": "Pitcairn",
		"pl": "Poland",
		"pt": "Portugal",
		"pr": "Puerto Rico",
		"qa": "Qatar",
		"re": "Reunion",
		"ro": "Romania",
		"ru": "Russian Federation",
		"rw": "Rwanda",
		"bl": "Saint Barthelemy",
		"sh": "Saint Helena",
		"kn": "Saint Kitts and Nevis",
		"lc": "Saint Lucia",
		"mf": "Saint Martin",
		"pm": "Saint Pierre and Miquelon",
		"vc": "Saint Vincent and the Grenadines",
		"ws": "Samoa",
		"sm": "San Marino",
		"st": "Sao Tome and Principe",
		"sa": "Saudi Arabia",
		"sn": "Senegal",
		"rs": "Serbia",
		"sc": "Seychelles",
		"sl": "Sierra Leone",
		"sg": "Singapore",
		"sx": "Sint Maarten",
		"sk": "Slovakia",
		"si": "Slovenia",
		"sb": "Solomon Islands",
		"so": "Somalia",
		"za": "South Africa",
		"gs": "South Georgia and the South Sandwich Islands",
		"ss": "South Sudan",
		"es": "Spain",
		"lk": "Sri Lanka",
		"sd": "Sudan",
		"sr": "Suriname",
		"sj": "Svalbard and Jan Mayen",
		"sz": "Swaziland",
		"se": "Sweden",
		"ch": "Switzerland",
		"sy": "Syrian Arab Republic",
		"tw": "Taiwan",
		"tj": "Tajikistan",
		"tz": "Tanzania",
		"th": "Thailand",
		"tl": "Timor-Leste",
		"tg": "Togo",
		"tk": "Tokelau",
		"to": "Tonga",
		"tt": "Trinidad and Tobago",
		"tn": "Tunisia",
		"tr": "Turkey",
		"tm": "Turkmenistan",
		"tc": "Turks and Caicos Islands",
		"tv": "Tuvalu",
		"ug": "Uganda",
		"ua": "Ukraine",
		"ae": "United Arab Emirates",
		"gb": "United Kingdom",
		"um": "United States Minor Outlying Islands",
		"us": "United States",
		"uy": "Uruguay",
		"vi": "US Virgin Islands",
		"uz": "Uzbekistan",
		"vu": "Vanuatu",
		"va": "Vatican City State",
		"ve": "Venezuela",
		"vn": "Vietnam",
		"wf": "Wallis and Futuna",
		"eh": "Western Sahara",
		"ye": "Yemen",
		"zm": "Zambia",
		"zw": "Zimbabwe",
	}
}
