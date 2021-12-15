package errs

const (
	INTERNAL_SERVER_ERROR_TH = "เกิดข้อผิดพลาด"
	INTERNAL_SERVER_ERROR_EN = "Internal server error"

	SERVICE_UNAVAILABLE_ERROR_TH = "บริการไม่พร้อมใช้งาน"
	SERVICE_UNAVAILABLE_ERROR_EN = "Service unavailable"

	BAD_REQUEST_ERROR_TH = "คำร้องขอไม่ถูกต้อง"
	BAD_REQUEST_ERROR_EN = "Bad request"

	NOT_FOUND_ERROR_TH = "ไม่พบข้อมูล"
	NOT_FOUND_ERROR_EN = "Not found"
)

const (
	INSERT_ERROR_TH = "เกิดข้อผิดพลาดในการบันทึกข้อมูล"
	INSERT_ERROR_EN = "Saving data error"

	GET_ERROR_TH = "เกิดข้อผิดพลาดในการโหลดข้อมูล"
	GET_ERROR_EN = "Get data error"

	UPDATE_ERROR_TH = "เกิดข้อผิดพลาดในการอัพเดตข้อมูล"
	UPDATE_ERROR_EN = "Update data error"

	DELETE_ERROR_TH = "เกิดข้อผิดพลาดในการลบข้อมูล"
	DELETE_ERROR_EN = "Delete data error"
)

var (
	ErrInternalServerError = NewInternalServerError(
		INTERNAL_SERVER_ERROR_TH,
		INTERNAL_SERVER_ERROR_EN,
	)
	ErrServiceUnavailableError = NewServiceUnavailableError(
		SERVICE_UNAVAILABLE_ERROR_TH,
		SERVICE_UNAVAILABLE_ERROR_EN,
	)
	ErrBadRequestError = NewBadRequestError(
		BAD_REQUEST_ERROR_TH,
		BAD_REQUEST_ERROR_EN,
	)
	ErrNotFoundError = NewNotFoundError(
		NOT_FOUND_ERROR_TH,
		NOT_FOUND_ERROR_EN,
	)
)

var (
	ErrInsertError = NewInternalServerError(INSERT_ERROR_TH, INSERT_ERROR_EN)
	ErrGetError    = NewInternalServerError(GET_ERROR_TH, GET_ERROR_EN)
	ErrUpdateError = NewInternalServerError(UPDATE_ERROR_TH, UPDATE_ERROR_EN)
	ErrDeleteError = NewInternalServerError(DELETE_ERROR_TH, DELETE_ERROR_EN)
)
