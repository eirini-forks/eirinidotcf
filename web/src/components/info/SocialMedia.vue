<template>
  <div>
    <v-btn v-for="(item, i) in items.slice(0, large ? items.length : 3)" :key="i" fab class="ma-2" :large="large" color="info" @click="openEirini()">
      <v-icon dark v-text="item.icon" />
    </v-btn>
  </div>
</template>

<script>
  export default {
    props: {
      large: {
        type: Boolean,
        default: false
      }
    },

    data: () => ({
      items: [{
        name: 'Github',
        icon: 'mdi-git'
      }]
    }),

    methods: {
      openEirini: function() {
        window.open('https://github.com/cloudfoundry-incubator/eirini-release', '_blank')
      }
    }
  }
</script>
