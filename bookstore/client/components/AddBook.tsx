"use client";

import { Dialog, DialogContent, DialogTrigger } from "@/components/ui/dialog";
import { Button } from "./ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "./ui/card";
import { Label } from "./ui/label";
import { Input } from "./ui/input";
import { SubmitHandler, useForm } from "react-hook-form";
import { postBook } from "@/actions/get-book";
import { DialogClose } from "@radix-ui/react-dialog";

export type InputProps = {
  name: string;
  author: string;
  publication: string;
};

export default function AddBook() {
  const { register, handleSubmit } = useForm<InputProps>();

  const onSubmit: SubmitHandler<InputProps> = async (data) => {
    console.log(data);
    const res = await postBook(data)
    if (res) {
      window.location.reload()
    }
  }

  return (
    <Dialog>
      <DialogTrigger className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
        Add Book
      </DialogTrigger>
      <DialogContent>
        <Card>
          <form onSubmit={handleSubmit(onSubmit)}>
            <CardHeader>
              <CardTitle>Create project</CardTitle>
              <CardDescription>
                Deploy your new project in one-click.
              </CardDescription>
            </CardHeader>
            <CardContent>
              <div className="grid w-full items-center gap-4">
                <div className="flex flex-col space-y-1.5">
                  <Label htmlFor="name">Name</Label>
                  <Input
                    id="name"
                    placeholder="Name of Book"
                    {...register("name")}
                  />
                </div>
                <div className="flex flex-col space-y-1.5">
                  <Label htmlFor="author">Author</Label>
                  <Input
                    id="author"
                    placeholder="Name of Author"
                    {...register("author")}
                  />
                </div>
                <div className="flex flex-col space-y-1.5">
                  <Label htmlFor="publication">Publication</Label>
                  <Input
                    id="publication"
                    placeholder="Name of Publication"
                    {...register("publication")}
                  />
                </div>
              </div>
            </CardContent>
            <CardFooter className="flex justify-between">
              <Button variant="outline">
                <DialogClose>Cancel</DialogClose>
              </Button>
              <Button type="submit">Deploy</Button>
            </CardFooter>
          </form>
        </Card>
      </DialogContent>
    </Dialog>
  );
}
