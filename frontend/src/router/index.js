import Vue from "vue";
import Router from "vue-router";
import Home from "../views/Home";
import Auth from "../views/Auth";
import store from '../store/store'

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
  ]
});

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.authRequired)) {
    if (!store.state.auth.user) {
      console.log('push auth')
      next({ path: "/auth" });
    } else {
      next()
    }
  } else {
    next();
  }
});

export default router;