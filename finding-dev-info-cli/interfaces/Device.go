package interfaces

type Device struct {
    DeviceType       string `json:"device_type"`
    DeviceName       string `json:"device_name"`
    DeviceModelNumber string `json:"device_model_number"`
}