<template>
  <section>
    <v-container fluid grid-list-lg>
      <v-layout grey lighten-3>
        <v-flex>
          <v-layout align-center justify-center fill-height wrap>
            <v-flex justify-center layout md12 xs12>
              <v-divider light></v-divider>
            </v-flex>
            <v-flex v-for="(avatar, i) in avatars" justify-center layout>
              <v-card color="transparent" class="black--text" width="400px">
                <v-layout>
                  <v-flex xs5>
                    <v-img
                      :src="avatar.image"
                      height="125px"
                      contain
                    />
                  </v-flex>
                  <v-flex xs7>
                    <v-card-title primary-title>
                      <div>
                        <div class="headline">
                          {{ avatar.name }}
                        </div>
                        <div>{{ avatar.role }}</div>
                        <div>{{ avatar.company }}</div>
                        <div>{{ avatar.city }}</div>
                        <div>{{ avatar.email }}</div>
                      </div>
                    </v-card-title>
                  </v-flex>
                </v-layout>
                <v-divider light />
                <v-card-actions class="pa-3">
                  <v-btn
                    flat
                    icon
                    color="black"
                    @click="openPage('https://github.com/'+avatar.git)"
                  >
                    <v-icon>mdi-github-circle</v-icon>
                  </v-btn>
                  <v-btn
                    flat
                    icon
                    color="info"
                    @click="openPage('https://twitter.com/'+avatar.twitter)"
                  >
                    <v-icon>mdi-twitter</v-icon>
                  </v-btn>
                  <v-spacer />
                </v-card-actions>
              </v-card>
            </v-flex>
          </v-layout>
        </v-flex>
      </v-layout>
    </v-container>
  </section>
</template>

<script>
import * as axios from 'axios';

export default {
    created () {
       axios
         .get('https://raw.githubusercontent.com/JulzDiverse/eirinidotcf/data/team.json')
         .then(response => (this.avatars = response.data))
    },

    data: () => ({
      avatars: []
    }),

    methods: {
      openPage: function (link) {
        window.open(link, '_blank')
      }
    }
  }
</script>
