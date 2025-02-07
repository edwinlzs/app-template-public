//nextjs.org/docs/app/building-your-application/rendering/composition-patterns#using-context-providers
"use client";

import { useGetUser } from "@/lib/queries/user";
import { createContext, useContext } from "react";

export interface UserContext {
  user: AppUser | null;
}

export interface AppUser {
  alias: string;
  email: string;
}

export const UserContext = createContext<UserContext>({ user: null });

export const useUserContext = () => useContext(UserContext);

export default function UserProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const { data } = useGetUser();
  const user = data || null;

  return (
    <UserContext.Provider value={{ user }}>{children}</UserContext.Provider>
  );
}
