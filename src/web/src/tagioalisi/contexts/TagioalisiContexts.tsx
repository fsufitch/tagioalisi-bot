import React, { ReactNode } from "react";
import ColorMode from './ColorMode';

const CONTEXT_NODE_LIST: ((props: {children: ReactNode}) => JSX.Element)[] = [
    ColorMode,
]

export default (props: {children: ReactNode}) => {
    // Composite all the various contexts Tagioalisi requires

    let inner = <>{props.children}</>;

    for (const contextElement of CONTEXT_NODE_LIST) {
        inner = contextElement({children: inner});
    }

    return <>{inner}</>;
}