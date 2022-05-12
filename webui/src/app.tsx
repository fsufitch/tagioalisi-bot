// import 'jquery';
// import 'bootstrap';
import 'core-js';
import "regenerator-runtime/runtime";

import React from 'react';

Promise.resolve().then(async () => {
    const container = document.getElementById('app-wrapper');
    if (!container) {
        throw 'Could not find a container component';
    }

    const { createRoot } = await import('react-dom/client');
    const { Root } = await import('tagioalisi/components/Root');
    createRoot(container).render(<Root />);
});
