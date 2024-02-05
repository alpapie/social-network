import { makeRequest } from "$lib/api.js";
import { authenticateUser } from "$lib/auth/auth.js";
import { redirect } from "@sveltejs/kit";


export const load = async ({cookies})=>{
    const IsAuth= await authenticateUser(cookies)
    if (IsAuth) {
        redirect(302,"/")
    }
}

export const actions = {
	default: async ({request,}) => {
		const formDatas= await request.formData()
        let response= await makeRequest("register","POST",formDatas)
        console.log(response.data)
        if(response?.data?.success){
            redirect(301,"/login")
        }
        return response.data
	}
};

