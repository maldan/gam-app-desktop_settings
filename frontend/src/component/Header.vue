<template>
  <div class="block">
    <button @click="$router.back()">Back</button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../util/RestApi';

export default defineComponent({
  props: {},
  components: {},
  async mounted() {},
  methods: {},
  data: () => {
    return {};
  },
});
</script>

<style lang="scss" scoped>
.block {
  margin-bottom: 10px;

  button {
    background: #5a5a5a;
    border: 0;
    outline: none;
    color: #fefefe;
    padding: 5px 15px;
    border-radius: 4px;

    cursor: pointer;

    &:hover {
      opacity: 0.7;
    }

    &:active {
      position: relative;
      opacity: 0.6;
      top: 1px;
    }
  }
}
</style>
