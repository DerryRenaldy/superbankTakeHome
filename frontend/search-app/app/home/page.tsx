"use client"

import React, { useEffect, useState } from "react";
import CustomersList from "@/components/home_component/customerList";
import { Customer } from "@/types/customers";
import { dashboardApi } from "@/lib/api/axiosDashboard";
// import { useSearchParams } from "next/navigation";
import { useAuth } from "@/context/auth/authContext";
import { ApiResponseDashboard } from "@/types/dashboard";
import {
    Pagination,
    PaginationContent,
    PaginationEllipsis,
    PaginationItem,
    PaginationLink,
    PaginationNext,
    PaginationPrevious,
} from "@/components/ui/pagination"
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select"
import { ScrollArea } from "@/components/ui/scroll-area"
import { Card } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useDebounce } from "use-debounce";

const CustomersPage: React.FC = () => {
    const [customers, setCustomers] = useState<Customer[]>([]);
    const [pageSize, setPageSize] = useState<number>(5);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);
    const [currentPage, setCurrentPage] = useState<number>(1); // New state for current page
    const [totalCount, setTotalCount] = useState<number>(0); // New state for total count
    // const searchParams = useSearchParams();
    const { user, logout } = useAuth();
    const [search, setSearch] = useState<string>("");
    const [debouncedSearch] = useDebounce(search, 500);
    const [value, setValue] = useState("");

    const handlePaste = (e: React.ClipboardEvent<HTMLInputElement>) => {
        e.preventDefault();
        const pasteText = e.clipboardData.getData("text").replace(/[^a-zA-Z0-9 ]/g, ""); // Allow spaces
        setValue((prev) => prev + pasteText);
    };

    useEffect(() => {
        const fetchCustomers = async () => {
            try {
                const response = await dashboardApi.get(`/accounts?customerName=${search}&page=${currentPage}&count=${pageSize}`, { // Updated to use currentPage
                    headers: {
                        "Authorization": `Bearer ${localStorage.getItem("auth_token")}`,
                        "Content-Type": "application/json",
                    },
                });

                const customers: ApiResponseDashboard = response.data || [];
                setTotalCount(customers.pagination.totalCount || 0);
                setCustomers(customers.data);
            } catch (err) {
                setError("Failed to fetch customers. Please try again.");
            } finally {
                setLoading(false);
            }
        };

        fetchCustomers();
    }, [currentPage, pageSize, search]);
    //searchParams, 

    const handlePageChange = (page: number) => {
        const maxPage = Math.ceil(totalCount / pageSize);
        if (page >= 1 && page <= maxPage) {
            setCurrentPage(page);
        }
    };

    const handleSearchChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const newValue = e.target.value.replace(/[^a-zA-Z0-9 ]/g, ""); // Allow letters, numbers, and spaces
        setValue(newValue);

        setSearch(e.target.value);
        setCurrentPage(1);
    };

    return (
        <div className="container mx-auto min-h-screen p-4">
            {loading ? (
                <div className="flex justify-center items-center h-40">
                    <p className="text-lg font-semibold text-gray-500">Loading...</p>
                </div>
            ) : (
                <>
                    <div className="flex justify-end items-center mt-4">
                        <h1 className="text-2xl font-bold text-center select-none">{user || "No user found"}</h1>
                    </div>
                    <div className="flex justify-between items-center mt-4 my-1.5">
                        <Input className="w-[200px] h-[40px] rounded-md" placeholder="Search customer" onChange={handleSearchChange}
                            type="text"
                            value={value}
                            onPaste={handlePaste} />
                        <div className="flex justify-center items-center h-full">
                            <Button className="bg-red-500 text-white hover:bg-red-600 cursor-pointer" onClick={() => {
                                logout();
                            }}>Logout</Button>
                        </div>
                    </div>
                    <Card className="shadow-lg">
                        {
                            !customers || customers.length === 0 ? (
                                <p className="text-center text-gray-500">No customers found.</p>
                            ) :
                                <ScrollArea className="h-[350px]">
                                    <CustomersList customers={customers} />
                                </ScrollArea>
                        }
                    </Card>
                    <div className="flex justify-between items-center mt-4">
                        <Select value={String(pageSize)} onValueChange={(value) => { setPageSize(Number(value)); handlePageChange(1); }}>
                            <SelectTrigger className="w-[80px] select-none">
                                <SelectValue placeholder="5" />
                            </SelectTrigger>
                            <SelectContent>
                                <SelectItem value="5">5</SelectItem>
                                <SelectItem value="15">15</SelectItem>
                                <SelectItem value="25">25</SelectItem>
                            </SelectContent>
                        </Select>
                        <Pagination>
                            <PaginationContent>
                                <PaginationItem className="select-none cursor-pointer">
                                    <PaginationPrevious onClick={() => handlePageChange(currentPage - 1)} />
                                </PaginationItem>
                                {Array.from({ length: Math.ceil(totalCount / pageSize) }, (_, index) => (
                                    <PaginationItem className="select-none cursor-pointer" key={index}>
                                        <PaginationLink onClick={() => handlePageChange(index + 1)} className={`${currentPage === index + 1 ? "bg-gray-700 text-white hover:bg-gray-500 hover:text-white" : ""}`}>{index + 1}</PaginationLink>
                                    </PaginationItem>
                                ))}
                                {
                                    Math.ceil(totalCount / pageSize) > 4 &&
                                    <PaginationItem className="select-none cursor-pointer">
                                        <PaginationEllipsis />
                                    </PaginationItem>
                                }
                                <PaginationItem className="select-none cursor-pointer">
                                    <PaginationNext onClick={() => handlePageChange(currentPage + 1)} />
                                </PaginationItem>
                            </PaginationContent>
                        </Pagination>
                    </div>
                </>
            )}
        </div>
    );
};

export default CustomersPage;
