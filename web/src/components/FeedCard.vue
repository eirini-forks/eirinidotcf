<template>
  <v-flex
    xs12
    :class="classes"
  >
    <base-card
      :height="value.prominent ? 450 : 350"
      color="grey lighten-1"
      dark
      href="#!"
    >
      <v-img
        :lazy-src="require('@/assets/logo.png')"
        :src="`${apiUrl}/photos/${value.hero}`"
        height="100%"
        gradient="rgba(0, 0, 0, 0), rgba(0, 0, 0, 0)"
      >
        <v-layout
          v-if="!value.prominent"
          fill-height
          wrap
          text-xs-right
          ma-0
        >
          <v-flex xs12>
            <h3 class="title font-weight-bold mb-2">
              {{ value.title }}
            </h3>
            <div class="caption">
              {{ value.author }}
            </div>
          </v-flex>
        </v-layout>
      </v-img>
    </base-card>
  </v-flex>
</template>

<script>
  import {
    mapState
  } from 'vuex'

  export default {
    props: {
      size: {
        type: Number,
        required: true
      },
      value: {
        type: Object,
        default: () => ({})
      }
    },

    computed: {
      ...mapState(['config']),

      classes () {
        return {
          'md6': this.size === 2,
          'md4': this.size === 3
        }
      },

      apiUrl () {
        return this.config.api_url
      }
    }
  }
</script>

<style>
.v-image__image {
  transition: .3s linear;
}
</style>
