<template>
  <div>
    <h2>Jackpot</h2>
    <div class="users">
      <v-card max-width="500" class="mx-auto">
        <v-toolbar color="blue darken-3" dark>
          <v-toolbar-title class="float-left">Users online</v-toolbar-title>
          <v-spacer></v-spacer>
        </v-toolbar>
        <v-list subheader>
          <v-list-item v-for="user in currentUsers" :key="user.Email" class="mt-1">
            <v-list-item-avatar>
              <v-img :src="getAvatar(user.Email)"></v-img>
            </v-list-item-avatar>

            <v-list-item-content>
              <v-list-item-title v-text="user.Email"></v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list>
        <v-divider></v-divider>
      </v-card>
    </div>
  </div>
</template>

<script>
export default {
  name: "jackpot",
  methods: {
    placeBet() {
      this.$store.dispatch("sendSocket", {
        name: "place-bet",
        data: { amount: Math.floor(Math.random() * 50) + 1 }
      });
    },

    getAvatar(x) {
      return `https://api.adorable.io/avatars/${x}`;
    }
  },

  computed: {
    currentUsers() {
      return this.$store.state.$game.currentUsers;
    }
  }
};
</script>

<style scoped></style>
