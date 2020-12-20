package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"go-rest-api/pkg/db"
	"go-rest-api/pkg/utils"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Entry type
type Entry struct {
	Date  string   `json:"date,omitempty"`
	Notes string   `json:"notes"`
	Tags  []string `json:"tags,omitempty"`
	Verb  string   `json:"verb,omitempty"`
	// ID    primitive.ObjectID `bson:"_id,omitempty"`
}

func entryCol() *mongo.Collection {
	return db.Client.Database("go").Collection("entries")
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
	col := entryCol()
	cursor, err := col.Find(context.TODO(), bson.M{})
	utils.CheckErr(err)

	var res []bson.M
	cursor.All(context.TODO(), &res)

	json.NewEncoder(w).Encode(res)
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte(`{"message": "get called"}`))
}

func parseDateFromPath(pathParams map[string]string, w http.ResponseWriter) time.Time {
	yyyy, err := strconv.Atoi(pathParams["yyyy"])
	if utils.CheckErr(err) {
		APIResponse(http.StatusBadRequest, "invalid year", w)
	}
	mm, err := strconv.Atoi(pathParams["mm"])
	if utils.CheckErr(err) || mm > 12 {
		APIResponse(http.StatusBadRequest, "invalid month of the year", w)
	}
	dd, err := strconv.Atoi(pathParams["dd"])
	if utils.CheckErr(err) || dd > 31 {
		APIResponse(http.StatusBadRequest, "invalid day of the month", w)
	}
	timezone, _ := time.LoadLocation("America/New_York")
	return time.Date(yyyy, utils.GetMonthVal(mm), dd, 12, 0, 0, 0, timezone)
}

// GetByDateHandler Get Entries
func GetByDateHandler(w http.ResponseWriter, r *http.Request) {
	d := parseDateFromPath(mux.Vars(r), w)

	fmt.Println(d)
	res := []Entry{
		{d.String(), "did some immportant things", []string{"project 1", "value 1"}, "Implemented"},
	}
	json.NewEncoder(w).Encode(res)
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte(`{"message": "get called"}`))
}

// CreateByDateHandler Get Entries
func CreateByDateHandler(w http.ResponseWriter, r *http.Request) {
	d := parseDateFromPath(mux.Vars(r), w)

	var e Entry
	e.Date = d.String()
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if utils.CheckErr(err) {
		APIResponse(http.StatusBadRequest, "invalid entry data", w)
	}

	collection := entryCol()
	res, err := collection.InsertOne(context.TODO(), e)
	utils.CheckErr(err)
	json.NewEncoder(w).Encode(res)
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
	APIResponse(http.StatusNotFound, "not found", w)
}

// APIResponse
func APIResponse(status int, message string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)
	w.Write([]byte(message))
}
