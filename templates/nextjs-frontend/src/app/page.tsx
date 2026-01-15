import Link from "next/link";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";

export default function HomePage() {
  return (
    <div className="container mx-auto py-10">
      <div className="flex flex-col items-center justify-center space-y-8">
        <div className="text-center space-y-4">
          <h1 className="text-4xl font-bold tracking-tight">
            Welcome to My App
          </h1>
          <p className="text-xl text-muted-foreground max-w-2xl">
            A modern web application built with Next.js and Go
          </p>
        </div>

        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3 w-full max-w-4xl">
          <Card>
            <CardHeader>
              <CardTitle>Item Management</CardTitle>
              <CardDescription>
                Create, read, update, and delete items
              </CardDescription>
            </CardHeader>
            <CardContent>
              <Button asChild className="w-full">
                <Link href="/items">View Items</Link>
              </Button>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Authentication</CardTitle>
              <CardDescription>
                Secure API with JWT authentication
              </CardDescription>
            </CardHeader>
            <CardContent>
              <Button asChild variant="outline" className="w-full">
                <Link href="/login">Login</Link>
              </Button>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>API Documentation</CardTitle>
              <CardDescription>
                Explore the API endpoints and responses
              </CardDescription>
            </CardHeader>
            <CardContent>
              <Button asChild variant="secondary" className="w-full">
                <a href={`${process.env.NEXT_PUBLIC_API_URL}/docs`} target="_blank" rel="noopener noreferrer">
                  View API Docs
                </a>
              </Button>
            </CardContent>
          </Card>
        </div>

        <div className="text-center text-sm text-muted-foreground">
          <p>Built with Next.js 15, TypeScript, and Tailwind CSS</p>
          <p>Connected to backend API</p>
        </div>
      </div>
    </div>
  );
}