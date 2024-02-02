import { authenticateUser } from "$lib/auth/auth.js";
import { redirect } from "@sveltejs/kit";


export const load =(event)=>{
    authenticateUser(event)
}

export const actions = {
	default: async ({request}) => {
		const formData= await request.formData()
        console.log(formData.getAll('password'));
	}
};

