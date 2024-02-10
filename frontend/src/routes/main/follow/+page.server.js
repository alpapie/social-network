import { makeRequest } from '$lib/api.js'
import { authenticateUser } from '$lib/auth/auth.js'

export const load = async ({cookies})=>{
    const IsAuth=await authenticateUser(cookies)
    if (!IsAuth) {
        redirect(302,"/login")
    }
    const response= await makeRequest("home","get",{},{},cookies)
    return response?.data;
}
