import React from 'react';
import { useAuthentication } from '../auth';
import { usePromiseEffect } from '../async';
import { useAPIConnection } from '../api';
import { useSynchronizedState } from '../state';
import { appendFileSync } from 'fs';


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

const CACHE_KEY = 'tagioalisi/api/hello/cache';
const CACHE_TTL_MS = 1000;

export const useHelloQuery = (deps: React.DependencyList = []) => {
    const [cachedResponse, setCachedResponse] = useSynchronizedState<CachedHelloResponse>(CACHE_KEY, { timeUnixMillis: 0, ttlMillis: 0 }, JSON.stringify, JSON.parse);
    const [api] = useAPIConnection();

    // XXX: improve multi-query on page load? idk lol.
    
    React.useEffect(() => {
        const nowUnixMillis = Date.now();
        if (nowUnixMillis - cachedResponse.timeUnixMillis < cachedResponse.ttlMillis) {
            return; // cache hit, nothing to do
        }
        console.log(`CACHE MISS @ ${nowUnixMillis}`, cachedResponse);
        setCachedResponse({ timeUnixMillis: nowUnixMillis, ttlMillis: 5000, response: { pending: true, done: false, ok: false } });
        fetch(`${api.baseUrl}`)
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

    }, [api, ...deps]);

    return cachedResponse.response ?? { done: false, ok: false, pending: false };
}