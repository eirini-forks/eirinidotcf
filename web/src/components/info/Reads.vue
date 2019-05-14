<template>
  <section id="services" class="hide-overflow">
    <v-layout grey lighten-3>
      <v-flex xs12 md6 text-xs-center pa-5>
        <v-flex xs12>
          <base-heading class="info--text">
            Reads
          </base-heading>
        </v-flex>
        <v-card>
          <v-list>
            <v-list-tile v-for="item in items" :key="item.title" avatar @click="openPage(item.url)">
              <v-list-tile-content>
                <v-list-tile-title v-text="item.title"></v-list-tile-title>
              </v-list-tile-content>
            </v-list-tile>
          </v-list>
        </v-card>
      </v-flex>
      <v-flex hidden-sm-and-down md6>
        <v-img src="https://www.iwantwallpaper.co.uk/images/muriva-books-faux-effect-wallpaper-j43027-p968-1441_image.jpg" height="100%" />
      </v-flex>
    </v-layout>
  </section>
</template>

<script>
import * as axios from 'axios';

export default {
    mounted () {
       axios
         .get('https://raw.githubusercontent.com/JulzDiverse/eirinidotcf/data/reads.json')
         .then(response => (this.items = response.data))
    },

    data() {
      return {
        items: []
      }
    },
    methods: {
      openPage: function(link) {
        window.open(link, '_blank')
      }
    }
  }
</script>
