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
  export default {
    data: () => ({
      avatars: [{
        name: 'Julian Friedman (Julz)',
        role: 'Product Manager',
        company: 'IBM',
        email: 'julz.friedman@uk.ibm.com',
        city: 'London, UK',
        image: 'https://avatars.githubusercontent.com/julz',
        twitter: 'doctor_julz',
        git: 'julz'
      }, {
        name: 'Steffen Uhlig',
        role: 'Anchor',
        company: 'IBM',
        email: 'steffen.uhlig@de.ibm.com',
        city: 'Boeblingen, Germany',
        image: 'https://avatars.githubusercontent.com/suhlig',
        git: 'suhlig',
        twitter: 'suhlig'
      }, {
        name: 'Julian Skupnjak (HerrJulz)',
        role: 'Eiringineer',
        company: 'IBM',
        email: 'skupnjak@de.ibm.com',
        city: 'Boeblingen, Germany',
        image: 'https://avatars.githubusercontent.com/JulzDiverse',
        git: 'JulzDiverse',
        twitter: 'herrjulz'
      }, {
        name: 'Mario Nitchev',
        role: 'Eiringineer',
        company: 'SAP',
        email: 'mario.nitchev@sap.com',
        city: 'Sofia, Bulgaria',
        twitter: 'mnitchev',
        git: 'mnitchev',
        image: 'https://avatars.githubusercontent.com/mnitchev'
      }, {
        name: 'Georgi Dankov',
        role: 'Eiringineer',
        company: 'SAP',
        twitter: 'gdankov',
        email: 'georgi.dankov@sap.com',
        city: 'Sofia, Bulgaria',
        git: 'gdankov',
        image: 'https://avatars.githubusercontent.com/gdankov'
      }, {
        name: 'Oleksandr Slynko',
        role: 'Eiringineer',
        company: 'Pivotal',
        email: 'oslynko@pivotal.io',
        twitter: 'alex_sly',
        city: 'Dublin, Ireland',
        git: 'alex-slynko',
        image: 'https://avatars.githubusercontent.com/alex-slynko'
      }, {
        name: 'Akshay Mankar',
        role: 'Eiringineer',
        company: 'Pivotal',
        email: 'amankar@pivotal.io',
        twitter: 'AkshayMankar',
        city: 'Dublin, Ireland',
        git: 'akshaymankar',
        image: 'https://avatars.githubusercontent.com/akshaymankar'
      }, {
        name: 'Matthew Sweeps',
        role: 'Eiringineer',
        company: 'Google',
        email: 'mattysweeps@google.com',
        city: 'Seattle, USA',
        twitter: '',
        git: 'mattysweeps',
        image: 'https://avatars.githubusercontent.com/mattysweeps'
      }, {
        name: 'Nima Kaviani',
        role: 'Eiringineer',
        company: 'IBM',
        email: 'nima.kaviani@us.ibm.com',
        city: 'San Francisco, USA',
        twitter: 'nimak',
        git: 'nimakaviani',
        image: 'https://avatars.githubusercontent.com/nimakaviani'
      }]
    }),

    methods: {
      openPage: function (link) {
        window.open(link, '_blank')
      }
    }
  }
</script>
