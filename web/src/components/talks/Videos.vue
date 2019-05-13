<template>
  <v-container fluid grid-list-lg>
    <v-layout wrap>
      <v-flex justify-center layout>
        <base-heading class="info--text">
          Videos
        </base-heading>
      </v-flex>
      <v-flex justify-center layout md12 xs12>
        <v-divider light></v-divider>
      </v-flex>
    </v-layout>
    <v-layout wrap>
      <v-flex v-for="(vid, i) in videos" justify-center align-center layout>
        <v-card width="800">
          <iframe width="100%" height="315" :src="embedUrl(vid.id)" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
          <v-card-title primary-title>
            <div>
              <div class="headline">{{ vid.title }}</div>
              <span class="grey--text">{{ vid.speaker }}</span>
            </div>
          </v-card-title>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
export default {
  data: () => ({
    videos: [{
      title: 'CF Summit NA - Keynote by Julz Friedman',
      speaker: 'Julz Friedman',
      id: 'XZgW2E6Xw-s?start=1726',
    }, {
      title: '"Kube" Your Enthusiasm: Bringing Cloud Foundry and Kubernetes Together with Project Eirini',
      speaker: 'Julz Friedman (IBM) & Mario Nitchev (SAP)',
      id: 'b9d8Xiapzr0',
    }, {
      title: 'Diego vs Eirini: How to choose the best fitting container scheduler',
      speaker: 'Herr Julz Skupnjak (IBM) & Georgi Dankov (SAP) ',
      id: 'mcoSFWm_U6A&t',
    }, {
      title: 'Microservices With Cloud Foundry and Kubernetes',
      speaker: 'Herr Julz Skupnjak (IBM) & Georgi Dankov (SAP) ',
      id: 'B2AII9RkPRk',
    }, {
      title: 'Kube Your Enthusiasm — Bringing the `CF Push` Experience to Kubernetes Operators (Project Eirini)',
      speaker: 'Doctor Julz Friedman & Herr Julz Skupnjak, IBM',
      id: 'dvLCT19I4QY',
    }, {
      title: 'CF^3 - Putting a Kubernetes Behind CF',
      speaker: 'Doctor Julz Friedman, Andrew Edgar & Herr Julz Skupnjak, IBM',
      id: '9l3GgW95GmQ',
    }]
  }),

  methods: {
    embedUrl: function(id) {
      return "https://www.youtube.com/embed/" + id
    }
  }
}
</script>
