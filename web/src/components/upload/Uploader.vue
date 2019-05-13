<template>
  <v-layout row justify-center>
    <v-dialog v-model="dialog" persistent max-width="600px">
      <template v-slot:activator="{ on }">
        <v-btn color="light-blue lighten-1" dark v-on="on">Upload Photo</v-btn>
      </template>
      <v-card>
        <v-card-title>
          <span class="headline">Upload Photo</span>
        </v-card-title>
        <v-card-text>
          <v-container grid-list-md>
            <v-layout wrap>
              <v-flex xs12 sm12 md12>
                <input id="fileChooser" class="ml-0" type="file" ref="uploadFile" accept="image/*" @change="setPhoto($event)" hidden>
                <div class="text-sm-center">
                  <v-btn flat icon color="light-blue lighten-1" dark @click="chooseFile()">
                    <v-icon x-large>mdi-camera</v-icon>
                  </v-btn>
                </div>
              </v-flex>
              <v-flex xs12 sm12 md12>
                <v-text-field label="Photo Title" v-model="title" :rules="defaultRule" required></v-text-field>
              </v-flex>
              <v-flex xs12 sm12 md12>
                <v-text-field label="Author" hint="name of the photographer" v-model="author" :rules="defaultRule" required></v-text-field>
              </v-flex>
            </v-layout>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue darken-1" flat @click="dialog = false">Cancel</v-btn>
          <v-btn color="blue darken-1" flat @click="upload()">Upload</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-layout>
</template>

<script>
  import * as axios from 'axios';
  import {
    mapState
  } from 'vuex'

  export default {
    data: () => ({
      dialog: false,
      author: "",
      title: "",
      photo: "",
      defaultRule: [
        v => !!v || 'Required field'
      ]
    }),
    computed: {
      ...mapState(['config'])
    },
    methods: {

      setPhoto(event) {
        this.photo = event.target.files[0]
      },
      chooseFile() {
        document.getElementById("fileChooser").click()
      },
      upload() {
        if (validateForm(this)) {
          this.dialog = false;
          var formData = new FormData();

          formData.append("photo", this.photo)
          formData.append("title", this.title)
          formData.append("author", this.author)

          axios.post(this.config.api_url + "/photos", formData)
          clearForm(this);
        }
      }
    }
  }

  function clearForm(that) {
    that.author = "";
    that.title = "";
  }

  function validateForm(that) {
    if (that.author == '' || that.title == '' || that.photo == '') {
      return false
    }
    return true
  }
</script>
