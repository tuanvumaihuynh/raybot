<script setup lang="ts">
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from '@/components/ui/card'
</script>

<template>
  <div class="container py-12 mx-auto">
    <div class="mb-12 text-center">
      <h1 class="mb-4 text-4xl font-bold">
        Welcome to Our App
      </h1>
      <p class="text-lg text-muted-foreground">
        A modern Vue 3 application with Tailwind CSS
      </p>
    </div>

    <div class="grid grid-cols-1 gap-6 md:grid-cols-3">
      <Card>
        <CardHeader>
          <CardTitle>Feature One</CardTitle>
          <CardDescription>Powerful and intuitive interface</CardDescription>
        </CardHeader>
        <CardContent>
          <p>Our app provides a seamless experience with a modern UI built on Vue 3 and Tailwind CSS.</p>
        </CardContent>
        <CardFooter>
          <Button class="w-full">
            Learn More
          </Button>
        </CardFooter>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Feature Two</CardTitle>
          <CardDescription>Fast and responsive</CardDescription>
        </CardHeader>
        <CardContent>
          <p>Built with performance in mind, our app loads quickly and responds instantly to user interactions.</p>
        </CardContent>
        <CardFooter>
          <Button class="w-full" variant="secondary">
            Explore
          </Button>
        </CardFooter>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Feature Three</CardTitle>
          <CardDescription>Customizable themes</CardDescription>
        </CardHeader>
        <CardContent>
          <p>Switch between light and dark mode with our built-in theme support powered by VueUse.</p>
        </CardContent>
        <CardFooter>
          <Button class="w-full" variant="outline">
            Try Now
          </Button>
        </CardFooter>
      </Card>
    </div>

    <div class="mt-12 text-center">
      <RouterLink to="/home">
        <Button size="lg">
          Go to Home
        </Button>
      </RouterLink>
    </div>
  </div>
</template>
