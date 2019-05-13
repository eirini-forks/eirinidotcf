<template>
  <section id="services" class="hide-overflow">
    <v-layout grey lighten-3>
      <v-flex hidden-sm-and-down md6>
        <v-img src="https://cdn.wallpapersafari.com/6/93/c8xaAH.jpg" height="100%" />
      </v-flex>
      <v-flex xs12 md6 text-xs-center pa-5>
        <v-layout align-center justify-center fill-height wrap>
          <v-flex xs12>
            <base-heading class="info--text">
              Need Help?
            </base-heading>
            <base-text class="mb-5">
              Find us on the Cloud Foundry Slack if you have any question.
              <br>
            </base-text>
          </v-flex>
          <v-flex v-for="(service, i) in services" :key="i" md6 text-xs-center mb-3>
            <v-avatar class="elevation-6 mb-2" color="red" size="64" tile>
              <v-icon dark size="52" @click="openPage(service.link)" v-text="service.icon" />
            </v-avatar>
            <base-text>
              <div class="mb-2" v-text="service.name" />
            </base-text>
          </v-flex>
          <v-flex xs12>
            <h2 class="display-5">Office Hours</h2>
            <base-text class="mb-5">
              We are giving office hours every first monday in the month at 4:30pm (CET)
              <br>
            </base-text>
          </v-flex>
        </v-layout>
      </v-flex>
    </v-layout>
  </section>
</template>

<script>
  export default {
    data: () => ({
      services: [{
        name: '#eirini-dev',
        icon: 'mdi-slack',
        link: 'https://cloudfoundry.slack.com/messages/C8RU3BZ26'
      }, ]
    }),

    methods: {
      openPage: function(link) {
        window.open(link, '_blank')
      }
    }
  }
</script>
