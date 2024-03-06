import { makeRequest } from "$lib/api.js"
import { authenticateUser } from "$lib/auth/auth"
import { error, redirect } from "@sveltejs/kit"

export const load = async ({ cookies, params }) => {
    const IsAuth = await authenticateUser(cookies)
    if (!IsAuth) {
        redirect(302, "/login")
    }
    const response = await makeRequest("profile_follow?user_id=" + params.id, "get", {}, {}, cookies)

    if (response?.data.success) {
        return response?.data
    }

    throw error(response?.data?.error?.Code, response?.data?.error?.Message)
}