<template>
  <base-card dark>
    <v-img
      class="grey lighten-2"
      height="400"
      width="100%"
      gradient="rgba(0, 0, 0, .42), rgba(0, 0, 0, .42)"
      :src="require('@/assets/eiriniteam.png')"
    >
      <v-layout
        fill-height
        align-center
        pa-3
      >
        <v-flex
          xs12
          md7
          offset-md5
        >
          <h1 class="display-3 font-weight-light">
            Meet The Team
          </h1>
        </v-flex>
      </v-layout>
    </v-img>
  </base-card>
</template>
