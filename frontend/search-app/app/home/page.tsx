"use client";

import { useEffect, useState } from "react";
import { Input } from "@/components/ui/inputSearchBar";
import { Card, CardContent } from "@/components/ui/cardList";
import { useAuth } from "@/context/authContext";
import { useRouter } from "next/navigation";

interface Customer {
    id: string;
    name: string;
    bankAccount: string;
    pocket: string;
    termDeposit: string;
}

const customers: Customer[] = [
    {
        id: "C001",
        name: "John Doe",
        bankAccount: "123-456-789",
        pocket: "$500",
        termDeposit: "$2000",
    },
    {
        id: "C002",
        name: "Jane Smith",
        bankAccount: "987-654-321",
        pocket: "$700",
        termDeposit: "$1500",
    },
];

export default function Home() {
    const [search, setSearch] = useState("");
    const filteredCustomers = customers.filter((customer) =>
        customer.name.toLowerCase().includes(search.toLowerCase())
    );

    const { user, logout } = useAuth();
    const router = useRouter();

    useEffect(() => {
        if (!user) {
            router.push("/login");
        }
    }, [user, router]);

    return user ? (
        <div className="p-4 max-w-lg mx-auto">
            <h1 className="text-xl font-bold mb-4">Customer Search</h1>
            <Input
                placeholder="Search customers..."
                value={search}
                onChange={(e) => setSearch(e.target.value)}
            />
            <div className="mt-4">
                {filteredCustomers.map((customer) => (
                    <Card key={customer.id} className="mb-2">
                        <CardContent>
                            <p><strong>Name:</strong> {customer.name}</p>
                            <p><strong>Bank Account:</strong> {customer.bankAccount}</p>
                            <p><strong>Pocket:</strong> {customer.pocket}</p>
                            <p><strong>Term Deposit:</strong> {customer.termDeposit}</p>
                        </CardContent>
                    </Card>
                ))}

                <button onClick={logout} className="mt-4 bg-red-500 text-white px-4 py-2">Logout</button>
            </div>
        </div>
    ) : null;
}
