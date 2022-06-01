import { ContactPageOutlined } from '@mui/icons-material';
import React, { ReactNode } from 'react';
import { useCookies } from 'react-cookie';


export const useDefaultBaseURL = () => {
    const [cookies] = useCookies(['BOT_EXTERNAL_BASE_URL']);
    const defaultBaseUrl = `${cookies.BOT_EXTERNAL_BASE_URL}`;
    return defaultBaseUrl;
}

export interface APIConfiguration {
    baseURL?: string;
}

interface APIConfigurationContextValue {
    configuration: APIConfiguration,
    setBaseURL: (url: string) => void,
}

export const APIConfigurationContext = React.createContext<APIConfigurationContextValue>({
    configuration: {},
    setBaseURL: (url: string) => {},
});

export default (props: {children: ReactNode}) => {
    const [configuration, setConfiguration] = React.useState<APIConfiguration>({});

    const defaultBaseURL = useDefaultBaseURL();
    React.useEffect(() => {
        console.log('set default base url', defaultBaseURL);
        setConfiguration({baseURL: defaultBaseURL});
    }, [defaultBaseURL]);


    const setBaseURL = (baseURL: string) => {
        console.log('set base url ???', {...configuration, baseURL})
        setConfiguration({...configuration, baseURL})
    };

    React.useEffect(() => {
        console.log('config', configuration);
    }, [configuration]);

    return <APIConfigurationContext.Provider value={{configuration, setBaseURL}}>
        {props.children}
    </APIConfigurationContext.Provider>

}