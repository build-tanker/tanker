package requester

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Requester - inteface to http.client
type Requester interface {
	Get(url string) ([]byte, error)
	Post(url string, body io.Reader) ([]byte, error)
	Put(url string) ([]byte, error)
	Delete(url string) ([]byte, error)
	Upload(url string, file string) ([]byte, error)
}

type requester struct {
	c *http.Client
}

// NewRequester - create a new requester which provides http.client
func NewRequester(timeout time.Duration) Requester {
	c := &http.Client{
		Timeout: timeout,
	}
	return &requester{
		c: c,
	}
}

func (r *requester) Get(url string) ([]byte, error) {
	return r.call(http.MethodGet, url, "", nil)
}

func (r *requester) Post(url string, body io.Reader) ([]byte, error) {
	// #TODO add body for post
	return r.call(http.MethodPost, url, "", body)
}

func (r *requester) Put(url string) ([]byte, error) {
	// #TODO add body for put
	return r.call(http.MethodPut, url, "", nil)
}

func (r *requester) Delete(url string) ([]byte, error) {
	return r.call(http.MethodDelete, url, "", nil)
}

func (r *requester) Upload(url string, file string) ([]byte, error) {
	return r.call(http.MethodPut, url, file, nil)
}

func (r *requester) call(method string, url string, filePath string, body io.Reader) ([]byte, error) {

	var request *http.Request
	var err error

	if filePath != "" {
		// Trying to upload a file
		file, err := os.Open(filePath)
		if err != nil {
			return []byte{}, err
		}
		defer file.Close()

		pr, pw := io.Pipe()
		bufin := bufio.NewReader(file)

		go func() {
			_, err := bufin.WriteTo(pw)
			if err != nil {
				fmt.Println(err)
			}
			pw.Close()
		}()

		request, err = http.NewRequest(method, url, pr)
	} else if body != nil {
		request, err = http.NewRequest(method, url, body)
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	} else {
		request, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return []byte{}, err
	}

	response, err := r.c.Do(request)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}
