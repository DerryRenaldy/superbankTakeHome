package store

const (
	QueryGetListAccount = `SELECT 
    c.name, 
    COALESCE(
        jsonb_agg(
            jsonb_build_object(
                'account_number', ba.account_number,
                'balance', ba.balance,
                'currency', ba.currency
            )
        ) FILTER (WHERE ba.account_number IS NOT NULL), '[]'::jsonb
    ) AS bank_accounts,
    
    COALESCE(
        jsonb_agg(
            jsonb_build_object(
                'name', p.name,
                'balance', p.balance,
                'currency', p.currency
            )
        ) FILTER (WHERE p.name IS NOT NULL), '[]'::jsonb
    ) AS pockets,
    
    COALESCE(
        jsonb_agg(
            jsonb_build_object(
                'amount', td.amount,
                'currency', td.currency,
                'interest_rate', td.interest_rate,
                'maturity_date', td.maturity_date
            )
        ) FILTER (WHERE td.amount IS NOT NULL), '[]'::jsonb
    ) AS term_deposits

FROM 
    account_dashboard.customers c
LEFT JOIN 
    account_dashboard.bank_accounts ba ON c.customer_id = ba.customer_id
LEFT JOIN 
    account_dashboard.pockets p ON c.customer_id = p.customer_id
LEFT JOIN 
    account_dashboard.term_deposits td ON c.customer_id = td.customer_id
WHERE TRUE %s
GROUP BY 
    c.customer_id, c.name
LIMIT %d
OFFSET %d;`

	QueryGetTotalCount = `SELECT COUNT(*) FROM account_dashboard.customers c WHERE TRUE %s;`
)
