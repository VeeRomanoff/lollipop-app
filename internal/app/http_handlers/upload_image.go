package httphandlers

import "net/http"

const bucketName = "lollipop-images-storage"

func (h *HTTPHandler) UploadImage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Failed to parse multipart form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Failed to get file", http.StatusBadRequest)
	}
	defer file.Close()

	fileName := header.Filename

	if err = h.minioStore.UploadImage(ctx, bucketName, fileName, file); err != nil {
		http.Error(w, "Failed to upload file to media storage", http.StatusInternalServerError)
	}
}
