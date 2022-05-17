import { useEffect, useState } from "react";

import process from 'process';

const API_ENDPOINT_UPDATE_EVENT = 'tagioalisi-api-endpoint-update';

const trimEnd = (s: string, suffix: string) => {
    while (s.endsWith(suffix)) {
        s = s.slice(0, s.length - suffix.length);
    }
    return s;
}

let API_ENDPOINT = trimEnd(process.env.API_ENDPOINT || '', '/');

export function useAPIEndpoint(): [string, (endpoint: string) => void] {
    const [endpoint, setEndpoint] = useState(API_ENDPOINT);

    const setEndpointAndNotify = (endpoint: string) => {
        API_ENDPOINT = trimEnd(endpoint, '/');
        window.dispatchEvent(new Event(API_ENDPOINT_UPDATE_EVENT));
    }

    useEffect(() => {
        window.addEventListener(API_ENDPOINT_UPDATE_EVENT, () => {
            setEndpoint(API_ENDPOINT);
        });
    }, []);

    return [endpoint, setEndpointAndNotify];

}