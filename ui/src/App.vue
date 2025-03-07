<script setup lang="ts">
import { Button } from '@/components/ui/button'
import { useColorMode } from '@vueuse/core'
import { Moon, Sun } from 'lucide-vue-next'

const mode = useColorMode()

function toggleDarkMode() {
  mode.value = mode.value === 'dark' ? 'light' : 'dark'
}
</script>

<template>
  <div class="flex flex-col min-h-screen">
    <header class="py-4 border-b">
      <div class="container flex items-center justify-between mx-auto">
        <div class="flex items-center gap-2">
          <img alt="Vue logo" class="w-8 h-8" src="@/assets/logo.svg">
          <span class="text-lg font-bold">Vue App</span>
        </div>

        <nav class="flex items-center gap-6">
          <RouterLink to="/" class="transition-colors hover:text-primary">
            Home
          </RouterLink>
          <RouterLink to="/about" class="transition-colors hover:text-primary">
            About
          </RouterLink>
          <Button variant="ghost" size="icon" @click="toggleDarkMode">
            <Sun v-if="mode === 'dark'" class="w-5 h-5" />
            <Moon v-else class="w-5 h-5" />
          </Button>
        </nav>
      </div>
    </header>

    <main class="flex-1">
      <RouterView />
    </main>

    <footer class="py-6 mt-12 border-t">
      <div class="container mx-auto text-sm text-center text-muted-foreground">
        &copy; {{ new Date().getFullYear() }} Vue App. All rights reserved.
      </div>
    </footer>
  </div>
</template>
