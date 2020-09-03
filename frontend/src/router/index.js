import Vue from "vue";
import Router from "vue-router";
import Home from "../views/Home";
import Auth from "../views/Auth";
import Jackpot from "../views/Jackpot";
import store from '../store'

Vue.use(Router);

const router = new Router({
  routes: [
    {
      path: "/",
      name: "Home",
      component: Home,
      meta: {
        authRequired: true
      }
    },
    {
      path: "/auth",
      name: "Auth",
      component: Auth,
      meta: {
        authRequired: false
      }
    },
    {
      path: "/jackpot",
      name: "jackpot",
      component: Jackpot,
      meta: {
        authRequired: true
      }
    },
  ]
});

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.authRequired)) {
    if (!store.state.$auth.user) {
      let user = JSON.parse(window.localStorage.getItem('user'));
      if (user) {
          store.commit("$auth/SET_USER", user)
          return next();
      }
      next({ path: "/auth" });
    } else {
      next()
    }
  } else {
    next();
  }
});

export default router;