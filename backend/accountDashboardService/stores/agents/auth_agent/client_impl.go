package authagent

import (
	respdto "accountDashboardService/dto/response"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (a *authClientImpl) VerifyToken(ctx context.Context, accessToken string) (*respdto.VerifyTokenResponse, error) {
	uri := fmt.Sprintf("%s/%s?access_token=%s", a.cfg.Auth.Address, "v1/auth/verify-token", accessToken)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// req.Header.Set("Origin", "http://localhost:8090")

	// Generate and print the cURL command
	curlCmd := fmt.Sprintf("curl -X GET \"%s\"", uri)
	fmt.Println("Generated cURL command:", curlCmd)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	a.l.Infof("resp : ", resp)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var response respdto.VerifyTokenResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return &response, nil
}