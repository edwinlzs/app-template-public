"use client"

import AppSidebarTrigger from "./app-sidebar-trigger";
import ProfileIcon from "../profile-icon";
import {
    Sidebar,
    SidebarContent,
    SidebarFooter,
    SidebarGroup,
    SidebarGroupContent,
    SidebarGroupLabel,
    SidebarHeader,
    SidebarMenu,
    SidebarMenuButton,
    SidebarMenuItem,
    useSidebar,
} from "../ui/sidebar";
import LogoutButton from "../auth/logout-button";
import { useGetAuthUser } from "@/lib/queries/authUser";
import { navItems } from "@/lib/nav";

export default function AppSidebar() {
    const user = useGetAuthUser().data;
    const { isMobile } = useSidebar();
    if (!isMobile) return <></>

    return (
        <Sidebar side="right" className="outline-none">
            <SidebarHeader className="p-6">
                <div className="flex justify-between">
                    <ProfileIcon authUser={user} />
                    <AppSidebarTrigger close />
                </div>
            </SidebarHeader>
            <hr className="mx-4" />
            <SidebarContent>
                <SidebarGroup>
                    <SidebarGroupLabel className="text-md">
                        Links
                    </SidebarGroupLabel>
                    <SidebarGroupContent>
                        <SidebarMenu>
                            {navItems.map((item) => (
                                <SidebarMenuItem key={item.title}>
                                    <SidebarMenuButton asChild>
                                        <a href={item.url}>
                                            {item.icon && <item.icon />}
                                            <span>{item.title}</span>
                                        </a>
                                    </SidebarMenuButton>
                                </SidebarMenuItem>
                            ))}
                        </SidebarMenu>
                    </SidebarGroupContent>
                </SidebarGroup>
            </SidebarContent>
            <SidebarFooter className="mb-10 mx-6">
                {user && <LogoutButton />}
            </SidebarFooter>
        </Sidebar>
    );
}
