"use client";

import { AuthUser } from "@supabase/supabase-js";
import { Avatar, AvatarFallback, AvatarImage } from "./ui/avatar";

export interface NavProfile {
    authUser?: AuthUser | null;
}

export default function ProfileIcon({ authUser }: NavProfile) {
    return (
        <Avatar data-testid="profile-avatar">
            {authUser && (
                <AvatarImage src={authUser.user_metadata.avatar_url} />
            )}
            <AvatarFallback>?</AvatarFallback>
        </Avatar>
    );
}
