package system

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/seakee/cpa-manager-plus/apps/manager-server/internal/http/middleware"
	"github.com/seakee/cpa-manager-plus/apps/manager-server/internal/http/response"
)

const (
	managerLatestReleaseURL       = "https://api.github.com/repos/seakee/CPA-Manager-Plus/releases/latest"
	managerLatestReleaseUserAgent = "CPA-Manager-Plus"
)

type managerReleaseInfo struct {
	TagName string `json:"tag_name"`
	Name    string `json:"name"`
}

// LatestManagerVersion proxies the GitHub latest-release lookup for CPA Manager Plus
// so the browser doesn't hit GitHub's anonymous rate limit directly.
func (h *Handler) LatestManagerVersion(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.MethodNotAllowed(w)
		return
	}
	if !middleware.AuthorizePanel(w, r, h.App.AdminAuthService) {
		return
	}

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, managerLatestReleaseURL, nil)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", managerLatestReleaseUserAgent)

	resp, err := client.Do(req)
	if err != nil {
		response.Error(w, http.StatusBadGateway, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))
		response.Error(w, http.StatusBadGateway, fmt.Errorf("unexpected_status: status %d: %s", resp.StatusCode, strings.TrimSpace(string(body))))
		return
	}

	var info managerReleaseInfo
	if err = json.NewDecoder(resp.Body).Decode(&info); err != nil {
		response.Error(w, http.StatusBadGateway, err)
		return
	}

	tagName := strings.TrimSpace(info.TagName)
	if tagName == "" {
		tagName = strings.TrimSpace(info.Name)
	}
	if tagName == "" {
		response.Error(w, http.StatusBadGateway, fmt.Errorf("invalid_response: missing release version"))
		return
	}

	response.JSON(w, http.StatusOK, map[string]any{"tag_name": tagName})
}
