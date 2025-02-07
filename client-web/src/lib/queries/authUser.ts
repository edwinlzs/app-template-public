import { useQuery } from "@tanstack/react-query";
import { getBrowserClient } from "@/utils/supabase/client";

/** Retrieves auth user details from auth provider */
export const useGetAuthUser = () => {
    const supabase = getBrowserClient();

    return useQuery({
        queryKey: ["authUser"],
        queryFn: async () => (await supabase.auth.getUser()).data.user,
    });
}
