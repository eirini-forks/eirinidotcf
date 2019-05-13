<template>
  <section id="welcome" class="hide-overflow">
    <v-layout>
      <v-flex hidden-sm-and-down md6>
        <v-img :src="require('@/assets/welcome.png')" height="100vh" />
      </v-flex>
      <v-flex xs12 md6 align-content-space-between layout :pa-5="$vuetify.breakpoint.smAndDown" wrap>
        <base-bubble-1 style="transform: rotate(180deg) translateY(25%)" />
        <v-layout align-center justify-center>
          <v-flex xs10 md6>
            <base-heading>Welcome!</base-heading>
            <base-text>
              Eirini is a Kubernetes backend for Cloud Foundry. It deploys CF apps to a kube backend, using OCI images and Kube deployments.
              <br> Eirini gives you the nice integrated cf push flow, with CF Apps mapped directly to kube StatefulSet.
            </base-text>
          </v-flex>
        </v-layout>
        <base-bubble-2 style="transform: rotate(180deg) translate(-200px, -15%)" />
      </v-flex>
    </v-layout>
  </section>
</template>