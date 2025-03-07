import type { Component } from 'vue'
import {
  Bot,
  Settings2,
} from 'lucide-vue-next'

interface Item {
  name: string
  path: string
  icon: Component
}

const items: Item[] = [
  {
    name: 'Home',
    path: '/home',
    icon: Bot,
  },
  {
    name: 'System config',
    path: '/system',
    icon: Settings2,
  },
]

export { items as routes }
