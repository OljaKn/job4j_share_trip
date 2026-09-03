package api_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"job4j.ru/go-share-trip/internal/api"
)

func TestServer_CreateTrip(t *testing.T) {
	t.Run("success - создание поездки", func(t *testing.T) {
		payload := api.CreateTripRequest{
			DriverId:      uuid.New(),
			FromPoint:     "Казань",
			ToPoint:       "Москва",
			DepartureTime: time.Now().Add(1 * time.Hour).UTC().Truncate(time.Second),
			Seats:         2,
		}

		body, err := json.Marshal(payload)
		require.NoError(t, err)

		req, err := http.NewRequest(
			http.MethodPost,
			"/trip/create",
			bytes.NewReader(body),
		)
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		resp, err := testApp.Test(req, -1)
		require.NoError(t, err)
		defer func() {
			if err := resp.Body.Close(); err != nil {
				t.Errorf("close response body: %v", err)
			}
		}()

		require.Equal(t, http.StatusCreated, resp.StatusCode)

		respBody, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var got api.CreateTripResponse
		err = json.Unmarshal(respBody, &got)
		require.NoError(t, err)

		require.NotEqual(t, uuid.Nil, got.ID)

		require.Equal(t, payload.DriverId, got.DriverId)
		require.Equal(t, payload.FromPoint, got.FromPoint)
		require.Equal(t, payload.ToPoint, got.ToPoint)
		require.Equal(t, payload.Seats, got.Seats)
		require.Equal(t, "draft", got.Status)
		require.WithinDuration(t, payload.DepartureTime, got.DepartureTime, time.Second)
		require.False(t, got.CreatedAt.IsZero())
		require.WithinDuration(t, time.Now(), got.CreatedAt, time.Minute)
	})
}
