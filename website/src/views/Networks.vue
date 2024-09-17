<script lang="ts" setup>
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import { useNetworksStore } from '@/stores/networks'

const networksStore = useNetworksStore()
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
          <TableHead>Id</TableHead>
          <TableHead class="text-right">Created</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        <TableRow v-for="network in networksStore.networks" :key="network.Name">
          <TableCell class="font-medium">{{ network.Name }}</TableCell>
          <TableCell>{{ network.Driver }}</TableCell>
          <TableCell>{{ network.Id }}</TableCell>
          <TableCell class="text-right">{{ network.Created }}</TableCell>
        </TableRow>
      </TableBody>
    </Table>
  </div>
</template>