package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/GauravJ3/go-service/external"
	"github.com/GauravJ3/go-service/models"
	"github.com/GauravJ3/go-service/utils"

	"github.com/gorilla/mux"
)

func ReportHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	url := fmt.Sprintf("http://localhost:5007/api/v1/students/%s", id)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-CSRF-Token", external.CsrfToken)
	req.AddCookie(&http.Cookie{Name: "accessToken", Value: external.AuthToken})
	req.AddCookie(&http.Cookie{Name: "refreshToken", Value: external.RefreshToken})
	req.AddCookie(&http.Cookie{Name: "csrfToken", Value: external.CsrfToken})

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to fetch student data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		http.Error(w, string(body), resp.StatusCode)
		return
	}

	var student models.Student
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &student)

	pdfBytes, err := utils.GeneratePDF(student)
	if err != nil {
		http.Error(w, "PDF generation failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=report.pdf")
	w.Write(pdfBytes)
}
