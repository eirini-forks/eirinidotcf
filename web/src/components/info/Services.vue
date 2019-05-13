<template>
  <section id="services" class="hide-overflow">
    <v-layout grey lighten-3>
      <v-flex xs12 md6 text-xs-center pa-5>
        <v-layout align-center justify-center fill-height wrap>
          <v-flex xs12>
            <base-heading class="info--text">
              Explore
            </base-heading>
          </v-flex>
          <v-flex v-for="(service, i) in services" :key="i" md6 text-xs-center mb-3>
            <v-avatar class="elevation-6 mb-2" color="#69A1BB" size="64" tile>
              <v-icon dark size="52" @click="openPage(service.link)" v-text="service.icon" />
            </v-avatar>
            <base-text>
              <div class="mb-2" v-text="service.name" />
            </base-text>
          </v-flex>
        </v-layout>
      </v-flex>
      <v-flex hidden-sm-and-down md6>
        <v-img src="https://images.unsplash.com/photo-1482398650355-d4c6462afa0e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&w=1000&q=80" height="100%" />
      </v-flex>
    </v-layout>
  </section>
</template>

<script>
  export default {
    data: () => ({
      services: [{
        name: 'About',
        icon: 'mdi-information-outline',
        link: 'https://www.cloudfoundry.org/project-eirini/'
      }, {
        name: 'Helm Release',
        icon: 'mdi-clipboard-text-outline',
        link: 'https://github.com/cloudfoundry-incubator/eirini-release'
      }, {
        name: 'Source Code',
        icon: 'mdi-code-braces',
        link: 'https://github.com/cloudfoundry-incubator/eirini-release'
      }, {
        name: 'Eirini CI',
        icon: 'mdi-airplane-takeoff',
        link: 'https://ci.flintstone.cf.cloud.ibm.com/teams/eirini/pipelines/ci'
      }]
    }),

    methods: {
      openPage: function(link) {
        window.open(link, '_blank')
      }
    }
  }
</script>
