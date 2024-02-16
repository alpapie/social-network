import { makeRequest } from "$lib/api.js";
import { authenticateUser } from "$lib/auth/auth.js";
import {  localStorageObj } from "$lib/db.js";
import { redirect } from "@sveltejs/kit";


export const load = async ({cookies})=>{

    console.log(localStorageObj?.data);
    if (localStorageObj?.data?.user) {
        const IsAuth= await authenticateUser(cookies)
        if (IsAuth) {
            redirect(302,"/")
        }
    }
}

const validateFormData = (formData) => {
    const emailRegex = /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/;
    const fields = ['email', 'password', 'firstname', 'lastname', 'birthdate'];

    for (let field of fields) {
        const value = formData.get(field);
        if (!value || value.length < 2) {
            return [false,field];
        }
        if (field === 'email' && !emailRegex.test(value)) {
            return [false, field];
        }
    }

    return [true,''];
}

export const actions = {
	default: async ({request,}) => {
		const formDatas= await request.formData()
        // data={
        //     email:formDatas.get()
        // }

        // console.log(validateFormData(formDatas));
        let errorMsg = '';
        let [bool, err] = validateFormData(formDatas)
        if (bool == true) {
            let response = await makeRequest("register","POST",formDatas)
            if (response.data.success == true) {
                redirect(302, "/login")
            }
            errorMsg = response.data.error
        }else{
            errorMsg = err.toUpperCase() +' invalid. Veuillez verifier et reesayer.'
        }
        
        let res = {
            error: errorMsg,
            email: formDatas.get('email'),
            firstname: formDatas.get('firstname'),
            lastname: formDatas.get('lastname'),
            birth: formDatas.get('birthdate'),
            nickname: formDatas.get('nickname'),
            bio: formDatas.get('bio')
            
        }

        return res;
        
	}
};

