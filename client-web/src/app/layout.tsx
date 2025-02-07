import type { Metadata } from "next";
import localFont from "next/font/local";
import "./globals.css";
import Providers from "./providers";
import PageContainer from "@/components/page-container";
import Navbar from "@/components/nav/navbar";
import AppSidebar from "@/components/nav/app-sidebar";

const geistSans = localFont({
    src: "./fonts/GeistVF.woff",
    variable: "--font-geist-sans",
    weight: "100 900",
});
const geistMono = localFont({
    src: "./fonts/GeistMonoVF.woff",
    variable: "--font-geist-mono",
    weight: "100 900",
});

export const metadata: Metadata = {
    title: "Web App",
    description: "A web application",
};

export default function RootLayout({
    children,
}: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <html lang="en">
            <body
                className={`${geistSans.variable} ${geistMono.variable} antialiased`}
            >
                <Providers>
                    <PageContainer>
                        <Navbar />
                        <AppSidebar />
                        {children}
                    </PageContainer>
                </Providers>
            </body>
        </html>
    );
}
