package utils

const (
	STATUS_SUCCESS = "success"
	STATUS_ERROR   = "error"
)

type ResponseCallback struct {
	Status     string
	StatusCode int
	Data       interface{}
	Message    string
}

func NewResponseCallback(params *ResponseCallback) *ResponseCallback {
	if params == nil {
		params = &ResponseCallback{} // Inisialisasi objek baru jika params nil
	}
	if params.Status == "" {
		params.Status = STATUS_SUCCESS
	}
	if params.StatusCode == 0 {
		params.StatusCode = 200
	}
	response := &ResponseCallback{
		Status:     params.Status,
		StatusCode: params.StatusCode,
		Data:       params.Data,
		Message:    params.Message,
	}
	return response
}
