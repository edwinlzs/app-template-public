"use client";

import { logout } from "@/app/actions";
import { Button } from "../ui/button";
import { useSidebar } from "../ui/sidebar";

export default function LogoutButton() {
    const { setOpenMobile, setOpen, isMobile } = useSidebar();
    return (
        <Button
            onClick={async () => {
                await logout();
                isMobile ? setOpenMobile(false) : setOpen(false);
            }}
        >
            Logout
        </Button>
    );
}
