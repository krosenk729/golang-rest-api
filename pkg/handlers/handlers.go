package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"go-rest-api/pkg/db"
	"go-rest-api/pkg/utils"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Entry type
type Entry struct {
	Date  string             `json:"date,omitempty"`
	Notes string             `json:"notes"`
	Tags  []string           `json:"tags,omitempty"`
	Verb  string             `json:"verb,omitempty"`
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
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
	log.Println("GetHandler", err)
	utils.CheckErr(err)

	var res []bson.M
	cursor.All(context.TODO(), &res)

	json.NewEncoder(w).Encode(res)
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

func formatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// GetByDateHandler Get Entries
func GetByDateHandler(w http.ResponseWriter, r *http.Request) {
	d := parseDateFromPath(mux.Vars(r), w)

	col := entryCol()
	cursor, err := col.Find(context.TODO(), bson.D{{Key: "date", Value: formatDate(d)}})
	utils.CheckErr(err)

	var res []bson.M
	cursor.All(context.TODO(), &res)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// CreateByDateHandler Create single entry
func CreateByDateHandler(w http.ResponseWriter, r *http.Request) {
	d := parseDateFromPath(mux.Vars(r), w)

	var e Entry
	e.Date = formatDate(d)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if utils.CheckErr(err) {
		APIResponse(http.StatusBadRequest, "invalid entry data", w)
	}

	collection := entryCol()
	res, err := collection.InsertOne(context.TODO(), e)
	utils.CheckErr(err)
	e.ID = res.InsertedID.(primitive.ObjectID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(e)
}

// CreateByDateBulkHandler create multiple entries
func CreateByDateBulkHandler(w http.ResponseWriter, r *http.Request) {
	d := parseDateFromPath(mux.Vars(r), w)

	decoder := json.NewDecoder(r.Body)
	// pull off opening array
	_, err := decoder.Token()
	if utils.CheckErr(err) {
		APIResponse(http.StatusBadRequest, "invalid array body", w)
	}
	var insertMany []interface{}
	for decoder.More() {
		var e Entry
		err := decoder.Decode(&e)
		utils.CheckErr(err)
		e.Date = formatDate(d)
		e.ID = primitive.NewObjectID()
		insertMany = append(insertMany, e)
	}

	collection := entryCol()
	_, err = collection.InsertMany(context.TODO(), insertMany)
	utils.CheckErr(err)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(insertMany)
}

// GetDbURL Create Entry
func GetDbURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(os.Getenv("MONGO_URI")))
}

// NotFound Create Entry
func NotFound(w http.ResponseWriter, r *http.Request) {
	APIResponse(http.StatusNotFound, "not found", w)
}

// APIResponse responds to api with exception code/message
func APIResponse(status int, message string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)
	w.Write([]byte(message))
}
