import { makeRequest } from "$lib/api.js"
import { authenticateUser } from "$lib/auth/auth"
import { error, redirect } from "@sveltejs/kit"
import { generateRandom, saveImage } from "$lib/index.js";
import { fail } from '@sveltejs/kit';

import { createGroup } from "$lib/groups/createGroup";
import { localStorageObj } from "../db.js";


export const load = async ({cookies})=>{
    const IsAuth=await authenticateUser(cookies)
    if (!IsAuth) {
        redirect(302,"/login")
    }

    const response= await makeRequest("home","get",{},{},cookies)
    if (response?.data?.success) {
        console.log("DATA FROM HOME" , response.data.user)
        return response?.data;
    }
    throw error(400,"bad request")
}

export const actions = {
  createGroup: async ({request,cookies}) => {
        const formDatas= await request.formData()
        let data={
            title:formDatas.get("title"),
            description:formDatas.get("description"),
        }
          try {
          const response = await createGroup(data,cookies);
        } catch (error) {
          console.error('Failed to create group:', error);

        }
    },
    createPost: async ({ request, cookies }) => {
        // TODO log the user in
        let data = await request.formData()
        let content = data.get('content')
        if (content == "") {
            return fail(400, { content, missing: true })
        }
        let post = {
            titre: "",
            content: content,
            image: await saveImage(data.get("avatar")),
            privacy: data.get("privacy"),
            allowedusers: data.getAll("allowedusers").map((v) => Number(v))
        }
        let response = await makeRequest("addPost", "POST", post, {}, cookies)
    },

    createComment: async ({ request, cookies }) => {
        // TODO log the user in
        let data = await request.formData()
        let content = data.get("comment")

        let postId = data.get("postId")
        // console.log("formcomment  ", content, postId)
        let user = JSON.parse(localStorageObj?.data?.user)

        let comment = {
            postId: Number(postId),
            comment: content,
            image: await saveImage(data.get("avatar")),
            firstName: user.FirstName,
            lastName: user.LastName,
        }

        const response = await makeRequest("addComment", "POST", comment, {}, cookies)
        if (response.status == 200) {
            return {succes : true , data : comment}
        } else {
            return { success: false }
        }
       
    }
};
