import { makeRequest } from "$lib/api.js";
import { authenticateUser } from "$lib/auth/auth.js";
import { redirect } from "@sveltejs/kit";


export const load = async ({cookies})=>{
    const IsAuth=await authenticateUser(cookies)
    if (!IsAuth) {
        redirect(302,"/login")
    }
}