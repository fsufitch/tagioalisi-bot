import React from 'react';
import {grpc} from '@improbable-eng/grpc-web';

import { GreeterClient } from 'tagioalisi/proto/hello_pb_service';
import { SockpuppetClient } from 'tagioalisi/proto/sockpuppet_pb_service';
import { AuthenticationContext } from 'tagioalisi/contexts/Authentication';
import { APIConfigurationContext } from 'tagioalisi/contexts/APIConfiguration';

export const useGreeterClient = () => {
    const { configuration } = React.useContext(APIConfigurationContext)
    const { authentication } = React.useContext(AuthenticationContext);
    
    const transport = grpc.WebsocketTransport();
    const client = React.useMemo(() => new GreeterClient(configuration.baseURL ?? '', {transport}), [configuration, authentication]);
    
    return client;
}

export const useSockpuppetClient = () => {
    const { configuration } = React.useContext(APIConfigurationContext)
    const { authentication } = React.useContext(AuthenticationContext);
    
    const transport = grpc.WebsocketTransport();
    const client = React.useMemo(() => new SockpuppetClient(configuration.baseURL ?? '', {transport}), [configuration, authentication]);
    
    return client;

}
