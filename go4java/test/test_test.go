package main_test

import "testing"

func TestOneEqualsOne(t *testing.T) {
	if 1 != 1 {
		t.FailNow()
	}
}

func TestDisabled(t *testing.T) {
	t.Skip("this test is disabled")
	t.Run("subtest", func(t *testing.T) {
		t.Fail()
	})
}
