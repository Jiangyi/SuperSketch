import Vue from 'vue';
import Router from 'vue-router';
import Index from '@/components/index/Index';
import Play from '@/components/play/Play';
import Register from '@/components/auth/Register';
import Login from '@/components/auth/Login';

Vue.use(Router);

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Index',
      component: Index,
    },
    {
      path: '/play',
      name: 'Play',
      component: Play,
    },
    {
      path: '/register',
      name: 'Register',
      component: Register,
    },
    {
      path: '/login',
      name: 'Login',
      component: Login,
    },
  ],
});
