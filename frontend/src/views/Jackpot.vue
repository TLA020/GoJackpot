<template>
  <v-container fluid>
    <v-row>
      <v-col cols="12" lg="8">
        <v-card class="mx-auto">
          <v-toolbar color="blue darken-3" dark>
            <v-toolbar-title>Game info</v-toolbar-title>
            <v-spacer></v-spacer>
          </v-toolbar>

          <div class="bet-controls pa-8">
            <Game></Game>
          </div>
        </v-card>
      </v-col>

      <v-col cols="12" lg="4">
        <v-card class="mx-auto">
          <v-toolbar color="blue darken-3" dark>
            <v-toolbar-title>Users online</v-toolbar-title>
            <v-spacer></v-spacer>
          </v-toolbar>
          <UsersOnline></UsersOnline>
          <v-divider></v-divider>
        </v-card>
      </v-col>
    </v-row>

    <v-row>
      <v-col cols="12" lg="8">
        <v-card class="mx-auto">
          <v-toolbar color="blue darken-3" dark>
            <v-toolbar-title>Play</v-toolbar-title>
            <v-spacer></v-spacer>
          </v-toolbar>

          <div class="bet-controls pa-8">
            <Bet v-model="bet.amount"></Bet>
          </div>
        </v-card>
      </v-col>

      <v-col cols="12" lg="4">
        <v-card class="mx-auto">
          <v-toolbar color="blue darken-3" dark>
            <v-toolbar-title>Game history</v-toolbar-title>
            <v-spacer></v-spacer>
          </v-toolbar>

          <div>
          <LastGames></LastGames>
          </div>
        </v-card>
      </v-col>

    </v-row>

  </v-container>
</template>

<script>
import Game from "../components/Game.vue"
import LastGames from "../components/LastGames.vue"
import Bet from "../components/Bet.vue";
import UsersOnline from "../components/UsersOnline";
export default {
  name: "jackpot",
  components: { Bet, UsersOnline, Game, LastGames },
  data() {
    return {
      bet: {
        amount: 0
      }
    };
  },

  methods: {},

  computed: {
    currentUsers() {
      return this.$store.state.$game.currentUsers;
    }
  }
};
</script>

<style scoped></style>
