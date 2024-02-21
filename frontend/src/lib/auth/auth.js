import { browser } from '$app/environment'
import { makeRequest } from '$lib/api'
import { DB, localStorageObj } from '$lib/db'


export const authenticateUser =  async (Cookies) => {

    let resp=await makeRequest("","GET",{},{},Cookies)
    if(resp?.data?.isauth && !localStorageObj?.data?.user) {

        DB("set","user",resp?.data?.user)
    }
    return resp?.data?.isauth
}