package tests

import (
	"errors"
	"fmt"
	"go_server/helper"
	"regexp"
	"testing"
)

func TestGetCurrentDateTime(t *testing.T) {
	resultDate := helper.GetCurrentDateTime()
	// Regular expression to match the "dd-mm-yyyy hh:mm:ss" format.
	re := regexp.MustCompile(`^\d{2}-\d{2}-\d{4} \d{2}:\d{2}:\d{2}$`)

	match := re.MatchString(resultDate)

	if !match {
		t.Errorf("GetCurrentDateTime() = %v, want match", resultDate)
	}

	fmt.Println(resultDate)
}

func TestGetEnvVariable(t *testing.T) {
	// tests edge cases for env variables returns a string instance or an error
	result := helper.GetEnvVariable("TEST_ENV_VAR")
	if result == "" {
		t.Errorf("GetEnvVariable() = %v, want match", result)
	}

}

func TestValidateEmailAddressEdgeCases(t *testing.T) {
	testCases := []struct {
		email string
		want  bool
	}{
		{"", false},
		{"@", false},
		{"test@", false},
		{"@test.com", false},
		{"test@.com", false},
		{"test.test@test.com", true},
		{"test+filter@test.com", true},
		{"test.test@test.co.uk", true},
		{"test.test@test.co.za", true},
	}

	for _, tc := range testCases {
		isEmail, err := helper.ValidateEmailAddress(tc.email)
		if errors.Is(err, fmt.Errorf("invalid email format")) {
			t.Fatalf("ValidateEmailAddress returned an error: %v", err)
		}
		if isEmail != tc.want {
			t.Errorf("ValidateEmailAddress(%q) = %v, want %v", tc.email, isEmail, tc.want)
		}
	}
}
