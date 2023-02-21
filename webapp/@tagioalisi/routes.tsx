import React from 'react';

import { Home } from '@tagioalisi/components/Home';
import { PageOne } from '@tagioalisi/components/PageOne';
import { PageTwo } from '@tagioalisi/components/PageTwo';
import { ApplicationLayout } from './components/ApplicationLayout';
import { ApplicationError } from './components/ApplicationError';

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
