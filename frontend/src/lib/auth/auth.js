import { makeRequest } from '$lib/api'

export const authenticateUser =  async (Cookies) => {

    let resp=await makeRequest("","GET",{},{},Cookies) 
    return resp?.data?.isauth
}