<script setup lang="ts">
import { ref } from 'vue'
import EnvironmentSwitcher from '@/components/EnvironmentSwitcher.vue'
import Nav, { type LinkProp } from '@/components/Nav.vue'
import { cn } from '@/lib/utils'
import { Separator } from '@/components/ui/separator'
import { TooltipProvider } from '@/components/ui/tooltip'
import { ResizableHandle, ResizablePanel, ResizablePanelGroup } from '@/components/ui/resizable'

const isCollapsed = ref(false)

const environmentLinks: LinkProp[] = [
  {
    title: 'Dashboard',
    icon: 'lucide:gauge',
    variant: 'default',
  },
  {
    title: 'Containers',
    icon: 'lucide:container',
    variant: 'ghost',
  },
  {
    title: 'Stacks',
    icon: 'lucide:layers',
    variant: 'ghost',
  },
  {
    title: 'Images',
    label: '',
    icon: 'lucide:file-image',
    variant: 'ghost',
  },
  {
    title: 'Volumes',
    icon: 'lucide:hard-drive',
    variant: 'ghost',
  },
  {
    title: 'Networks',
    label: '',
    icon: 'lucide:network',
    variant: 'ghost',
  },
]

const globalLinks: LinkProp[] = [
  {
    title: 'Settings',
    icon: 'lucide:settings',
    variant: 'ghost',
  },
  {
    title: 'Documentation',
    icon: 'lucide:book-open-text',
    variant: 'ghost',
  },
]
</script>

<template>
  <Toaster />

  <div class="w-screen h-screen">
    <TooltipProvider :delay-duration="0">
      <ResizablePanelGroup direction="horizontal">
        <ResizablePanel collapsible :min-size="5" :max-size="15" :class="cn(isCollapsed && 'min-w-[50px] transition-all duration-300 ease-in-out')" @expand="isCollapsed = false" @collapse="isCollapsed = true">
          <div :class="cn('flex h-[52px] items-center justify-center', isCollapsed ? 'h-[52px]' : 'px-2')">
            <EnvironmentSwitcher :is-collapsed="isCollapsed"  />
          </div>
          <Nav :is-collapsed="isCollapsed" :links="environmentLinks"/>
          <Separator />
          <Nav class="" :is-collapsed="isCollapsed" :links="globalLinks"/>
        </ResizablePanel>

        <ResizableHandle />
  
        <ResizablePanel :min-size="30">
          <Suspense>
            <RouterView />
          </Suspense>
        </ResizablePanel>
  
      </ResizablePanelGroup>
    </TooltipProvider>
  </div>
</template>
