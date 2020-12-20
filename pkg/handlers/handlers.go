package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Entry type
type Entry struct {
	// ID	primitive.ObjectID	`bson:"_id,omitempty"`
	Date  string   `json:"date"`
	Notes string   `json:"notes"`
	Tags  []string `json:"tags,omitempty"`
	Verb  string   `json:"verb,omitempty"`
}

// GetHandler Get Entries
func GetHandler(w http.ResponseWriter, r *http.Request) {
	/* isValid := validators.IsDateRangeValid(req.QueryStringParameters["startDate"], req.QueryStringParameters["endDate"])
	if !isValid {
		return apiResponse(http.StatusBadRequest, ErrorBody{"Start date must be prior to end date"})
	}

	result, err := entries.FetchEntries(req)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusOK, result) */
	res := []Entry{
		Entry{"2012-12-12", "did some immportant things", []string{"project 1", "value 1"}, "Implemented"},
	}
	json.NewEncoder(w).Encode(res)
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte(`{"message": "get called"}`))
}

func getMonthVal(m int) time.Month {
	months := []time.Month{
		time.January,
		time.February,
		time.March,
		time.April,
		time.May,
		time.June,
		time.July,
		time.August,
		time.September,
		time.October,
		time.November,
		time.December,
	}
	return months[m-1]
}

// GetByDateHandler Get Entries
func GetByDateHandler(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)

	yyyy, err := strconv.Atoi(pathParams["yyyy"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "invalid date year"}`))
	}
	mm, err := strconv.Atoi(pathParams["mm"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "invalid month of the year"}`))
	}
	dd, err := strconv.Atoi(pathParams["dd"])
	if err != nil || dd > 31 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "invalid day year"}`))
	}
	timezone, _ := time.LoadLocation("America/New_York")
	d := time.Date(yyyy, getMonthVal(mm), dd, 12, 0, 0, 0, timezone)
	fmt.Println(d)
	res := []Entry{
		{d.String(), "did some immportant things", []string{"project 1", "value 1"}, "Implemented"},
	}
	json.NewEncoder(w).Encode(res)
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte(`{"message": "get called"}`))
}

// PostHandler Create Entry
func PostHandler(w http.ResponseWriter, r *http.Request) {
	res := []Entry{
		Entry{"2012-12-12", "did some immportant things", []string{"project 1", "value 1"}, "Implemented"},
	}
	json.NewEncoder(w).Encode(res)
}

// NotFound Create Entry
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}
