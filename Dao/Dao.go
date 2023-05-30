package Dao

import (
	"context"
	"encoding/json"
	"fmt"
	tracer "github.com/MihajloJankovic/Alati/tracer"
	"github.com/hashicorp/consul/api"
	"os"
)

type Dao struct {
	cli *api.Client
}

func New() (*Dao, error) {
	db := os.Getenv("DB")
	dbport := os.Getenv("DBPORT")

	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%s", db, dbport)
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &Dao{
		cli: client,
	}, nil
}

func (ps *Dao) Get(ctx context.Context, id string, version string) (*Config, error) {
	span := tracer.StartSpanFromContext(ctx, "Get")
	defer span.Finish()

	kv := ps.cli.KV()

	data, _, err := kv.List(constructKey(id, version, ""), nil)
	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	post := &Config{}
	for _, pair := range data {
		err = json.Unmarshal(pair.Value, post)
		if err != nil {
			tracer.LogError(span, err)
			return nil, err
		}

	}
	return post, nil
}

func (ps *Dao) GetAll(ctx context.Context) ([]*Config, error) {
	span := tracer.StartSpanFromContext(ctx, "GetAll")
	defer span.Finish()

	kv := ps.cli.KV()
	data, _, err := kv.List(all, nil)
	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	posts := []*Config{}
	for _, pair := range data {
		post := &Config{}
		err = json.Unmarshal(pair.Value, post)
		if err != nil {
			tracer.LogError(span, err)
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (ps *Dao) Delete(ctx context.Context, id string, version string) (map[string]string, error) {
	span := tracer.StartSpanFromContext(ctx, "Delete")
	defer span.Finish()

	kv := ps.cli.KV()
	_, err := kv.DeleteTree(constructKey(id, version, ""), nil)
	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	return map[string]string{"Deleted": id}, nil
}

func (ps *Dao) Create(ctx context.Context, post *Config) (*Config, error) {
	span := tracer.StartSpanFromContext(ctx, "Create")
	defer span.Finish()

	kv := ps.cli.KV()

	sid, rid := generateKey(post.Version, post.Labels)
	post.Id = rid

	data, err := json.Marshal(post)
	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	p := &api.KVPair{Key: sid, Value: data}
	_, err = kv.Put(p, nil)
	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	return post, nil
}

func (ps *Dao) GetPostsByLabels(ctx context.Context, id string, version string, labels string) ([]*Config, error) {
	span := tracer.StartSpanFromContext(ctx, "GetPostsByLabels")
	defer span.Finish()

	kv := ps.cli.KV()

	data, _, err := kv.List(constructKey(id, version, labels), nil)
	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	posts := []*Config{}

	for _, pair := range data {
		post := &Config{}
		err = json.Unmarshal(pair.Value, post)
		if err != nil {
			tracer.LogError(span, err)
			return nil, err
		}
		posts = append(posts, post)
	}

	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	return posts, nil
}
