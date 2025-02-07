"use client";

import { AlignJustify, X } from "lucide-react";
import { useSidebar } from "../ui/sidebar";

export default function AppSidebarTrigger({ close }: { close?: boolean }) {
    const { setOpenMobile } = useSidebar();

    return (
        <div className="flex justify-center items-center cursor-pointer">
            {close ? (
                <X data-testid="sidebar-close-button" onClick={() => setOpenMobile(false)} />
            ) : (
                <AlignJustify onClick={() => setOpenMobile(true)} />
            )}
        </div>
    );
}
