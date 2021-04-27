package routes

import (
	"encoding/json"
	"go-alcochange-dtx-ga-ga/go-alcochange-dtx-ga/dbcon/mssqlcon"
	"go-alcochange-dtx-ga-ga/go-alcochange-dtx-ga/errs"
	"io"
	"io/ioutil"
	lg "log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/FenixAra/go-util/log"
	"github.com/julienschmidt/httprouter"
)

const (
	ERR_MSG = "ERROR_MESSAGE"
	MSG     = "MESSAGE"
)

type ResStruct struct {
	Status   string `json:"status" example:"SUCCESS" example:"FAILED"`
	HTTPCode int    `json:"httpCode" example:"200" example:"500"`
	Message  string `json:"message" example:"pong" example:"could not connect to db"`
}

type Res500Struct struct {
	Status   string `json:"status" example:"FAILED"`
	HTTPCode int    `json:"httpCode" example:"500"`
	Message  string `json:"message" example:"could not connect to db"`
}

type Res400Struct struct {
	Status   string `json:"status" example:"FAILED"`
	HTTPCode int    `json:"httpCode" example:"400"`
	Message  string `json:"message" example:"Invalid param"`
}

type RequestData struct {
	l           *log.Logger
	Start       time.Time
	w           http.ResponseWriter
	r           *http.Request
	dbConnMSSQL *mssqlcon.DBConn
}

type RenderData struct {
	Data  interface{}
	Paths []string
}

type TemplateData struct {
	Data interface{}
}

func (t *TemplateData) SetConstants() {

}

func logAndGetContext(w http.ResponseWriter, r *http.Request) *RequestData {
	w.Header().Add("X-Content-Type-Options", "nosniff")
	w.Header().Add("X-Frame-Options", "DENY")

	//url := strings.TrimSpace("http://0.0.0.0:9008/v1/logs")

	//Set config according to the use case
	cfg := log.NewConfig("alcochange-dtx")
	//cfg.SetRemoteConfig(url, "", "admin")
	cfg.SetLevelStr("Debug")
	cfg.SetFilePathSizeStr("Full")
	cfg.SetReference(r.Header.Get("ReferenceID"))

	l := log.New(cfg)
	dbConn := new(mssqlcon.DBConn)
	dbConn.Init(l)

	//pgdbConn := new(pgsqldb.Conn)
	//pgdbConn.Init(l)

	start := time.Now()
	l.LogAPIInfo(r, 0, 0)

	return &RequestData{
		l:           l,
		Start:       start,
		r:           r,
		w:           w,
		dbConnMSSQL: dbConn,
	}
}

func redirectTo(path string, rd *RequestData) {
	rd.l.Info("Status Code:", http.StatusFound, ", Response time:", time.Since(rd.Start), ", Response: url redirect - ", path)
	rd.l.LogAPIInfo(rd.r, time.Since(rd.Start).Seconds(), http.StatusFound)
	http.Redirect(rd.w, rd.r, path, http.StatusFound)
}

func jsonifyMessage(msg string, msgType string, httpCode int) ([]byte, int) {
	var data []byte
	var Obj struct {
		Status   string `json:"status"`
		HTTPCode int    `json:"code"`
		Message  string `json:"message"`
		Err      error  `json:"error"`
	}
	Obj.Message = msg
	Obj.HTTPCode = httpCode
	switch msgType {
	case ERR_MSG:
		Obj.Status = "FAILED"

	case MSG:
		Obj.Status = "SUCCESS"
	}
	data, _ = json.Marshal(Obj)
	return data, httpCode
}

func writeJSONMessage(msg string, msgType string, httpCode int, rd *RequestData) {
	d, code := jsonifyMessage(msg, msgType, httpCode)
	writeJSONResponse(d, code, rd)
}

func writeJSONStruct(v interface{}, code int, rd *RequestData) {
	d, err := json.Marshal(v)
	if err != nil {
		writeJSONMessage("Unable to marshal data. Err: "+err.Error(), ERR_MSG, http.StatusInternalServerError, rd)
		return
	}
	writeJSONResponse(d, code, rd)
}

