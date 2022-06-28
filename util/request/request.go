package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fehu/conf"
	"fehu/model/http_error"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Bind(c *gin.Context, i interface{}) {

	err := c.ShouldBind(i)
	if err != nil && err != io.EOF {

		terr := http_error.MissingParametersError
		if e, ok := err.(*json.UnmarshalTypeError); ok {
			terr.ErrorMsg += ", 参数类型错误: " + e.Field
		}

		fmt.Println(err.Error())
		panic(terr)
	}

}

func Curl(url string, data interface{}) (ret []byte, err error) {
	j, err := json.Marshal(data)
	if err != nil {
		return ret, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(j))
	if err != nil {
		return ret, err
	}
	defer resp.Body.Close()
	ret, err = ioutil.ReadAll(resp.Body)
	return ret, err
}

type CurlFileData struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

func CurlFile(url, folder string, header *multipart.FileHeader, reader io.Reader) (data *CurlFileData, err error) {

	var file_writer io.Writer

	body_buf := &bytes.Buffer{}
	send_writer := multipart.NewWriter(body_buf)

	if file_writer, err = send_writer.CreateFormFile("file", header.Filename); err != nil {
		return nil, err
	}

	if _, err = io.Copy(file_writer, reader); err != nil {
		return nil, err
	}

	err = send_writer.WriteField("type", folder)
	if err != nil {
		return nil, err
	}

	from_type := send_writer.FormDataContentType()
	send_writer.Close()

	req, err := http.NewRequest("POST", url, body_buf)

	req.Header.Set("Content-Type", from_type)
	req.Header.Set("AUTH", conf.Http.UploadAuth)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	s, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	type Res struct {
		Data    *CurlFileData `json:"data"`
		Code    int           `json:"code"`
		Url     string        `json:"url"`
		Message string        `json:"message"`
	}

	res := &Res{}

	err = json.Unmarshal(s, res)
	if err != nil {
		return nil, err
	}

	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}

	return res.Data, err
}
