import { makeRequest } from "$lib/api";

export const load = async ({cookies}) => {
    let response = await makeRequest("groups", "GET", {}, {}, cookies);
    if (response.status ===  200) {
        
        let res = {
            result : response.data
        }
        return{ res }
    }
};

