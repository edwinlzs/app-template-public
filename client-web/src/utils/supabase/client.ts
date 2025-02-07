import { createBrowserClient } from "@supabase/ssr";

export function getBrowserClient() {
  return createBrowserClient(
    process.env.NEXT_PUBLIC_SUPABASE_URL!,
    process.env.NEXT_PUBLIC_SUPABASE_ANON_KEY!,
  );
}

/** Get auth user session on the client side. */
export const getBrowserUserSession = async () => {
  const supabase = getBrowserClient();
  const { data: sessionData } = await supabase.auth.getSession();
  return sessionData.session;
};
