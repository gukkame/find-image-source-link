<script setup>
import Results from '../components/Results.vue'
</script>

<template>
  <div class="">
    <div class="bg-sky-700 m-4 p-4 gap-4 flex justify-center flex-col rounded-lg">
      <div class="flex flex-col justify-around text-center my-2">
        <h2 class="text-3xl text-white">Image source search by image url using Vertex AI API</h2>
      </div>

      <label for="message" class="block pl-2 text-sm font-medium text-gray-900 dark:text-white"
        >Paste your link here:</label
      >
      <textarea
        class="block p-2.5 w-full h-28 text-sm text-white bg-gray-900 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 border-blue-500"
        placeholder="Paste link to Image to recieve links where it comes from "
        v-model="photoURL"
      ></textarea>
      <div class="flex flex-col justify-around items-center my-2">
              <img v-if="photoURL" :src="photoURL" class="max-w-40 w-auto" />
      </div>

      <button
        type="button"
        class="relative inline-flex items-center justify-center p-0.5 mb-2 overflow-hidden text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-cyan-500 to-blue-500 group-hover:from-cyan-500 group-hover:to-blue-500 hover:text-white dark:text-white"
        @click="submit()"
      >
        <span
          class="w-full relative px-5 py-2.5 transition-all ease-in duration-75 bg-gray-900 rounded-md group-hover:bg-opacity-0"
        >
          Find links!
        </span>
      </button>
      <Results class="rounded-lg" :text="text" :arrayOfLinks="arrayOfLinks" />
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      photoURL: null,
      arrayOfLinks: null,
      text: null
    }
  },

  methods: {
    submit() {
      if (!this.photoURL) {
        console.log('No Link given')
        this.arrayOfLinks = ['Please add image link!']
        return
      }
      axios
        .post('http://localhost:8080/', {
          PhotoURL: this.photoURL
        })
        .then((res) => {
          if (res.data.FullMatchingLinks != null) {
            this.arrayOfLinks = res.data.FullMatchingLinks
          } else if (res.data.ImageMatchingLinks != null) {
            this.arrayOfLinks = res.data.ImageMatchingLinks
          }
          this.text = res.data.Text
        })
        .catch((error) => {
          console.error('Error at login data fetching: ' + error)
        })
    }
  }
}
</script>
