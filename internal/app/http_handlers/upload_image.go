package httphandlers

import (
	"net/http"
	"path/filepath"
	"strings"
)

const bucketName = "lollipop-images-storage"

func (h *HTTPHandler) UploadImage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Парсим форму для ограничения размера файла в 10 мб
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Failed to parse multipart form", http.StatusBadRequest)
		return
	}

	// Вытаскиваем файл из формы + хэдеры для проверки на валидность картинки
	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Failed to get file", http.StatusBadRequest)
		file.Close()
		return
	}
	defer file.Close()

	// Проверяем что нам пришла именно картинка используя хэдеры которые мы получили из FromFile
	contentType := header.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		http.Error(w, "Only images are supported", http.StatusBadRequest)
		return
	}

	// Проверяем на валидность расширений
	extension := strings.ToLower(filepath.Ext(header.Filename))
	switch extension {
	case ".jpg", ".jpeg", ".png", ".webp":
	default:
		http.Error(w, "Unsupported image format", http.StatusBadRequest)
		return
	}

	fileName := header.Filename

	// Загружаем файл в хранилище
	if err = h.minioStore.UploadImage(ctx, bucketName, fileName, file); err != nil {
		http.Error(w, "Failed to upload file to media storage", http.StatusInternalServerError)
	}
}
