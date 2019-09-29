package car

// YarisBuilder defines toyota yaris builder property
type YarisBuilder struct {
	Model              string
	FrontBrake         string
	FrontDiscBrakeType string
	RearBrake          string
	RearDiscBrakeType  string
	Transmission       string
	AutomaticType      string
}

// NewYarisBuilder will create new yaris builder instance
func NewYarisBuilder() *YarisBuilder {
	return &YarisBuilder{}
}

// SetModel will set yaris model type
func (y *YarisBuilder) SetModel(model string) *YarisBuilder {
	y.Model = model
	return y
}

// SetFrontBrake will set front brake type
func (y *YarisBuilder) SetFrontBrake(brake string) *YarisBuilder {
	y.FrontBrake = brake
	return y
}

// SetFrontDiscBrakeType will set front disc brake type
func (y *YarisBuilder) SetFrontDiscBrakeType(discType string) *YarisBuilder {
	y.FrontDiscBrakeType = discType
	return y
}

// SetRearBrake will set rear brake type
func (y *YarisBuilder) SetRearBrake(brake string) *YarisBuilder {
	y.RearBrake = brake
	return y
}

// SetRearDiscBrakeType will set read disc brake type
func (y *YarisBuilder) SetRearDiscBrakeType(discType string) *YarisBuilder {
	y.RearDiscBrakeType = discType
	return y
}

// SetTransmission will set transmission type
func (y *YarisBuilder) SetTransmission(gear string) *YarisBuilder {
	y.Transmission = gear
	return y
}

// SetAutomaticType will set automatic transmission type
func (y *YarisBuilder) SetAutomaticType(maticType string) *YarisBuilder {
	y.AutomaticType = maticType
	return y
}

// Build will create toyota yaris instance
func (y *YarisBuilder) Build() *Yaris {
	return &Yaris{
		Model:              y.Model,
		FrontBrake:         y.FrontBrake,
		FrontDiscBrakeType: y.FrontDiscBrakeType,
		RearBrake:          y.RearBrake,
		RearDiscBrakeType:  y.RearDiscBrakeType,
		Transmission:       y.Transmission,
		AutomaticType:      y.AutomaticType,
	}
}
