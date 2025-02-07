interface NavItem {
    title: string;
    url: string;
    icon?: React.ComponentType;
}

export const navItems: NavItem[] = [
    {
        title: "Home",
        url: "/",
    },
];