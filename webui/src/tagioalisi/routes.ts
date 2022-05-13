interface Route {
    title: string,
    navText: string,
    url: string,
    isHome?: boolean,
}

export const ROUTES = [
    {title: 'Home', navText: 'Home', url: '/', isHome: true},
    {title: 'Sockpuppet', navText: 'Sockpuppet', url: '/sockpuppet'},
];