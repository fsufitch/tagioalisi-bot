import React from 'react';
import { APIConfigurationContext } from '../../contexts/APIConfiguration';
import { StorageContext } from '../../contexts/Storage';


export interface HelloResponse {
    pending: boolean;
    done: boolean;
    ok: boolean;
    result?: {
        debugMode: boolean;
        discordClientId: string;
        botModuleBlacklist: string[];
        groupPrefix: string;
        uptimeSeconds: number;
    }
    error?: string;
}

interface CachedHelloResponse {
    timeUnixMillis: number,
    ttlMillis: number,
    response?: HelloResponse,
}

const CACHE_KEY = 'hello/cache';
const CACHE_TTL_MS = 1000;

export const useHelloQuery = (deps: React.DependencyList = []) => {
    const { state } = React.useContext(StorageContext);
    const [ cachedResponse, setCachedResponse ]  = state?.useJSON<CachedHelloResponse>(CACHE_KEY) ?? [null, () => {}];
    const { configuration } = React.useContext(APIConfigurationContext);

    // XXX: improve multi-query on page load? idk lol.
    
    React.useEffect(() => {
        if (!setCachedResponse) return;
        const cacheTimeUnixMillis = cachedResponse?.timeUnixMillis || 0;
        const cacheTTLMillis = cachedResponse?.ttlMillis || 0;
        const nowUnixMillis = Date.now()
        if (nowUnixMillis - cacheTimeUnixMillis < cacheTTLMillis) {
            return; // cache hit, nothing to do
        }
        console.log(`CACHE MISS @ ${nowUnixMillis}`, cachedResponse);
        setCachedResponse({ timeUnixMillis: nowUnixMillis, ttlMillis: 5000, response: { pending: true, done: false, ok: false } });
        fetch(`${configuration.baseURL}`)
            .then(response => response.json())
            .then(responseJSON => setCachedResponse({
                timeUnixMillis: nowUnixMillis,
                ttlMillis: CACHE_TTL_MS,
                response: {
                    pending: false, done: true, ok: true,
                    result: {
                        debugMode: !!responseJSON.debug_mode,
                        discordClientId: `${responseJSON.discord_client_id}`,
                        botModuleBlacklist: responseJSON.bot_module_blacklist || [],
                        groupPrefix: `${responseJSON.group_prefix}`,
                        uptimeSeconds: responseJSON.uptime_seconds || 0.0,
                    },
                },
            }))
            .catch(err => setCachedResponse({
                timeUnixMillis: nowUnixMillis,
                ttlMillis: CACHE_TTL_MS,
                response: {
                    pending: false, done: true, ok: false, error: `${err}`,
                },
            }));

    }, [configuration, ...deps]);

    return cachedResponse?.response ?? { done: false, ok: false, pending: false };
}