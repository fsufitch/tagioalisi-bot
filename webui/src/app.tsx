import 'jquery';
import 'bootstrap';

import React from 'react';
import ReactDOM from 'react-dom';

import { Root } from 'tagioalisi/components/Root';

const wrapper = document.getElementById('app-wrapper');

const appPlaceholder = (<Root />);

if (wrapper) {
    ReactDOM.render(appPlaceholder, wrapper);
} else {
    console.error("No wrapper element found");
}