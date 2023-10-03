import axios from "axios";


export default async function Home() {

  
  const getBooks = async () => {
    const res = await axios.get("http://localhost:8080/books");
    return res;
  }
  const res = await getBooks();

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      {res.data.map((book: any) => (
        <div
          key={book.ID}
          className="flex flex-col items-center justify-center"
        >
          <h1 className="text-2xl font-bold">{book.name}</h1>
          <p className="text-xl">{book.author}</p>
          <p className="text-xl">{book.publication}</p>
        </div>
      ))}
    </main>
  );
}
