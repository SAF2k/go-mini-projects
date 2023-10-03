import AddBook from "@/components/AddBook";
import DataTable from "@/components/DataTable";

export default function Home() {
  return (
    <main className="">
      <div className="mt-10 mx-16 flex justify-around">
        <h1 className="font-bold text-5xl">Books Store</h1>
        <AddBook />
      </div>
      <DataTable />
    </main>
  );
}
