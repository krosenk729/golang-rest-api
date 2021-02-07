import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router';
import Dashboard from '@/views/Dashboard.vue';

import Calendar from '@/components/Calendar.vue';
import Review from '@/components/Review.vue';
import Track from '@/components/Track.vue';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Dashboard',
    component: Dashboard,
    children: [
      { path: 'calendar', name: 'Calendar', component: Calendar },
      { path: 'review', name: 'Review', component: Review },
      { path: 'track', name: 'Track', component: Track },
    ],
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
