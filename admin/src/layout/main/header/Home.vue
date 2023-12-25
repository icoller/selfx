<template>
  <a-button @click="open">
    <template #icon>
      <icon-home :size="18" />
    </template>
  </a-button>
</template>

<script setup>
  import {useOpenLink} from "@/hooks/utils";
  import {useSiteURL} from "@/hooks/app"
  import {useStore} from "@/store";


  const store = useStore()
  function open(){
    useOpenLink(useSiteURL(store),'')
  }

</script>