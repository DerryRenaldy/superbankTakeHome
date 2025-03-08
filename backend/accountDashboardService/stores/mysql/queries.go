package store

const (
	QueryGetListAccount = `SELECT 
    c.name, 
    COALESCE(
        JSON_ARRAYAGG(
            JSON_OBJECT(
                'account_number', ba.account_number,
                'balance', ba.balance,
                'currency', ba.currency
            )
        ), '[]'
    ) AS bank_accounts,
    
    COALESCE(
        JSON_ARRAYAGG(
            JSON_OBJECT(
                'name', p.name,
                'balance', p.balance,
                'currency', p.currency
            )
        ), '[]'
    ) AS pockets,
    
    COALESCE(
        JSON_ARRAYAGG(
            JSON_OBJECT(
                'amount', td.amount,
                'currency', td.currency,
                'interest_rate', td.interest_rate,
                'maturity_date', td.maturity_date
            )
        ), '[]'
    ) AS term_deposits

FROM 
    account_dashboard.customers c
LEFT JOIN 
    account_dashboard.bank_accounts ba ON c.customer_id = ba.customer_id
LEFT JOIN 
    account_dashboard.pockets p ON c.customer_id = p.customer_id
LEFT JOIN 
    account_dashboard.term_deposits td ON c.customer_id = td.customer_id
WHERE 1 %s
GROUP BY 
    c.customer_id,  c.name
LIMIT %d
OFFSET %d;`

	QueryGetTotalCount = `SELECT COUNT(*) FROM account_dashboard.customers c WHERE 1 %s;`
)

