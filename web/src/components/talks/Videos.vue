<template>
  <v-container fluid grid-list-lg>
    <v-layout wrap>
      <v-flex justify-center layout>
        <base-heading class="info--text">
          Videos
        </base-heading>
      </v-flex>
      <v-flex justify-center layout md12 xs12>
        <v-divider light></v-divider>
      </v-flex>
    </v-layout>
    <v-layout wrap>
      <v-flex
        v-for="(vid, i) in videos"
        justify-center
        align-center
        layout
      >
        <v-card width="800">
          <iframe width="100%" height="315" :src="embedUrl(vid.id)" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
          <v-card-title primary-title>
            <div>
              <div class="headline">{{ vid.title }}</div>
              <span class="grey--text">{{ vid.speaker }}</span>
            </div>
          </v-card-title>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import * as axios from 'axios';

export default {
  data: () => ({
    videos: null
  }),

  created () {
     console.log("mounting")
     axios
       .get('https://raw.githubusercontent.com/JulzDiverse/eirinidotcf/data/videos.json')
       .then(response => (this.videos = response.data))
       .catch(error => console.log(error))
  },

  methods: {
    embedUrl: function (id) {
      return "https://www.youtube.com/embed/" + id
    }
  }
}
</script>
