"use client";

import { Button } from "../ui/button";
import { Input } from "../ui/input";

export default function CreateUserForm({
  createUser,
}: {
  createUser: (formData: FormData) => void;
}) {
  return (
    <form action={createUser}>
      <Input name="alias" />
      <Button type="submit">Create User</Button>
    </form>
  );
}
