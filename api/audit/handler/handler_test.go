package handler

import (
	"github.com/gorilla/mux"
	"github.com/zopdev/zopdev/api/audit/store"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gofr.dev/pkg/gofr"
	gofrHttp "gofr.dev/pkg/gofr/http"
)

func TestHandler_RunAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := NewMockService(ctrl)
	handler := New(mockService)

	testCases := []struct {
		name          string
		pathParam     string
		expectedError error
		mockResponse  map[string][]*store.Result
		mockError     error
	}{
		{
			name:          "Missing ID",
			pathParam:     "",
			expectedError: gofrHttp.ErrorMissingParam{Params: []string{"id"}},
		},
		{
			name:          "Invalid ID",
			pathParam:     "abc",
			expectedError: gofrHttp.ErrorInvalidParam{Params: []string{"id"}},
		},
		{
			name:         "Success",
			pathParam:    "123",
			mockResponse: map[string][]*store.Result{},
			mockError:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, "/api/audit/all", nil)
			r = mux.SetURLVars(r, map[string]string{"id": tc.pathParam})

			ctx := &gofr.Context{
				Request: gofrHttp.NewRequest(r),
			}

			if tc.mockResponse != nil || tc.mockError != nil {
				mockService.EXPECT().RunAll(ctx, int64(123)).
					Return(tc.mockResponse, tc.mockError)
			}

			resp, err := handler.RunAll(ctx)

			assert.Equal(t, tc.expectedError, err)
			if err == nil {
				assert.Equal(t, tc.mockResponse, resp)
			}
		})
	}
}

func TestHandler_RunById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := NewMockService(ctrl)
	handler := New(mockService)

	testCases := []struct {
		name          string
		ruleId        string
		cloudAccId    string
		expectedError error
		mockResponse  *store.Result
		mockError     error
	}{
		{
			name:          "Missing ID",
			cloudAccId:    "",
			expectedError: gofrHttp.ErrorMissingParam{Params: []string{"id"}},
		},
		{
			name:          "Invalid ID",
			cloudAccId:    "abc",
			expectedError: gofrHttp.ErrorInvalidParam{Params: []string{"id"}},
		},
		{
			name:          "Missing Rule ID",
			cloudAccId:    "123",
			ruleId:        "",
			expectedError: gofrHttp.ErrorMissingParam{Params: []string{"ruleId"}},
		},
		{
			name:         "Success",
			cloudAccId:   "123",
			ruleId:       "rule-1",
			mockResponse: &store.Result{},
			mockError:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, "/api/audit/cloud-accounts/{id}/rule/{ruleId}", nil)
			r = mux.SetURLVars(r, map[string]string{"id": tc.cloudAccId, "ruleId": tc.ruleId})

			ctx := &gofr.Context{
				Request: gofrHttp.NewRequest(r),
			}

			if tc.mockResponse != nil || tc.mockError != nil {
				cId, _ := strconv.ParseInt(tc.cloudAccId, 10, 64)
				mockService.EXPECT().RunById(ctx, tc.ruleId, cId).
					Return(tc.mockResponse, tc.mockError)
			}

			resp, err := handler.RunById(ctx)

			assert.Equal(t, tc.expectedError, err)
			if err == nil {
				assert.Equal(t, tc.mockResponse, resp)
			}
		})
	}
}

func TestHandler_RunByCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := NewMockService(ctrl)
	handler := New(mockService)

	testCases := []struct {
		name          string
		categoryId    string
		cloudAccId    string
		expectedError error
		mockResponse  []*store.Result
		mockError     error
	}{
		{
			name:          "Missing ID",
			cloudAccId:    "",
			expectedError: gofrHttp.ErrorMissingParam{Params: []string{"id"}},
		},
		{
			name:          "Invalid ID",
			cloudAccId:    "overprovision",
			expectedError: gofrHttp.ErrorInvalidParam{Params: []string{"id"}},
		},
		{
			name:          "Missing Category",
			cloudAccId:    "123",
			categoryId:    "",
			expectedError: gofrHttp.ErrorMissingParam{Params: []string{"category"}},
		},
		{
			name:         "Success",
			cloudAccId:   "123",
			categoryId:   "overprovision",
			mockResponse: []*store.Result{},
			mockError:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, "/api/audit/cloud-account/{id}/category/{category}", nil)
			r = mux.SetURLVars(r, map[string]string{"id": tc.cloudAccId, "category": tc.categoryId})

			ctx := &gofr.Context{
				Request: gofrHttp.NewRequest(r),
			}

			if tc.mockResponse != nil || tc.mockError != nil {
				cId, _ := strconv.ParseInt(tc.cloudAccId, 10, 64)
				mockService.EXPECT().RunByCategory(ctx, tc.categoryId, cId).
					Return(tc.mockResponse, tc.mockError)
			}

			resp, err := handler.RunByCategory(ctx)

			assert.Equal(t, tc.expectedError, err)
			if err == nil {
				assert.Equal(t, tc.mockResponse, resp)
			}
		})
	}
}
