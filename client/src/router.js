import {createRouter, createWebHistory} from "vue-router";
import NotFound from "./pages/NotFound.vue";
import UsersList from "@/pages/users/UsersList.vue";
import UserInfo from "@/pages/users/UserInfo.vue";
import RolesList from "@/pages/roles/RolesList.vue";
import RoleInfo from "@/pages/roles/RoleInfo.vue";
import UserRoles from "@/pages/user-roles/UserRoles.vue";
import LoginPage from "@/pages/auth/LoginPage.vue";
import RegisterPage from "@/pages/auth/RegisterPage.vue";
import UserProfile from "@/pages/users/UserProfile.vue";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {path: '/', redirect: '/users'},
    {path: '/users', component: UsersList},
    {path: '/users/:id', component: UserInfo, props: true},
    {path: '/profile', component: UserProfile, props: true},
    {path: '/roles', component: RolesList},
    {path: '/roles/:id', component: RoleInfo, props: true},
    {path: '/user-roles', component: UserRoles},
    {path: '/login', component: LoginPage},
    {path: '/register', component: RegisterPage},
    {path: "/:notFound(.*)", component: NotFound},
  ],
});

export default router;
