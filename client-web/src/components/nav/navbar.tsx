"use client"

import {
    NavigationMenu,
    NavigationMenuItem,
    NavigationMenuLink,
    NavigationMenuList,
} from "../ui/navigation-menu";
import AppSidebarTrigger from "./app-sidebar-trigger";
import { ModeToggle } from "../mode-toggle";
import { useSidebar } from "../ui/sidebar";
import { navItems } from "@/lib/nav";

export default function Navbar() {
    const { isMobile } = useSidebar();

    return (
        <NavigationMenu className="fixed top-0 left-0 w-screen max-w-screen py-5 px-7 sm:px-14 [&>*]:w-full">
            <NavigationMenuList className="justify-between">
                <NavigationMenuItem>
                    <NavigationMenuLink className="cursor-pointer" href="/">
                        Logo
                    </NavigationMenuLink>
                </NavigationMenuItem>
                <NavigationMenuItem>
                    <div className="flex gap-4 justify-center items-center">
                        {!isMobile && navItems.map((item) =>
                            <a href={item.url} key={item.title}>
                                <span>{item.title}</span>
                            </a>
                        )}
                        <ModeToggle />
                        {isMobile && <AppSidebarTrigger />}
                    </div>
                </NavigationMenuItem>
            </NavigationMenuList>
        </NavigationMenu>
    );
}
