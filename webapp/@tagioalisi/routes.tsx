import React, { lazy } from 'react';

const ApplicationLayout = lazy(() => import('@tagioalisi/components/ApplicationLayout'));
const Home = lazy(() => import('@tagioalisi/components/Home'));
const PageOne = lazy(() => import('@tagioalisi/components/PageOne'));
const PageTwo = lazy(() => import('@tagioalisi/components/PageTwo'));
const ApplicationError = lazy(() => import('@tagioalisi/components/ApplicationError'));

export default [
  {
    id: 'app-layout',
    path: '/',
    element: <ApplicationLayout />,
    children: [
      {
        id: 'home',
        path: '/',
        element: <Home />,
      },
      {
        id: 'page-one',
        path: 'page-one',
        element: <PageOne />,
      },
      {
        id: 'page-two',
        path: 'page-two',
        element: <PageTwo />,
      },
    ],
    errorElement: <ApplicationError />,
  },
];
