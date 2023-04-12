package dto

type ExcelizeUploadItem struct {
	Name     string
	Category string
}

type ExcelizeUploadRequest struct {
	Item []ExcelizeUploadItem
}
