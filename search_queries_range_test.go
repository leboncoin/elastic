// Copyright 2012-present Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

import (
	"encoding/json"
	"testing"
)

func TestRangeQuery(t *testing.T) {
	q := NewRangeQuery("postDate").From("2010-03-01").To("2010-04-01").Boost(3).Relation("within")
	q = q.QueryName("my_query")
	src, err := q.Source()
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(src)
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"range":{"_name":"my_query","postDate":{"boost":3,"gte":"2010-03-01","lte":"2010-04-01","relation":"within"}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func TestRangeQueryWithTimeZone(t *testing.T) {
	q := NewRangeQuery("born").
		Gte("2012-01-01").
		Lte("now").
		TimeZone("+1:00")
	src, err := q.Source()
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(src)
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"range":{"born":{"gte":"2012-01-01","lte":"now","time_zone":"+1:00"}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func TestRangeQueryWithFormat(t *testing.T) {
	q := NewRangeQuery("born").
		Gte("2012/01/01").
		Lte("now").
		Format("yyyy/MM/dd")
	src, err := q.Source()
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(src)
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"range":{"born":{"format":"yyyy/MM/dd","gte":"2012/01/01","lte":"now"}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func TestRangeQueryWithExclusiveBounds(t *testing.T) {
	q := NewRangeQuery("score").Gt(10).Lt(100)
	src, err := q.Source()
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(src)
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"range":{"score":{"gt":10,"lt":100}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func TestRangeQueryWithMixedBounds(t *testing.T) {
	q := NewRangeQuery("price").Gt(50.0).Lte(200.0)
	src, err := q.Source()
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(src)
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"range":{"price":{"gt":50,"lte":200}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func TestRangeQueryWithSingleBound(t *testing.T) {
	q := NewRangeQuery("age").Gte(18)
	src, err := q.Source()
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(src)
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"range":{"age":{"gte":18}}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
