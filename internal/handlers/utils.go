package handlers

import (
	"encoding/json"
	"errors"
	"github.com/kumato/kumato/internal/logger"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"strings"
)

func parseJSON(r io.Reader, v interface{}, omitempty []string) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, v)
	if err != nil {
		return err
	}

	return validJSON(v, omitempty)
}

func handleOK(w http.ResponseWriter, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		handleErr(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func newResponseMessage(m responseMessage) []byte {
	b, err := json.Marshal(m)
	if err != nil {
		logger.Fatal(err)
		return nil
	}

	return b
}

func handleErr(w http.ResponseWriter, status int, e error) {
	logger.Fatal(e)
	w.WriteHeader(status)
	w.Write(newResponseMessage(responseMessage{"error", e.Error()}))
}

func validJSON(i interface{}, omitempty []string) error {
	v := reflect.ValueOf(i)

	// Get real value if passed i is a pointer.
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		tagName := v.Type().Field(i).Tag.Get("json")

		// trim suffix in json tag: ",omitempty"
		tagName = strings.Split(tagName, ",")[0]
		if tagName == "-" {
			continue
		}
		if isInStringSlice(tagName, omitempty) {
			continue
		}
		if v.Field(i).IsZero() {
			return errors.New(tagName + " is required")
		}
	}

	return nil
}

func isInStringSlice(x string, a []string) bool {
	sort.Strings(a)
	i := sort.SearchStrings(a, x)

	return 0 <= i && i < len(a) && a[i] == x
}
