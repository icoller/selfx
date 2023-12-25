import layoutMain from '@/layout/main/index.vue'
import { IconNav,IconDashboard,IconLayers,IconFile,IconTag,IconRelation,IconComputer,IconSettings,IconStorage } from '@arco-design/web-vue/es/icon';

export default [
    {
        path: '/login',
        name: 'login',
        component: () => import('@/views/admin/login.vue')
    },
    {
        path: '/createAdmin',
        name: 'createAdmin',
        component: () => import('@/views/admin/createAdmin.vue')
    },
    {
        path: '/',
        name:"index",
        redirect: '/dashboard',
        component: layoutMain,
        children: [
            {
                path: '/dashboard',
                name: 'dashboard',
                meta: { icon:IconDashboard },
                component: () => import('@/views/dashboard/index.vue'),
            },
            {
                path: '/article',
                name: 'article',
                meta: {icon:IconFile},
                component: () => import('@/views/article/index.vue'),
            },
            {
                path: '/category',
                name: 'category',
                meta: {icon:IconNav},
                component:() => import('@/views/category/index.vue'),
            },
            {
                path: '/tag',
                name: 'tags',
                meta: {icon:IconTag},
                component: () => import('@/views/tag/index.vue'),
            },
            {
                path: '/link',
                name: 'link',
                meta: {icon:IconRelation},
                component: () => import('@/views/link/index.vue'),
            },
            {
                path: '/crawl',
                name: 'crawl',
                meta: {icon:IconStorage},
                component: () => import('@/views/crawl/index.vue'),
            },
            {
                path: '/log',
                name: 'log',
                redirect: { name: 'logs-item', params:{id:''} },
                meta: {icon:IconComputer},
                component: () => import('@/views/log/layout.vue'),
                children: [
                    {
                        path: '/log/:id?',
                        name: 'logs-item',
                        meta:{ lang:'log'},
                        component: () => import('@/views/log/index.vue'),
                    }
                ]
            },
            {
                path: '/plugin',
                name: 'plugin',
                meta: {icon:IconLayers},
                component: () => import('@/views/plugin/index.vue'),
            },
            {
                path: '/configure',
                name: 'config',
                redirect: { name: 'config-item', params:{id:'site'} },
                meta: {icon:IconSettings},
                component: () => import('@/views/config/layout.vue'),
                children: [
                    {
                        path: '/config/:id',
                        name: 'config-item',
                        meta:{ lang:'config'},
                        component: () => import('@/views/config/index.vue'),
                    }
                ]
            },
        ]
    }
]