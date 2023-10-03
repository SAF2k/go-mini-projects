"use client";

import { useEffect, useState } from "react";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "./ui/table";
import getBooks from "@/actions/get-book";

export interface Book {
  ID: number;
  name: string;
  author: string;
  publication: string;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  // Add other book properties here
}

export default function DataTable() {
  const [data, setData] = useState<Book[]>([]); // Use state with the Book type
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function fetchData() {
      try {
        const booksData: Book[] = await getBooks();
        setData(booksData);
        setLoading(false); // Set loading to false when data is ready
      } catch (error) {
        console.error("Error fetching books:", error);
        setLoading(false); // Set loading to false on error
      }
    }

    fetchData();
  }, []);

  return (
    <div className="mt-10 mx-14 border rounded-xl">
      <Table>
        <TableCaption>A list of your recent invoices.</TableCaption>
        <TableHeader>
          <TableRow>
            <TableHead className="w-[100px]">Invoice</TableHead>
            <TableHead>Name</TableHead>
            <TableHead>Author</TableHead>
            <TableHead>Publication</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
            {loading ? (
              <TableRow>
                <TableCell colSpan={4}>Loading...</TableCell>
              </TableRow>
            ) : (
              data.map((book, index) => (
                <TableRow key={index}>
                  <TableCell>#{book.ID}</TableCell>
                  <TableCell>{book.name}</TableCell>
                  <TableCell>{book.author}</TableCell>
                  <TableCell>{book.publication}</TableCell>
                </TableRow>
              ))
            )}
        </TableBody>
      </Table>
    </div>
  );
}
