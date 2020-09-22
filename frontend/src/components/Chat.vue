<template>
  <div>
    <div>
      <v-toolbar color="grey darken-4" dark>
        <v-icon class="ml-2">mdi-account-group</v-icon> 12 online
      </v-toolbar>

      <v-list>
        <template v-for="(item, index) in items">
          <v-list-item :key="index">
            <v-list-item-avatar>
              <v-img :src="item.avatar"></v-img>
            </v-list-item-avatar>
            <v-list-item-content>
              <v-list-item-title v-html="item.userName"></v-list-item-title>
              <p class="text--secondary subtitle-2">{{ item.message }}</p>
            </v-list-item-content>
          </v-list-item>
          <v-divider :key="index + 10" :inset="item.inset"></v-divider>
        </template>
      </v-list>
    </div>

    <v-text-field placeholder="..." class="ml-1 mr-1"></v-text-field>
  </div>
</template>
<script>
export default {
  name: "Chat",
  data: () => ({
    items: [
      {
        avatar: "https://api.adorable.io/avatars/64/",
        userName: "Darkangel_1337",
        message: "scared as fuck huh?"
      },
      {
        avatar:
          "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/48/4830c1e5abc9cc8d8b02ec9679de7346e9514ef4_medium.jpg",
        userName: "TTV/xBrenni",
        message: "still wait since 1 year for the dc call pog"
      },
      {
        avatar: "https://api.adorable.io/avatars/64/5",
        userName: "LIVE|Pret0xX",
        message: "gogogo start bets"
      },
      {
        avatar:
          "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/48/4830c1e5abc9cc8d8b02ec9679de7346e9514ef4_medium.jpg",
        userName: "TTV/xBrenni",
        message: "lol wtf go all in then pussy"
      }
    ]
  })
};
</script>
