import { Homepage } from './components/Homepage/Homepage';
import { Configuration } from './components/Configuration/Configuration';

interface Route {
    title: string,
    path: string,
    isHome?: boolean,
    component: () => JSX.Element, 
}

export const ROUTES: {[id: string]: Route} = {
    id: {
        title: 'Home', 
        path: '/', 
        isHome: true,
        component: Homepage,
    },
    config: {
        title: 'Configuration', 
        path: '/config', 
        component: Configuration,
    },
};

export const getRoute = (id: string) => {
    if (!!ROUTES[id]) {
        return ROUTES[id];
    }
    console.error(`No route with id: ${id}`)
    return {title: "INVALID ROUTE", path:"INVALID", component: () => null}
}