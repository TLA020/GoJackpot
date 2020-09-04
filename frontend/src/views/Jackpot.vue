<template>
  <v-container fluid>
    <v-row>
      <v-col cols="2">
        <v-card max-width="500" class="mx-auto">
          <v-toolbar color="blue darken-3" dark>
            <v-toolbar-title class="float-left">Users online</v-toolbar-title>
            <v-spacer></v-spacer>
          </v-toolbar>
          <UsersOnline></UsersOnline>
          <v-divider></v-divider>
        </v-card>
      </v-col>

      <v-col cols="6">
        <v-card class="mx-auto">
          <v-toolbar color="blue darken-3" dark>
            <v-toolbar-title class="float-left">Game</v-toolbar-title>
            <v-spacer></v-spacer>
          </v-toolbar>

          <div class="bet-controls pa-8">
            <Bet v-model="bet.amount"></Bet>
          </div>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import Bet from "../components/Bet.vue";
import UsersOnline from "../components/UsersOnline";
export default {
  name: "jackpot",
  components: { Bet, UsersOnline },
  data() {
    return {
      bet: {
        amount: 0
      }
    };
  },

  methods: {

  },

  computed: {
    currentUsers() {
      return this.$store.state.$game.currentUsers;
    }
  }
};
</script>

<style scoped></style>
