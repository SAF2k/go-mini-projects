
import { Book } from "@/components/DataTable";
import axios, { AxiosResponse } from "axios";

const getBooks = async () => {
  const res: AxiosResponse<Book[]> = await axios.get(
    "http://localhost:8080/books"
  );
  return res.data
}

export default getBooks