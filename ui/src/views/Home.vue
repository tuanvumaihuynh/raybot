<script setup lang="ts">
import { Button } from '@/components/ui/button'
</script>

<template>
  <div class="container py-12 mx-auto">
    <div class="max-w-3xl mx-auto">
      <h1 class="mb-6 text-3xl font-bold">
        Home
      </h1>

      <p class="mb-4">
        We are a team of passionate developers building modern web applications with Vue 3,
        TypeScript, and Tailwind CSS.
      </p>

      <p class="mb-6">
        Our mission is to create intuitive, performant, and beautiful user interfaces that
        provide an exceptional user experience.
      </p>

      <RouterLink to="/system">
        <Button variant="outline">
          Go to System
        </Button>
      </RouterLink>
    </div>
  </div>
</template>
