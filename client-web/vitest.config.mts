import { defineConfig } from "vitest/config";
import react from "@vitejs/plugin-react";
import tsconfigPaths from "vite-tsconfig-paths";

export default defineConfig({
    plugins: [tsconfigPaths(), react()],
    test: {
        globals: true,
        environment: "jsdom",
        poolOptions: {
            forks: {
                singleFork: true,
            }
        },
        setupFiles: ["vitest.setup.ts"],
    },
});
