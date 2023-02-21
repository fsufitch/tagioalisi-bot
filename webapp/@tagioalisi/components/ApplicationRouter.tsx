import * as React from 'react';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { PageOne } from './PageOne';
import { PageTwo } from './PageTwo';
import { Home } from './Home';

// const ROUTER = createBrowserRouter(
//   createRoutesFromElements(
//     <Route>
//       <Route path="/" element={<Home />}>
//         <PageOneRoute />
//         <PageTwoRoute />
//       </Route>
//     </Route>,
//   ),
// );

const ROUTER = createBrowserRouter([
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
]);

export const ApplicationRouter: React.FC = () => {
  return <RouterProvider router={ROUTER} />;
};
