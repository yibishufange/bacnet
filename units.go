package bacnet

type Unit uint16

//go:generate stringer -type=Unit
const (
	SquareMeters                    Unit = 0x00
	SquareFeet                      Unit = 0x01
	Milliamperes                    Unit = 0x02
	Amperes                         Unit = 0x03
	Ohms                            Unit = 0x04
	Volts                           Unit = 0x05
	Kilovolts                       Unit = 0x06
	Megavolts                       Unit = 0x07
	VoltAmperes                     Unit = 0x08
	KilovoltAmperes                 Unit = 0x09
	MegavoltAmperes                 Unit = 0x0A
	VoltAmperesReactive             Unit = 0x0B
	KilovoltAmperesReactive         Unit = 0x0C
	MegavoltAmperesReactive         Unit = 0x0D
	DegreesPhase                    Unit = 0x0E
	PowerFactor                     Unit = 0x0F
	Joules                          Unit = 0x10
	Kilojoules                      Unit = 0x11
	WattHours                       Unit = 0x12
	KilowattHours                   Unit = 0x13
	Btus                            Unit = 0x14
	Therms                          Unit = 0x15
	TonHours                        Unit = 0x16
	JoulesPerKilogramDryAir         Unit = 0x17
	BtusPerPoundDryAir              Unit = 0x18
	CyclesPerHour                   Unit = 0x19
	CyclesPerMinute                 Unit = 0x1A
	Hertz                           Unit = 0x1B
	GramsOfWaterPerKilogramDryAir   Unit = 0x1C
	PercentRelativeHumidity         Unit = 0x1D
	Millimeters                     Unit = 0x1E
	Meters                          Unit = 0x1F
	Inches                          Unit = 0x20
	Feet                            Unit = 0x21
	WattsPerSquareFoot              Unit = 0x22
	WattsPerSquareMeter             Unit = 0x23
	Lumens                          Unit = 0x24
	Luxes                           Unit = 0x25
	FootCandles                     Unit = 0x26
	Kilograms                       Unit = 0x27
	PoundsMass                      Unit = 0x28
	Tons                            Unit = 0x29
	KilogramsPerSecond              Unit = 0x2A
	KilogramsPerMinute              Unit = 0x2B
	KilogramsPerHour                Unit = 0x2C
	PoundsMassPerMinute             Unit = 0x2D
	PoundsMassPerHour               Unit = 0x2E
	Watts                           Unit = 0x2F
	Kilowatts                       Unit = 0x30
	Megawatts                       Unit = 0x31
	BtusPerHour                     Unit = 0x32
	Horsepower                      Unit = 0x33
	TonsRefrigeration               Unit = 0x34
	Pascals                         Unit = 0x35
	Kilopascals                     Unit = 0x36
	Bars                            Unit = 0x37
	PoundsForcePerSquareInch        Unit = 0x38
	CentimetersOfWater              Unit = 0x39
	InchesOfWater                   Unit = 0x3A
	MillimetersOfMercury            Unit = 0x3B
	CentimetersOfMercury            Unit = 0x3C
	InchesOfMercury                 Unit = 0x3D
	DegreesCelsius                  Unit = 0x3E
	DegreesKelvin                   Unit = 0x3F
	DegreesFahrenheit               Unit = 0x40
	DegreeDaysCelsius               Unit = 0x41
	DegreeDaysFahrenheit            Unit = 0x42
	Years                           Unit = 0x43
	Months                          Unit = 0x44
	Weeks                           Unit = 0x45
	Days                            Unit = 0x46
	Hours                           Unit = 0x47
	Minutes                         Unit = 0x48
	Seconds                         Unit = 0x49
	MetersPerSecond                 Unit = 0x4A
	KilometersPerHour               Unit = 0x4B
	FeetPerSecond                   Unit = 0x4C
	FeetPerMinute                   Unit = 0x4D
	MilesPerHour                    Unit = 0x4E
	CubicFeet                       Unit = 0x4F
	CubicMeters                     Unit = 0x50
	ImperialGallons                 Unit = 0x51
	Liters                          Unit = 0x52
	UsGallons                       Unit = 0x53
	CubicFeetPerMinute              Unit = 0x54
	CubicMetersPerSecond            Unit = 0x55
	ImperialGallonsPerMinute        Unit = 0x56
	LitersPerSecond                 Unit = 0x57
	LitersPerMinute                 Unit = 0x58
	USGallonsPerMinute              Unit = 0x59
	DegreesAngular                  Unit = 0x5A
	DegreesCelsiusPerHour           Unit = 0x5B
	DegreesCelsiusPerMinute         Unit = 0x5C
	DegreesFahrenheitPerHour        Unit = 0x5D
	DegreesFahrenheitPerMinute      Unit = 0x5E
	NoUnits                         Unit = 0x5F
	PartsPerMillion                 Unit = 0x60
	PartsPerBillion                 Unit = 0x61
	Percent                         Unit = 0x62
	PercentPerSecond                Unit = 0x63
	PerMinute                       Unit = 0x64
	PerSecond                       Unit = 0x65
	PsiPerDegreeFahrenheit          Unit = 0x66
	Radians                         Unit = 0x67
	RevolutionsPerMinute            Unit = 0x68
	SquareInches                    Unit = 0x73
	SquareCentimeters               Unit = 0x74
	BtusPerPound                    Unit = 0x75
	Centimeters                     Unit = 0x76
	PoundsMassPerSecond             Unit = 0x77
	DeltaDegreesFahrenheit          Unit = 0x78
	DeltaDegreesKelvin              Unit = 0x79
	Kilohms                         Unit = 0x7A
	Megohms                         Unit = 0x7B
	Millivolts                      Unit = 0x7C
	KilojoulesPerKilogram           Unit = 0x7D
	Megajoules                      Unit = 0x7E
	JoulesPerDegreeKelvin           Unit = 0x7F
	JoulesPerKilogramDegreeKelvin   Unit = 0x80
	Kilohertz                       Unit = 0x81
	Megahertz                       Unit = 0x82
	PerHour                         Unit = 0x83
	Milliwatts                      Unit = 0x84
	Hectopascals                    Unit = 0x85
	Millibars                       Unit = 0x86
	CubicMetersPerHour              Unit = 0x87
	LitersPerHour                   Unit = 0x88
	KwHoursPerSquareMeter           Unit = 0x89
	KwHoursPerSquareFoot            Unit = 0x8A
	MegajoulesPerSquareMeter        Unit = 0x8B
	MegajoulesPerSquareFoot         Unit = 0x8C
	WattsPerSquareMeterDegreeKelvin Unit = 0x8D
	CubicFeetPerSecond              Unit = 0x8E
	PercentObscurationPerFoot       Unit = 0x8F
	PercentObscurationPerMeter      Unit = 0x90
	Milliohms                       Unit = 0x91
	MegawattHours                   Unit = 0x92
	KiloBtus                        Unit = 0x93
	MegaBtus                        Unit = 0x94
	KilojoulesPerKilogramDryAir     Unit = 0x95
	MegajoulesPerKilogramDryAir     Unit = 0x96
	KilojoulesPerDegreeKelvin       Unit = 0x97
	MegajoulesPerDegreeKelvin       Unit = 0x98
	Newton                          Unit = 0x99
	GramsPerSecond                  Unit = 0x9A
	GramsPerMinute                  Unit = 0x9B
	TonsPerHour                     Unit = 0x9C
	KiloBtusPerHour                 Unit = 0x9D
	HundredthsSeconds               Unit = 0x9E
	Milliseconds                    Unit = 0x9F
	NewtonMeters                    Unit = 0xA0
	MillimetersPerSecond            Unit = 0xA1
	MillimetersPerMinute            Unit = 0xA2
	MetersPerMinute                 Unit = 0xA3
	MetersPerHour                   Unit = 0xA4
	CubicMetersPerMinute            Unit = 0xA5
	MetersPerSecondPerSecond        Unit = 0xA6
	AmperesPerMeter                 Unit = 0xA7
	AmperesPerSquareMeter           Unit = 0xA8
	AmpereSquareMeters              Unit = 0xA9
	Farads                          Unit = 0xAA
	Henrys                          Unit = 0xAB
	OhmMeters                       Unit = 0xAC
	Siemens                         Unit = 0xAD // 1 mho equals 1 siemens
	SiemensPerMeter                 Unit = 0xAE
	Teslas                          Unit = 0xAF
	VoltsPerDegreeKelvin            Unit = 0xB0
	VoltsPerMeter                   Unit = 0xB1
	Webers                          Unit = 0xB2
	Candelas                        Unit = 0xB3
	CandelasPerSquareMeter          Unit = 0xB4
	DegreesKelvinPerHour            Unit = 0xB5
	DegreesKelvinPerMinute          Unit = 0xB6
	JouleSeconds                    Unit = 0xB7
	RadiansPerSecond                Unit = 0xB8
	SquareMetersPerNewton           Unit = 0xB9
	KilogramsPerCubicMeter          Unit = 0xBA
	NewtonSeconds                   Unit = 0xBB
	NewtonsPerMeter                 Unit = 0xBC
	WattsPerMeterPerDegreeKelvin    Unit = 0xBD
	Microsiemens                    Unit = 0xBE
	CubicFeetPerHour                Unit = 0xBF
	USGallonsPerHour                Unit = 0xC0
	Kilometers                      Unit = 0xC1
	Micrometers                     Unit = 0xC2
	Grams                           Unit = 0xC3
	Milligrams                      Unit = 0xC4
	Milliliters                     Unit = 0xC5
	MillilitersPerSecond            Unit = 0xC6
	Decibels                        Unit = 0xC7
	DecibelsMillivolt               Unit = 0xC8
	DecibelsVolt                    Unit = 0xC9
	Millisiemens                    Unit = 0xCA
	WattHoursReactive               Unit = 0xCB
	KilowattHoursReactive           Unit = 0xCC
	MegawattHoursReactive           Unit = 0xCD
	MillimetersOfWater              Unit = 0xCE
	PerMille                        Unit = 0xCF
	GramsPerGram                    Unit = 0xD0
	KilogramsPerKilogram            Unit = 0xD1
	GramsPerKilogram                Unit = 0xD2
	MilligramsPerGram               Unit = 0xD3
	MilligramsPerKilogram           Unit = 0xD4
	GramsPerMilliliter              Unit = 0xD5
	GramsPerLiter                   Unit = 0xD6
	MilligramsPerLiter              Unit = 0xD7
	MicrogramsPerLiter              Unit = 0xD8
	GramsPerCubicMeter              Unit = 0xD9
	MilligramsPerCubicMeter         Unit = 0xDA
	MicrogramsPerCubicMeter         Unit = 0xDB
	NanogramsPerCubicMeter          Unit = 0xDC
	GramsPerCubicCentimeter         Unit = 0xDD
	Becquerels                      Unit = 0xDE
	Megabecquerels                  Unit = 0xE0
	Gray                            Unit = 0xE1
	Milligray                       Unit = 0xE2
	Microgray                       Unit = 0xE3
	Sieverts                        Unit = 0xE4
	Millisieverts                   Unit = 0xE5
	Microsieverts                   Unit = 0xE6
	MicrosievertsPerHour            Unit = 0xE7
	DecibelsA                       Unit = 0xE8
	NephelometricTurbidityUnit      Unit = 0xE9
	Ph                              Unit = 0xEA
	GramsPerSquareMeter             Unit = 0xEB
	MinutesPerDegreeKelvin          Unit = 0xEC
)
