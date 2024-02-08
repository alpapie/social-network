import { makeRequest } from "$lib/api";
export async function createGroup(groupData) {
    try {
        const response = await makeRequest('createGroup', 'POST', groupData, {});
        console.log(response);
        if (response.status === 201) {
            console.log('Group created successfully:', response.data);
            return response.data;
        } else {
            console.error('Failed to create group:', response.data);
            throw new Error(response.data.message || 'Failed to create group');
        }
    } catch (error) {
        console.error('Error creating group:', error);
        throw error;
    }
}