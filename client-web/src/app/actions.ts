"use server";

import { revalidatePath } from "next/cache";
import { redirect } from "next/navigation";

import { getServerClient, getServerUserSession } from "@/utils/supabase/server";

export async function login(formData: FormData) {
  const supabase = await getServerClient();

  // type-casting here for convenience
  // in practice, you should validate your inputs
  const data = {
    email: formData.get("email") as string,
    password: formData.get("password") as string,
  };

  const { error } = await supabase.auth.signInWithPassword(data);

  if (error) {
    // TODO: Handle error
    console.error("Error logging in: ", error);
    redirect("/error");
  }

  revalidatePath("/", "layout");
  redirect("/");
}

export async function signup(formData: FormData) {
  const supabase = await getServerClient();

  // type-casting here for convenience
  // in practice, you should validate your inputs
  const data = {
    email: formData.get("email") as string,
    password: formData.get("password") as string,
  };

  const { error } = await supabase.auth.signUp(data);

  if (error) {
    // TODO: Handle error
    console.error("Error signing up: ", error);
    redirect("/error");
  }

  revalidatePath("/", "layout");
  redirect("/");
}

export async function loginWithOAuth() {
  const supabase = await getServerClient();
  const { data, error } = await supabase.auth.signInWithOAuth({
    provider: "google",
    options: {
      redirectTo: "http://localhost:3000/auth/callback",
    },
  });

  if (data.url) {
    redirect(data.url);
  }

  if (error) {
    // TODO: Handle error
    redirect("/error");
  }
}

export async function logout() {
  const supabase = await getServerClient();
  const { error } = await supabase.auth.signOut();

  if (error) {
    // TODO: Handle error
    redirect("/error");
  }
}

export const createUser = async (formData: FormData) => {
  const accessToken = (await getServerUserSession())?.access_token;
  const alias = formData.get("alias");
  if (!alias) {
    console.log("Failed to get alias: ", alias);
  }

  accessToken &&
    (await fetch("http://localhost:8080/auth/user", {
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
      method: "POST",
      body: JSON.stringify({
        alias,
      }),
    }));
};
