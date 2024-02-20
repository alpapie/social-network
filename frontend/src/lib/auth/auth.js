import { browser } from '$app/environment'
import { makeRequest } from '$lib/api'
import { InitialiseSocket, NotifSocket } from '$lib/socket'


export const authenticateUser =  async (Cookies) => {

    let resp=await makeRequest("","GET",{},{},Cookies)
    return resp?.data?.isauth
}