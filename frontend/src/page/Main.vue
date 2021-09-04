<template>
  <div class="main">
    <div class="item clickable" v-for="x in list" :key="x.title">
      {{ x.title }}
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  components: {},
  async mounted() {},
  methods: {},
  data: () => {
    return {
      list: [
        {
          title: 'Background',
        },
        {
          title: 'Backup',
        },
      ],
    };
  },
});
</script>

<style lang="scss" scoped>
.main {
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
  gap: 5px;
  padding: 5px;

  .item {
    background: #1b1b1b;
    padding: 10px;
    color: #b3b3b3;
  }
}
</style>
