<script lang="ts" setup>
import { format, fromUnixTime } from 'date-fns'
import { Badge } from '@/components/ui/badge'
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import { useContainersStore } from '@/stores/containers'

const containersStore = useContainersStore()
</script>

<template>
  <div class="items-center px-4 py-2">
    <h1 class="text-xl font-bold mb-4">
      Containers
    </h1>

    <Table>
      <TableHeader>
        <TableRow>
          <TableHead>Name</TableHead>
          <TableHead>Status</TableHead>
          <TableHead>Image</TableHead>
          <TableHead>Ports</TableHead>
          <TableHead class="text-right">Created</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        <TableRow v-for="container in containersStore.containers" :key="container.Id">
          <TableCell class="font-medium">
            {{ container.Names[0].replace('/', '') }}
          </TableCell>
          <TableCell>
            <Badge :variant="container.State === 'running' ? 'default' : 'destructive'">
              {{ container.Status }}
            </Badge>
          </TableCell>
          <TableCell>{{ container.Image }}</TableCell>
          <TableCell>{{ container.Ports }}</TableCell>
          <TableCell class="text-right">{{ format(fromUnixTime(container.Created), "yyyy-MM-dd'T'HH:mm:ss.SSSxxx") }}</TableCell>
        </TableRow>
      </TableBody>
    </Table>
  </div>
</template>