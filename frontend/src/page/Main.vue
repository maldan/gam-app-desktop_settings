<template>
  <div class="main">
    <div @click="$router.push(x.path)" class="item clickable" v-for="x in list" :key="x.title">
      <img :src="x.image" :alt="x.title" />
      <div>{{ x.title }}</div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
declare const require: (path: string) => void;

export default defineComponent({
  components: {},
  async mounted() {},
  methods: {},
  data: () => {
    return {
      list: [
        {
          title: 'Background',
          image: require('../asset/background.svg'),
          path: '/background',
        },
        {
          title: 'Backup',
          image: require('../asset/backup.svg'),
          path: '/backup',
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
    display: flex;
    align-items: center;

    img {
      display: block;
      margin-right: 10px;
    }
  }
}
</style>
