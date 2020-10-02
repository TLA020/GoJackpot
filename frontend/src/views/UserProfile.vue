<template>
  <v-container fill-height fluid>
    <v-card class="elevation-12" style="width: 100%">
      <v-toolbar light color="primary">
        <v-toolbar-title>Edit profile</v-toolbar-title>
      </v-toolbar>
      <v-card-text>
        <v-form ref="form">
          <v-text-field
            prepend-icon="mdi-account"
            name="email"
            label="Email"
            type="email"
            v-model="user.email"
            disabled
          >
          </v-text-field>
          <v-text-field
            prepend-icon="mdi-account"
            name="username"
            label="Username"
            disabled
            type="text"
            v-model="user.username"
          >
          </v-text-field>
          <v-row class="ml-0">
            <v-file-input
              v-model="avatar"
              show-size
              truncate-length="15"
              label="Avatar"
            ></v-file-input>
            <v-avatar color="primary" size="56"
              ><img v-if="user.avatar" alt="Avatar" :src="userAvatar"/>
              <v-icon v-else color="green" v-text="'?'"></v-icon
            ></v-avatar>
          </v-row>
        </v-form>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="primary" :loading="loading" @click="submit"
          >Opslaan</v-btn
        >
      </v-card-actions>
    </v-card>
    <div></div>
  </v-container>
</template>

<script>
import { mapState } from "vuex";
import { mapGetters } from "vuex";
import axios from "axios";
export default {
  name: "UserProfile",
  computed: {
    ...mapState("$auth", ["error", "user"]),
    ...mapGetters("$auth", ["userAvatar"])
  },

  data() {
    return {
      avatar: [],
      loading: false
    };
  },
  methods: {
    submit() {
      this.uploadAvatar();
    },

    uploadAvatar() {
      this.loading = true;
      const API_URL =
        (process.env.VUE_APP_API_BASE_URL || "/api/v1") + "/accounts";
      let formData = new FormData();

      formData.append("avatar", this.avatar);

      axios
        .post(`${API_URL}/avatar`, formData, {
          headers: {
            "Content-Type": "multipart/form-data",
            Authorization: "Bearer " + this.user.token
          }
        })
        .then(res => {
          this.$store.commit("$auth/SET_USER", res.data);
          this.loading = false;
        })
        .catch(e => {
          console.log(e);
        });
    }
  }
};
</script>
