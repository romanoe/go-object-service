import {ref} from 'vue'

interface GoelandObject {
    created_at: Date;
    fk_type: number;
    id: number
}


export function useFetch<T>(url: string): Promise<T> {

    const data = ref<T[]>([]);

    return fetch(url, {
        method: "GET"
    })
        .then(response => {
            if (!response.ok) {
                throw new Error(response.statusText)
            }
            return response.json() as Promise<T>

        })


}

