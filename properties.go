package bacnet

//PropertyType is the type of an object property
type PropertyType uint32

//go:generate stringer -type=PropertyType
const (
	AckedTransitions                 PropertyType = 0x00
	AckRequired                      PropertyType = 0x01
	Action                           PropertyType = 0x02
	ActionText                       PropertyType = 0x03
	ActiveText                       PropertyType = 0x04
	ActiveVtSessions                 PropertyType = 0x05
	AlarmValue                       PropertyType = 0x06
	AlarmValues                      PropertyType = 0x07
	All                              PropertyType = 0x08
	AllWritesSuccessful              PropertyType = 0x09
	ApduSegmentTimeout               PropertyType = 0x0A
	ApduTimeout                      PropertyType = 0x0B
	ApplicationSoftwareVersion       PropertyType = 0x0C
	Archive                          PropertyType = 0x0D
	Bias                             PropertyType = 0x0E
	ChangeOfStateCount               PropertyType = 0x0F
	ChangeOfStateTime                PropertyType = 0x10
	NotificationClassProp            PropertyType = 0x11
	ControlledVariableReference      PropertyType = 0x13
	ControlledVariableUnits          PropertyType = 0x14
	ControlledVariableValue          PropertyType = 0x15
	CovIncrement                     PropertyType = 0x16
	DateList                         PropertyType = 0x17
	DaylightSavingsStatus            PropertyType = 0x18
	Deadband                         PropertyType = 0x19
	DerivativeAnt                    PropertyType = 0x1A
	DerivativeAntUnits               PropertyType = 0x1B
	Description                      PropertyType = 0x1C
	DescriptionOfHalt                PropertyType = 0x1D
	DeviceAddressBinding             PropertyType = 0x1E
	DeviceType                       PropertyType = 0x1F
	EffectivePeriod                  PropertyType = 0x20
	ElapsedActiveTime                PropertyType = 0x21
	ErrorLimit                       PropertyType = 0x22
	EventEnable                      PropertyType = 0x23
	EventState                       PropertyType = 0x24
	EventType                        PropertyType = 0x25
	ExceptionSchedule                PropertyType = 0x26
	FaultValues                      PropertyType = 0x27
	FeedbackValue                    PropertyType = 0x28
	FileAccessMethod                 PropertyType = 0x29
	FileSize                         PropertyType = 0x2A
	FileType                         PropertyType = 0x2B
	FirmwareRevision                 PropertyType = 0x2C
	HighLimit                        PropertyType = 0x2D
	InactiveText                     PropertyType = 0x2E
	InProcess                        PropertyType = 0x2F
	InstanceOf                       PropertyType = 0x30
	IntegralAnt                      PropertyType = 0x31
	IntegralAntUnits                 PropertyType = 0x32
	IssueConfirmedNotifications      PropertyType = 0x33
	LimitEnable                      PropertyType = 0x34
	ListOfGroupMembers               PropertyType = 0x35
	ListOfObjectPropertyReferences   PropertyType = 0x36
	ListOfSessionKeys                PropertyType = 0x37
	LocalDate                        PropertyType = 0x38
	LocalTime                        PropertyType = 0x39
	Location                         PropertyType = 0x3A
	LowLimit                         PropertyType = 0x3B
	ManipulatedVariableReference     PropertyType = 0x3C
	MaximumOutput                    PropertyType = 0x3D
	MaxApduLengthAccepted            PropertyType = 0x3E
	MaxInfoFrames                    PropertyType = 0x3F
	MaxMaster                        PropertyType = 0x40
	MaxPresValue                     PropertyType = 0x41
	MinimumOffTime                   PropertyType = 0x42
	MinimumOnTime                    PropertyType = 0x43
	MinimumOutput                    PropertyType = 0x44
	MinPresValue                     PropertyType = 0x45
	ModelName                        PropertyType = 0x46
	ModificationDate                 PropertyType = 0x47
	NotifyType                       PropertyType = 0x48
	NumberOfApduRetries              PropertyType = 0x49
	NumberOfStates                   PropertyType = 0x4A
	ObjectIdentifier                 PropertyType = 0x4B
	ObjectList                       PropertyType = 0x4C
	ObjectName                       PropertyType = 0x4D
	ObjectPropertyReference          PropertyType = 0x4E
	ObjectTypeProp                   PropertyType = 0x4F
	Optional                         PropertyType = 0x50
	OutOfService                     PropertyType = 0x51
	OutputUnits                      PropertyType = 0x52
	EventParameters                  PropertyType = 0x53
	Polarity                         PropertyType = 0x54
	PresentValue                     PropertyType = 0x55
	Priority                         PropertyType = 0x56
	PriorityArray                    PropertyType = 0x57
	PriorityForWriting               PropertyType = 0x58
	ProcessIdentifier                PropertyType = 0x59
	ProgramChange                    PropertyType = 0x5A
	ProgramLocation                  PropertyType = 0x5B
	ProgramState                     PropertyType = 0x5C
	ProportionalAnt                  PropertyType = 0x5D
	ProportionalAntUnits             PropertyType = 0x5E
	ProtocolConformanceClass         PropertyType = 0x5F // Deleted In Version 1 revision 2
	ProtocolObjectTypesSupported     PropertyType = 0x60
	ProtocolServicesSupported        PropertyType = 0x61
	ProtocolVersion                  PropertyType = 0x62
	ReadOnly                         PropertyType = 0x63
	ReasonForHalt                    PropertyType = 0x64
	Recipient                        PropertyType = 0x65
	RecipientList                    PropertyType = 0x66
	Reliability                      PropertyType = 0x67
	RelinquishDefault                PropertyType = 0x68
	Required                         PropertyType = 0x69
	Resolution                       PropertyType = 0x6A
	SegmentationSupported            PropertyType = 0x6B
	Setpoint                         PropertyType = 0x6C
	SetpointReference                PropertyType = 0x6D
	StateText                        PropertyType = 0x6E
	StatusFlags                      PropertyType = 0x6F
	SystemStatus                     PropertyType = 0x70
	TimeDelay                        PropertyType = 0x71
	TimeOfActiveTimeReset            PropertyType = 0x72
	TimeOfStateCountReset            PropertyType = 0x73
	TimeSynchronizationRecipients    PropertyType = 0x74
	Units                            PropertyType = 0x75
	UpdateInterval                   PropertyType = 0x76
	UtcOffset                        PropertyType = 0x77
	VendorIdentifier                 PropertyType = 0x78
	VendorName                       PropertyType = 0x79
	VtClassesSupported               PropertyType = 0x7A
	WeeklySchedule                   PropertyType = 0x7B
	AttemptedSamples                 PropertyType = 0x7C
	AverageValue                     PropertyType = 0x7D
	BufferSize                       PropertyType = 0x7E
	ClientCovIncrement               PropertyType = 0x7F
	CovResubscriptionInterval        PropertyType = 0x80
	CurrentNotifyTime                PropertyType = 0x81
	EventTimeStamps                  PropertyType = 0x82
	LogBuffer                        PropertyType = 0x83
	LogDeviceObjectProperty          PropertyType = 0x84
	Enable                           PropertyType = 0x85
	LogInterval                      PropertyType = 0x86
	MaximumValue                     PropertyType = 0x87
	MinimumValue                     PropertyType = 0x88
	NotificationThreshold            PropertyType = 0x89
	PreviousNotifyTime               PropertyType = 0x8A
	ProtocolRevision                 PropertyType = 0x8B
	RecordsSinceNotification         PropertyType = 0x8C
	RecordCount                      PropertyType = 0x8D
	StartTime                        PropertyType = 0x8E
	StopTime                         PropertyType = 0x8F
	StopWhenFull                     PropertyType = 0x90
	TotalRecordCount                 PropertyType = 0x91
	ValidSamples                     PropertyType = 0x92
	WindowInterval                   PropertyType = 0x93
	WindowSamples                    PropertyType = 0x94
	MaximumValueTimestamp            PropertyType = 0x95
	MinimumValueTimestamp            PropertyType = 0x96
	VarianceValue                    PropertyType = 0x97
	ActiveCovSubscriptions           PropertyType = 0x98
	BackupFailureTimeout             PropertyType = 0x99
	ConfigurationFiles               PropertyType = 0x9A
	DatabaseRevision                 PropertyType = 0x9B
	DirectReading                    PropertyType = 0x9C
	LastRestoreTime                  PropertyType = 0x9D
	MaintenanceRequired              PropertyType = 0x9E
	MemberOf                         PropertyType = 0x9F
	Mode                             PropertyType = 0xA0
	OperationExpected                PropertyType = 0xA1
	Setting                          PropertyType = 0xA2
	Silenced                         PropertyType = 0xA3
	TrackingValue                    PropertyType = 0xA4
	ZoneMembers                      PropertyType = 0xA5
	LifeSafetyAlarmValues            PropertyType = 0xA6
	MaxSegmentsAccepted              PropertyType = 0xA7
	ProfileName                      PropertyType = 0xA8
	AutoSlaveDiscovery               PropertyType = 0xA9
	ManualSlaveAddressBinding        PropertyType = 0xAA
	SlaveAddressBinding              PropertyType = 0xAB
	SlaveProxyEnable                 PropertyType = 0xAC
	LastNotifyRecord                 PropertyType = 0xAD
	ScheduleDefault                  PropertyType = 0xAE
	AcceptedModes                    PropertyType = 0xAF
	AdjustValue                      PropertyType = 0xB0
	Count                            PropertyType = 0xB1
	CountBeforeChange                PropertyType = 0xB2
	CountChangeTime                  PropertyType = 0xB3
	CovPeriod                        PropertyType = 0xB4
	InputReference                   PropertyType = 0xB5
	LimitMonitoringInterval          PropertyType = 0xB6
	LoggingObject                    PropertyType = 0xB7
	LoggingRecord                    PropertyType = 0xB8
	Prescale                         PropertyType = 0xB9
	PulseRate                        PropertyType = 0xBA
	Scale                            PropertyType = 0xBB
	ScaleFactor                      PropertyType = 0xBC
	UpdateTime                       PropertyType = 0xBD
	ValueBeforeChange                PropertyType = 0xBE
	ValueSet                         PropertyType = 0xBF
	ValueChangeTime                  PropertyType = 0xC0
	AlignIntervals                   PropertyType = 0xC1
	IntervalOffset                   PropertyType = 0xC3
	LastRestartReason                PropertyType = 0xC4
	LoggingType                      PropertyType = 0xC5
	RestartNotificationRecipients    PropertyType = 0xCA
	TimeOfDeviceRestart              PropertyType = 0xCB
	TimeSynchronizationInterval      PropertyType = 0xCC
	Trigger                          PropertyType = 0xCD
	UTCTimeSynchronizationRecipients PropertyType = 0xCE
	NodeSubtype                      PropertyType = 0xCF
	NodeType                         PropertyType = 0xD0
	StructuredObjectList             PropertyType = 0xD1
	SubordinateAnnotations           PropertyType = 0xD2
	SubordinateList                  PropertyType = 0xD3
	ActualShedLevel                  PropertyType = 0xD4
	DutyWindow                       PropertyType = 0xD5
	ExpectedShedLevel                PropertyType = 0xD6
	FullDutyBaseline                 PropertyType = 0xD7
	RequestedShedLevel               PropertyType = 0xDA
	ShedDuration                     PropertyType = 0xDB
	ShedLevelDescriptions            PropertyType = 0xDC
	ShedLevels                       PropertyType = 0xDD
	StateDescription                 PropertyType = 0xDE
	DoorAlarmState                   PropertyType = 0xE2
	DoorExtendedPulseTime            PropertyType = 0xE3
	DoorMembers                      PropertyType = 0xE4
	DoorOpenTooLongTime              PropertyType = 0xE5
	DoorPulseTime                    PropertyType = 0xE6
	DoorStatus                       PropertyType = 0xE7
	DoorUnlockDelayTime              PropertyType = 0xE8
	LockStatus                       PropertyType = 0xE9
	MaskedAlarmValues                PropertyType = 0xEA
	SecuredStatus                    PropertyType = 0xEB
	AbsenteeLimit                    PropertyType = 0xF4
	AccessAlarmEvents                PropertyType = 0xF5
	AccessDoors                      PropertyType = 0xF6
	AccessEvent                      PropertyType = 0xF7
	AccessEventAuthenticationFactor  PropertyType = 0xF8
	AccessEventCredential            PropertyType = 0xF9
	AccessEventTime                  PropertyType = 0xFA
	AccessTransactionEvents          PropertyType = 0xFB
	Accompaniment                    PropertyType = 0xFC
	AccompanimentTime                PropertyType = 0xFD
	ActivationTime                   PropertyType = 0xFE
	ActiveAuthenticationPolicy       PropertyType = 0xFF
	AssignedAccessRights             PropertyType = 0x100
	AuthenticationFactors            PropertyType = 0x101
	AuthenticationPolicyList         PropertyType = 0x102
	AuthenticationPolicyNames        PropertyType = 0x103
	AuthenticationStatus             PropertyType = 0x104
	AuthorizationMode                PropertyType = 0x105
	BelongsTo                        PropertyType = 0x106
	CredentialDisable                PropertyType = 0x107
	CredentialStatus                 PropertyType = 0x108
	Credentials                      PropertyType = 0x109
	CredentialsInZone                PropertyType = 0x10A
	DaysRemaining                    PropertyType = 0x10B
	EntryPoints                      PropertyType = 0x10C
	ExitPoints                       PropertyType = 0x10D
	ExpiryTime                       PropertyType = 0x10E
	ExtendedTimeEnable               PropertyType = 0x10F
	FailedAttemptEvents              PropertyType = 0x110
	FailedAttempts                   PropertyType = 0x111
	FailedAttemptsTime               PropertyType = 0x112
	LastAccessEvent                  PropertyType = 0x113
	LastAccessPoint                  PropertyType = 0x114
	LastCredentialAdded              PropertyType = 0x115
	LastCredentialAddedTime          PropertyType = 0x116
	LastCredentialRemoved            PropertyType = 0x117
	LastCredentialRemovedTime        PropertyType = 0x118
	LastUseTime                      PropertyType = 0x119
	Lockout                          PropertyType = 0x11A
	LockoutRelinquishTime            PropertyType = 0x11B
	MasterExemption                  PropertyType = 0x11C
	MaxFailedAttempts                PropertyType = 0x11D
	Members                          PropertyType = 0x11E
	MusterPoint                      PropertyType = 0x11F
	NegativeAccessRules              PropertyType = 0x120
	NumberOfAuthenticationPolicies   PropertyType = 0x121
	OccupancyCount                   PropertyType = 0x122
	OccupancyCountAdjust             PropertyType = 0x123
	OccupancyCountEnable             PropertyType = 0x124
	OccupancyExemption               PropertyType = 0x125
	OccupancyLowerLimit              PropertyType = 0x126
	OccupancyLowerLimitEnforced      PropertyType = 0x127
	OccupancyState                   PropertyType = 0x128
	OccupancyUpperLimit              PropertyType = 0x129
	OccupancyUpperLimitEnforced      PropertyType = 0x12A
	PassbackExemption                PropertyType = 0x12B
	PassbackMode                     PropertyType = 0x12C
	PassbackTimeout                  PropertyType = 0x12D
	PositiveAccessRules              PropertyType = 0x12E
	ReasonForDisable                 PropertyType = 0x12F
	SupportedFormats                 PropertyType = 0x130
	SupportedFormatClasses           PropertyType = 0x131
	ThreatAuthority                  PropertyType = 0x132
	ThreatLevel                      PropertyType = 0x133
	TraceFlag                        PropertyType = 0x134
	TransactionNotificationClass     PropertyType = 0x135
	UserExternalIdentifier           PropertyType = 0x136
	UserInformationReference         PropertyType = 0x137
	UserName                         PropertyType = 0x13D
	UserType                         PropertyType = 0x13E
	UsesRemaining                    PropertyType = 0x13F
	ZoneFrom                         PropertyType = 0x140
	ZoneTo                           PropertyType = 0x141
	AccessEventTag                   PropertyType = 0x142
	GlobalIdentifier                 PropertyType = 0x143
	VerificationTime                 PropertyType = 0x146
	BaseDeviceSecurityPolicy         PropertyType = 0x147
	DistributionKeyRevision          PropertyType = 0x148
	DoNotHide                        PropertyType = 0x149
	KeySets                          PropertyType = 0x14A
	LastKeyServer                    PropertyType = 0x14B
	NetworkAccessSecurityPolicies    PropertyType = 0x14C
	PacketReorderTime                PropertyType = 0x14D
	SecurityPduTimeout               PropertyType = 0x14E
	SecurityTimeWindow               PropertyType = 0x14F
	SupportedSecurityAlgorithm       PropertyType = 0x150
	UpdateKeySetTimeout              PropertyType = 0x151
	BackupAndRestoreState            PropertyType = 0x152
	BackupPreparationTime            PropertyType = 0x153
	RestoreCompletionTime            PropertyType = 0x154
	RestorePreparationTime           PropertyType = 0x155
	BitMask                          PropertyType = 0x156
	BitText                          PropertyType = 0x157
	IsUTC                            PropertyType = 0x158
	GroupMembers                     PropertyType = 0x159
	GroupMemberNames                 PropertyType = 0x15A
	MemberStatusFlags                PropertyType = 0x15B
	RequestedUpdateInterval          PropertyType = 0x15C
	CovuPeriod                       PropertyType = 0x15D
	CovuRecipients                   PropertyType = 0x15E
	EventMessageTexts                PropertyType = 0x15F
	EventMessageTextsConfig          PropertyType = 0x160
	EventDetectionEnable             PropertyType = 0x161
	EventAlgorithmInhibit            PropertyType = 0x162
	EventAlgorithmInhibitRef         PropertyType = 0x163
	TimeDelayNormal                  PropertyType = 0x164
	ReliabilityEvaluationInhibit     PropertyType = 0x165
	FaultParameters                  PropertyType = 0x166
	FaultType                        PropertyType = 0x167
	LocalForwardingOnly              PropertyType = 0x168
	ProcessIdentifierFilter          PropertyType = 0x169
	SubscribedRecipients             PropertyType = 0x16A
	PortFilter                       PropertyType = 0x16B
	AuthorizationExemptions          PropertyType = 0x16C
	AllowGroupDelayInhibit           PropertyType = 0x16D
	ChannelNumber                    PropertyType = 0x16E
	ControlGroups                    PropertyType = 0x16F
	ExecutionDelay                   PropertyType = 0x170
	LastPriority                     PropertyType = 0x171
	WriteStatus                      PropertyType = 0x172
	PropertyList                     PropertyType = 0x173
	SerialNumber                     PropertyType = 0x174
	BlinkWarnEnable                  PropertyType = 0x175
	DefaultFadeTime                  PropertyType = 0x176
	DefaultRampRate                  PropertyType = 0x177
	DefaultStepIncrement             PropertyType = 0x178
	EgressTime                       PropertyType = 0x179
	InProgress                       PropertyType = 0x17A
	InstantaneousPower               PropertyType = 0x17B
	LightingCommand                  PropertyType = 0x17C
	LightingCommandDefaultPriority   PropertyType = 0x17D
	MaxActualValue                   PropertyType = 0x17E
	MinActualValue                   PropertyType = 0x17F
	Power                            PropertyType = 0x180
	Transition                       PropertyType = 0x181
	EgressActive                     PropertyType = 0x182
)

//PropertyNullValue is the NULL value of one property
const PropertyNullValue = "NULL"
