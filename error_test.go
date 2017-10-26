package patreon

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErrorResponse(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/oauth2/api/current_user", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusForbidden)
		fmt.Fprint(writer, errorResp)
	})

	_, err := client.FetchUser()
	require.Error(t, err)
	require.Equal(t, "The server could not verify that you are authorized to access the URL requested.", err.Error())

	errResp, ok := err.(ErrorResponse)
	require.True(t, ok)
	require.Equal(t, 1, len(errResp.Errors))
	require.Equal(t, 1, errResp.Errors[0].Code)
	require.Equal(t, "Unauthorized", errResp.Errors[0].CodeName)
	require.Equal(t, "bb16af2c-c11e-4796-af4d-19fa60007709", errResp.Errors[0].ID)
	require.Equal(t, "401", errResp.Errors[0].Status)
	require.Equal(t, "Unauthorized", errResp.Errors[0].Title)
	require.Equal(t, "The server could not verify that you are authorized to access the URL requested.", errResp.Errors[0].Detail)
}

func TestDefaultErrorString(t *testing.T) {
	err := ErrorResponse{}
	require.Equal(t, "(ERR)", err.Error())
}

const errorResp = `
{
    "errors": [
        {
            "code": 1,
            "code_name": "Unauthorized",
            "detail": "The server could not verify that you are authorized to access the URL requested.",
            "id": "bb16af2c-c11e-4796-af4d-19fa60007709",
            "status": "401",
            "title": "Unauthorized"
        }
    ]
}
`
