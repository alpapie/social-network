import { makeRequest } from "$lib/api.js";
import { generateRandom, saveImage } from "$lib/index.js";
import { fail } from '@sveltejs/kit';
import { error, redirect } from "@sveltejs/kit"
import { authenticateUser } from "$lib/auth/auth"
import { localStorageObj } from "$lib/db.js";

export const load = async ({ cookies }) => {
    const IsAuth = await authenticateUser(cookies)
    if (!IsAuth) {
        redirect(302, "/login")
    }
    const response = await makeRequest("home", "get", {}, {}, cookies)
    if (response?.data?.success) {
        return response?.data;
    }
    throw error(400, "bad request")
}

export const actions = {
    createPost: async ({ request, cookies }) => {
        // TODO log the user in
        let data = await request.formData()
        let content = data.get('content')
        if (content == "") {
            console.log("fail content")
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
            firstName: user.FirstName,
            lastName: user.LastName,
        }

        const response = await makeRequest("addComment", "POST", comment, {}, cookies)
        console.log("comment value", comment);

        if (response.status == 200) {
            return comment
        } else {
            return { success: false }
        }
        // let responseData = response.json()
        // console.log("COMMENT HANDLER RESPONSE ", response)
        return response?.data
    }
};