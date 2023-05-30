package main

import (
	"bytes"
	"encoding/json"
	"github.com/MarkoKondic/ARS-2022-23Projekat/Alati/Dao"
	"io"
	"net/http"

	"github.com/google/uuid"
)

func StreamToByte(ctx context.Context, stream io.Reader) []byte {
	span := tracer.StartSpanFromContext(ctx, "StreamToByte")
	defer span.Finish()

	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}
func decodeBody(ctx context.Context, r io.Reader) (*Dao.Config, error) {
	span := tracer.StartSpanFromContext(ctx, "decodeBody")
	defer span.Finish()

	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var rt Dao.Config
	if err := json.Unmarshal(StreamToByte(r), &rt); err != nil {
		tracer.LogError(span, err)
		return nil, err
	}
	return &rt, nil
}

func decodeGroupBody(ctx context.Context, r io.Reader) (*Dao.ConfigGroup, error) {
	span := tracer.StartSpanFromContext(ctx, "decodeGroupBody")
	defer span.Finish()

	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var rt Dao.ConfigGroup
	if err := json.Unmarshal(StreamToByte(r), &rt); err != nil {
		tracer.LogError(span, err)
		return nil, err
	}
	return &rt, nil
}

func renderJSON(ctx context.Context, w http.ResponseWriter, v interface{}) {
	span := tracer.StartSpanFromContext(ctx, "renderJSON") //moguce da ide decodeBody opet pod navodnike
	defer span.Finish()

	js, err := json.Marshal(v)
	if err != nil {
		tracer.LogError(span, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func createId(ctx context.Context) string {
	span := tracer.StartSpanFromContext(ctx, "createId")
	defer span.Finish()

	return uuid.New().String()
}
