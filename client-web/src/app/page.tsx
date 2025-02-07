import { getServerClient } from "@/utils/supabase/server";
import LoggedInPage from "./logged-in-page";
import { redirect } from "next/navigation";

export default async function Home() {
    const supabase = await getServerClient();
    const { data, error } = await supabase.auth.getUser();

    if (error || !data?.user) redirect("/login");

    return <LoggedInPage />;
}
