package api

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/nfnt/resize"
)

type PhotoBase interface {
	List() ([]Photo, error)
	Add(Photo) error
	Get(string) (Photo, error)
}

type Photo struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Data   []byte `json:"-"`
	Hero   string `json:"hero"`
}

func New(photobase PhotoBase) http.Handler {
	handler := httprouter.New()

	photoHandler := PhotoHandler{photobase}

	handler.GET("/photos", photoHandler.listPhotos)
	handler.GET("/photos/:id", photoHandler.getPhoto)
	handler.POST("/photos", photoHandler.addPhoto)

	return handler
}

type PhotoHandler struct {
	photobase PhotoBase
}

func (p *PhotoHandler) listPhotos(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	photoList, err := p.photobase.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err = json.NewEncoder(w).Encode(&photoList)
	if err != nil {
		fmt.Println(err)
	}
}

func (p *PhotoHandler) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	photo, err := p.photobase.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	buffer := bytes.NewBuffer(photo.Data)
	buffer2 := bytes.NewBuffer(photo.Data)
	compressed := bytes.NewBuffer([]byte{})

	mime, err := guessImageFormat(buffer2)
	if err != nil {
		log.Fatal(err)
	}

	var img image.Image

	switch mime {
	case "png":
		img, err = png.Decode(buffer)
	case "jpeg":
		img, err = jpeg.Decode(buffer)
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err != nil {
		log.Fatal(err)
	}

	resized := resize.Resize(uint(1024), uint(0), img, resize.Lanczos3)

	err = jpeg.Encode(compressed, resized, &jpeg.Options{Quality: 30})
	if err != nil {
		log.Fatal(err)
	}

	buffer = compressed

	w.Header().Set("Etag", fmt.Sprintf(`"%s"`, id))
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

func (p *PhotoHandler) addPhoto(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	file, _, err := r.FormFile("photo")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	raw, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	buffer := bytes.NewBuffer(raw)

	mime, err := guessImageFormat(buffer)
	if err != nil {
		log.Fatal(err)
	}

	if mime != "png" && mime != "jpeg" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	hero := md5.Sum([]byte(fmt.Sprintf("%s-%s", title, author)))

	photo := Photo{
		Title:  title,
		Author: author,
		Data:   raw,
		Hero:   fmt.Sprintf("%x", hero),
	}

	err = p.photobase.Add(photo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusAccepted)
}

// Guess image format from gif/jpeg/png/webp
func guessImageFormat(r io.Reader) (format string, err error) {
	_, format, err = image.DecodeConfig(r)
	format = strings.TrimSpace(format)
	return
}
