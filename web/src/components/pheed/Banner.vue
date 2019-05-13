<template>
  <base-card dark>
    <v-img class="grey lighten-2" height="400" width="100%" gradient="rgba(0, 0, 0, .42), rgba(0, 0, 0, .42)" :src="require('@/assets/banner.jpeg')">
      <v-layout fill-height align-center pa-3>
        <v-flex xs12 md7 offset-md5>
          <h1 class="display-3 font-weight-light">
            Photo Feed
          </h1>
          <div class="subheading text-uppercase pl-2 mb-4">
            {{ getTitle }}
          </div>
        </v-flex>
      </v-layout>
    </v-img>
  </base-card>
</template>

<script>
  // Utilities
  import {
    mapState
  } from 'vuex'

  export default {
    computed: {
      ...mapState(['config']),
      getTitle() {
        return this.config.title
      }
    }
  }
</script>
