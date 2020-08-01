import { useState, useEffect } from "react";

const LOCAL_STORAGE_UPDATED_EVENT = 'storage-local';
interface LocalStorageUpdatedEventDetail<T> {
    key: string;
    oldValue: T;
    newValue: T;
}

export function useLocalStorage<T>(key: string, defaultValue: T, from?: string): [T, (val: T | (() => T)) => void] {
    const serialize = (val: T) => JSON.stringify(val);
    const deserialize = (val: string | null) => JSON.parse(val ?? 'null') as T;
    const save = (val: T) => window.localStorage.setItem(key, serialize(val));
    const load = () => deserialize(window.localStorage.getItem(key));

    const [value, setValue] = useState<T>(() => {
        let initialValue = defaultValue;
        try {
            initialValue = load();
        } catch (err) {
            console.error(`Failed loading extant value from localStorage: ${err}`);
            save(initialValue);
        }
        return initialValue;
    });

    const setValueAndNotify = (_newValue: T | (() => T)) => {
        const oldValue = load();
        const newValue = (_newValue instanceof Function) ? _newValue() : _newValue;
        save(newValue);
        setValue(newValue);
        console.log('storage DIRECT setValue', {from, key, newValue});
        window.dispatchEvent(new CustomEvent(LOCAL_STORAGE_UPDATED_EVENT, {detail: {key, oldValue, newValue} as LocalStorageUpdatedEventDetail<T>}));
    }

    useEffect(() => {
        const handleLocalStorageUpdated = (ev: Event) => {
            if (!(ev instanceof CustomEvent)) {
                return;
            }
            const detail = ev.detail as LocalStorageUpdatedEventDetail<T>;
            if (detail.key !== key || detail.oldValue === detail.newValue) {
                return;
            }
            setValue(detail.newValue);
            console.log('storage EVENT setValue', {from, key, newValue: detail.newValue});
        }
        window.addEventListener(LOCAL_STORAGE_UPDATED_EVENT, handleLocalStorageUpdated);

        setTimeout(() => {
            setValueAndNotify(load());
        }, 10000);

        return () => window.removeEventListener(LOCAL_STORAGE_UPDATED_EVENT, handleLocalStorageUpdated);
    }, []);

    return [value, setValueAndNotify];
}

