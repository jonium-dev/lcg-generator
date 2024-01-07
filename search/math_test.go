package search

import "testing"

func TestGCDIsSomewhatCorrect(t *testing.T) {
	if gcd := GCD(10, 9); gcd != 1 {
		t.Fatalf("Failed to produce correct GCD (%d)", gcd)
	}
	if gcd := GCD(10, 30); gcd != 10 {
		t.Fatalf("Failed to produce correct GCD (%d)", gcd)
	}
	if gcd := GCD(9, 10); gcd != 1 {
		t.Fatalf("Failed to produce correct GCD (%d)", gcd)
	}
	if gcd := GCD(30, 10); gcd != 10 {
		t.Fatalf("Failed to produce correct GCD (%d)", gcd)
	}
}
