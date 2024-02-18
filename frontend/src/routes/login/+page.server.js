import { makeRequest } from "$lib/api.js";
import { authenticateUser } from "$lib/auth/auth.js";
import { redirect } from "@sveltejs/kit";
import { DB, localStorageObj } from "$lib/db.js";
import { error } from '@sveltejs/kit';

export const load = async ({cookies})=>{
    if (localStorageObj?.data?.user) {
        const IsAuth= await authenticateUser(cookies)
        if (IsAuth) {
            redirect(302,"/")
        }
    }
}

export const actions = {
    default: async ({ request, cookies }) => {
        const formDatas = await request.formData()
        let response = await makeRequest("login", "POST", formDatas, {}, cookies)
        if (response?.data?.success) {
            DB("set","user",response?.data?.user)            
     
            cookies.set('sessionId',response?.data?.data, {
                httpOnly: false,
                sameSite: 'strict',
                secure: false,
                path: '/',
                maxAge: 3600 * 24 * 3,
            });
            redirect(302, "/")
        }
        
        if (response?.data?.error?.Code) {
            throw  error(response?.data?.error?.Code, response?.data?.error?.Message);
        }
        return { error: response?.data?.error, email: formDatas.get("email") }
    }
};
