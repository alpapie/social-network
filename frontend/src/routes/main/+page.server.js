import { makeRequest } from "$lib/api.js";
import { generateRandom,saveImage } from "$lib/index.js";
import { fail } from '@sveltejs/kit';
import { writeFileSync } from "fs";

export async function load({cookies}) {
    // let response = await makeRequest("getPosts","get",{},{},cookies)
    // console.log("response status", await response.status)
    const response = await fetch('http://localhost:8080/server/home', {
        method: 'GET',
        headers: {
            'content-type': 'application/json'
        }
    });

   let  total = await response.json();
    return total
}

export const actions = {
	default : async ({request }) => {
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
        let response = await makeRequest("addPost" , "POST",post)
	}
};