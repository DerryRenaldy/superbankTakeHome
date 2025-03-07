"use client"
import React from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Customer } from "@/types/customers";

interface CustomerCardProps {
    customer: Customer;
}

interface CustomersListProps {
    customers: Customer[];
}

const CustomerCard: React.FC<CustomerCardProps> = ({ customer }) => {
    return (
        <Card className="p-4 m-4 w-full max-w-xl">
            <CardHeader>
                <CardTitle className="text-lg font-bold">{customer.name}</CardTitle>
            </CardHeader>
            <CardContent>
                <div className="space-y-4">
                    {/* Bank Accounts */}
                    <div>
                        <h3 className="font-semibold">Bank Accounts</h3>
                        <ul className="list-disc pl-4">
                            {customer.bank_accounts.map((account, index) => (
                                <li key={index} className="text-sm">
                                    {account.account_number} - {account.balance.toLocaleString()} {account.currency}
                                </li>
                            ))}
                        </ul>
                    </div>

                    {/* Pockets */}
                    <div>
                        <h3 className="font-semibold">Pockets</h3>
                        <ul className="list-disc pl-4">
                            {customer.pockets.map((pocket, index) => (
                                <li key={index} className="text-sm">
                                    {pocket.name}: {pocket.balance.toLocaleString()} {pocket.currency}
                                </li>
                            ))}
                        </ul>
                    </div>

                    {/* Term Deposits */}
                    <div>
                        <h3 className="font-semibold">Term Deposits</h3>
                        <ul className="list-disc pl-4">
                            {customer.term_deposits.map((deposit, index) => (
                                <li key={index} className="text-sm">
                                    {deposit.amount.toLocaleString()} {deposit.currency} - {deposit.interest_rate}% (Maturity: {deposit.maturity_date})
                                </li>
                            ))}
                        </ul>
                    </div>
                </div>
            </CardContent>
        </Card>
    );
};

const CustomersList: React.FC<CustomersListProps> = ({ customers }) => {
    return (
        <div className="flex flex-wrap justify-center gap-4">
            {customers.map((customer, index) => (
                <CustomerCard key={index} customer={customer} />
            ))}
        </div>
    );
};

export default CustomersList;
