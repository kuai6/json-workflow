package jsonworkflow

// EventType general events dictionary
type EventType string

const (
	// EventContactClosure ...
	EventContactClosure EventType = "ContactClosure"
	// EventADC ...
	EventADC EventType = "ADC"
)

// Direction general signal changes direction dictionary
type Direction uint8

const (
	// DirectionOnFall ...
	DirectionOnFall Direction = iota
	// DirectionOnRise ...
	DirectionOnRise
	// DirectionOnChange ...
	DirectionOnChange
)

// ActivityClass general activity classes
type ActivityClass string

const (
	// ActivityClassAssign ...
	ActivityClassAssign ActivityClass = "assign"
	// ActivityClassCondition ...
	ActivityClassCondition ActivityClass = "condition"
)

// DeviceBank general device banks dictionary
type DeviceBank uint8

const (
	// BANK0 ...
	BANK0 DeviceBank = iota
	// BANK1 ...
	BANK1
	// BANK2 ...
	BANK2
	// BANK3 ...
	BANK3
)

// DeviceBit general device bits
type DeviceBit uint8

const (
	// BIT0 ...
	BIT0 DeviceBit = iota
	// BIT1 ...
	BIT1
	// BIT2 ...
	BIT2
	// BIT3 ...
	BIT3
	// BIT4 ...
	BIT4
	// BIT5 ...
	BIT5
	// BIT6 ...
	BIT6
	// BIT7 ...
	BIT7
)
