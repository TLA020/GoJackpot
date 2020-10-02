<template>
  <div >
    <div>
      <v-toolbar color="grey darken-4" dark>
        <v-icon class="ml-2 mr-1">mdi-account-group</v-icon> Chat
      </v-toolbar>
      <v-list>
        <template v-for="(item, index) in items">
          <v-list-item :key="index">
            <v-list-item-avatar v-show="item.avatar">
              <v-img :src="item.avatar"></v-img>
            </v-list-item-avatar>
            <v-list-item-content>
              <v-list-item-title v-html="item.userName"></v-list-item-title>
              <p class="text--secondary subtitle-2">{{ item.msg }}</p>
            </v-list-item-content>
          </v-list-item>
          <v-divider :key="index + 10" :inset="item.inset"></v-divider>
        </template>
      </v-list>
    </div>

    <v-footer fixed>
      <v-row no-gutters>
        <v-col cols="8">
          <v-text-field
            placeholder="enter message"
            class="ml-1 mr-4"
            @keyup.enter="sendChat"
            v-model="message"
          ></v-text-field>
        </v-col>
        <v-col cols="4">
          <v-btn class="mt-4" @click="sendChat"><v-icon>mdi-send</v-icon></v-btn>
        </v-col>
      </v-row>
    </v-footer>
  </div>
</template>
<script>
import Vue from "vue";

export default {
  name: "Chat",
  mounted() {
    this.$socketEvent("chat-message", res => {
      this.items.push(res);
    });

    this.$socketEvent("chat-snapshot", res => {
      this.items = res.messages;
    });
  },

  data: () => ({
    items: [],
    message: ""
  }),

  methods: {
    sendChat() {
      Vue.$socket.sendObj({
        type: "chat-message",
        data: { message: this.message }
      });
      this.message = "";
    }
  }
};
</script>
