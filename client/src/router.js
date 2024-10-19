import {createRouter, createWebHistory} from "vue-router";
import NotFound from "./pages/NotFound.vue";
import UsersList from "@/pages/users/UsersList.vue";
import UserInfo from "@/pages/users/UserInfo.vue";
import RolesList from "@/pages/roles/RolesList.vue";
import RoleInfo from "@/pages/roles/RoleInfo.vue";
import UserRoles from "@/pages/user-roles/UserRoles.vue";
import LoginPage from "@/pages/auth/LoginPage.vue";
import RegisterPage from "@/pages/auth/RegisterPage.vue";
import ActivateUser from "@/pages/auth/ActivateUser.vue";
import UserProfile from "@/pages/users/UserProfile.vue";
import store from "./store/index.js"

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {path: '/', redirect: '/users'},
    {path: '/users', component: UsersList},
    {path: '/users/:id', component: UserInfo, props: true},
    {path: '/profile', component: UserProfile, props: true, meta: {requiresAuth: true}},
    {path: '/roles', component: RolesList, meta: {requiresAdmin: true}},
    {path: '/roles/:id', component: RoleInfo, props: true, meta: {requiresAdmin: true}},
    {path: '/user-roles', component: UserRoles, meta: {requiresAdmin: true}},
    {path: '/login', component: LoginPage, meta: {requiresAuth: false}},
    {path: '/register', component: RegisterPage, meta: {requiresAuth: false}},
    {path: '/activate-user/:token', component: ActivateUser, props: true, meta: {requiresAuth: false}},
    {path: "/:notFound(.*)", component: NotFound},
  ],
});

router.beforeEach((to, from, next) => {
  const isAuthenticated = store.getters['auth/isAuthenticated'];
  const isAdmin = store.getters['auth/isAdmin'];
  if (to.meta.requiresAuth && !isAuthenticated) {
    next('/login');
  } else if (to.meta.requiresAdmin && !isAdmin) {
    next('/');
  } else {
    next();
  }
});

export default router;
