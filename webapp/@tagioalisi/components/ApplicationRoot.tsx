import * as React from 'react';
import { ReactNode } from 'react';
// import { createBrowserRouter } from 'react-router-dom';
import { ApplicationRouter } from './ApplicationRouter';

export const ApplicationRoot: React.FC<{ reactStrict: boolean }> = ({ reactStrict }) => {
  return <StrictModeToggler reactStrict={reactStrict}> hello </StrictModeToggler>;
};

const StrictModeToggler: React.FunctionComponent<{ reactStrict: boolean; children: ReactNode }> = ({
  reactStrict,
  children,
}) =>
  reactStrict ? (
    <React.StrictMode>
      <ApplicationRouter />
    </React.StrictMode>
  ) : (
    <> {children}</>
  );
