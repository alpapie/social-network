import { makeRequest } from "$lib/api.js";
import { authenticateUser } from "$lib/auth/auth.js";
import { localStorageObj } from "$lib/db.js";
import { redirect } from "@sveltejs/kit";
import { generateRandom, saveImage } from "$lib/index.js";


export const load = async ({ cookies }) => {

    console.log(localStorageObj?.data);
    if (localStorageObj?.data?.user) {
        const IsAuth = await authenticateUser(cookies)
        if (IsAuth) {
            redirect(302, "/")
        }
    }
}

const validateFormData = (formData) => {
    const emailRegex = /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/;
    const fields = ['email', 'password', 'firstname', 'lastname', 'birthdate'];

    for (let field of fields) {
        const value = formData.get(field);
        if (!value || value.length < 2) {
            return [false, field];
        }
        if (field === 'email' && !emailRegex.test(value)) {
            return [false, field];
        }
    }

    return [true, ''];
}

export const actions = {
    default: async ({ request, }) => {
        const formDatas = await request.formData()
        let errorMsg = '';
        let [bool, err] = validateFormData(formDatas)
        if (bool == true) {
            let imagePath = '';
            let imageName = formDatas.get('avatar').name

            if (imageName != 'undefined' && imageName != "") {
                let image = await saveImage(formDatas.get('avatar'));
                if (image !== '') {
                    imagePath = image;
                    formDatas.set('avatar', imagePath);
                } else {
                    errorMsg = 'Invalid image size or format';
                }
            }
            if (errorMsg.length == 0) {
                let response = await makeRequest("register", "POST", formDatas)
                if (response.data.success == true) {
                    redirect(302, "/login")
                }
                errorMsg = response.data.error
            }
        } else {
            errorMsg = err.toUpperCase() + ' invalid. Veuillez verifier et reesayer.'
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

