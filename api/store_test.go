package api

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mock_db "github.com/ZoengYu/order-fast-project/db/mock"
	mockdb "github.com/ZoengYu/order-fast-project/db/mock"
	db "github.com/ZoengYu/order-fast-project/db/sqlc"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetStoreAPI(t *testing.T) {
	store := randonStore()

	testCases := []struct {
		name 			string
		body 			[]byte
		buildStubs 		func(mock_db *mockdb.MockDBService)
		checkResponse 	func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body : []byte(fmt.Sprintf(`{"name": "%s"}`, store.StoreName)),
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
					GetStoreByName(gomock.Any(), gomock.Eq(store.StoreName)).
					Times(1).
					Return(store, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "NotFound",
			body : []byte(fmt.Sprintf(`{"name": "%s"}`, store.StoreName)),
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
					GetStoreByName(gomock.Any(), gomock.Eq(store.StoreName)).
					Times(1).
					Return(db.Store{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusNotFound, recoder.Code)
			},
		},
		{
			name: "BadRequestPayload",
			body : []byte(fmt.Sprintf(`{"wrong": "%s"}`, store.StoreName)),
			buildStubs: func(mockdb *mockdb.MockDBService){
				mockdb.EXPECT().
					GetStoreByName(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockdb_service := mock_db.NewMockDBService(ctrl)
			tc.buildStubs(mockdb_service)
			server := newTestServer(t, mockdb_service)
			recorder := httptest.NewRecorder()
			url := "/v1/store"
			request, err := http.NewRequest(http.MethodGet, url, bytes.NewReader(tc.body))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randonStore() db.Store {
	return db.Store{
		ID:	1,
		StoreName: "Store A",
		StoreAddress: "fake address",
		StorePhone: "0123456789",
		StoreOwner: "Harry",
		StoreManager: "Alex",
		CreatedAt: time.Now(),
	}
}
