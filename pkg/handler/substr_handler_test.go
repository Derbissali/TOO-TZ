package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"tidy/pkg/model"
	"tidy/pkg/service"
	mock_api "tidy/pkg/service/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Substr(t *testing.T) {
	type mockBehavior func(s *mock_api.MockSubstringService, substr model.Substring)

	testTable := []struct {
		name                string
		inputBody           string
		inputSub            model.Substring
		mockBehavior        mockBehavior
		expectedStatus      int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"substring":"Test"}`,
			inputSub: model.Substring{
				Substring: "Test",
			},
			mockBehavior: func(s *mock_api.MockSubstringService, substr model.Substring) {
				s.EXPECT().MaxLength(&substr.Substring).Return("Test", nil)
			},
			expectedStatus:      200,
			expectedRequestBody: `"substring":"Test"`,
		},
		{
			name:      "BadRequest",
			inputBody: `{"substring":"фыв"}`,
			inputSub: model.Substring{
				Substring: "фыв",
			},
			mockBehavior: func(s *mock_api.MockSubstringService, substr model.Substring) {
				s.EXPECT().MaxLength(&substr.Substring).Return("", errors.New("bad req"))
			},
			expectedStatus:      400,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			sub := mock_api.NewMockSubstringService(c)

			testCase.mockBehavior(sub, testCase.inputSub)
			services := &service.Service{SubstringService: sub}

			mux := http.NewServeMux()
			handler := NewHandler(services)
			handler.Register(mux)

			w := httptest.NewRecorder()

			req := httptest.NewRequest("POST", "/rest/substr/find", bytes.NewBufferString(testCase.inputBody))

			mux.ServeHTTP(w, req)
			//Assert

			assert.Equal(t, testCase.expectedStatus, w.Code)
		})
	}
}
