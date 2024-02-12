
import { makeRequest } from "$lib/api.js";
import { authenticateUser } from "$lib/auth/auth.js";
export const load = async ({cookies})=>{
  const IsAuth= await authenticateUser(cookies)
}
import { redirect } from "@sveltejs/kit";
import { DB, localStorageObj } from "../../db.js";
import { error } from '@sveltejs/kit';
import { createGroup } from "$lib/groups/createGroup";



export const actions = {
  'createGroup': async ({request,cookies}) => {
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