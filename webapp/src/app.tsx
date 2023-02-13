// Common styles/fonts, only import once
import 'tagioalisi/styles/common.scss';

import React from 'react';

Promise.resolve().then(async () => {
    const container = document.getElementById('app-wrapper');
    if (!container) {
        throw 'Could not find a container component';
    }

    const { createRoot } = await import('react-dom/client');
    const { lazy } = await import('react');
    const ApplicationRoot = lazy(() => import('tagioalisi/components/ApplicationRoot'));
    createRoot(container).render(<ApplicationRoot />);
});