package entity

type LoginRequest struct {
	Email    string
	Password string
}

type RegisterRequest struct {
	Email       string
	Password    string
	BluetoothID string
}
