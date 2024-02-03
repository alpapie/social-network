import { makeRequest } from "$lib/api.js";
import { authenticateUser } from "$lib/auth/auth.js";
import { redirect } from "@sveltejs/kit";


export const load =(event)=>{
    authenticateUser(event)
}

export const actions = {
	default: async ({request,}) => {
		const formDatas= await request.formData()
        // data={
        //     email:formDatas.get()
        // }
        console.log(formDatas);
        let response= await makeRequest("login","POST",formDatas)
        console.log("eeeerrrrrrrrrrrrrrrr alpapie " +response)
        // return response
	}
};

