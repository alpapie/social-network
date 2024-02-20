import { makeRequest } from "$lib/api.js"
import { authenticateUser } from "$lib/auth/auth"
import { error, redirect } from "@sveltejs/kit"

export const load= async ({cookies})=>{
    const IsAuth=await authenticateUser(cookies)
    if (!IsAuth) {
        redirect(302,"/login")
    }

    const response= await makeRequest("notificatons","get",{},{},cookies)
    if (response?.data?.success) {
        return response?.data;
    }

    throw error(400,"bad request")
}