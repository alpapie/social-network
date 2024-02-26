import { makeRequest } from "$lib/api.js"
import { authenticateUser } from "$lib/auth/auth"
import { localStorageObj } from "../db.js"

import { error, redirect } from "@sveltejs/kit"


export const load= async ({cookies})=>{
    const IsAuth=await authenticateUser(cookies)
    if (!IsAuth) {
        redirect(302,"/login")
    }

    const response= await makeRequest("notificatons","get",{},{},cookies)
    if (response?.data?.success) {
        return{notif: response?.data, user:localStorageObj.data.user}
    }

    throw error(400,"bad request")
}