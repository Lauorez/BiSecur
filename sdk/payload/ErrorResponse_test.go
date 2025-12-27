package payload

import (
	"encoding/hex"
	"testing"
)

func TestErrorResponseDecode(t *testing.T) {
	testCases := []struct {
		Name                 string
		EncodedInput         string
		ExpectedDecodedInput ErrorResponse
	}{
		{
			Name:         "PERMISSION_DENIED error decode",
			EncodedInput: "0C", // Full response message: 5410EC8528BB000000000006000A03A53BBCE3010C9955
			ExpectedDecodedInput: ErrorResponse{
				Payload{
					data: []byte{ERROR_PERMISSION_DENIED},
				},
			},
		},
		{
			Name:         "PORT_ERROR error decode",
			EncodedInput: "0A", // Full response message: 5410EC8528BB000000000006000A0194313F4B010A6522
			ExpectedDecodedInput: ErrorResponse{
				Payload{
					data: []byte{ERROR_PORT_ERROR},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(test *testing.T) {
			encodedInput := testCase.EncodedInput
			encodedInputBytes, err := hex.DecodeString(encodedInput)
			if err != nil {
				test.Logf("Failed to decode test data. %v", err)
				test.Fail()
				return
			}

			decoded, err := DecodeErrorPayload(encodedInputBytes)
			if err != nil {
				test.Logf("Failed to decode response. %v", err)
				test.Fail()
				return
			}

			expected := &testCase.ExpectedDecodedInput

			if !expected.Equal(decoded.(*ErrorResponse)) {
				test.Logf("Expected value: %v, Actual value: %v", expected, decoded)
				test.Fail()
			}
			test.Logf("Decoded error response: %v", decoded)
		})
	}
}
