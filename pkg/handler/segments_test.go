package handler

import (
	"bytes"
	"dynamic-segmentation-service/pkg/model"
	"dynamic-segmentation-service/pkg/service"
	service_mocks "dynamic-segmentation-service/pkg/service/mocks"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler_createSegment(t *testing.T) {
	type mockBehavior func(s *service_mocks.MockSegment, segment model.Segment)

	tests := []struct {
		name                 string
		inputBody            string
		inputSegment         model.Segment
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"slug": "test"}`,
			inputSegment: model.Segment{
				Slug: "test",
			},
			mockBehavior: func(r *service_mocks.MockSegment, segment model.Segment) {
				r.EXPECT().CreateSegment(segment).Return(model.Segment{
					Id:        1,
					Slug:      "test",
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
					Users:     nil,
				}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"segment":{"id":1,"slug":"test","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","users":null}}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `{}`,
			inputSegment:         model.Segment{},
			mockBehavior:         func(r *service_mocks.MockSegment, segment model.Segment) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"slug": "test"}`,
			inputSegment: model.Segment{
				Id:   0,
				Slug: "test",
			},
			mockBehavior: func(r *service_mocks.MockSegment, segment model.Segment) {
				r.EXPECT().CreateSegment(segment).Return(model.Segment{Id: 0, Slug: "test"}, errors.New("some error"))
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"some error"}`,
		},
	}
	const (
		urlSegmentsHandler    = "/segments"
		methodSegmentsHandler = "POST"
	)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := service_mocks.NewMockSegment(c)
			test.mockBehavior(repo, test.inputSegment)

			services := &service.Service{Segment: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST(urlSegmentsHandler, handler.createSegment)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(methodSegmentsHandler, urlSegmentsHandler,
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_deleteSegment(t *testing.T) {
	type mockBehavior func(s *service_mocks.MockSegment, id int)

	tests := []struct {
		name                 string
		inputBody            string
		inputId              int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: "1",
			inputId:   1,
			mockBehavior: func(r *service_mocks.MockSegment, id int) {
				r.EXPECT().DeleteSegment(id).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `null`,
		},
		{
			name:                 "Wrong Url ID",
			inputBody:            "test",
			inputId:              1,
			mockBehavior:         func(r *service_mocks.MockSegment, id int) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"id is of invalid type"}`,
		},
	}
	const (
		urlSegmentsHandler    = "/segments"
		methodSegmentsHandler = "DELETE"
	)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := service_mocks.NewMockSegment(c)
			test.mockBehavior(repo, test.inputId)

			services := &service.Service{Segment: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.DELETE(urlSegmentsHandler+"/:id", handler.deleteSegment)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(methodSegmentsHandler, urlSegmentsHandler+"/"+test.inputBody,
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
