package api

//func TestTransferAPI(t *testing.T) {
//	amount := int64(10)
//
//	user1, _ := randomUser(t)
//	user2, _ := randomUser(t)
//	user3, _ := randomUser(t)
//
//	account1 := randomAccount(user1.Username)
//	account2 := randomAccount(user2.Username)
//	account3 := randomAccount(user3.Username)
//
//	account1.Currency = util.USD
//	account2.Currency = util.USD
//	account3.Currency = util.EUR
//
//	testCase := []struct {
//		name          string
//		body          gin.H
//		buildStubs    func(store *mockdb.MockStore)
//		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
//	}{
//		{
//			name: "OK",
//			body: gin.H{
//				"from_account_id": account1.ID,
//				"to_account_id":   account2.ID,
//				"amount":          amount,
//				"currency":        util.USD,
//			},
//			buildStubs: func(store *mockdb.MockStore) {
//				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account1.ID)).Times(1).Return(account1, nil)
//				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account2.ID)).Times(1).Return(account2, nil)
//				arg := db.TransferTxParams{
//					FromAccountID: account1.ID,
//					ToAccountID:   account2.ID,
//					Amount:        amount,
//				}
//				store.EXPECT().
//					TranferTx(gomock.Any(), gomock.Eq(arg)).
//					Times(1)
//			},
//			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
//				require.Equal(t, http.StatusOK, recorder.Code)
//			},
//		},
//		{
//			name: "FromAccountNotFound",
//			body: gin.H{
//				"from_account_id": account1.ID,
//				"to_account_id":   account2.ID,
//				"amount":          amount,
//				"currency":        util.USD,
//			},
//			buildStubs: func(store *mockdb.MockStore) {
//				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account1.ID)).Times(1).Return(db.Account{})
//			},
//			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
//				require.Equal(t, http.StatusNotFound, recorder.Code)
//			},
//		},
//		{
//			name: "GetAccountErr",
//			body: gin.H{
//				"from_account_id": account1.ID,
//				"to_account_id":   account2.ID,
//				"amount":          amount,
//				"currency":        util.USD,
//			},
//
//			buildStubs: func(store *mockdb.MockStore) {
//				store.EXPECT().GetAccount(gomock.Any(), gomock.Any()).Times(1).Return(db.Account{}, sql.ErrConnDone)
//				store.EXPECT().GetAccount(gomock.Any(), gomock.Any()).Times(0)
//			},
//			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
//				require.Equal(t, http.StatusForbidden, recorder.Code)
//			},
//		},
//		{
//			name: "InvalidUsername",
//			body: gin.H{
//				"username":  "invalid-user#1",
//				"password":  password,
//				"full_name": user.FullName,
//				"email":     user.Email,
//			},
//			buildStubs: func(store *mockdb.MockStore) {
//				store.EXPECT().
//					CreateUser(gomock.Any(), gomock.Any()).
//					Times(0)
//			},
//			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
//				require.Equal(t, http.StatusBadRequest, recorder.Code)
//			},
//		},
//		{
//			name: "InvalidEmail",
//			body: gin.H{
//				"username":  user.Username,
//				"password":  password,
//				"full_name": user.FullName,
//				"email":     "invalid-email",
//			},
//			buildStubs: func(store *mockdb.MockStore) {
//				store.EXPECT().
//					CreateUser(gomock.Any(), gomock.Any()).
//					Times(0)
//			},
//			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
//				require.Equal(t, http.StatusBadRequest, recorder.Code)
//			},
//		},
//		{
//			name: "TooShortPassword",
//			body: gin.H{
//				"username":  user.Username,
//				"password":  "123",
//				"full_name": user.FullName,
//				"email":     user.Email,
//			},
//			buildStubs: func(store *mockdb.MockStore) {
//				store.EXPECT().
//					CreateUser(gomock.Any(), gomock.Any()).
//					Times(0)
//			},
//			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
//				require.Equal(t, http.StatusBadRequest, recorder.Code)
//			},
//		},
//	}
//
//	for i := range testCase {
//		tc := testCase[i]
//
//		t.Run(tc.name, func(t *testing.T) {
//			ctrl := gomock.NewController(t)
//			defer ctrl.Finish()
//
//			store := mockdb.NewMockStore(ctrl)
//			tc.buildStubs(store)
//
//			//start server and send request
//			server := newTestServer(t, store)
//			recorder := httptest.NewRecorder()
//
//			// Marshal Body data to JSON
//			data, err := json.Marshal(tc.body)
//			require.NoError(t, err)
//
//			url := "/users"
//			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
//			require.NoError(t, err)
//
//			server.router.ServeHTTP(recorder, request)
//
//			//Check Response
//			tc.checkResponse(t, recorder)
//		})
//	}
//}
