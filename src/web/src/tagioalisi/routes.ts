interface Route {
    title: string,
    navText: string,
    path: string,
    isHome?: boolean,
    loadComponent: () => Promise<() => JSX.Element>,  // use as: FooComponent = await route.loadComponent(); return <FooComponent />
}

export const ROUTES: Route[] = [
    {title: 'Home', navText: 'Home', path: '/', isHome: true, loadComponent: () => import('tagioalisi/components/Homepage').then(({Homepage}) => Homepage)},
    {title: 'Configuration', navText: 'Configuration', path: '/config', loadComponent: () => import('tagioalisi/components/Configuration').then(({Configuration}) => Configuration)},
];

export const asyncLoadRoute = async (route: Route) => {
    return {...route, component: await route.loadComponent()};
}