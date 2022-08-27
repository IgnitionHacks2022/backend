package entity

type LoginRequest struct {
	Email    string
	Password string
}

type RegisterRequest struct {
	Email       string
	Name        string
	Password    string
	BluetoothID string
}
