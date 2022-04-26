package multi_process

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	url string
	t   = 5
	//wg sync.WaitGroup
)

type information struct {
	r   *http.Response
	err error
}

func connect(c context.Context) error {
	defer wg.Done()
	info := make(chan information, 1)

	tr := &http.Transport{}
	httpClient := &http.Client{Transport: tr}
	request, _ := http.NewRequest("GET", url, nil)
	request = request.WithContext(c)

	go func() {
		request, err := httpClient.Do(request)
		if err != nil {
			fmt.Println(err)
			info <- information{nil, err}
		} else {
			info <- information{request, err}
		}
	}()

	select {
	case <-c.Done():
		fmt.Println("request is cancelled!!")
	case ok := <-info:
		err := ok.err
		r := ok.r

		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		defer r.Body.Close()
		realInfo, err := ioutil.ReadAll(r.Body)

		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		fmt.Printf("Response:%s\n", string(realInfo))
	}

	return nil
}

func TestContext() {
	url = "http://www.baidu.com"
	ctx := context.Background()
	ctx, cancelFunc := context.WithTimeout(ctx, time.Duration(t)*time.Second)
	defer cancelFunc()
	fmt.Printf("connecting to %s \n", url)

	wg.Add(1)
	go connect(ctx)
	wg.Wait()

	fmt.Println("End")

}
