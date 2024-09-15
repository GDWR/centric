<script lang="ts" setup>
import { ref, watch } from 'vue'
import { Icon } from '@iconify/vue'
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select'
import { cn } from '@/lib/utils'

import { useEnvironmentStore } from '@/stores/environment'

const environmentStore = useEnvironmentStore();
const selected = ref(environmentStore.selectedEnvironment);

watch(() => environmentStore.selectedEnvironment, (value) => {
  selected.value = value;
});

defineProps<{ isCollapsed: boolean }>()
</script>

<template>
  <Select v-model="selected" @update:modelValue="environmentStore.setSelectedEnvironmentId($event)">
    <SelectTrigger aria-label="Select environment"
      :class="cn(
        'flex items-center gap-2 [&>span]:line-clamp-1 [&>span]:flex [&>span]:w-full [&>span]:items-center [&>span]:gap-1 [&>span]:truncate [&_svg]:h-4 [&_svg]:w-4 [&_svg]:shrink-0',
        { 'flex h-9 w-9 shrink-0 items-center justify-center p-0 [&>span]:w-auto [&>svg]:hidden': isCollapsed },
      )"
    >
      <SelectValue placeholder="Select an environment">
        <div class="flex items-center gap-3">
          <Icon class="size-4" :icon="environmentStore.selectedEnvironment?.icon" />
          <span v-if="!isCollapsed">
            {{ environmentStore.selectedEnvironment?.name }}
          </span>
        </div>
      </SelectValue>
    </SelectTrigger>
    <SelectContent>
      <SelectItem v-for="environment of environmentStore.environments.values()" :key="environment.id" :value="environment.id" :value-text="environment.name">
        <div class="flex items-center gap-3 [&_svg]:size-4 [&_svg]:shrink-0 [&_svg]:text-foreground">
          <Icon class="size-4" :icon="environment.icon" />
          {{ environment.name }}
        </div>
      </SelectItem>
    </SelectContent>
  </Select>
</template>