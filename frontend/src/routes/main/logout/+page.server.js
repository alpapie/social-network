import { makeRequest } from "$lib/api"
import { redirect } from "@sveltejs/kit"
export const load=async ({cookies})=>{
   let response= await makeRequest("logout","get",{},{},cookies)
   if (response?.data?.success) {
       cookies.delete("sessionId", { path: '/' })
       redirect(302,"/login")
   }
}