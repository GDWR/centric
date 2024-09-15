<script setup lang="ts">
import * as z from 'zod'
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { Button } from '@/components/ui/button'
import { AutoForm } from '@/components/ui/auto-form'

const router = useRouter()
const schema = z.object({
  adminAccount: z.object({
    username: z.string(),
    password: z.string(),
    confirmPassword: z.string(),
  }),
  environment: z.object({
    uri: z.string(),
    name: z.string(),
  }).default({
    uri: 'unix:///var/run/docker.sock',
    name: 'Main',
  }),
});

function onSubmit(values: Record<string, any>) {
  fetch('/api/v1/system/initial-setup', {
    method: 'POST',
    body: JSON.stringify(values),
  })
}

onMounted(() => {
  fetch('/api/v1/system')
    .then((res) => res.json())
    .then((data) => {
      if (data.initialSetupComplete) {
        router.push('/')
      }
    })
})
</script>

<template>
  <div class="flex items-center justify-center h-screen bg-gray-100">
    <div class="w-[38rem] max-w-full p-6 bg-white rounded-lg shadow-lg">
      <h1 class="w-fit text-2xl font-semibold text-gray-800 mb-4">First time setup</h1>
      <AutoForm class="w-2/3 space-y-6" @submit="onSubmit" :schema="schema"
        :field-config="{
          adminAccount: {
            password: { inputProps: {type: 'password'} },
            confirmPassword: { inputProps: {type: 'password'} },
          }
        }"
      >
        <Button type="submit">
          Submit
        </Button>
      </AutoForm>
    </div>
  </div>
</template>