import { makeRequest } from '$lib/api.js'
import { authenticateUser } from '$lib/auth/auth.js'
import { error, redirect } from '@sveltejs/kit'

export const load = async ({cookies})=>{
    const IsAuth=await authenticateUser(cookies)
    if (!IsAuth) {
        redirect(302,"/login")
    }
    const response= await makeRequest("unfollower","get",{},{},cookies)
    if (response?.data?.success) {
        return response?.data;
    }
    throw error(400,"bad request")
}

