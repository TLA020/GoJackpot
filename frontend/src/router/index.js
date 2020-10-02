import Vue from "vue";
import Router from "vue-router";
import Auth from "../views/Auth";
import UserProfile from "../views/UserProfile";
import Jackpot from "../views/Jackpot";
import store from "../store";

Vue.use(Router);

const router = new Router({
  routes: [
    {
      path: "/auth",
      name: "Auth",
      component: Auth,
      meta: {
        authRequired: false
      }
    },
    {
      path: "/",
      name: "jackpot",
      component: Jackpot,
      meta: {
        authRequired: true
      }
    },
    {
      path: "/account/edit",
      name: "userProfile",
      component: UserProfile,
      meta: {
        authRequired: true
      }
    }
  ]
});

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.authRequired)) {
    if (!store.state.$auth.user) {
      let user = JSON.parse(window.localStorage.getItem("user"));
      if (user) {
        store.commit("$auth/SET_USER", user);
        return next();
      }
      next({ path: "/auth" });
    } else {
      next();
    }
  } else {
    next();
  }
});

export default router;
