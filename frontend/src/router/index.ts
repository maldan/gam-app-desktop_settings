import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router';
import Main from '../page/Main.vue';
import Background from '../page/Background.vue';
import Backup from '../page/Backup.vue';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Main',
    component: Main,
  },
  {
    path: '/background',
    name: 'Background',
    component: Background,
  },
  {
    path: '/backup',
    name: 'Backup',
    component: Backup,
  },
];

const router = createRouter({
  history: createWebHashHistory(process.env.BASE_URL),
  routes,
});

export default router;
