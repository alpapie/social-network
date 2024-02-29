import { makeRequest } from '$lib/api.js'

export const load = async ({cookies , params})=>{
    let groupid = params.id;
    let response = await makeRequest(`chatgroup?groupid=${groupid}`, "GET", {}, {}, cookies);
    if (response.status ===  200) {
        console.log('gondia : ', response.data);
        
        let res = {
            result : response.data
        }
        return{ res }
    }
}
