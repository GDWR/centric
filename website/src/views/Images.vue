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
import { useImagesStore } from '@/stores/images'
import { formatBytes } from '@/lib/format'

const imagesStore = useImagesStore()
</script>

<template>
  <div class="items-center px-4 py-2">
    <h1 class="text-xl font-bold mb-4">
      Images
    </h1>

    <Table>
      <TableHeader>
        <TableRow>
          <TableHead>Id</TableHead>
          <TableHead>Tags</TableHead>
          <TableHead>Size</TableHead>
          <TableHead>Use by</TableHead>
          <TableHead class="text-right">Created</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        <TableRow v-for="image in imagesStore.images" :key="image.Id">
          <TableCell class="font-medium">{{ image.Id }}</TableCell>
          <TableCell>{{ image.RepoTags }}</TableCell>
          <TableCell>{{ formatBytes(image.Size) }}</TableCell>
          <TableCell>{{ image.Containers === -1 ? "None" : image.Containers }}</TableCell>
          <TableCell class="text-right">{{ format(fromUnixTime(image.Created), "yyyy-MM-dd'T'HH:mm:ss.SSSxxx") }}</TableCell>
        </TableRow>
      </TableBody>
    </Table>
  </div>
</template>