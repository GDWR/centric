<script lang="ts" setup>
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import { useVolumesStore } from '@/stores/volumes'

const volumesStore = useVolumesStore()
</script>

<template>
  <div class="items-center px-4 py-2">
    <h1 class="text-xl font-bold mb-4">
      Volumes 
    </h1>

    <Table>
      <TableHeader>
        <TableRow>
          <TableHead>Name</TableHead>
          <TableHead>Driver</TableHead>
          <TableHead class="text-right">Created</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        <TableRow v-for="volume in volumesStore.volumes" :key="volume.Name">
          <TableCell class="font-medium">{{ volume.Name }}</TableCell>
          <TableCell>{{ volume.Driver }}</TableCell>
          <TableCell class="text-right">{{ volume.CreatedAt }}</TableCell>
        </TableRow>
      </TableBody>
    </Table>
  </div>
</template>