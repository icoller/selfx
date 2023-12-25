<template>
  <a-card class="w-full" :title="$t('about') + ' SelfX'" :bordered="false">
    <div v-if="appInfo" class="mb-2 leading-8 italic">
      version {{ appInfo?.version }}
    </div>
    <div class="mb-2 leading-8">
      {{ $t('selfxAbout') }}<br>{{ $t('selfxInstructions') }}
    </div>
    <a-link @click="useOpenLink">
      <template #icon><icon-github /></template>https://github.com/icoller/selfx
    </a-link>
  </a-card>
</template>

<script setup>
  import {useRequest} from "vue-request";
  import {useOpenLink} from '@/hooks/utils.js'
  import {dashboardData} from "@/api/index.js";


  const { data:appInfo } = useRequest(dashboardData, {defaultParams:['appInfo']})
</script>