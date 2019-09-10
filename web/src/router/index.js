import Vue from 'vue'
import Router from 'vue-router'
import i18n from '@/lang'
import priv from '@/lib/priv'

Vue.use(Router)

const _import = file => () => import('../view/' + file + '.vue')

const appMap = [
    {
        path: '/login',
        name: 'login',
        component: _import('Login'),
    },
]

const routerMap = [
    {
        path: '/',
        component: _import('Layer'),
        name: 'main',
        meta: {},
        redirect: { name: 'dashboard' },
        children: [
            {
                path: 'dashboard',
                name: 'dashboard',
                meta: {
                    title: i18n.t('dashboard'),
                    icon: 'icon-dashboard',
                    single: true,
                },
                component: _import('Dashboard'),
            },
        ],
    },
    {
        path: '/user',
        name: 'user',
        component: _import('Layer'),
        meta: {
            title: i18n.t('user'),
            icon: 'icon-group',
        },
        children: [
            {
                path: 'group',
                name: 'userGroup',
                meta: {
                    title: i18n.t('role_manage'),
                    role: [priv.USER_ROLE_VIEW],
                },
                component: _import('user/Group'),
            },
            {
                path: 'list',
                name: 'userList',
                meta: {
                    title: i18n.t('user_manage'),
                    role: [priv.USER_VIEW],
                },
                component: _import('user/User'),
            },
        ],
    },
]

const router = new Router({
    routes: appMap.concat(routerMap),
    scrollBehavior: () => ({ y: 0 }),
    mode: 'history',
})

export { routerMap }

export default router