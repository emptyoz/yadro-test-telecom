package parser

import "testing"

func TestParseLine_Gold(t *testing.T) {
	tests := []struct {
		name      string
		line      string
		wantOK    bool
		wantErr   bool
		wantInt   int
		wantExtra string
	}{
		{"three_fields", "[14:00:00] 1 2", true, false, 0, ""},
		{"int_extra", "[14:27:00] 2 11 60", true, false, 60, ""},
		{"multi_string_extra", "[14:27:00] 2 9 network timeout on gateway", true, false, 0, "network timeout on gateway"},
		{"invalid_time", "[14:77:00] 1 2", false, true, 0, ""},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ev, ok, err := parseLine(tt.line, i+1)
			if tt.wantErr {
				if err == nil {
					t.Fatal("want error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if ok != tt.wantOK {
				t.Fatalf("want ok=%v, got %v", tt.wantOK, ok)
			}
			if !ok {
				return
			}
			if ev.IntParam != tt.wantInt || ev.ExtraParam != tt.wantExtra {
				t.Fatalf("unexpected extra: int=%d extra=%q", ev.IntParam, ev.ExtraParam)
			}
		})
	}
}
