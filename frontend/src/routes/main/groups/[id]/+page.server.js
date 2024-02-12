

export const actions = {
	default: async ({request,cookies}) => {
        
		const formDatas= await request.formData()

        console.log(formDatas);
        return {
            title: true,
        };

        // let response= await makeRequest("login","POST",formDatas,{},cookies)

        // return {error:response?.data?.error,email:formDatas.get("email")}
	}
};