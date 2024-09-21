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
import { useNetworksStore } from '@/stores/networks'
import { useVolumesStore } from '@/stores/volumes'
import { formatNumber } from './lib/format'

const containersStore = useContainersStore()
const imagesStore = useImagesStore()
const networksStore = useNetworksStore()
const volumesStore = useVolumesStore()
const isCollapsed = ref(false)
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
          <Nav :is-collapsed="isCollapsed" :links="[
            {
              title: 'Dashboard',
              path: '/dashboard',
              icon: 'lucide:gauge',
            },
            {
              title: 'Containers',
              path: '/containers',
              label: computed(() => formatNumber(containersStore.count, 99)),
              icon: 'lucide:container',
            },
            {
              title: 'Images',
              path: '/images',
              label: computed(() => formatNumber(imagesStore.count, 99)),
              icon: 'lucide:file-box',
            },
            {
              title: 'Volumes',
              path: '/volumes',
              label: computed(() => formatNumber(volumesStore.count, 99)),
              icon: 'lucide:hard-drive',
            },
            {
              title: 'Networks',
              path: '/networks',
              label: computed(() => formatNumber(networksStore.count, 99)),
              icon: 'lucide:network',
            },
            {
              title: 'Configs',
              path: '/configs',
              icon: 'lucide:file-cog',
            },
            {
              title: 'Secrets',
              path: '/secrets',
              icon: 'lucide:lock',
            },
          ]"/>
          <Separator />
          <Nav :is-collapsed="isCollapsed" :links="[
            {
              title: 'Swarm',
              path: '/swarm',
              icon: 'lucide:cloud',
            },
            {
              title: 'Services',
              path: '/services',
              icon: 'lucide:shield-ellipsis',
            },
            {
              title: 'Stacks',
              path: '/stacks',
              icon: 'lucide:layers',
            },
          ]"/>
          <Separator />
          <Nav :is-collapsed="isCollapsed" :links="[
            {
              title: 'Settings',
              path: '/settings',
              icon: 'lucide:settings',
            },
            {
              title: 'Documentation',
              path: '/documentation',
              icon: 'lucide:book-open-text',
            },
          ]"/>
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
