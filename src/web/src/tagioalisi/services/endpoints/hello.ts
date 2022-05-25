import React from 'react';
import { useAuthentication } from '../auth';
import { usePromiseEffect } from '../async';
import { useAPIConnection } from '../api';

// {"debug_mode":true,"discord_client_id":"674763618008694795","bot_module_blacklist":[""],"group_prefix":"g-"}

export interface HelloResponse {
    pending: boolean;
    done: boolean;
    ok: boolean;
    result?: {
        debugMode: boolean;
        discordClientId: string;
        botModuleBlacklist: string[];
        groupPrefix: string;
    }
    error?: string;
}

export const useHelloQuery = (deps: React.DependencyList = []) => {
    const [response, setResponse] = React.useState<HelloResponse>({
        pending: false, done: false, ok: false,
    });

    const [api] = useAPIConnection();
    React.useEffect(() => {
        setResponse({ pending: true, done: false, ok: false });
        fetch(`${api.baseUrl}`)
            .then(response => response.json())
            .then(responseJSON => setResponse({
                pending: false, done: true, ok: true,
                result: {
                    debugMode: !!responseJSON.debug_mode,
                    discordClientId: `${responseJSON.discord_client_id}`,
                    botModuleBlacklist: responseJSON.bot_module_blacklist || [],
                    groupPrefix: `${responseJSON.group_prefix}`,
                },
            }))
            .catch(err => setResponse({
                pending: false, done: true, ok: false, error: `${err}`,
            }));

    }, [api, ...deps]);

    return response;
}