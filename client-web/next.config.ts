import type { NextConfig } from "next";

const nextConfig: NextConfig = {
    compiler: {
        reactRemoveProperties: { properties: ["^data-testid$"] },
    },
};

export default nextConfig;
