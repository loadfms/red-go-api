package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"

	. "../dao"
	. "../models"
)

var widgetDao = NewWidgetsDAO(GetSession())

// GET list of widgets
var AllWidgetsEndPoint = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	widgets, err := widgetDao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, widgets)
})

// GET a widget by its ID
var FindWidgetsEndpoint = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	widget, err := widgetDao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Widget ID")
		return
	}
	respondWithJson(w, http.StatusOK, widget)
})

// POST a new widget
var CreateWidgetsEndPoint = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var widget Widget
	if err := json.NewDecoder(r.Body).Decode(&widget); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	widget.Id = bson.NewObjectId()
	if err := widgetDao.Insert(widget); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, widget)
})

// PUT update an existing widget
var UpdateWidgetsEndPoint = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var widget Widget
	if err := json.NewDecoder(r.Body).Decode(&widget); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := widgetDao.Update(params["id"], widget); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
})

// DELETE an existing widget
var DeleteWidgetsEndPoint = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var widget Widget
	if err := json.NewDecoder(r.Body).Decode(&widget); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := widgetDao.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
})
