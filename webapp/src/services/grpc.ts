import React from 'react';

import * as grpc from 'nice-grpc-web';

import { AuthenticationContext } from '@tagioalisi/contexts/Authentication';
import { APIConfigurationContext } from '@tagioalisi/contexts/APIConfiguration';


export const useClient = <SD extends grpc.CompatServiceDefinition> (service: SD) => {
    const { configuration } = React.useContext(APIConfigurationContext)
    const { authentication } = React.useContext(AuthenticationContext);

    const client = React.useMemo(() => {
        if (!configuration.grpcBaseURL) {
            throw "No GRPC base URL";
        }
        const channel = grpc.createChannel(configuration.grpcBaseURL);
        const client = grpc.createClient(service, channel);
        return client;


    }, [configuration, authentication]);

    return client;

}
