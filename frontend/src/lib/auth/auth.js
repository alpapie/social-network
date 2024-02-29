import { makeRequest } from '$lib/api'

import { redirect } from '@sveltejs/kit'
import axios from 'axios'
import { DB ,localStorageObj} from '../../routes/db'


export const authenticateUser =  async (Cookies) => {

    let resp=await makeRequest("","GET",{},{},Cookies)
    // if(resp?.data?.isauth && !localStorageObj?.data?.user) {
    //     console.log("locale data user ",localStorageObj?.data);
        DB("set","user",resp?.data?.user)
    // }
    return resp?.data?.isauth
}
