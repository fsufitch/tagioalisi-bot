import { useEffect, useState } from "react";

const API_ENDPOINT_UPDATE_EVENT = 'tagioalisi-api-endpoint-update';

let API_ENDPOINT = process.env.API_ENDPOINT || '';

export function useAPIEndpoint(): [string, (endpoint: string) => void] {
    const [endpoint, setEndpoint] = useState(API_ENDPOINT);

    const setEndpointAndNotify = (endpoint: string) => {
        API_ENDPOINT = endpoint;
        window.dispatchEvent(new Event(API_ENDPOINT_UPDATE_EVENT));
    }

    useEffect(() => {
        window.addEventListener(API_ENDPOINT_UPDATE_EVENT, () => {
            setEndpoint(API_ENDPOINT);
        });
    }, []);

    return [endpoint, setEndpointAndNotify];

}