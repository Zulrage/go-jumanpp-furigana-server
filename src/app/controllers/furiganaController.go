package controllers

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/julienschmidt/httprouter"
    "app/models"
    "app/jumanpp"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type (
    // UserController represents the controller for operating on the User resource
    FuriganaController struct{
      session *mgo.Session
    }
)

func NewFuriganaController(s *mgo.Session) *FuriganaController {
    return &FuriganaController{s}
}

func (fc FuriganaController) GetFurigana(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  // Grab id
  id := p.ByName("id")
  fmt.Println("Get text" + id)

  // Stub blog
  f := models.Furigana{}

  // Fetch blog
  if err := fc.session.DB("furigana").C("text").Find(bson.M{"id": id}).One(&f); err != nil {
      w.WriteHeader(404)
      return
  }

  // Marshal provided interface into JSON structure
  fj, _ := json.Marshal(f)

  // Write content-type, statuscode, payload
  w.WriteHeader(200)
  fmt.Fprintf(w, "%s", fj)
}

func (fc FuriganaController) GetAllFurigana(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

  fmt.Println("Get all texts")

  // Stub blog
  f := []models.Furigana{}

  // Fetch blog
  if err := fc.session.DB("furigana").C("text").Find(nil).All(&f); err != nil {
      w.WriteHeader(404)
      return
  }

  // Marshal provided interface into JSON structure
  fj, _ := json.Marshal(f)

  // Write content-type, statuscode, payload
  w.WriteHeader(200)
  fmt.Fprintf(w, "%s", fj)
}

// CreateFurigana creates a new blog resource
func (fc FuriganaController) CreateFurigana(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  // Stub an blog to be populated from the body
  f := models.Furigana{}
  fmt.Println("Create text")
  // Populate the blog data
  json.NewDecoder(r.Body).Decode(&f)

  // Add an Id
  f.MongoId = bson.NewObjectId()
  f.Content = jumanpp.ToFuriganaText(f.Content)
  // Write the blog to mongo
  fc.session.DB("furigana").C("text").Insert(f)

  // Marshal provided interface into JSON structure
  fj, _ := json.Marshal(f)

  // Write content-type, statuscode, payload
  w.WriteHeader(201)
  fmt.Fprintf(w, "%s", fj)
}

// RemoveFurigana removes an existing user resource
func (fc FuriganaController) RemoveFurigana(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  // Grab id
  id := p.ByName("id")
  fmt.Println("Remove text" + id)
  // Verify id is ObjectId, otherwise bail
  if !bson.IsObjectIdHex(id) {
    w.WriteHeader(404)
    return
  }

  // Grab id
  oid := bson.ObjectIdHex(id)

  // Remove blog
  if err := fc.session.DB("furigana").C("text").RemoveId(oid); err != nil {
    w.WriteHeader(404)
    return
  }

  // Write status
  w.WriteHeader(200)
}
