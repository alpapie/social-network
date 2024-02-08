import { authenticateUser } from "$lib/auth/auth"

export const load = async ({cookies})=>{
    const IsAuth=await authenticateUser(cookies)
    if (!IsAuth) {
        redirect(302,"/login")
    }
    
}