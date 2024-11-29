package main

import "testing"

func TestParseHost(t *testing.T) {
	tests := []struct {
		host       string
		originUser string
		originPort string
		wantUser   string
		wantIP     string
		wantPort   string
	}{
		{
			host:       "user@ip",
			originUser: "",
			originPort: "",
			wantUser:   "user",
			wantIP:     "ip",
			wantPort:   "22",
		},
		{
			host:       "ip",
			originUser: "root",
			originPort: "",
			wantUser:   "root",
			wantIP:     "ip",
			wantPort:   "22",
		},
		{
			host:       "ip:port",
			originUser: "",
			originPort: "",
			wantUser:   "",
			wantIP:     "ip",
			wantPort:   "port",
		},
		{
			host:       "ip:port",
			originUser: "foo",
			originPort: "",
			wantUser:   "foo",
			wantIP:     "ip",
			wantPort:   "port",
		},
		{
			host:       "user@ip:port",
			originUser: "",
			originPort: "",
			wantUser:   "user",
			wantIP:     "ip",
			wantPort:   "port",
		},
	}

	for _, test := range tests {
		user, ip, port := parseHost(test.host, test.originUser, test.originPort)
		if user != test.wantUser {
			t.Errorf("parseHost(%q, %q, %q) = user %q, want %q", test.host, test.originUser, test.originPort, user, test.wantUser)
		}
		if ip != test.wantIP {
			t.Errorf("parseHost(%q, %q, %q) = ip %q, want %q", test.host, test.originUser, test.originPort, ip, test.wantIP)
		}
		if port != test.wantPort {
			t.Errorf("parseHost(%q, %q, %q) = port %q, want %q", test.host, test.originUser, test.originPort, port, test.wantPort)
		}
	}
}
