<template>
  <v-container
    grid-list-xl
  >
    <v-layout wrap>
      <v-flex xs12>
        <slot />
      </v-flex>

      <feed-card
        v-for="(photo, i) in paginatedPhotos"
        :key="photo.title"
        :size="layout[i]"
        :value="photo"
      />
    </v-layout>

    <v-layout align-center>
      <v-flex xs3>
        <base-btn
          v-if="page !== 1"
          class="ml-0"
          title="Previous page"
          square
          @click="page--"
        >
          <v-icon>mdi-chevron-left</v-icon>
        </base-btn>
      </v-flex>

      <v-flex
        xs6
        text-xs-center
        subheading
      >
        PAGE {{ page }} OF {{ pages }}
      </v-flex>

      <v-flex
        xs3
        text-xs-right
      >
        <base-btn
          v-if="pages > 1 && page < pages"
          class="mr-0"
          title="Next page"
          square
          @click="page++"
        >
          <v-icon>mdi-chevron-right</v-icon>
        </base-btn>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
  // Utilities
  import {
    mapState
  } from 'vuex'
  import * as axios from 'axios';


  export default {
    name: 'Feed',

    components: {
      FeedCard: () => import('@/components/FeedCard')
    },

    mounted: function() {
       this.downloadContents();
    },

    data: () => ({
      layout: [2, 2, 1, 2, 2, 3, 3, 3, 3, 3, 3],
      page: 1,
      contents: []
    }),

    methods: {
      downloadContents(){
       axios.get(this.config.api_url+"/photos")
       .then((response) => {
         console.log(response);
         this.contents = response.data;
       })
       .catch(function (error) {
         console.log(error);
       })
      }
    },

    computed: {
      ...mapState(['config']),

      pages () {
        return Math.ceil(this.contents.length / 11)
      },
      paginatedPhotos () {
        const start = (this.page - 1) * 11
        const stop = this.page * 11

        return this.contents.slice(start, stop)
      }
    },

    watch: {
      page () {
        window.scrollTo(0, 0)
      }
    }
  }
</script>
