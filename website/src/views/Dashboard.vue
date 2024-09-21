<script lang="ts" setup>
import { RouterLink } from 'vue-router'
import { DonutChart } from '@/components/ui/chart-donut'
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { useEnvironmentStore } from '@/stores/environment'

const { selectedEnvironmentId } = useEnvironmentStore()
const environmentData = await fetch(`/api/v1/environments/${selectedEnvironmentId}/system-information`)
  .then((res) => res.json());
</script>

<template>
  <div class="items-center px-4 py-2">
    <h1 class="text-xl font-bold mb-4">
      Dashboard
    </h1>

    <div class="grid gap-2 lg:grid-cols-3 lg:gap-4">
      <Card>
        <CardHeader>
          <CardTitle>System Information</CardTitle>
          <CardDescription>
            Information about the system running the Docker daemon.
          </CardDescription>
        </CardHeader>
        <CardContent class="flex flex-col gap-2">
          <span>Cpus: {{ environmentData["NCPU"] }}</span>
          <span>Memory: {{ Math.round(environmentData["MemTotal"]/1073741824) }} GB</span>
          <span>HostName: {{ environmentData["Name"] }}</span>
          <span>Operating System: {{ environmentData["OperatingSystem"] }}</span>
          <span>Kernel Version: {{ environmentData["KernelVersion"] }}</span>
          <span>Architecture: {{ environmentData["Architecture"] }}</span>
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Containers</CardTitle>
        </CardHeader>
        <CardContent>
          <DonutChart
            index="name"
            :category="'total'"
            :data="[
              { name: 'Running', total: environmentData['ContainersRunning'] },
              { name: 'Stopped', total: environmentData['ContainersStopped'] },
              { name: 'Paused', total: environmentData['ContainersPaused'] },
            ]"
            :colors="['green', 'grey', 'orange']"
          />
        </CardContent>
        <CardFooter>
          <Button as-child>
            <RouterLink class="w-full" to="/containers">
              More Details
            </RouterLink>
          </Button>
        </CardFooter>
      </Card>
    </div>
  </div>
</template>