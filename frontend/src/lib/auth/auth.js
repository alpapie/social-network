import { makeRequest } from '$lib/api'
import { DB, localStorageObj } from '$lib/db'
import { redirect } from '@sveltejs/kit'
import axios from 'axios'


export const authenticateUser =  async (Cookies) => {

    let resp=await makeRequest("","GET",{},{},Cookies)
    if(resp?.data?.isauth && !localStorageObj?.data?.user) {
        console.log("locale data user ",localStorageObj?.data);
        DB("set","user",resp?.data?.user)
    }
    return resp?.data?.isauth
}

export const logOutUser = async (cookie) => {
    let url = `http://localhost:8080/server/logout`;
    let header={
        cookie:cookie
    }
    const config = { method:"get",withCredentials: true , header,mode: 'no-cors', }
   
    const response = await axios(url,config);
    
    if (response?.data?.success) {
        // cookies.delete("sessionId", { path: '/' })
        redirect(302,"/login")
    }
}