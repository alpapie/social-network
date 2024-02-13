import { makeRequest } from "$lib/api.js";
import { generateRandom,saveImage } from "$lib/index.js";
import { fail } from '@sveltejs/kit';
import { error, redirect } from "@sveltejs/kit"
import { authenticateUser } from "$lib/auth/auth"
import { writeFileSync } from "fs";

export const load = async ({cookies})=>{
    const IsAuth=await authenticateUser(cookies)
    if (!IsAuth) {
        redirect(302,"/login")
    }
    const response= await makeRequest("home","get",{},{},cookies)
    if (response?.data?.success) {
        return response?.data;
    }
    throw error(400,"bad request")
}


export const actions = {
	default : async ({request , cookies}) => {
		// TODO log the user in
        let data = await  request.formData()
        let content = data.get('content')
        if (content == "") {
			console.log("fail content")
			return fail(400, {content , missing: true})
		}

        let post = {
            titre : "",
            content : content,
            image : await saveImage(data.get("avatar")),
            privacy : data.get("privacy"),
            allowedusers : data.getAll("allowedusers").map((v) => Number(v))
        }
        let response = await makeRequest("addPost" , "POST",post , {} ,cookies)
	}
};