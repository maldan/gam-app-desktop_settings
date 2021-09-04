<template>
  <div class="main">
    <Header />
    <div class="item" v-for="x in list" :key="x.id">
      <div style="flex: none; width: 40px">
        <img :src="`${$root.API_URL}/application/icon?path=${x.path}`" alt="" />
      </div>
      <div>{{ x.author }} / {{ x.name }}</div>
      <div>{{ x.path }}</div>
      <div class="version">{{ x.version }}</div>
      <div>
        <Select :items="backup" v-model="x.backup" @update:modelValue="change(x, $event)" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Header from '../component/Header.vue';
import Select from '../component/Select.vue';
import { RestApi } from '../util/RestApi';

export default defineComponent({
  components: { Header, Select },
  async mounted() {
    const list = await RestApi.application.getList();
    for (let i = 0; i < list.length; i++) {
      list[i].backup = { label: list[i].backup, value: list[i].backup };
    }
    this.list = list;
  },
  methods: {
    change(app: any, value: any) {
      console.log(app, value);
    },
  },
  data: () => {
    return {
      list: [] as any[],
      backup: [
        { label: 'Never', value: 'never' },
        { label: 'Each day', value: '1day' },
        { label: 'Each 3 day', value: '3day' },
        { label: 'Each 7 day', value: '7day' },
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
    padding: 10px;
    align-items: center;
    background: #2b2b2b;
    margin-bottom: 5px;

    img {
      display: block;
      margin-right: 10px;
    }

    > div {
      padding: 5px 10px;
      flex: 1;
    }

    .version {
      flex: none;
      background: #1b1b1b;
      border-radius: 4px;
      width: max-content;
    }
  }
}
</style>
