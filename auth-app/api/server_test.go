package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"auth-app/api"
	"auth-app/model"
	"auth-app/store"
	"auth-app/token"
	"auth-app/util"
	"auth-app/util/sample"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestRegister(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		checkResponse func(recoder *httptest.ResponseRecorder)
		name          string
		user          *model.User
	}{}

	for _, role := range util.AvailableRoles {
		cases := []struct {
			checkResponse func(recoder *httptest.ResponseRecorder)
			name          string
			user          *model.User
		}{
			{
				checkResponse: func(recorder *httptest.ResponseRecorder) {
					require.Equal(t, http.StatusOK, recorder.Code)
				},
				name: "success_" + role,
				user: sample.NewUser(role),
			},
			{
				checkResponse: func(recorder *httptest.ResponseRecorder) {
					require.Equal(t, http.StatusBadRequest, recorder.Code)
				},
				name: "failure_invalid_argument_case_" + role,
				user: sample.NewInvalidUser(),
			},
		}

		testCases = append(testCases, cases...)
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			apiServer := createApiServer(t)
			recorder := httptest.NewRecorder()

			body := gin.H{
				"name":  tc.user.Name,
				"phone": tc.user.Phone,
				"role":  tc.user.Role,
			}

			// Marshal body data to JSON
			data, err := json.Marshal(body)
			require.NoError(t, err)

			url := api.REGISTER_ROUTE
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			apiServer.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

// ================================ utilities ================================
func createApiServer(t *testing.T) *api.ApiServer {
	// load environment variables from .env file
	config, err := util.LoadConfig("../.")
	require.NoError(t, err)

	tokenManager, err := token.NewJWTManager(config.Secret)
	require.NoError(t, err)

	userStore := store.NewInMemoryUserStore("../data/user_data")

	err = userStore.PopulateDataFromFile()
	if err != nil {
		return nil
	}

	apiServer := api.NewApiServer(config, userStore, tokenManager)

	return apiServer
}
