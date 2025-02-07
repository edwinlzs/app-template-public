import { vi } from "vitest";

const userEmail = "user@email.com";

vi.mock("@/utils/supabase/server", () => ({
    getServerClient: vi.fn(() => ({
        auth: {
            getUser: () => ({
                data: {
                    user: {
                        email: userEmail,
                        user_metadata: {
                            avatar_url: "avatar.com",
                        },
                    },
                    user_metadata: {
                        avatar_url: "avatar.com",
                    },
                },
            }),
        },
    })),
}));

vi.mock("@/app/auth/provider", () => ({
    useUserContext: vi.fn(() => ({
        user: {
            alias: "testuser",
            email: userEmail,
        },
    })),
}));

// https://jestjs.io/docs/manual-mocks#mocking-methods-which-are-not-implemented-in-jsdom
Object.defineProperty(window, "matchMedia", {
    writable: true,
    value: vi.fn().mockImplementation((query) => ({
        matches: false,
        media: query,
        onchange: null,
        addListener: vi.fn(), // deprecated
        removeListener: vi.fn(), // deprecated
        addEventListener: vi.fn(),
        removeEventListener: vi.fn(),
        dispatchEvent: vi.fn(),
    })),
});
