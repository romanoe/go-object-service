// Inspired by Create a Basic useFetch Hook in Vue.js (https://javascript.plainenglish.io/create-a-basic-usefetch-hook-in-vue-b3ff113872d7), Medium, March 2022
import {reactive, toRefs} from 'vue';

interface State<T> {
    isLoading: boolean;
    hasError: boolean;
    errorMessage: string;
    data: T | null;
}


export const useFetch = async <T>(url: string, options?: Record<string, unknown>) =>
{

    const state = reactive<State<T>>({
            isLoading: true,
            hasError: false,
            errorMessage: '',
            data: null,
        })

    const fetchData = async() => {

        state.isLoading = true;

        try {
            const res = await fetch(url, options);

            if (!res.ok) {
                throw new Error(res.statusText);
            }

            state.data = await res.json();

        } catch (err: unknown) {
            const typedError = err as Error;
            state.hasError = true;
            state.errorMessage = typedError.message;
        } finally {
            state.isLoading = false
        }
    };
        await fetchData()

        return {...toRefs(state)}

}

