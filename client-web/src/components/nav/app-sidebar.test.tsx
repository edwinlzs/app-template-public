import { render, screen, within } from "@testing-library/react";
import AppSidebar from "./app-sidebar";
import { SidebarProvider } from "../ui/sidebar";

describe("AppSidebar", () => {
    describe("when user is logged in", () => {
        it("should render the logout button", async () => {
            render(<SidebarProvider>{await AppSidebar()}</SidebarProvider>);

            expect(
                screen.getByRole("button", { name: "Logout" })
            ).toBeDefined();
        });
    });

    it("should display the user avatar and navigation links", async () => {
        render(<SidebarProvider>{await AppSidebar()}</SidebarProvider>);

        expect(screen.getByTestId('profile-avatar')).toBeDefined();
        expect(screen.getByTestId('sidebar-close-button')).toBeDefined();

        const menu = screen.getByRole('list');
        expect(within(menu).getByRole('link', { name: 'Home'})).toBeDefined();
    });
});
