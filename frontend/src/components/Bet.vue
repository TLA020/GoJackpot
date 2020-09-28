<template>
  <v-card shaped class="darken-1">
    <div class="pa-4">
      <v-card-title>Place a bet</v-card-title>
      <v-form ref="form" class="mt-4">
        <v-row>
          <v-slider
            v-model="amount"
            append-icon="mdi-currency-usd"
            prepend-icon="mdi-currency-usd-off"
            class="align-center col-12 p-4"
            :max="250"
            thumb-label="always"
            :min="1"
            hide-details
          >
          </v-slider>
        </v-row>
        <div class="text-right mr-4">
          <v-btn
            :disabled="false"
            color="secondary"
            class="mr-4"
            :loading="placingBet"
            @click="placeBet"
          >
            Place bet
          </v-btn>
          <v-btn color="red" @click="fakeBet">
            Fake bets
          </v-btn>
        </div>
      </v-form>
    </div>
  </v-card>
</template>

<script>
export default {
  name: "Bet",

  data() {
    return {
      placingBet: false,
      amount: 0
    };
  },

  methods: {
    placeBet() {
      this.placingBet = true;
      this.$store.dispatch("sendSocket", {
        type: "place-bet",
        data: { amount: this.amount }
      });
      this.placingBet = false;
      this.amount = 0;
    },

    fakeBet() {
      this.$store.dispatch("$game/fakeBets");
    }
  }
};
</script>

<style scoped></style>
