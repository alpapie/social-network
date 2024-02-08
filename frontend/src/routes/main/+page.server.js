import { makeRequest } from "$lib/api.js";

export async function load({cookies}) {
    // let response = await makeRequest("getPosts","get",{},{},cookies)
    // console.log("response status", await response.status)
    const response = await fetch('http://localhost:8080/server/getPosts', {
        method: 'GET',
        headers: {
            'content-type': 'application/json'
        }
    });

   let  total = await response.json();
    return {
        
        data : total
    }
}

export const actions = {
	"createPost": async ({request , cookies}) => {
		// TODO log the user in
        let data = await  request.formData()
        console.log("here is the event" ,data )
	}
};