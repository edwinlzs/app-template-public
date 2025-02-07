import { getBrowserUserSession } from "@/utils/supabase/client";

export interface GetUserResponse {
  id: string;
  alias: string;
  email: string;
}

export const getUser = async (): Promise<GetUserResponse | null> => {
  const accessToken = (await getBrowserUserSession())?.access_token;

  if (!accessToken) return null;

  return (
    await fetch("http://localhost:8080/auth/user", {
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
      method: "GET",
    })
  )
    .json()
    .catch(() => null);
};
