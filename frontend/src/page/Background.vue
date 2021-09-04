<template>
  <div class="main">
    <Header />
    <div class="item" v-for="x in list" :key="x.id">
      <div>Desktop {{ x.id + 1 }}</div>
      <div><input type="file" /></div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Header from '../component/Header.vue';

export default defineComponent({
  components: { Header },
  async mounted() {},
  methods: {},
  data: () => {
    return {
      list: [
        { id: 0, image: '' },
        { id: 1, image: '' },
        { id: 2, image: '' },
        { id: 3, image: '' },
      ],
    };
  },
});
</script>

<style lang="scss" scoped>
.main {
  padding: 5px;
  color: #b3b3b3;

  .item {
    display: flex;
  }
}
</style>
