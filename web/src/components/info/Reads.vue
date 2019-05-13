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
  export default {
    data() {
      return {
        items: [
          { title: 'The Fresh Prince of Cloud Native: Bringing Cloud Foundry & Kubernetes Together', url: 'https://medium.com/@doctor_julz/the-fresh-prince-of-cloud-native-bringing-cloud-foundry-kubernetes-together-93eedb2d1dec' },
          { title: 'Project Eirini: Choose the best fitting container scheduler', url: 'https://medium.com/@julian.skupnjak/project-eirini-choose-the-best-fitting-container-orchestrator-11bc31866f9e' },
          { title: 'Cloud Foundry, Kubernetes & The Fundamental Emptiness of Labels', url: 'https://medium.com/@doctor_julz/cloud-foundry-kubernetes-the-fundamental-emptiness-of-labels-bf9bec688c3e' },
          { title: 'With Project Eirini, Cloud Foundry Adapts to a New Open Source Ecosystem', url: 'https://thenewstack.io/with-project-eirini-cloud-foundry-adapts-to-a-new-open-source-ecosystem/' },
        ]
      }
    },
    methods: {
      openPage: function(link) {
        window.open(link, '_blank')
      }
    }
  }
</script>
