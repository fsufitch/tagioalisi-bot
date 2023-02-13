import React, { ReactNode } from "react";
import ColorModeProvider from './ColorMode';
import AuthenticationProvider from './Authentication';
import APIConfigurationProvider from './APIConfiguration';
import StorageProvider from './Storage';

// List of contexts to create providers for, in reverse order of nesting (innermost first)
const CONTEXT_NODE_LIST: ((props: {children: ReactNode}) => JSX.Element)[] = [
    AuthenticationProvider,
    ColorModeProvider,
    APIConfigurationProvider,
    StorageProvider,
]

export default (props: {children: ReactNode}) => {
    // Composite all the various contexts Tagioalisi requires

    let inner = <>{props.children}</>;

    for (const contextElement of CONTEXT_NODE_LIST) {
        inner = React.createElement(contextElement, {children: inner});
    }

    return <>{inner}</>;
}