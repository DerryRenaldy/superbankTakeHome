"use client"
"use client"
import React, { useEffect, useState } from "react";
import CustomersList from "@/components/home_component/customerList";
import { Customer } from "@/types/customers";
import axios from "axios";

interface CustomerCardProps {
    customer: Customer;
}

interface CustomersListProps {
    customers: Customer[];
}

const customersData = [
    {
        name: "John Doe",
        bank_accounts: [
            { account_number: "1234567890", balance: 5000000, currency: "IDR" },
            { account_number: "9876543210", balance: 7500000, currency: "IDR" },
        ],
        pockets: [
            { name: "Vacation Savings", balance: 2000000, currency: "IDR" },
            { name: "Vacation Savings", balance: 2000000, currency: "IDR" },
        ],
        term_deposits: [
            { amount: 10000000, currency: "IDR", interest_rate: 5.5, maturity_date: "2026-03-07" },
            { amount: 10000000, currency: "IDR", interest_rate: 5.5, maturity_date: "2026-03-07" },
        ],
    },
    {
        name: "Jane Smith",
        bank_accounts: [
            { account_number: "1122334455", balance: 12000000, currency: "USD" },
        ],
        pockets: [
            { name: "Emergency Fund", balance: 5000000, currency: "IDR" },
            { name: "Shopping Wallet", balance: 1500000, currency: "USD" },
        ],
        term_deposits: [
            { amount: 25000000, currency: "USD", interest_rate: 4.8, maturity_date: "2027-06-15" },
        ],
    },
];

const CustomersPage: React.FC = () => {
    const [customers, setCustomers] = useState<Customer[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchCustomers = async () => {
            try {
                const response = await axios.get<Customer[]>("http://localhost:8090/v1/dashboard/accounts", {
                    headers: {
                        Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2LCJlbWFpbCI6InVzZXJAZXhhbXBsZS5jb20iLCJyb2xlIjoidXNlciIsInN1YiI6InVzZXJAZXhhbXBsZS5jb20iLCJleHAiOjE3NDEzODc5NjUsImlhdCI6MTc0MTM4NzA2NSwianRpIjoiMTM2NTcyOTItMjFjYy00NzgyLWI5NTMtZjk1NTQzY2ZhMTNhIn0.aJw6Zh-CLfaInVZFe8ZITJ4jMAs7mhb9ji0YKtIjCf0`,
                        "Content-Type": "application/json",
                    },
                });
                setCustomers(response.data);
            } catch (err) {
                setError("Failed to fetch customers. Please try again.");
            } finally {
                setLoading(false);
            }
        };

        fetchCustomers();
    }, []);

    return (
        <div className="container mx-auto p-4">
            <h1 className="text-2xl font-bold text-center mb-6">Customers</h1>
            {loading && <p className="text-center">Loading...</p>}
            {error && <p className="text-center text-red-500">{error}</p>}
            {!loading && !error && <CustomersList customers={customers} />}
        </div>
    );
};

export default CustomersPage;
