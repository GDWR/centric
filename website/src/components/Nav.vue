<script lang="ts" setup>
import { useRoute, RouterLink } from 'vue-router'
import { Icon } from '@iconify/vue'
import { cn } from '@/lib/utils'
import { buttonVariants } from '@/components/ui/button'
import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from '@/components/ui/tooltip'

export interface LinkProp {
  title: string
  path: string
  label?: string
  icon: string
}

interface NavProps {
  isCollapsed: boolean
  links: LinkProp[]
}

defineProps<NavProps>()
const route = useRoute()
</script>

<template>
  <div :data-collapsed="isCollapsed" class="group flex flex-col gap-4 py-2 data-[collapsed=true]:py-2">
    <nav class="grid gap-1 px-2 group-[[data-collapsed=true]]:justify-center group-[[data-collapsed=true]]:px-2">
      <template v-for="(link, index) of links">
        <Tooltip v-if="isCollapsed" :key="`1-${index}`" :delay-duration="0">
          <TooltipTrigger as-child>
            <a
              href="#"
              :class="cn(
                buttonVariants({ variant: link.path == route.path, size: 'icon' }),
                'h-9 w-9',
                link.path === route.path
                  && 'dark:bg-muted dark:text-muted-foreground dark:hover:bg-muted dark:hover:text-white',
              )"
            >
              <Icon :icon="link.icon" class="size-4" />
              <span class="sr-only">{{ link.title }}</span>
            </a>
          </TooltipTrigger>
          <TooltipContent side="right" class="flex items-center gap-4">
            {{ link.title }}
            <span v-if="link.label" class="ml-auto text-muted-foreground">
              {{ link.label }}
            </span>
          </TooltipContent>
        </Tooltip>

        <RouterLink
          v-else
          :key="`2-${index}`"
          :to="link.path"
          :class="cn(
            buttonVariants({ variant: link.path === route.path ? 'default' : 'ghost', size: 'sm' }),
            link.path === route.path
              && 'dark:bg-muted dark:text-white dark:hover:bg-muted dark:hover:text-white',
            'justify-start',
          )"
        >

          <Icon :icon="link.icon" class="mr-2 size-4" />
          {{ link.title }}
          <span
            v-if="link.label"
            :class="cn(
              'ml-auto',
              link.path === route.path
                && 'text-background dark:text-white',
            )"
          >
            {{ link.label }}
          </span>
        </RouterLink>
      </template>
    </nav>
  </div>
</template>