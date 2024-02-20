import { makeRequest } from "$lib/api";

export const load = async ({cookies}) => {
    let response = await makeRequest("groups", "GET", {}, {}, cookies);
    console.log(response);
    if (response.status ===  200) {
        console.log('gondia : ', response.data);
        
        let res = {
            result : response.data
        }
        return{ res }
    }
};

