<template>
  <div class="main">
    <Header />
    <div class="item" v-for="x in list" :key="x.id">
      <div>Desktop {{ x.id + 1 }}</div>
      <div>
        <img style="width: 320px" :src="`${$root.API_URL}/main/background?id=${x.id}`" alt="" />
      </div>
      <div>
        <input @change="change($event, x.id)" :ref="`file_${x.id}`" type="file" accept="image/*" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Header from '../component/Header.vue';
import { RestApi } from '../util/RestApi';

export default defineComponent({
  components: { Header },
  async mounted() {},
  methods: {
    async change(e: any, id: number) {
      await RestApi.main.uploadBackground(id, e.target.files[0]);
      window.location.reload();
    },
  },
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
    align-items: center;

    > div {
      flex: 1;
    }
  }
}
</style>
