package reqdto

type AccountListRequest struct {
	CustomerName string `json:"customerName"`
	Page         int    `json:"page"`
	Count        int    `json:"count"`
}
