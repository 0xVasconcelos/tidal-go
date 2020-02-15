package tidal

type Client struct {
	SessionID string
}

type Media struct {
	Media         []byte
	EncryptionKey []byte
	Key           []byte
	Nonce         []byte
}
