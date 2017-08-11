// import Vue from 'vue'
// import App from './App.vue'
//
// new Vue({
//   el: '#app',
//   render: h => h(App)
// })
import Vue from 'vue'

// Ленивая загрузка
// const Hairstyles = resolve => require(['./components/Hairstyles.vue'], resolve);

import Header from './components/index/Header.vue'
import Footer from './components/index/Footer.vue'
import Hairstyles from './components/Hairstyles.vue'
import Hairstyle from './components/Hairstyle.vue'
import Haircuts from './components/Haircuts.vue'
import Haircut from './components/Haircut.vue'
import Diplomas from './components/Diplomas.vue'
import Diploma from './components/Diploma.vue'
import Contacts from './components/Contacts.vue'



// vue-resource
import VueResource from 'vue-resource'
Vue.use(VueResource);

// vue-router
import VueRouter from 'vue-router'
Vue.use(VueRouter);


const scrollBehavior = (to, from, savedPosition) => {
  if (savedPosition) {
    // savedPosition is only available for popstate navigations.
    return savedPosition
  } else {
    const position = {};
    // new navigation.
    // scroll to anchor by returning the selector
    if (to.hash) {
      position.selector = to.hash
    }
    // check if any matched route config has meta that requires scrolling to top
    if (to.matched.some(m => m.meta.scrollToTop)) {
      // cords will be used if no selector is provided,
      // or if the selector didn't match any element.
      position.x = 0;
      position.y = 0
    }
    // if the returned position is falsy or an empty object,
    // will retain current scroll position.
    return position
  }
};


const router = new VueRouter({
  mode: 'history',
  base: '/',
  scrollBehavior,
  routes: [
    // { path: '/hairstyles', name: 'Hairstyles', component: Hairstyles, meta: { scrollToTop: true } },
    { path: '/hairstyles', name: 'Hairstyles', component: Hairstyles},
    { path: '/haircuts', name: 'Haircuts', component: Haircuts},
    { path: '/diplomas', name: 'Diplomas', component: Diplomas},
    { path: '/contacts', name: 'Contacts', component: Contacts},
    { path: '/hairstyles/:id', name: 'Hairstyle', component: Hairstyle},
    { path: '/haircuts/:id', name: 'Haircut', component: Haircut},
    { path: '/diplomas/:id', name: 'Diploma', component: Diploma},
    { path: '*', redirect: { name: 'Hairstyles' }},
  ],
});

new Vue({
  router: router,
  components: {
    IndexHeader: Header,
    IndexFooter: Footer,
  }
}).$mount('#app');

