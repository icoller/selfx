<template>
  <a-tabs type="rounded" animation>
    <a-tab-pane key="article" :title="$t('article')">
      <Item :data="article" />
    </a-tab-pane>
    <a-tab-pane key="category" :title="$t('category')">
      <Item :data="category" />
    </a-tab-pane>
    <a-tab-pane key="tag" :title="$t('tag')">
      <Item :data="tag" />
    </a-tab-pane>
    <a-tab-pane key="link" :title="$t('link')">
      <Item :data="link" />
    </a-tab-pane>
    <a-tab-pane key="crawl" :title="$t('crawl')">
      <Item :data="crawl" />
    </a-tab-pane>
  </a-tabs>

</template>

<script setup>
  import Item from "./Item.vue";
  import {article,category,tag,link, crawl} from "./data"
</script>