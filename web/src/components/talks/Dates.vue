<template>
  <v-container fluid grid-list-lg>
    <v-layout wrap>
      <v-flex justify-center layout>
        <base-heading class="info--text">
          Dates
        </base-heading>
      </v-flex>
      <v-flex justify-center layout md12 xs12>
        <v-divider light></v-divider>
      </v-flex>
    </v-layout>
    <template>
      <v-layout row>
        <v-flex xs12 sm6 offset-sm3>
          <v-card>
            <v-list two-line>
              <template v-for="(item, index) in items">
                <v-subheader v-if="item.header" :key="item.header">
                  {{ item.header }}
                </v-subheader>
                <v-divider v-else-if="item.divider" :key="index" :inset="item.inset"></v-divider>
                <v-list-tile v-else :key="item.title" avatar @click="">
                  <v-list-tile-avatar>
                    <img :src="item.avatar">
                  </v-list-tile-avatar>
                  <v-list-tile-content>
                    <v-list-tile-title>{{ item.date }} / {{ item.location }}</v-list-tile-title>
                    <v-list-tile-title v-html="item.title"></v-list-tile-title>
                    <v-list-tile-sub-title v-html="item.speaker"></v-list-tile-sub-title>
                  </v-list-tile-content>
                </v-list-tile>
              </template>
            </v-list>
          </v-card>
        </v-flex>
      </v-layout>
    </template>
  </v-container>
</template>

<script>
import * as axios from 'axios';

export default {
  created () {
     axios
       .get('https://raw.githubusercontent.com/JulzDiverse/eirinidotcf/data/dates.json')
       .then(response => (this.items = response.data))
  },

  data: () => ({
    items: null
  })
}
</script>
