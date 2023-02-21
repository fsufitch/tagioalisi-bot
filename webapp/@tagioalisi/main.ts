import * as React from 'react';
import { createRoot } from 'react-dom/client';
import { ApplicationRoot } from './components/ApplicationRoot';

(() =>
  createRoot(document.getElementById('app-container') as HTMLElement).render(
    React.createElement(ApplicationRoot, { reactStrict: true }),
  ))();
