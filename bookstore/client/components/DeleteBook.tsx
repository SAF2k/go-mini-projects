"use client";

import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Button } from "./ui/button";
import { DialogClose } from "@radix-ui/react-dialog";
import { deleteBook } from "@/actions/get-book";

export default function DeleteBook({ id }: { id: number }) {
  const onDelete = async () => {
    console.log(id);
    await deleteBook(id);
    window.location.reload();
  };

  return (
    <Dialog>
      <DialogTrigger className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded">
        Delete
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Are you sure you want to Delete?</DialogTitle>
        </DialogHeader>
        <DialogFooter className="flex justify-center">
          <DialogClose>
            <Button>Cancel</Button>
          </DialogClose>
          <Button
            className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded"
            onClick={onDelete}
          >
            Delete
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
