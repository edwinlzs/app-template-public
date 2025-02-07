"use client";

import CreateUserForm from "@/components/auth/create-user-form";
import { createUser } from "./actions";
import { User as AuthUser } from "@supabase/supabase-js";
import ProfilePage from "./profile-page";
import { invalidateUser, useGetUser } from "../lib/queries/user";

export interface LoggedInPage {
  authUser: AuthUser;
}

export default function LoggedInPage() {
  const { data: user } = useGetUser();

  return (
    <div>
      {user ? (
        <ProfilePage user={user} />
      ) : (
        <CreateUserForm
          createUser={async (formData) => {
            await createUser(formData);
            await invalidateUser();
          }}
        />
      )}
    </div>
  );
}
