package mecomo

type Device struct {
	Id int `xml:"DEVICE>DEVICE_ID"`
	Name string `xml:"DEVICE>DEVICE_NAME"`
	RegDate string `xml:"DEVICE>DEVICE_REGDATE"`
	Description string `xml:"DEVICE>DEVICE_DESCRIPTION"`
	Serial string `xml:"DEVICE>DEVICE_SERIAL"`
	GpsId string `xml:"DEVICE>DEVICE_GPS_ID"`
	Provider int64 `xml:"DEVICE>DEVICE_PROVIDER"`
	ProviderName string `xml:"DEVICE>DEVICE_PROVIDER_NAME"`
	Msisdn string `xml:"DEVICE>DEVICE_MSISDN"`
	Simcard string `xml:"DEVICE>DEVICE_SIMCARD"`
	NetworkOperator string `xml:"DEVICE>DEVICE_NETWORK_OPERATOR"`
	ObjectId int `xml:"OBJECT>OBJECT_ID"`
	ObjectName string `xml:"OBJECT>OBJECT_NAME"`
}

type DeviceHeader struct {
	Id int `xml:"DEVICE_ID"`
	Name string `xml:"DEVICE_NAME"`
}

type Telemetry struct {

}

