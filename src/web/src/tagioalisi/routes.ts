import React from 'react';

interface Route {
    title: string,
    path: string,
    isHome?: boolean,
    component: React.ExoticComponent, 
}

export const ROUTES: {[id: string]: Route} = {
    home: {
        title: 'Home', 
        path: '/', 
        isHome: true,
        component: React.lazy(() => import('tagioalisi/components/Homepage')),
    },
    config: {
        title: 'Configuration', 
        path: '/config', 
        component: React.lazy(() => import('tagioalisi/components/Configuration')),
    },
};

export const getRoute = (id: string) => {
    if (!!ROUTES[id]) {
        return ROUTES[id];
    }
    console.error(`No route with id: ${id}`);
    return {title: "INVALID ROUTE", path:"INVALID", component: () => null}
}