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
	"job4j.ru/go-share-trip/internal/domain"
)

func TestServer_PublishTrip(t *testing.T) {
	t.Parallel()
	t.Run("success - перевод поездки в статус Опубликована", func(t *testing.T) {
		t.Parallel()
		tripId := uuid.New()
		driverId := uuid.New()
		trip := domain.Trip{
			Id:            tripId,
			DriverId:      driverId,
			FromPoint:     "Казань",
			ToPoint:       "Москва",
			DepartureTime: time.Now().Add(1 * time.Hour).UTC().Truncate(time.Second),
			Seats:         3,
			Status:        domain.Draft,
			CreatedAt:     time.Now().UTC().Truncate(time.Microsecond),
		}
		_, err := testPool.Exec(testCtx, `insert into trips (id, driver_id, from_point, to_point, departure_time, seats, status, created_at)
		values ($1,$2,$3,$4,$5,$6,$7,$8)`,
			tripId, driverId, trip.FromPoint, trip.ToPoint, trip.DepartureTime, trip.Seats, trip.Status, trip.CreatedAt)
		require.NoError(t, err)

		payload := api.PublishTripRequest{
			TripId:   tripId,
			DriverId: driverId,
		}
		body, err := json.Marshal(payload)
		require.NoError(t, err)

		req, err := http.NewRequest(
			http.MethodPost,
			"/trip/publish",
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

		require.Equal(t, http.StatusOK, resp.StatusCode)
		respBody, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var got api.PublishTripResponse
		err = json.Unmarshal(respBody, &got)
		require.NoError(t, err)

		expected := api.PublishTripResponse{
			ID:            tripId,
			DriverId:      driverId,
			FromPoint:     "Казань",
			ToPoint:       "Москва",
			DepartureTime: trip.DepartureTime,
			Seats:         3,
			Status:        "published",
			CreatedAt:     trip.CreatedAt,
		}
		require.Equal(t, expected.ID, got.ID)
		require.Equal(t, expected.DriverId, got.DriverId)
		require.Equal(t, expected.FromPoint, got.FromPoint)
		require.Equal(t, expected.ToPoint, got.ToPoint)
		require.WithinDuration(t, expected.DepartureTime, got.DepartureTime, time.Second)
		require.Equal(t, expected.Seats, got.Seats)
		require.Equal(t, expected.Status, got.Status)
		require.WithinDuration(t, expected.CreatedAt, got.CreatedAt, time.Second)
	})

}
