import { createRouter, createWebHashHistory } from 'vue-router';
import GoelandObjectList from '@/components/GoelandObjectList.vue';
import GoelandObjectForm from '@/components/GoelandObjectForm.vue'



// Routes
const routes = [
  { path: '/', name : 'GoelandObjectList', component: GoelandObjectList },
  { path: '/:id', name: 'GoelandObjectForm', component: GoelandObjectForm}

];

// Create the router instance and pass the `routes` option
const router =  createRouter({
    // Provide the history implementation to use. We are using the hash history for simplicity here.
    history: createWebHashHistory(),
    routes, // short for `routes: routes`
  });
  

export default router; 