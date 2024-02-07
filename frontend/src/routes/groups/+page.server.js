import { makeRequest } from "$lib/api";

export const load = async ({cookies}) => {
    let response = await makeRequest("groups", "GET", {}, {}, cookies);
    if (response.status ===  200) {
        console.log('les donnees : ', response.data);

        // groups = response.data
        return {
            groups: response.data
        };
    }
};
