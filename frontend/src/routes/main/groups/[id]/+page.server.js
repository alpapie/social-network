import { makeRequest } from "$lib/api";
import axios from 'axios';


export const actions = {
    default: async ({request, cookies, params}) => {
        const formDatas = await request.formData();
        const groupId = params.id;
        formDatas.append('groupId', groupId);
        console.log(formDatas);
        

         const title = formDatas.get('title');
         const description = formDatas.get('description');
         const date = formDatas.get('date');
         const time = formDatas.get('time');

 
         let errorMsg = '';
 
         if (title.length <=  2 || description.length <=  2) {
             errorMsg = 'Invalid input entered';
         }
 
         const today = new Date();
         const eventDate = new Date(date);
         if (eventDate <= today) {
             errorMsg = 'Invalid Date';
         }
 
         const timeRegex = /^([01]?[0-9]|2[0-3]):[0-5][0-9]$/;
         if (!timeRegex.test(time)) {
             errorMsg = 'invalid Time (format HH:MM).';
         }
        

         if (errorMsg.length < 1) {
            let response = await makeRequest("createEvent","POST",formDatas,{},cookies)
            if (response.data.success == false) {
                errorMsg = response.data.error
            }else{
                errorMsg = 'no'
            }
         }

        let res = {
            error: errorMsg,
            title: formDatas.get('title'),
            description: formDatas.get('description'),
            date: formDatas.get('date'),
            time: formDatas.get('time'),            
        }

        return res;
    }

};

// export const load = async ({cookies, params}) => {
    

    
// };
export const load = async ({cookies, params}) => {
    const groupId = params.id;
    console.log("grID => ", groupId);
    const data = {groupId: groupId}
    let response = await makeRequest(`getgroupdetail/${groupId}`, "GET", data, {}, cookies);
    if (response.status ===  200) {
        console.log('les donnees : ', response.data);
        
        let res = {
            result : response.data
        }
        return{ res }
    }
};
// export const load = async ({cookies, params}) => {
//     const groupId = params.id;
//     let data;
//     try {
//         const response = await axios.get(`http://localhost:8080/server/groupdetail/${groupId}`);
//         data = response.data; 
//     } catch (error) {
//         console.error('Error fetching data:', error);

//     }
//     return { data };
// };
// export const load = async ({cookies, params}) => {
//     const groupId = params.id;
//     let data;
//     try {
//         const response = await axios.get(`http://your-api-url/endpoint/${groupId}`, {
//             headers: {
//                 Cookie: `session=${cookies.session}`
//             }
//         });
//         data = response.data;
//     } catch (error) {
//         console.error('Error fetching data:', error);
//     }
//     return { data };
// };