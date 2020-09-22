<template>
  <v-app id="inspire">
    <v-navigation-drawer v-model="drawer" clipped app>
      <Chat></Chat>
    </v-navigation-drawer>
    <v-app-bar
      :clipped-left="$vuetify.breakpoint.lgAndUp"
      app
      dark
    >
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title style="width: 300px" class="ml-0 pl-4">
        <span class="hidden-sm-and-down">GoJackPot</span>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <span v-if="user"> {{ user.email }} </span>
      <v-btn icon>
        <v-icon>mdi-apps</v-icon>
      </v-btn>
      <v-btn icon>
        <v-icon>mdi-bell</v-icon>
      </v-btn>
      <v-btn icon large>
        <!--        <v-avatar-->
        <!--                size="32px"-->
        <!--                item-->
        <!--        >-->
        <!--          <v-img-->
        <!--                  src=""-->
        <!--                  alt=""-->
        <!--          ></v-img></v-avatar>-->
      </v-btn>
    </v-app-bar>
    <v-main>
      <v-container fluid>
        <router-view></router-view>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import Chat from "@/components/Chat.vue";
export default {
  components: { Chat },
  data: () => ({
    dialog: false,
    drawer: null,
    items: [
      { icon: "mdi-home", text: "Home", url: "/" },
      { icon: "mdi-cash-usd", text: "Jackpot", url: "/jackpot" },
      { icon: "mdi-cog", text: "Settings" },
      { icon: "mdi-help-circle", text: "Help" }
    ],
    source: {}
  }),
  created() {
    this.$vuetify.theme.dark = true;
  },
  computed: {
    user() {
      return this.$store.state.$auth.user;
    }
  }
};
</script>
