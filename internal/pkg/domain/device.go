package domain

type DeviceType string

const (
	DeviceTypeIOS               DeviceType = "ios"
	DeviceTypeAndroid           DeviceType = "android"
	DeviceTypeAndroidS3         DeviceType = "android-s3"
	DeviceTypeAndroidGooglePlay DeviceType = "android-google-play"
	DeviceTypeAndroidHuawei     DeviceType = "android-huawei"
	DeviceTypeUnknown           DeviceType = "unknown"
)

type Device struct {
	GUID  string
	Token string
	Type  DeviceType
}
