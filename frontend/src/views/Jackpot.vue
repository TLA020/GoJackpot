<template>
  <v-container fluid>
    <v-col>
      <GameStatsBar></GameStatsBar>
    </v-col>
    <v-col>
      <Game></Game>
    </v-col>
    <v-col> <Bet></Bet></v-col>
    <v-btn color="primary" @click="testChat">test</v-btn>
  </v-container>
</template>

<script>
import GameStatsBar from "@/components/GameStatsBar";
import Game from "@/components/Game";
import Bet from "@/components/Bet";
import Vue from "vue";

export default {
  name: "jackpot",
  components: { GameStatsBar, Game, Bet },
  data() {
    return {
      bet: {
        amount: 0
      }
    };
  },

  methods: {
    testChat() {
      Vue.$socket.sendObj({ type: "chat-message", data: { message: "duitsland" } });
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
