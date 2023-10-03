
import { InputProps } from "@/components/AddBook";
import { Book } from "@/components/DataTable";
import axios, { AxiosResponse } from "axios";

const getBooks = async () => {
  const res: AxiosResponse<Book[]> = await axios.get(
    "http://localhost:8080/books"
  );
  return res.data
}

const postBook = async (data: InputProps) => {
  const res: AxiosResponse<Book> = await axios.post(
    "http://localhost:8080/books", data
  );
  return res.data
}

const deleteBook = async (id: number) => {
  const res: AxiosResponse<Book> = await axios.delete(
    `http://localhost:8080/books/${id}`
  );
  return res.data
}

export { postBook, getBooks, deleteBook }