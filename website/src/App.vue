<script setup lang="ts">
import { ref, computed } from 'vue'
import EnvironmentSwitcher from '@/components/EnvironmentSwitcher.vue'
import Nav, { type LinkProp } from '@/components/Nav.vue'
import { cn } from '@/lib/utils'
import { Separator } from '@/components/ui/separator'
import { TooltipProvider } from '@/components/ui/tooltip'
import { ResizableHandle, ResizablePanel, ResizablePanelGroup } from '@/components/ui/resizable'
import { useContainersStore } from '@/stores/containers'
import { useImagesStore } from '@/stores/images'
import { useVolumesStore } from '@/stores/volumes'

const containersStore = useContainersStore()
const imagesStore = useImagesStore()
const volumesStore = useVolumesStore()
const isCollapsed = ref(false)

const environmentLinks: LinkProp[] = [
  {
    title: 'Dashboard',
    path: '/dashboard',
    icon: 'lucide:gauge',
  },
  {
    title: 'Containers',
    path: '/containers',
    label: computed(() => containersStore.count <= 99 ? containersStore.count : '99+'),
    icon: 'lucide:container',
  },
  {
    title: 'Images',
    path: '/images',
    label: computed(() => imagesStore.count <= 99 ? imagesStore.count : '99+'),
    icon: 'lucide:file-image',
  },
  {
    title: 'Volumes',
    path: '/volumes',
    label: computed(() => volumesStore.count <= 99 ? volumesStore.count : '99+'),
    icon: 'lucide:hard-drive',
  },
  {
    title: 'Networks',
    path: '/networks',
    icon: 'lucide:network',
  },
]

const globalLinks: LinkProp[] = [
  {
    title: 'Settings',
    icon: 'lucide:settings',
  },
  {
    title: 'Documentation',
    icon: 'lucide:book-open-text',
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
          <Nav :is-collapsed="isCollapsed" :links="globalLinks"/>
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
