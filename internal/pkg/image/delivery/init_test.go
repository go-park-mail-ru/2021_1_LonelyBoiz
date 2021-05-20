package delivery

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	usecaseMocks "server/internal/pkg/image/usecase/mocks"
	"server/internal/pkg/models"
	"strings"

	"testing"

	"server/internal/pkg/utils/metrics"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestAddToSecreteAlbum(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	metrics.New()

	imageUseCaseMock := usecaseMocks.NewMockImageUsecaseInterface(mockCtrl)

	handlerTest := ImageHandler{
		Usecase: imageUseCaseMock,
	}

	user := models.User{
		Id:     1,
		Photos: []string{"1", "2"},
	}

	murl, er := url.Parse("/images")
	if er != nil {
		t.Error(er)
	}

	body := ioutil.NopCloser(strings.NewReader("Image bytes"))

	req := &http.Request{
		Method: "POST",
		URL:    murl,
		Body:   body,
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	rw := httptest.NewRecorder()

	imageUseCaseMock.EXPECT().GetIdFromContext(ctx).Return(user.Id, true)
	imageUseCaseMock.EXPECT().AddImage(user.Id, gomock.Any()).Return(models.Image{}, nil)
	imageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.UploadImage(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestAddToSecreteAlbum_ReadBody_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	imageUseCaseMock := usecaseMocks.NewMockImageUsecaseInterface(mockCtrl)

	handlerTest := ImageHandler{
		Usecase: imageUseCaseMock,
	}

	user := models.User{
		Id:     1,
		Photos: []string{"1", "2"},
	}

	murl, er := url.Parse("/images")
	if er != nil {
		t.Error(er)
	}

	body := ioutil.NopCloser(strings.NewReader(""))

	req := &http.Request{
		Method: "POST",
		URL:    murl,
		Body:   body,
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	rw := httptest.NewRecorder()

	imageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.UploadImage(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestAddToSecreteAlbum_GetIdFromContext_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	imageUseCaseMock := usecaseMocks.NewMockImageUsecaseInterface(mockCtrl)

	handlerTest := ImageHandler{
		Usecase: imageUseCaseMock,
	}

	user := models.User{
		Id:     1,
		Photos: []string{"1", "2"},
	}

	murl, er := url.Parse("/images")
	if er != nil {
		t.Error(er)
	}

	body := ioutil.NopCloser(strings.NewReader("Image bytes"))

	req := &http.Request{
		Method: "POST",
		URL:    murl,
		Body:   body,
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	rw := httptest.NewRecorder()

	imageUseCaseMock.EXPECT().GetIdFromContext(ctx).Return(user.Id, false)
	imageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.UploadImage(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 403, response.StatusCode)
}

func TestAddToSecreteAlbum_AddImage_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	imageUseCaseMock := usecaseMocks.NewMockImageUsecaseInterface(mockCtrl)

	handlerTest := ImageHandler{
		Usecase: imageUseCaseMock,
	}

	user := models.User{
		Id:     1,
		Photos: []string{"1", "2"},
	}

	murl, er := url.Parse("/images")
	if er != nil {
		t.Error(er)
	}

	body := ioutil.NopCloser(strings.NewReader("Image bytes"))

	req := &http.Request{
		Method: "POST",
		URL:    murl,
		Body:   body,
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	rw := httptest.NewRecorder()

	imageUseCaseMock.EXPECT().GetIdFromContext(ctx).Return(user.Id, true)
	imageUseCaseMock.EXPECT().AddImage(user.Id, gomock.Any()).Return(models.Image{}, errors.New("Some error"))
	imageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.UploadImage(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestDeleteImage(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	imageUseCaseMock := usecaseMocks.NewMockImageUsecaseInterface(mockCtrl)

	handlerTest := ImageHandler{
		Usecase: imageUseCaseMock,
	}

	user := models.User{
		Id:     1,
		Photos: []string{"1", "2"},
	}

	murl, er := url.Parse("/images")
	if er != nil {
		t.Error(er)
	}

	body := ioutil.NopCloser(strings.NewReader(""))

	req := &http.Request{
		Method: "POST",
		URL:    murl,
		Body:   body,
	}
	vars := map[string]string{
		"uuid": "3f35821d-e101-4d7f-8e60-3a78bad1f951",
	}
	req = mux.SetURLVars(req, vars)

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	rw := httptest.NewRecorder()

	imageUseCaseMock.EXPECT().GetIdFromContext(gomock.Any()).Return(user.Id, true)
	imageUseCaseMock.EXPECT().DeleteImage(user.Id, gomock.Any()).Return(nil)
	imageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.DeleteImage(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 204, response.StatusCode)
}

func TestDeleteImage_GetIdFromContext_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	imageUseCaseMock := usecaseMocks.NewMockImageUsecaseInterface(mockCtrl)

	handlerTest := ImageHandler{
		Usecase: imageUseCaseMock,
	}

	user := models.User{
		Id:     1,
		Photos: []string{"1", "2"},
	}

	murl, er := url.Parse("/images")
	if er != nil {
		t.Error(er)
	}

	body := ioutil.NopCloser(strings.NewReader(""))

	req := &http.Request{
		Method: "POST",
		URL:    murl,
		Body:   body,
	}
	vars := map[string]string{
		"uuid": "3f35821d-e101-4d7f-8e60-3a78bad1f951",
	}
	req = mux.SetURLVars(req, vars)

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	rw := httptest.NewRecorder()

	imageUseCaseMock.EXPECT().GetIdFromContext(gomock.Any()).Return(user.Id, false)
	imageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.DeleteImage(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 403, response.StatusCode)
}

func TestDeleteImage_DeleteImage_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	imageUseCaseMock := usecaseMocks.NewMockImageUsecaseInterface(mockCtrl)

	handlerTest := ImageHandler{
		Usecase: imageUseCaseMock,
	}

	user := models.User{
		Id:     1,
		Photos: []string{"1", "2"},
	}

	murl, er := url.Parse("/images")
	if er != nil {
		t.Error(er)
	}

	body := ioutil.NopCloser(strings.NewReader(""))

	req := &http.Request{
		Method: "POST",
		URL:    murl,
		Body:   body,
	}
	vars := map[string]string{
		"uuid": "3f35821d-e101-4d7f-8e60-3a78bad1f951",
	}
	req = mux.SetURLVars(req, vars)

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	rw := httptest.NewRecorder()

	imageUseCaseMock.EXPECT().GetIdFromContext(gomock.Any()).Return(user.Id, true)
	imageUseCaseMock.EXPECT().DeleteImage(user.Id, gomock.Any()).Return(errors.New("Some error"))
	imageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.DeleteImage(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestDeleteImage_Vars_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	imageUseCaseMock := usecaseMocks.NewMockImageUsecaseInterface(mockCtrl)

	handlerTest := ImageHandler{
		Usecase: imageUseCaseMock,
	}

	user := models.User{
		Id:     1,
		Photos: []string{"1", "2"},
	}

	murl, er := url.Parse("/images")
	if er != nil {
		t.Error(er)
	}

	body := ioutil.NopCloser(strings.NewReader(""))

	req := &http.Request{
		Method: "POST",
		URL:    murl,
		Body:   body,
	}
	vars := map[string]string{
		"notUuid": "3f35821d-e101-4d7f-8e60-3a78bad1f951",
	}
	req = mux.SetURLVars(req, vars)

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	rw := httptest.NewRecorder()

	imageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.DeleteImage(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}
