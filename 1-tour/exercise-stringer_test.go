package main

import "testing"

func TestIPAddr_String(t *testing.T) {
	tests := []struct {
		name string
		ip   IPAddr
		want string
	}{
		{"loopback", [4]byte{127, 0, 0, 1}, "127.0.0.1"},
		{"googleDNS", [4]byte{8, 8, 8, 8}, "8.8.8.8"},
		{"test", [4]byte{1, 2, 3, 4}, "1.2.3.4"},
		{"gateway", [4]byte{192, 168, 1, 254}, "192.168.1.254"},
	}
	for _, tt := range tests {
		if got := tt.ip.String(); got != tt.want {
			t.Errorf("%q. IPAddr.String() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
