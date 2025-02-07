import { render, screen } from "@testing-library/react";
import Home from "./page";
import { useUserContext } from "./auth/provider";
import { Mock } from "vitest";

describe("Home", () => {
    describe("when user is not created", () => {
        it("renders the create user form", async () => {
            (useUserContext as Mock).mockReturnValueOnce({ user: null });

            render(await Home());
            expect(screen.getByRole('textbox').getAttribute('name')).toEqual('alias');
            expect(screen.getByRole('button', { name: 'Create User' })).toBeDefined();
        });
    });

    describe("when user is created", () => {
        it("welcomes the user", async () => {
            render(await Home());
            expect(screen.getByText("Welcome, testuser")).toBeDefined();
        });
    });
});
