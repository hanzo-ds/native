package proto

import "strings"

import "testing"

// Name is sent in ClientHello and lands in the server's system.query_log, so it
// is wire visible and must not carry the upstream vendor's name.
func TestClientHelloName(t *testing.T) {
	if Name != "datastore/native" {
		t.Errorf("Name = %q, want %q", Name, "datastore/native")
	}
	if strings.Contains(strings.ToLower(Name), "clickhouse") {
		t.Errorf("Name %q still carries the vendor brand", Name)
	}
}