func writeJSONResponse(d []byte, code int, rd *RequestData) {
	rd.l.LogAPIInfo(rd.r, time.Since(rd.Start).Seconds(), code)
	if code == http.StatusInternalServerError {
		rd.l.Info(rd.r.URL, "Status Code:", code, ", Response time:", time.Since(rd.Start), rd.r.URL, " Response:", string(d))
	} else {
		rd.l.Info(rd.r.URL, "Status Code:", code, ", Response time:", time.Since(rd.Start))
	}
	rd.w.Header().Set("Access-Control-Allow-Origin", "*")
	rd.w.Header().Set("Content-Type", "application/json; charset=utf-8")
	rd.w.WriteHeader(code)
	rd.w.Write(d)
}

func writeJSONMessageWithData(msg string, msgType string, httpCode int, rd *RequestData, functionName string, requestData string) {
	d, code := jsonifyMessage(msg, msgType, httpCode)
	writeJSONResponseWithData(d, code, rd, functionName, requestData)
}

func writeJSONResponseWithData(d []byte, code int, rd *RequestData, functionName string, requestData string) {
	rd.l.LogAPIInfo(rd.r, time.Since(rd.Start).Seconds(), code)
	//if code == http.StatusInternalServerError {
	//	rd.l.Info("Status Code:", code, ", Response time:", time.Since(rd.Start), " Response:", string(d))
	//} else if code == http.StatusBadRequest {
	//	rd.l.Info("Status Code:", code, ", Response time:", time.Since(rd.Start), " Response:", string(d))
	//} else {
	///	rd.l.Info("Service name : ", functionName, "Request Data : ", requestData,
	//		"Response data : ", string(d), "Status Code : ", code, "Response time : ", time.Since(rd.Start))
	//}

	rd.l.Info("Service name : ", functionName, ", Request Data : ", requestData,
		", Response data : ", string(d), ", Status Code : ", code, ", Response time : ", time.Since(rd.Start))

	rd.w.Header().Set("Access-Control-Allow-Origin", "*")
	rd.w.Header().Set("Content-Type", "application/json; charset=utf-8")
	rd.w.WriteHeader(code)
	rd.w.Write(d)
}

func writeJSONStructWithData(v interface{}, code int, rd *RequestData, functionName string, requestData string) {
	d, err := json.Marshal(v)
	if err != nil {
		writeJSONMessageWithData("Unable to marshal data. Err: "+err.Error(), ERR_MSG, http.StatusInternalServerError, rd, functionName, requestData)
		return
	}
	writeJSONResponseWithData(d, code, rd, functionName, requestData)
}

func renderJSON(w http.ResponseWriter, status int, res interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	if status == http.StatusNoContent {
		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		lg.Printf("ERROR: renderJson - %q\n", err)
	}
}

func getandConvertToInt64(query url.Values, str string) int64 {
	intData, _ := strconv.ParseInt(query.Get(str), 10, 64)
	return intData
}

func getandConvertToInt(query url.Values, str string) int {
	intData, _ := strconv.Atoi(query.Get(str))
	return intData
}

//parseJson parse data to model
func parseJSON(w http.ResponseWriter, body io.ReadCloser, model interface{}) bool {
	defer body.Close()

	b, _ := ioutil.ReadAll(body)
	err := json.Unmarshal(b, model)

	if err != nil {
		e := &errs.Error{}
		e.Message = "Error in parsing json"
		e.Err = err
		renderERROR(w, e)
		return false
	}

	return true
}

func renderERROR(w http.ResponseWriter, err *errs.Error) {
	err.Set()
	renderJSON(w, err.Code, err)
}

func parseJSONWithError(w http.ResponseWriter, body io.ReadCloser, model interface{}) (bool, error) {
	defer body.Close()
	b, _ := ioutil.ReadAll(body)
	err := json.Unmarshal(b, model)
	if err != nil {
		e := &errs.Error{}
		e.Message = "Error in parsing json"
		e.Err = err
		//renderERROR(w, e)
		return false, err
	}
	return true, nil
}

//GetIDFromParams function is use to get ID from request
func GetIDFromParams(w http.ResponseWriter, r *http.Request, key string) (int64, error) {
	params, _ := r.Context().Value("params").(httprouter.Params)
	idStr := params.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return id, err
	}

	return id, nil
}
