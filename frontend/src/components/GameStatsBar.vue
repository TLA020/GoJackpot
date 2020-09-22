<template>
  <v-row>
    <v-col cols="4">
      <v-card shaped class="darken-1">
        <p class="text-center pa-4">
          <v-icon style="font-size: 64px">mdi-av-timer</v-icon>
          %{{ share.toFixed(2) }}
        </p>
      </v-card>
    </v-col>

    <v-col cols="4">
      <v-card shaped class="darken-1">
        <p class="text-center pa-4">
          <v-icon style="font-size: 64px">mdi-cash-usd-outline </v-icon>
          ${{ totalBet.toFixed(2) }}
        </p>
      </v-card>
    </v-col>

    <v-col cols="4">
      <v-card shaped class="darken-1">
        <p class="text-center pa-4">
          <v-icon style="font-size: 64px" class="mr-2"
            >mdi-gamepad-circle-up
          </v-icon>
          {{ players }} players
        </p>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: "GameStatsBar",
  data: () => ({}),
  computed: {
    userBet() {
      if (this.$store.state.$game.game.userBets) {
        return this.$store.state.$game.game.userBets.find(
          ub => ub.player.id === this.$store.state.$auth.user.ID
        );
      } else {
        return {};
      }
    },

    share() {
      if (this.userBet && this.userBet.share) {
        return this.userBet.share;
      }
      return 0;
    },

    totalBet() {
      if (!this.userBet || !this.userBet.bets) {
        return 0;
      }
      return this.userBet.bets
        .map(o => o.amount)
        .reduce((a, c) => {
          return a + c;
        });
    },

    players() {
      if (!this.$store.state.$game.game.userBets) {return 0;}
      return this.$store.state.$game.game.userBets.length;
    }
  }
};
</script>
