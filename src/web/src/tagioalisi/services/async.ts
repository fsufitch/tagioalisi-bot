import { useEffect, useState } from 'react';

type PromiseFunction<T> = () => Promise<T>;

interface usePromiseState<T> {
    result?: T,
    done: boolean,
    error?: any,
}

export function usePromiseEffect<T>(func: PromiseFunction<T>, deps?: React.DependencyList) {
    const [{ result, done, error }, setState] = useState<usePromiseState<T>>({ done: false });

    useEffect(() => {
        func()
            .then(value => setState({ done: true, result: value }))
            .catch(reason => setState({ done: true, error: reason || 'unknown failure'}))
    }, deps);
    return [result, done, error] as [T | undefined, boolean, any];

}