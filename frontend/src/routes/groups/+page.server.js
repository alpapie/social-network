import { makeRequest } from "$lib/api";

export const load = async ({params, cookies}) => {
    let response = await makeRequest("groups", "GET", {}, {}, cookies);
    if (response.status ===  200) {
        console.log('les donnees : ', response.data);
        
        let res = {
            groups : response.data
        }
        return{ res }
        // return {
        //     groups: response.data[0].title
        // };
    }
};
