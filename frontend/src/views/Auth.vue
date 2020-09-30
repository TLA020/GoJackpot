<template>
  <v-container fill-height>
    <v-layout align-center justify-center>
      <v-flex xs12 sm8 md4>
        <v-card class="elevation-12">
          <v-toolbar light color="primary">
            <v-toolbar-title>Login / Register</v-toolbar-title>
          </v-toolbar>
          <v-card-text>
            <v-form ref="form" v-model="valid">
              <v-text-field
                prepend-icon="mdi-account"
                name="email"
                label="Email"
                type="email"
                v-model="email"
                :rules="emailRules"
                required
                data-cy="signinEmailField"
              >
              </v-text-field>
              <v-text-field
                prepend-icon="mdi-account"
                name="username"
                label="Username"
                type="text"
                v-model="username"
                :rules="usernameRules"
                required
              >
              </v-text-field>
              <v-text-field
                prepend-icon="mdi-lock-question"
                name="password"
                label="Password"
                type="password"
                data-cy="signinPasswordField"
                v-model="password"
                :rules="passwordRules"
                required
              >
              </v-text-field>
            </v-form>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
              color="primary"
              :disabled="!valid"
              :loading="loading"
              @click="submit('register')"
              data-cy="signinSubmitBtn"
              >Register</v-btn
            >
            <v-btn
              color="secondary"
              :disabled="!valid"
              :loading="loading"
              @click="submit('login')"
              data-cy="signinSubmitBtn"
              >Login</v-btn
            >
          </v-card-actions>
        </v-card>
        <div>
          <v-alert v-model="error" v-if="error" dismissible type="error">
            {{ error }}
          </v-alert>
        </div>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import { mapState } from "vuex";

export default {
  name: "Signin",
  computed: {
    ...mapState("$auth", ["error", "loading"])
  },

  data() {
    return {
      valid: false,
      email: "",
      password: "",
      username: "",
      emailRules: [
        v => !!v || "E-mail is required",
        v => /.+@.+/.test(v) || "E-mail must be valid"
      ],
      passwordRules: [
        v => !!v || "Password is required",
        v => v.length >= 6 || "Password must be greater than 6 characters"
      ],
      usernameRules: [
        v => !!v || "Username",
        v => v.length >= 4 || "Username must be greater than 6 characters"
      ]
    };
  },
  methods: {
    submit(action) {
      if (this.$refs.form.validate()) {
        this.$store.dispatch(`$auth/${action}`, {
          email: this.email,
          password: this.password,
          username: this.username
        });
      }
    }
  }
};
</script>
