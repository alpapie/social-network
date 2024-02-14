import { makeRequest } from "$lib/api";
export const actions = {
    default: async ({request, cookies}) => {
        const formDatas = await request.formData();
        console.log(formDatas);
        let response = await makeRequest("createEvent","POST",formDatas,{},cookies)
        console.log(response);

        return {
            closed: true 
        };
    }
};