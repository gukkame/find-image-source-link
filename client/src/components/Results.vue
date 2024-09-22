<template>
  <div
    class="p-4 mt-4 md:px-12 bg-gray-900 text-white flex justify-center md:items-center md:text-center flex-wrap flex-col w-full"
  >
    <h2 v-show="text" class="text-2xl pb-6 w-full">{{ text }}</h2>
    <div v-if="arrayOfLinks" v-for="item in arrayOfLinks" class="truncate w-full py-2">
      <a :href="item" target="_blank" class="">
        {{ item }}
      </a>
    </div>
    <h2 v-else>No results!</h2>
  </div>
</template>
<script>
export default {
  data() {
    return {}
  },
  props: {
    text: {
      type: String,
      default: ''
    },
    arrayOfLinks: {
      type: Array,
      default: []
    }
  }
}
</script>
