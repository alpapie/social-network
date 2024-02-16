import { makeRequest } from "$lib/api.js"
import { authenticateUser } from "$lib/auth/auth"
import { error, redirect } from "@sveltejs/kit"

import { createGroup } from "$lib/groups/createGroup";

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
  createGroup: async ({request,cookies}) => {
        const formDatas= await request.formData()
        let data={
            title:formDatas.get("title"),
            description:formDatas.get("description"),
        }
          try {
          const response = await createGroup(data,cookies);
          console.log('Group created successfully:', response);
        } catch (error) {
          console.error('Failed to create group:', error);
        }
    }
};
