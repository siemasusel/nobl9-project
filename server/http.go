package server

import (
	"net/http"
	"net/url"
	"stddevapi"
	"strconv"
)

func NewHTTPHandler(service *stddevapi.Service) http.Handler {
	srv := HTTPServer{
		service: service,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/random/mean", srv.handleRandomMean)

	return mux
}

type HTTPServer struct {
	service *stddevapi.Service
}

func getIntFromURLQuery(u *url.URL, key string) (int, error) {
	valueStr := u.Query().Get(key)
	if valueStr == "" {
		return 0, stddevapi.NewValidationError(nil, "query string '%s' missing", key)
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, stddevapi.NewValidationError(err, "invalid value '%s' for query string '%s', expecting integer", valueStr, key)
	}

	return value, nil
}

func (s *HTTPServer) handleRandomMean(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	requests, err := getIntFromURLQuery(r.URL, "requests")
	if err != nil {
		respondWithError(err, w, r)
		return
	}

	length, err := getIntFromURLQuery(r.URL, "length")
	if err != nil {
		respondWithError(err, w, r)
		return
	}

	res, err := s.service.CalculateStdDev(r.Context(), requests, length)
	if err != nil {
		respondWithError(err, w, r)
		return
	}

	respondJSON(w, r, stdDevGroupResultToResponse(res))
}

func stdDevGroupResultToResponse(result stddevapi.StdDevGroupResult) []RandomMean {
	response := make([]RandomMean, 0, len(result.Results)+1)
	for _, r := range result.Results {
		response = append(response, RandomMean(r))
	}

	response = append(response, RandomMean(result.ResultSum))

	return response
}

type RandomMean struct {
	StdDev float64 `json:"stddev"`
	Data   []int   `json:"data"`
}
