<template>
  <v-app id="inspire">
    <v-navigation-drawer v-model="drawer" clipped app>
      <v-row class="fill-height" no-gutters>
        <v-col cols="3">
          <v-navigation-drawer darken-4 v-model="drawer" clipped>
            <v-list dense nav>
              <v-list-item
                v-for="item in items"
                :key="item.title"
                :to="item.url"
              >
                <v-list-item-action>
                  <v-icon class="ml-1">{{ item.icon }}</v-icon>
                </v-list-item-action>

                <v-list-item-title>{{ item.title }}</v-list-item-title>
              </v-list-item>
            </v-list>
          </v-navigation-drawer>
        </v-col>
        <v-col>
          <Chat></Chat>
        </v-col>
      </v-row>
    </v-navigation-drawer>
    <v-app-bar :clipped-left="$vuetify.breakpoint.lgAndUp" app dark>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title style="width: 300px" class="ml-0 pl-4">
        <span class="hidden-sm-and-down">GoJackPot</span>
      </v-toolbar-title>

      <v-spacer></v-spacer>
      <span v-if="user">
       <v-avatar color="primary" size="56"
       ><img v-if="avatar" alt="Avatar" :src="avatar"/>
          <v-icon v-else color="green" v-text="'?'"></v-icon
          ></v-avatar>
              {{user.username}}
      </span>
      <v-btn icon>
        <v-icon>mdi-bell</v-icon>
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
      //{ icon: "mdi-home", text: "Home", url: "/" },
      { icon: "mdi-cash-usd", text: "Jackpot", url: "/" },
      { icon: "mdi-cog", text: "auth", url: "/account/edit" },
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
    },

    avatar() {
      return this.$store.getters["$auth/userAvatar"];
    }
  }
};
</script>
