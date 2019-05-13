import Vue from 'vue'
import Router from 'vue-router'
import Pheed from './views/Pheed.vue'
import Info from './views/Info.vue'
import Team from './views/Team.vue'
import Talks from './views/Talks.vue'
import Upload from './views/Upload.vue'

Vue.use(Router)

export default new Router({
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Info
    },
    {
      path: '/pheed',
      name: 'pheed',
      component: Pheed
    },
    {
      path: '/talks',
      name: 'talks',
      component: Talks
    },
    {
      path: '/team',
      name: 'team',
      component: Team
    },
    {
      path: '/pheedupload',
      name: 'upload',
      component: Upload
    }
  ]
})
